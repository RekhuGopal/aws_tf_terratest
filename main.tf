###### root/main.tf ######
module "s3" {
  source                     = "./modules/s3"
  with_policy                = true
  tag_bucket_name            = "module-eks-${lower(random_string.suffix.result)}"
  tag_bucket_environment     = "Dev"
}