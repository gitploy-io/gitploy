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
* **gitploy_request_duration_seconds**<br/> The HTTP request latencies in seconds.
* **gitploy_deployment_count** <br/> The count of success deployment of the production environment.
* **gitploy_rollback_count**<br/> The count of rollback of the production environment.
* **gitploy_line_additions**<br/> The count of added lines from the latest deployment of the production environment.
* **gitploy_line_deletions**<br/> The count of deleted lines from the latest deployment of the production environment.
* **gitploy_line_changes**<br/> The count of changed lines from the latest deployment of the production environment.
* **gitploy_member_count**<br/> The total count of members.
* **gitploy_member_limit**<br/> The limit count of members.
