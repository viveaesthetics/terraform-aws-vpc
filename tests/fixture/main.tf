module "vpc" {
  source = "../../"

  name = "vpc-test012"
  cidr = "172.17.48.0/20"

  azs = ["us-east-1a", "us-east-1b", "us-east-1c"]

  private_subnets = [
    "172.17.48.0/23",
    "172.17.50.0/23",
    "172.17.52.0/23"
  ]

  public_subnets = [
    "172.17.54.0/23",
    "172.17.56.0/23",
    "172.17.58.0/23"
  ]


  enable_nat_gateway = true
  enable_vpn_gateway = true

  tags = {
    Terraform = "true"
  }
}

provider "aws" {
  region = "us-east-1"
  profile = "bb-dev-oaar"
}

output "vpc_cidr_block" {
  value       = module.vpc.vpc_cidr_block
  description = "VPC CIDR block"
}

output "vpc_id" {
  value       = module.vpc.vpc_id
  description = "VPC ID"
}



