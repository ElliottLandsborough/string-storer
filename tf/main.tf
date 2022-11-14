terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}
provider "aws" {
  region = "eu-west-2"

  default_tags {
    tags = {
      Name = "string-storer"
    }
  }
}

resource "aws_ecr_repository" "string_storer_repo" {
  name = "string-storer" # Naming my repository
}
