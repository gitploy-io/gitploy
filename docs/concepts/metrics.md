# Metrics

Gtiploy publishes and exposes metrics that Prometheus can consume at the standard `/metrics` endpoint. 

## Configuration

1\. Configure the server:

```
GITPLOY_PROMETHEUS_ENABLED=true
GITPLOY_PROMETHEUS_AUTH_SECRET=YOUR_SECRET
```

2\. Configure the prometheus scraper:

```
global:
  scrape_interval: 60s

scrape_configs:
  - job_name: gitploy
    authorization:
      credentials: YOUR_SECRET

    static_configs:
      - targets: ['domain.com']
```

## Gitploy Metrics

Gitploy provides the following Gitploy metrics. *Note that Some metrics are provided only for the production environment* (i.e. `production_environment: true` in the configuration file).

* **gitploy_requests_total** <br/> How many HTTP requests processed, partitioned by status code and HTTP method.
* **gitploy_request_duration_seconds**<br/> The HTTP request latencies in seconds
* **gitploy_deployment_count** <br/> The total deployment count of the production deployments.
* **gitploy_rollback_count**<br/> The total rollback count of the production deployments.
* **gitploy_line_additions**<br/> The total added lines of the production deployments.
* **gitploy_line_deletions**<br/> The total deleted lines of the production deployments.
* **gitploy_line_changes**<br/> The total changed lines of the production deployments.
* **gitploy_lead_time_seconds**<br/> The total amount of time it takes a commit to get into the production environments.
* **gitploy_commit_count**<br/> The total commit count of production deployments.
* **gitploy_member_count**<br/> The total count of members.
* **gitploy_member_limit**<br/> The limit count of members.
* **gitploy_total_active_repo_count**<br/> The count of active repositories.
* **gitploy_total_repo_count**<br/> The count of repositories.
