package env0

import (
	"context"
	"errors"
	"log"

	"github.com/env0/terraform-provider-env0/client"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,

		Importer: &schema.ResourceImporter{StateContext: resourceProjectImport},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Description:      "name to give the project",
				Required:         true,
				ValidateDiagFunc: ValidateNotEmptyString,
			},
			"id": {
				Type:        schema.TypeString,
				Description: "id of the project",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "description of the project",
				Optional:    true,
			},
			"force_destroy": {
				Type:        schema.TypeBool,
				Description: "Destroy the project even when environments exist",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(client.ApiClientInterface)

	var payload client.ProjectCreatePayload
	if err := readResourceData(&payload, d); err != nil {
		return diag.Errorf("schema resource data deserialization failed: %v", err)
	}

	project, err := apiClient.ProjectCreate(payload)
	if err != nil {
		return diag.Errorf("could not create project: %v", err)
	}

	d.SetId(project.Id)

	return nil
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(client.ApiClientInterface)

	project, err := apiClient.Project(d.Id())
	if err != nil {
		return ResourceGetFailure("project", d, err)
	}

	if err := writeResourceData(&project, d); err != nil {
		return diag.Errorf("schema resource data deserialization failed: %v", err)
	}

	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(client.ApiClientInterface)

	id := d.Id()
	var payload client.ProjectCreatePayload

	if err := readResourceData(&payload, d); err != nil {
		return diag.Errorf("schema resource data deserialization failed: %v", err)
	}

	if _, err := apiClient.ProjectUpdate(id, payload); err != nil {
		return diag.Errorf("could not update project: %v", err)
	}

	return nil
}

func resourceProjectAssertCanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	forceDestroy := d.Get("force_destroy").(bool)
	if forceDestroy {
		return nil
	}

	apiClient := meta.(client.ApiClientInterface)

	id := d.Id()

	envs, err := apiClient.ProjectEnvironments(id)
	if err != nil {
		return err
	}

	for _, env := range envs {
		if !env.IsArchived {
			return errors.New("has active environments (remove the environments or use the force_destroy flag)")
		}
	}

	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(client.ApiClientInterface)

	id := d.Id()

	if err := resourceProjectAssertCanDelete(ctx, d, meta); err != nil {
		return diag.Errorf("could not delete project: %v", err)
	}

	if err := apiClient.ProjectDelete(id); err != nil {
		return diag.Errorf("could not delete project: %v", err)
	}
	return nil
}

func resourceProjectImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()
	_, err := uuid.Parse(id)

	var project client.Project

	if err == nil {
		log.Println("[INFO] Resolving Project by id: ", id)
		if project, err = getProjectById(id, meta); err != nil {
			return nil, err
		}
	} else {
		log.Println("[INFO] Resolving Project by name: ", id)

		if project, err = getProjectByName(id, meta); err != nil {
			return nil, err
		}
	}

	if err := writeResourceData(&project, d); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
