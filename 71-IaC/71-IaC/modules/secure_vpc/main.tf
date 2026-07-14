# Standardizing the "Golden Network" Module
# Force compliance via IaC

module "secure_vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.0.0"

  name = "prod-vpc"
  cidr = "10.0.0.0/16"

  # Enforce best practices: Private subnets only
  public_subnets  = [] 
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]

  # Enforce logging (Auditability)
  enable_flow_log                      = true
  create_flow_log_cloudwatch_log_group = true
  
  tags = {
    Environment = "production"
    ManagedBy   = "terraform"
  }
}
