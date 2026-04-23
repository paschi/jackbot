resource "scaleway_function_namespace" "jackbot" {
  name = var.function_namespace_name
}

resource "scaleway_function" "jackbot" {
  namespace_id = scaleway_function_namespace.jackbot.id
  name         = var.function_name
  runtime      = "go124"
  handler      = "cmd/bot/Handle"
  privacy      = "public"
  timeout      = 10
  max_scale    = 2
  memory_limit = 128
  zip_file     = "../function.zip"
  zip_hash     = filesha256("../function.zip")
  deploy       = true
  secret_environment_variables = {
    DISCORD_BOT_TOKEN         = var.discord_bot_token
    DISCORD_PUBLIC_KEY        = var.discord_public_key
    DISCORD_GUILD_ID          = var.discord_guild_id
    DISCORD_STATUS_CHANNEL_ID = var.discord_status_channel_id
    GITHUB_ACCESS_TOKEN       = var.github_access_token
  }
}
