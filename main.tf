###### root/main.tf ######
module "s3" {
  source                     = "./modules/s3"
  with_policy                = var.with_policy
  tag_bucket_name            = var.tag_bucket_environment
  tag_bucket_environment     = var.tag_bucket_environment
}