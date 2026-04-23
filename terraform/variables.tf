variable "access_key" {
  type      = string
  sensitive = true
}

variable "secret_key" {
  type      = string
  sensitive = true
}

variable "project_id" {
  type      = string
  sensitive = true
}

variable "region" {
  type    = string
  default = "nl-ams"
}

variable "zone" {
  type    = string
  default = "nl-ams-1"
}

variable "discord_bot_token" {
  type      = string
  sensitive = true
}

variable "discord_public_key" {
  type      = string
  sensitive = true
}

variable "discord_guild_id" {
  type      = string
  sensitive = true
}

variable "discord_status_channel_id" {
  type      = string
  sensitive = true
}

variable "function_namespace_name" {
  type    = string
  default = "jackbot"
}

variable "function_name" {
  type    = string
  default = "jackbot"
}

variable "github_access_token" {
  type      = string
  sensitive = true
}
