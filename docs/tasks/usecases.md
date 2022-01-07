# Use Cases

## Deploying the head of the `main` branch

The general way of deployment is deploying the head of the `main` branch that has only verified commits. To constrain deploying only the head of the `main` branch, we can use the `auto_merge` parameter of GitHub [deployment API](https://docs.github.com/en/rest/reference/repos#create-a-deployment) to ensure that the deployment is the head of the branch, and set `deployable_ref` field with the `main`.

```yaml
envs:
  - name: production
    auto_merge: true
    deployable_ref: main
    production_environment: true
```

## Deploying with the version

The versioning is the general way to specify the artifact or the commit, and GitHub provides the release page for versioning. If your team (or organization) wants to constrain deploying with the version, you can use the `deployable_ref` field like below.

```yaml
envs:
  - name: production
    auto_merge: false
    deployable_ref: 'v.*\..*\..*'       # Semantic versioning
    production_environment: true
```

## Deploying the artifact

The artifact could be a binary file from compiling source codes or a docker image, which means we have to build the artifact before we deploy. The commit status is the best way to verify if the artifact exists or not. The builder, such as GitHub Action or Circle CI, posts the commit status after building the artifact, and we can verify it by the `required_contexts` parameter when we deploy. You can reference the `deploy.yml` of Gitploy.

```yaml
envs:
  - name: production
    auto_merge: false
    required_contexts:
      - "publish-image"         # The commit status of building the artifact.
    production_environment: true
```

## Auto deployment

If you want to enable the auto-deployment when the pull request is merged into the main branch, you should configure the `auto_deploy_on` field like the below.

```yaml
envs:
  - name: production
    auto_merge: true
    required_contexts: []
    auto_deploy_on: refs/heads/main
    deployable_ref: main
    production_environment: true
```
