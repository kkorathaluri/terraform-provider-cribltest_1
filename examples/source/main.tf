terraform {
  required_providers {
    cribl-terraform = {
      source = "kkorathaluri/cribltest-new"
    }
  }
}

provider "cribl-terraform" {
  # Configuration options
  server_url ="https://app.cribl-playground.cloud/organizations/beautiful-nguyen-y8y4azd/workspaces/main/app/api/v1/m/default"
}

data "cribl-terraform_source" "my_source" {
  id = "in_cribl_http"
}

output "source_details" {
  value = {
    id = data.cribl-terraform_source.my_source.id
    count_total = data.cribl-terraform_source.my_source.count_total
    items = data.cribl-terraform_source.my_source.items
  }
} 
