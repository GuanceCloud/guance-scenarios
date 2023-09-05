## Download the Grafana Dashboard JSON

Then, download the [Node Exporter Dashboard on Grafana](https://grafana.com/grafana/dashboards/1860-node-exporter-full/).

Or use `wget` to download it:

```shell
wget https://grafana.com/api/dashboards/1860/revisions/31/download -O grafana.json
```

After running the command, a Grafana dashboard JSON file named `grafana.json` will be downloaded to the local disk. We could use it to import this file on the Guance Cloud Console.
