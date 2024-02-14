output "bucket_id" {
  description = "S3 Bucket ID"
  value       = module.s3.bucket_id
}

output "bucket_arn" {
  description = "Bucket ARN"
  value       = module.s3.bucket_arn
}
