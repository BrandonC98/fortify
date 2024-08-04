resource "aws_s3_bucket" "fortify_bucket" {
  bucket = var.bucket_name

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_s3_bucket_versioning" "fortify_bucket_versioning" {
  bucket = aws_s3_bucket.fortify_bucket.id
  versioning_configuration {
    status = var.enable_bucket_versioning
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "fortify_bucket_encryption" {
  bucket = aws_s3_bucket.fortify_bucket.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "fortify_bucket_access" {
  bucket                  = aws_s3_bucket.fortify_bucket.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

