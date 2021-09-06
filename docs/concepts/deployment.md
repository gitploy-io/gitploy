# Deployment

Gitploy provides two types of deployment: Deploy and Rollback.

## Deploy

Deploying is the primary feature of Gitploy. When you deploy, you have to select the environment and the `ref`: for the environment, you can choose one of the environments listed in the `deploy.yml`, and for the `ref`, you can choose one of commit, branch, and tag.

With approval, Gitploy waits until it matches the required approving approvals. So you have to confirm to deploy after approval.

## Rollback

Rollback is the best way to recover while you fix the problems. Gitploy supports the rollback. You can choose one of the successful deployments to rollback. 

*Note that if the ref of the selected deployment is a branch, Gitploy automatically references the commit SHA to prevent deploying the head of the branch.*
