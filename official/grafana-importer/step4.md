## Run the DataKit to collect metrics from Prometheus

For example, the DataKit config is:

```toml
[[inputs.prom]]
  url = "http://node-exporter:9100/metrics"
  source = "prom"
  metric_types = []
  interval = "60s"
  measurement_name = "prom"
  [inputs.prom.tags]
    job = "DataKit"
```

Then run the Guance CLI to import the downloaded JSON.

```shell
docker compose up -d
```

Wait for a while, then you can see the metrics is uploaded and show on dashboard of Guance Cloud.
