# Gitploy

## What is Gitploy?

GitHub provides the [deployment API](https://docs.github.com/en/rest/reference/deployments#deployments) to deploy a specific ref(branch, SHA, tag). It enables your organization to build the deployment system loosely coupled tooling without worrying about the implementation details of delivering different types of applications (e.g., web, native). **But it takes a lot of resources to build the deployment system around GitHub deployment API from scratch.**

Gitploy enables your organization **to build the deployment system around GitHub in minutes.** Gitploy provides the place to manage all deployment and deploying in the same manner. Now, do not waste time building the deployment system.

![gitploy](./docs/images/gitploy.gif)

## Features

* Manage all deployments in one place.
* Provides the intuitive UI to deploy a specific `ref` (branch, SHA, tag).
* Build an event-driven deployment system for loosely coupled tooling.
* Provides both continuous delivery and continuous deployment.
* Provides advanced deployment features: Rollback, Review, Lock, Freeze Window.
* Provides DevOps metrics.