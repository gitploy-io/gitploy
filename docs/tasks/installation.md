# Installation

This article explains how to install the Gitploy server for GitHub.

## Step 1: Preparation

### Provision an instance

The server should be installed on a server or virtual machine with standard http and https ports open. 

### Create an OAuth Application

[Create a GitHub OAuth application.](https://docs.github.com/en/developers/apps/building-oauth-apps/creating-an-oauth-app) The Client Key and Client Secret are used to authorize access to GitHub resources.

*The authorization callback URL must match with the format:* `GITPLOY_SERVER_PROTO://GITPLOY_SERVER_HOST/signin`.

Figure) Github OAuth

![Github OAuth](../images/github-oauth.png)

## Step 2: Download

The server is distributed as a Docker image. The image is self-contained and does not have any external dependencies. We recommend to use the last version.

```
docker pull gitployio/gitploy:0.2
```

## Step 3: Configuration

The server is configured using environment variables. This article only configures with least environment. See [Configurations](../references/configurations.md) for a complete list of configuration options.

* **GITPLOY_SERVER_HOST**: 
Required string value configures the user-facing hostname. This value is used to create webhooks and redirect urls. 

* **GITPLOY_SERVER_PROTO**: 
Optional string value configures the user-facing protocol. This value is used to create webhooks and redirect urls. It can be one of them: `http` or `https`, and the default value is `https`.

* **GITPLOY_GITHUB_CLIENT_ID**:
Required string value configures the GitHub OAuth client id. This is used to authorize access to GitHub on behalf of a Gitploy user.

* **GITPLOY_GITHUB_CLIENT_SECRET**:
Required string value configures the GitHub OAuth client secret. This is used to authorize access to GitHub on behalf of a Gitploy user.

## Step 4: Start server

The server container can be started with the below command. The container is configured through environment variables.

```shell
docker run \
  --volume=/var/lib/gitploy:/data \
  --env=GITPLOY_SERVER_HOST={{GITPLOY_SERVER_HOST}} \
  --env=GITPLOY_SERVER_PROTO={{GITPLOY_SERVER_PROTO}} \
  --env=GITPLOY_GITHUB_CLIENT_ID={{GITPLOY_GITHUB_CLIENT_ID}} \
  --env=GITPLOY_GITHUB_CLIENT_SECRET={{GITPLOY_GITHUB_CLIENT_SECRET}} \
  --publish=80:80 \
  --publish=443:443 \
  --restart=always \
  --detach=true \
  --name=gitploy \
  gitployio/gitploy:0.2
```