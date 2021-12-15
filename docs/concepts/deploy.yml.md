# deploy.yml

Gitploy configures a pipeline with a simple, easy‑to‑read file that you commit to your git repository. *The configuration file must be at the head of the default branch.* The default path is `deploy.yml` at the root directory, but you can replace the file path in the settings tab of Gitploy. You can check the [documentation](../references/deploy.yml.md) for the specification of the configuration file and also references [Use Cases](../tasks/usecases.md).

## Quick Start

If you want to get started quickly, you should copy the `deploy.yml` file and push it into your git repository. Then you can find the environment in the Gitploy.

```yaml
# deploy.yml
envs:
  - name: production
    auto_merge: false
    # To avoid the context verification.
    # required_context: []
```

## Features
### Multi Environment

For multi-environment, Gitploy provides you can make different pipelines for each environment, respectively. The configuration has to be defined under the `envs` field. The following example shows each environment has a different configuration.

```yaml
envs:
  # Development environment
  - name: development
    auto_merge: false
    required_contexts: []

  # Production environment
  - name: production
    auto_merge: true
    required_contexts: 
      - test
      - docker-image
```

### Parameters of GitHub deployment API

Internally, Gitploy posts a deployment to GitHub [deployments API](https://docs.github.com/en/rest/reference/repos#create-a-deployment) with parameters from the configuration file. These parameters help you can verify the artifact before you start to deploy.  The configuration provides fields to set all parameters of GitHub deployment API. You can check the [document](../references/deploy.yml.md) for the detail.

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
### Review

Gitploy provides the review process. You can list up to users on the configuration file. You can check the [document](./review.md) for the detail.

```yaml
envs:
  - name: production
    review:
      enabled: true
      reviewers: ["ocotocat", "noah"]
```
