# Review

Gitploy has the review to require at least one approval for the deployment. You can list up to users on the configuration file. The reviewers must have at least read access to the repository. 

```yaml
envs:
  - name: production
    auto_merge: true
    review:
      enabled: true
      reviewers: ["octocat", "noah"]
```
