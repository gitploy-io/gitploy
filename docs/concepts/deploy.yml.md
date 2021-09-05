# deploy.yml

 Gitploy configures a pipeline with a simple, easy‑to‑read file that you commit to your git repository. 

## Features

### Parameters for Github deployment

Github deployments offer a few configurable [parameters](https://docs.github.com/en/rest/reference/repos#create-a-deployment--parameters). You can configure these parameters in the deploy.yml file. You can check details in the reference.

Here is the example to deploy based on tag:

```yaml
envs:
  - name: prod
    auto_merge: false
    required_contexts:
      - "go-test"
      - "react-test"
      - "publish-image"
    production_environment: true
```

### Approval

Gitploy supports the approval step to protect to deploy until it matches the required approving approvals.

```yaml
envs:
  - name: prod
    approval:
      enabled: true
      required_count: 1
```

## File Location

When you activate the repository in Gitploy, the default path is `deploy.yml` at the root. But you can replace the file path in the settings tab in Gitploy.

*Note that Gitploy always reads the file from the head of the default branch.*
