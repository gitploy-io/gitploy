# Integration

## GitHub Action

GitHub Actions help you automate tasks to run an actual deployment. GitHub Actions are event-driven, meaning that you can run a series of commands after a deployment event has occurred. 

You must specify `deployment` for the `on` field to listen for the deployment event. And you can use the `if` conditional to run a job for a specific environment. Here is the example below.

```yaml
# Listening the deployment event
on:
  deployment

jobs:
  deploy-dev:
    runs-on: ubuntu-latest
    # Run a job when the environment is 'production.
    if: ${{ github.event.deployment.environment == 'production' }}
    steps:
      -
        name: Checkout
        uses: actions/checkout@v2
      - 
        name: Start to deploy
        uses: chrnorm/deployment-status@releases/v1
        with:
          deployment_id: ${{ github.event.deployment.id }}
          description: Start to deploy ...
          state: "in_progress"
          token: "${{ github.token }}"
    # Run your deployment commands.
```

## Slack

Slack integration provides notifications for events.

### Step 1: Create App

Firstly, we have to create [Slack App](https://api.slack.com/apps). You should click the Create App button and fill out inputs.

### Step 2: Configure Permissions

After creating App, we move to the *OAuth & Permissions* page and set up *the redirect URLs* and *Bot Token scopes*on this page. Firstly, you should add a new redirect URL with the `GITPLOY_SERVER_PROTO://GITPLOY_SERVER_HOST/slack/signin` format; secondly, add `chat:write` scope into the Bot Token scopes.

Figure) Slack Bot Token Scopes

![Slack Bot Token Sceops](../images/slack-bot-token-scopes.png)

### Step 3: Run Server With App Credentials

To enable Slack integration, you have to set up these environments when you run the server: `GITPLOY_SLACK_CLIENT_ID` and `GITPLOY_SLACK_CLIENT_SECRET`. You can get these credentials from *App Credentials* section of *Basic Information* page. 
