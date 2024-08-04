variable "bucket_name" {
  type        = string
  description = "Name of the s3 bucket"
}

variable "enable_bucket_versioning" {
  type        = string
  description = "Enable bucket versioning(Enable or Disable)"
  default     = "Enable"
}

variable "prevent_bucket_destruction" {
  type        = bool
  description = "Set if the bucket can be destroyed"
  default     = true

}
