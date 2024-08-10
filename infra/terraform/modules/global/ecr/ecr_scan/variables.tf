variable "scan_type" {
  type        = string
  description = "Type of security scan to run"
  default     = "BASIC"
}
variable "scan_filter" {
  type        = string
  description = "String filter to filter out scans on matched names"
  default     = "test-image"
}
