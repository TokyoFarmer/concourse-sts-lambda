# ------------------------------------------------------------------------------
# Variables
# ------------------------------------------------------------------------------
variable "name_prefix" {
  description = "Prefix used for resource names."
  type        = string
}

variable "filename" {
  description = "Path to the handler zip-file."
  type        = string
  default     = null
}

variable "source_code_hash" {
  description = "Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either filename or s3_key."
  type        = string
  default     = null
}

variable "s3_bucket" {
  description = "The bucket where the lambda function is uploaded."
  type        = string
  default     = null
}

variable "s3_key" {
  description = "The s3 key for the Lambda artifact."
  type        = string
  default     = null
}

variable "config_bucket_arn" {
  description = "The ARN of the config bucket."
  type        = string
}

variable "role_prefix" {
  description = "Prefix of CI roles which the Lambda function will be allowed to assume. (should be the same in all accounts)."
  type        = string
}

variable "secrets_manager_prefix" {
  description = "Prefix used for secrets. The Lambda will be allowed to create and write secrets to any secret with this prefix."
  type        = string
  default     = "concourse"
}

variable "tags" {
  description = "A map of tags (key-value pairs) passed to resources."
  type        = map(string)
  default     = {}
}

