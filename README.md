<div align="center">

# jackbot

[![Go][badge-go]][uri-go]
[![OpenTofu][badge-opentofu]][uri-opentofu]
[![Scaleway][badge-scaleway]][uri-scaleway]

*A Discord bot written in **Go** that runs as a **Scaleway Function**, managed with **OpenTofu**.*

</div>

## ⚡️ Getting Started

First, install [OpenTofu][uri-opentofu-install] if you have not already.

To deploy this bot, you need a **Scaleway** account and a **Discord** application.

### 1. Discord Setup
Create a new application in the [Discord Developer Portal](https://discord.com/developers/applications) and:
- Create a Bot and copy the **Bot Token**.
- Copy the **Public Key** from the General Information page.
- Invite the bot to your server using the OAuth2 URL generator (select `applications.commands` scope).
- Copy the **Guild ID** where you want to use the bot.

### 2. Scaleway Setup
Create an **API Key** in your Scaleway account and ensure you have a project created.

### 3. Deployment
To configure the deployment, create a file named `terraform.tfvars` with the following content:

```hcl
project_id                = "<YOUR_SCALEWAY_PROJECT_ID>"
access_key                = "<YOUR_SCALEWAY_ACCESS_KEY>"
secret_key                = "<YOUR_SCALEWAY_SECRET_KEY>"
region                    = "nl-ams"
zone                      = "nl-ams-1"

discord_bot_token         = "<YOUR_DISCORD_BOT_TOKEN>"
discord_public_key        = "<YOUR_DISCORD_PUBLIC_KEY>"
discord_guild_id          = "<YOUR_DISCORD_GUILD_ID>"
discord_status_channel_id = "<YOUR_DISCORD_STATUS_CHANNEL_ID>"

github_access_token       = "<YOUR_GITHUB_ACCESS_TOKEN>"
```

*Note: The `github_access_token` is used by the `/server` command to trigger workflows in the `satisfactory-hcloud` repository.*

After that, initialize the OpenTofu configuration and deploy:
```shell
$ tofu init
$ tofu apply
```

The bot will be deployed as a Scaleway Function. You'll need to set the **Interactions Endpoint URL** in the Discord Developer Portal to the URL provided by the function deployment.

## 🎮 Commands

The bot provides the following slash commands:

- `/ping`: Simple command to check if the bot is alive.
- `/roll [sides]`: Rolls a dice with the specified number of sides (2-100).
- `/server [game] [action]`: Triggers a GitHub Action workflow to start or stop a game server (currently supports Satisfactory).

## ⚙️ Configuration

The following variables can be configured in your `terraform.tfvars` file:

| Variable | Description | Default |
| --- | --- | --- |
| `project_id` | Scaleway Project ID. | |
| `access_key` | Scaleway Access Key. | |
| `secret_key` | Scaleway Secret Key. | |
| `region` | Scaleway Region. | `nl-ams` |
| `zone` | Scaleway Zone. | `nl-ams-1` |
| `discord_bot_token` | Discord Bot Token. | |
| `discord_public_key` | Discord Public Key. | |
| `discord_guild_id` | Discord Guild ID. | |
| `discord_status_channel_id` | Discord Channel ID for status updates. | |
| `github_access_token` | GitHub Personal Access Token. | |
| `function_namespace_name` | Name of the Scaleway Function Namespace. | `jackbot` |
| `function_name` | Name of the Scaleway Function. | `jackbot` |

[badge-scaleway]: https://img.shields.io/badge/Scaleway-510099?style=for-the-badge&logo=scaleway&logoColor=white
[badge-opentofu]: https://img.shields.io/badge/OpenTofu-FF4E00.svg?style=for-the-badge&logo=OpenTofu&logoColor=white
[badge-go]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[uri-scaleway]: https://www.scaleway.com
[uri-opentofu]: https://opentofu.org
[uri-opentofu-install]: https://opentofu.org/docs/intro/install/
[uri-go]: https://go.dev
