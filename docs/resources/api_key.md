---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "env0_api_key Resource - terraform-provider-env0"
subcategory: ""
description: |-
  
---

# env0_api_key (Resource)



## Example Usage

```terraform
resource "env0_api_key" "api_key_example" {
  name = "api-key-example"
}

resource "env0_project" "project_resource" {
  name = "project-resource"
}

resource "env0_user_project_assignment" "api_key_project_assignment_example" {
  user_id    = env0_api_key.api_key_example.id
  project_id = env0_project.project_resource.id
  role       = "Viewer"
}

resource "env0_team" "team_resource" {
  name = "team-resource"
}

resource "env0_user_team_assignment" "api_key_team_assignment_example" {
  user_id = env0_api_key.api_key_example.id
  team_id = env0_team.team_resource.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) the api key name

### Optional

- `omit_api_key_secret` (Boolean) if set to 'true' will omit the api_key_secret from the state. This would mean that the api_key_secret cannot be used
- `organization_role` (String) the api key type. 'Admin' or 'User'. Defaults to 'Admin'. For more details check https://docs.env0.com/docs/api-keys

### Read-Only

- `api_key_secret` (String, Sensitive) the api key secret. This attribute is not computed for imported resources. Note that this will be written to the state file. To omit the secret: set 'omit_api_key_secret' to 'true'
- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import env0_api_key.by_id ddda7b30-6789-4d24-937c-22322754934e
terraform import env0_api_key.by_name api-key-name"
```
