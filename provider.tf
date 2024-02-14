terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.15.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }
  }

backend "remote" {
		hostname = "app.terraform.io"
		organization = "CloudQuickLabs"

		workspaces {
			name = "AWSEKS"
		}
	}
}

provider "aws" {
  region = "us-west-2"
}

resource "random_string" "suffix" {
  length  = 5
  special = false
}