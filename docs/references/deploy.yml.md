# deploy.yml

## Env

Field                    |Type                    |Tag       |Description
---                      |----                    |---       |---
`name`                   |*string*                |          |This parameter is the runtime environment such as `production`, `staging`, and `qa`. 
`task`                   |*string*                |`github`  |This parameter is used by the deployment system to distinguish the kind of deployment. The default values is `deploy`, but for rollback, the default value is `rollback`.
`description`            |*string*                |`github`  |This parameter is the short description of the deployment. The default value is `Gitploy starts to deploy.`
`auto_merge`             |*boolean*               |`github`  |This parameter is used to ensure that the requested ref is not behind the repository's default branch. If you deploy with the commit or the tag you need to set `false`. The default values is `true`.
`required_contexts`      |*[]string*              |`github`  |This parameter allows you to specify a subset of contexts that must be success. The default value is `nil`, it means every submitted context must be in a success state.
`production_environment` |*boolean*               |`github`  |This parameter specifies whether this runtime environment is production or not.
`approval`               |*[Approval](#approval)* |          |This parameter configures approval.

## Approval

Field            |Type      |Tag   |Description
---              |---       |---   |
`enabled`        |*boolean* |      |This parameter make to enable the approval feature. The default value is `false`.
`required_count` |*integer* |      |This parameter determine how many the required approving approvals is needs to deploy. The default value is `0`. 