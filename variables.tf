variable "tag_bucket_name" {
  description = "Name for the S3 bucket"
}

variable "tag_bucket_environment" {
  description = "Environment for the S3 bucket"
}

variable "with_policy" {
  description = "Flag indicating whether to include a policy"
  type        = bool
}