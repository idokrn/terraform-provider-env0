provider "random" {}

resource "random_string" "random" {
  length    = 8
  special   = false
  min_lower = 8
}

resource "env0_project" "test_project" {
  name          = "Test-Project-for-environment-${random_string.random.result}"
  force_destroy = true
}

resource "env0_template" "template" {
  name              = "Template for environment resource-${random_string.random.result}"
  type              = "terraform"
  repository        = "https://github.com/env0/templates"
  path              = "misc/null-resource"
  terraform_version = "0.15.1"
}

resource "env0_template_project_assignment" "assignment" {
  template_id = env0_template.template.id
  project_id  = env0_project.test_project.id
}

resource "env0_environment" "example" {
  depends_on    = [env0_template_project_assignment.assignment]
  force_destroy = true
  name          = "environment-${random_string.random.result}"
  project_id    = env0_project.test_project.id
  template_id   = env0_template.template.id
  configuration {
    name  = "environment configuration variable"
    value = "value"
  }
  approve_plan_automatically = true
  revision                   = "master"
  vcs_commands_alias         = "alias"
}

resource "env0_template" "terragrunt_template" {
  name               = "Terragrunt template for environment resource-${random_string.random.result}"
  type               = "terragrunt"
  repository         = "https://github.com/env0/templates"
  path               = "misc/null-resource"
  terraform_version  = "0.15.1"
  terragrunt_version = "0.35.0"
}

resource "env0_template_project_assignment" "terragrunt_assignment" {
  template_id = env0_template.terragrunt_template.id
  project_id  = env0_project.test_project.id
}

resource "env0_environment" "terragrunt_environment" {
  depends_on                       = [env0_template_project_assignment.terragrunt_assignment]
  force_destroy                    = true
  name                             = "environment-${random_string.random.result}"
  project_id                       = env0_project.test_project.id
  template_id                      = env0_template.terragrunt_template.id
  approve_plan_automatically       = true
  revision                         = "master"
  terragrunt_working_directory     = var.second_run ? "/second-dir" : "/first-dir"
  auto_deploy_on_path_changes_only = false
}

data "env0_environment" "test" {
  id = env0_environment.example.id
}

output "revision" {
  value = data.env0_environment.test.revision
}

output "terragrunt_working_directory" {
  value = env0_environment.terragrunt_environment.terragrunt_working_directory
}

data "env0_template" "github_template" {
  name = "Github Integrated Template"
}

resource "env0_environment" "environment-without-template" {
  force_destroy                    = true
  name                             = "environment-without-template-${random_string.random.result}"
  project_id                       = env0_project.test_project.id
  approve_plan_automatically       = true
  revision                         = "master"
  auto_deploy_on_path_changes_only = false

  without_template_settings {
    description                             = "Template description - GitHub"
    type                                    = "terraform"
    repository                              = data.env0_template.github_template.repository
    github_installation_id                  = data.env0_template.github_template.github_installation_id
    path                                    = var.second_run ? "second" : "misc/null-resource"
    retries_on_deploy                       = 3
    retry_on_deploy_only_when_matches_regex = "abc"
    retries_on_destroy                      = 1
    terraform_version                       = "0.15.1"
  }
}