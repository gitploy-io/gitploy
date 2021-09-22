# deploy.yml

## Env


Field                    |Type                     |Required  |Description
---                      |----                     |---       |---
`name`                   |*string*                 |`true`    |This field is the runtime environment such as `production`, `staging`, and `qa`. 
`task`                   |*string*                 |`false`   |This field is used by the deployment system to distinguish the kind of deployment. 
`description`            |*string*                 |`false`   |This field is the short description of the deployment. 
`auto_merge`             |*boolean*                |`false`   |This field is used to ensure that the requested ref is not behind the repository's default branch. If you deploy with the commit or the tag you need to set `false`. For rollback, Gitploy set the field `false`.
`required_contexts`      |*[]string*               |`false`   |This field allows you to specify a subset of contexts that must be success. 
`payload`                |*object* or *string*     |`false`   |This field is JSON payload with extra information about the deployment. 
`production_environment` |*boolean*                |`false`   |This field specifies whether this runtime environment is production or not.
`approval`               |*[Approval](#approval)*  |`false`   |This field configures approval.

## Approval

Field            |Type      |Tag     |Description
---              |---       |---     |---
`enabled`        |*boolean* |`true`  |This field make to enable the approval feature. The default value is `false`.
`required_count` |*integer* |`false` |This field determine how many the required approving approvals is needs to deploy. The default value is `0`.

## Variables

The following variables are available in `${ }` syntax when evaluating `deploy.yml` before deploy or rollback:

* `GITPLOY_DEPLOY_TASK`: Returns `deploy` for deploy, but rollback, it returns the empty string.
* `GITPLOY_ROLLBACK_TASK`: Returns `rollback` for rollback, but deploy, it returns the empty string.
* `GITPLOY_IS_ROLLBACK`: Returns `true` for rollback, but deploy, it returns `false`.

An example usage of this:

```yaml
envs:
  - name: prod
    task: "${GITPLOY_DEPLOY_TASK}${GITPLOY_ROLLBACK_TASK}:kubernetes"  # It returns "deploy:kubernetes" or "rollback:kubernetes"
```

And Gitploy provides the string operation to facilitate customized values. You can check supported functions at [here](https://github.com/drone/envsubst).

```yaml
envs:
  - name: prod
    task: "${GITPLOY_DEPLOY_TASK=rollback}:kubernetes" # It returns "deploy:kubernetes" or "rollback:kubernetes"
```
