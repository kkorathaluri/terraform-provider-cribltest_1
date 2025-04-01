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
# Example of using the cribl-terraform_pack resource
resource "cribl-terraform_pack" "example_pack" {
  id = "example-pack-id"
  name = "example-pack-name"
}

# Output the pack details to see the read-only attributes
output "pack_details" {
  value = cribl-terraform_pack.example_pack
} 