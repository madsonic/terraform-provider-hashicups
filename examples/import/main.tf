terraform {
  required_providers {
    hashicups = {
      versions = ["0.3"]
      source = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
  username = "education"
  password = "test123"
}

resource "hashicups_order" "first" {}

output "first_order" {
  value = hashicups_order.first
}