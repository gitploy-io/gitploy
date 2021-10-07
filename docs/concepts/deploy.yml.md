# deploy.yml

Gitploy configures a pipeline with a simple, easy‑to‑read file that you commit to your git repository. *The configuration file must be at the head of the default branch.* The default path is `deploy.yml` at the root directory, but you can replace the file path in the settings tab of Gitploy. You can check the [document](../references/deploy.yml.md) for the specification of the configuration file.

## Environments

The configuration file is configured for each environment, respectively. The following example is the fundamental structure of a configuration file.

<details>
<summary>Fundamental structure</summary>

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

</details>

## Parameters for Github deployment API

When Gitploy deploys, it posts a new deployment to GitHub [deployments API](https://docs.github.com/en/rest/reference/repos#create-a-deployment) with parameters from the configuration file. The configuration file provides fields to configure all parameters of GitHub deployment API. You can check the [document](../references/deploy.yml.md) for the detail.

<details>
<summary>GitHub parameter field</summary>

```yaml
envs:
  - name: production
    task: deploy:lambda
    description: Start to deploy to the production.
    auto_merge: false
    required_contexts:
      - test
      - integration-test
    production_environment: true
```

</details>

## Approval

Gitploy provides the approval step to protect to deploy until it matches the required approving approvals.

<details>
<summary>Enable Approval</summary>

```yaml
envs:
  - name: production
    approval:
      enabled: true
      required_count: 1
```

</details>
