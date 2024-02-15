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

  backend "s3" {
    bucket = "terratestbucketcloudquicklabs"
    region = "us-west-2"
    dynamodb_table = "backendstroageterraform"
    key = "terraform.tfstate"
  }
}

provider "aws" {
  region = "us-west-2"
}

resource "random_string" "suffix" {
  length  = 5
  special = false
}