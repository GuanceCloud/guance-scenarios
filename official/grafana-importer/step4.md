## Run the DataKit to collect metrics from Prometheus

The next step is to use the DataKit agent to upload real-world metrics data to Guance Cloud and display it on the dashboard that was just created.

**Note**: All the configuration files are in the `~/` folder. You don't need to create it again.

### Create DataKit configuration

Edit the `~/prom.conf` file, see the following content:

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

This configuration will do the following things:

1. Collect the metrics from `http://node-exporter:9100/metrics`.
2. (Optional) Set the source and measurement name to `prom`.
3. (Optional) Set the `job` tag to `DataKit`.

### Create a Docker Compose file

Edit the `~/docker-compose.yml` file. See the following content:

```yaml
version: '3'

services:
  node-exporter:
    image: prom/node-exporter
    ports:
      - "9100:9100"
    networks:
      - app-network
  datakit:
    image: guancecloud/datakit:1.6.3-alpha
    ports:
      - "8080:8080"
    networks:
      - app-network
    environment:
      - ENV_DATAWAY=https://openway.guance.com?token=${GUANCE_WORKSPACE_TOKEN}
    volumes:
      - ./prom.conf:/usr/local/datakit/conf.d/prom/prom.conf

networks:
  app-network:
```

This file will create two containers:

- `node-exporter`: The Prometheus exporter for hardware and OS metrics.
- `datakit`: The DataKit agent collects metrics from Prometheus and uploads them to Guance Cloud.

### Start to collect the metrics

Run the following command to start the DataKit agent:

```shell
# Change the current directory to $HOME
cd ~

# Start the DataKit agent
docker-compose up -d

# You can use the following command to check the status:
docker-compose ps
```

Wait for a while. Then, you can see the metrics are uploaded and shown on the dashboard of Guance Cloud.

