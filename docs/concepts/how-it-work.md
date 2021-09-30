# How it works

Gitploy builds the system around GitHub [deployment API](https://docs.github.com/en/rest/reference/repos#deployments). **It's an event-driven decoupled way to deploy your code.** Internally, Gitploy creates a new Github deployment resource, and Github dispatches a deployment event that external services can listen for and act. It enables developers and organizations to build loosely coupled tooling. 

This approach has several pros:

* Replace deployment tools easily without changing your deployment pipeline. 
* Easy to implement details of deploying different types of applications (e.g., web, native).

Below is a simple diagram for how these interactions would work:

```
+---------+             +--------+             +---------+         +-------------+
| Gitploy |             | GitHub |             |  Tools  |         | Your Server |
+---------+             +--------+             +---------+         +-------------+
     |                      |                       |                     |
     |  Create Deployment   |                       |                     |
     |--------------------->|                       |                     |
     |                      |                       |                     |
     |  Deployment Created  |                       |                     |
     |<---------------------|                       |                     |
     |                      |                       |                     |
     |                      |   Deployment Event    |                     |
     |                      |---------------------->|                     |
     |                      |                       |     SSH+Deploys     |
     |                      |                       |-------------------->|
     |                      |                       |                     |
     |                      |   Deployment Status   |                     |
     |                      |<----------------------|                     |
     |   Deployment Status  |                       |                     |
     |<---------------------|                       |                     |
     |                      |                       |                     |
     |                      |                       |   Deploy Completed  |
     |                      |                       |<--------------------|
     |                      |   Deployment Status   |                     |
     |                      |<----------------------|                     |
     |   Deployment Status  |                       |                     |
     |<---------------------|                       |                     |
     |                      |                       |                     |
```

Gitploy lets you can build the advanced deployment system so your team and organization enable to deploy the application with lower risk and faster.

*Keep in mind that Gitploy is never actually accessing your servers. It's up to your tools to interact with deployment events.*
