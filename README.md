# Gitploy 

<p align="center">
  <img src="./docs/images/logo_400.png"><br/>
  <img src="https://img.shields.io/github/v/release/gitploy-io/gitploy?display_name=release">
  <img src="https://img.shields.io/github/v/release/gitploy-io/gitploy?include_prereleases&label=pre-release">
  <img src="https://github.com/gitploy-io/gitploy/actions/workflows/test.yaml/badge.svg">
  <img src="https://github.com/gitploy-io/gitploy/actions/workflows/publish.yaml/badge.svg"><br/>
  <b>Gitploy helps your team build the deployment system around GitHub in minutes.</b><br/>
  <a href="https://docs.gitploy.io/">Documentation</a> | <a href="https://github.com/gitploy-io/gitploy/discussions">Community</a> | <a href="https://docs.gitploy.io/tasks/installation/">Installation Guide</a>
</p>

---

## What is Gitploy?

GitHub provides the [deployment API](https://docs.github.com/en/rest/reference/deployments#deployments) to deploy a specific ref(branch, SHA, tag). It offers strong features to make your team (or organization) can deploy fast and safety. **But it takes a lot of resources to build the deployment system around GitHub deployment API.**

Gitploy makes your team or organization **build the deployment system around GitHub in minutes.** Now, do not waste the time to build the deployment system.

![gitploy](./docs/images/gitploy.gif)


## Features

* Provides the intuitive UI to deploy a specific ref (branch, SHA, tag).
* Build an event-driven deployment system around GitHub. See GitHub [deployment event](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deployment).
* Integrate with GitHub [Action](https://github.com/features/actions) in minutes.
* Provides advanced deployment features: Rollback, Review, Lock.
* Provides various validations for the deployment.
* Provides DevOps metrics.

## Gitploy vs GitHub environment

Name        | Gitploy | GitHub environment
---         |---      |---
Manual deploy              | ✅ | ✅
Auto deploy                | ✅ | ✅
Review                     | ✅ | ✅
Rollback                   | ✅ | ❌
Lock environment           | ✅ | ❌
Commit statuses validation | ✅ | ❌
Display changed commmits   | ✅ | ❌
Private repositories for teams plan | ✅ | ❌ 

## Getting Started

To install Gitploy on your hosting, read this [documentation](https://docs.gitploy.io/tasks/installation/). 

For public repositories, we're providing the [free cloud](https://github.com/marketplace/gtiploy).

### Important Links

Documentation | Community | Installation Guide | Docker Image
--- |--- |--- |---
📚 [Documentation](https://docs.gitploy.io/) |💬 [Community](https://github.com/gitploy-io/gitploy/discussions) |📖 [Installation Guide](https://docs.gitploy.io/tasks/installation/) |🐋 [Docker Image](https://hub.docker.com/repository/docker/gitployio/gitploy)

## Contributors

Don't be afraid to contribute! We have many things you can do to help out. If you're trying to contribute but stuck, please tag @hanjunlee.

You can check the [contributing](./docs/contributing.md) for exact details on how to contribute.
