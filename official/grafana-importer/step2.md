## Convert the Grafana dashboard to Guance's

The next step is to use Guance CLI to convert the Grafana dashboard JSON file into a Terraform module.

### Run CLI

The CLI command is:

```shell
guance iac import grafana -f ./grafana.json -t terraform-module -o ./out -m prom
```

* **-f**: The Grafana dashboard JSON file path.
* **-t**: The output type. The value is `terraform-module`.
* **-o**: The output folder path.
* **-m**: The measurement. The default value is `prom`. If set to empty string, It will auto split metric name by `_`, the first part is the measurement.

    node_cpu_seconds_total -> node is the measurement

### Got result

You will get a Terraform module in the `./out` folder. The folder structure is:

```
.
├── main.tf
├── manifest.json
├── outputs.tf
├── variables.tf
└── versions.tf
```

The `manifest.json` file is used as a dashboard JSON file for Guance Cloud. Additionally, it can be utilized to import the dashboard directly to the Guance Cloud Console. 
