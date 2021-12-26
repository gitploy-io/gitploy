# deploy.yml


Field                    |Type                     |Required  |Description
---                      |----                     |---       |---
`envs`                   |*[][Env](#env)*          |`true`    |Thie field configures the pipeline for each environment, respectively.

## Env


Field                    |Type                     |Required  |Description
---                      |----                     |---       |---
`name`                   |*string*                 |`true`    |This field is the runtime environment such as `production`, `staging`, and `qa`. 
`task`                   |*string*                 |`false`   |This field is used by the deployment system to distinguish the kind of deployment. (*Only for `GitHub`*)
`description`            |*string*                 |`false`   |This field is the short description of the deployment. (*Only for `GitHub`*)
`auto_merge`             |*boolean*                |`false`   |This field is used to ensure that the requested ref is not behind the repository's default branch. If you deploy with the commit or the tag you need to set `false`. For rollback, Gitploy set the field `false`. (*Only for `GitHub`*)
`required_contexts`      |*[]string*               |`false`   |This field allows you to specify a subset of contexts that must be success. (*Only for `GitHub`*)
`payload`                |*object* or *string*     |`false`   |This field is JSON payload with extra information about the deployment. (*Only for `GitHub`*)
`production_environment` |*boolean*                |`false`   |This field specifies whether this runtime environment is production or not.
`deployable_ref`         |*string*                 |`false`   |This field specifies which the ref(branch, SHA, tag) is deployable or not. It supports the regular expression ([re2]((https://github.com/google/re2/wiki/Syntax))). 
`auto_deploy_on`         |*string*                 |`false`   |This field controls auto-deployment behaviour given a ref(branch, SHA, tag). If any new push events are detected on this event, the deployment will be triggered. It supports the regular expression ([re2](https://github.com/google/re2/wiki/Syntax)). E.g. `refs/heads/main` or `refs/tags/v.*`
`review`                 |*[Review](#review)*      |`false`   |This field configures reviewers.
`frozen_windows`         |*[][Frozen Window](#frozen-window)* |`false`   |This field configures to add a frozen window to prevent unintended deployment for the environment.

## Review

Field            |Type       |Required  |Description
---              |---        |---       |---
`enabled`        |*boolean*  |`false`    |This field makes to enables the review feature. The default value is `false`.
`reviewers`      |*[]string* |`false`  |This field list up reviewers. The default value is `[]`. You should specify the maintainers of the project.

## Frozen Window

Field            |Type       |Required     |Description
---              |---        |---          |---
`start`          |*string*   |`true`       |This field is a cron expression to indicate when the window starts. For example, `55 23 * * *` means it starts to freeze a window before 5 minutes of midnight. You can check the [documentation](https://github.com/gitploy-io/cronexpr) for details.
`duration`       |*string*   |`true`       |This field configures how long the window is frozen from the starting. The duration string is a possibly signed sequence of decimal numbers and a unit suffix such as `5m`, or `1h30m`. Valid time units are `ns`, `us`, `ms`, `s`, `m`, `h`.
`location`       |*string*   |`false`      |This field configures the location of the `start` time. The value is taken to be a location name corresponding to a file in the IANA Time Zone database, such as `America/New_York`. The default value is `UTC`. You can check the [document](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for the Time Zone database name.

## Variables

The following variables are available in `${ }` syntax when evaluating `deploy.yml` before deploy or rollback:

* `GITPLOY_DEPLOY_TASK`: Returns `deploy` for deploy, but rollback, it returns the empty string.
* `GITPLOY_ROLLBACK_TASK`: Returns `rollback` for rollback, but deploy, it returns the empty string.
* `GITPLOY_IS_ROLLBACK`: Returns `true` for rollback, but deploy, it returns `false`.

Example usage of this:

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
