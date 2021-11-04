# deploy.yml

Gitploy configures a pipeline with a simple, easy‑to‑read file that you commit to your git repository. *The configuration file must be at the head of the default branch.* The default path is `deploy.yml` at the root directory, but you can replace the file path in the settings tab of Gitploy. You can check the [document](../references/deploy.yml.md) for the specification of the configuration file.

## Environments

You can configure each environments, respectively. The configuration have to be under the `evns` field. The following example show each environments have different configuration.

```yaml
envs:
  - name: dev
    auto_merge: false
    required_contexts: []
  - name: production
    auto_merge: true
    required_contexts: 
      - test
      - docker-image
```

## Parameters of GitHub deployment API

Internally, Gitploy posts a deployment to GitHub [deployments API](https://docs.github.com/en/rest/reference/repos#create-a-deployment) with parameters from the configuration file. These parameter help you can verify the artifact before you start to deploy. The configuration file provides fields to configure all parameters of GitHub deployment API. You can check the [document](../references/deploy.yml.md) for the detail.

```yaml
envs:
  - name: production
    task: deploy:lambda
    description: Start deploying to the production.
    auto_merge: false
    required_contexts:
      - test
      - integration-test
    production_environment: true
```
## Review

Gitploy provides the review process. You can list up to users on the configuration file. You can check the [document](./review.md) for the detail.

```yaml
envs:
  - name: production
    review:
      enabled: true
      reviewers: ["ocotocat", "noah"]
```
