---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "env0_project_policy Data Source - terraform-provider-env0"
subcategory: ""
description: |-
  
---

# env0_project_policy (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **project_id** (String) id of the project

### Read-Only

- **continuous_deployment_default** (Boolean) Redeploy on every push to the git branch default value
- **disable_destroy_environments** (Boolean) Disallow destroying environment in the project
- **id** (String) id of the policy
- **include_cost_estimation** (Boolean) Enable cost estimation for the project
- **number_of_environments** (Number) Max number of environments a single user can have in this project, 0 indicates no limit
- **number_of_environments_total** (Number) Max number of environments in this project, 0 indicates no limit
- **requires_approval_default** (Boolean) Requires approval default value when creating a new environment in the project
- **run_pull_request_plan_default** (Boolean) Run Terraform Plan on Pull Requests for new environments targeting their branch default value
- **skip_apply_when_plan_is_empty** (Boolean) Skip apply when plan has no changes
- **skip_redundant_deployments** (Boolean) Automatically skip queued deployments when a newer deployment is triggered
- **updated_by** (String) updated by

