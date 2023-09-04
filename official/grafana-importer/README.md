Grafana Importer Introduction
=============================

Importer Grafana dashboard by Guance CLI tookit

Introduction
------------

![Preview](images/preview.png)

Taking Ubuntu as an example, this guide demonstrates how to use the command line interface to import the Grafana dashboard [Node Exporter Full](https://grafana.com/grafana/dashboards/1860-node-exporter-full/) using Guance CLI.

Step 0: Prepare
---------------

The first step is to install the Guance CLI by running the command:

```bash
echo "deb [trusted=yes] https://releases.guance.io/apt/ /" | sudo tee /etc/apt/sources.list.d/guance.list
sudo apt update
apt install guance
```

For other ways to install Guance CLI, please refer to [Guance CLI Installation](https://github.com/GuanceCloud/guance-cli).

Step 1: Download the Grafana Dashboard JSON
-------------------------------------------

Then download the [Node Exporter Dashboard on Grafana](https://grafana.com/grafana/dashboards/1860-node-exporter-full/).

Or use `wget` to download it:

```shell
wget https://grafana.com/api/dashboards/1860/revisions/31/download -O grafana.json
```

After running the command, a Grafana dashboard json file named `grafana.json` will be generated in the local disk.e used to import This file can b the dashboard to Guance Cloud.

Step 2: Use Guance CLI to convert grafana as Terraform module
-------------------------------------------------------------

The next step is to use Guance CLI to convert the Grafana dashboard JSON file into a Terraform module.

The CLI command is:

```shell
guance iac import grafana -f ./input.json -t terraform-module -o ./out
```

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

To use the Guance Cloud console, please follow these steps:

1.	Open [Create Dashboard Page](https://console.guance.com/scene/dashboard/createDashboard).
2.	Click `Import Template`.
3.	Import the generated JSON file.

In this example, we will use the Terraform module to create the dashboard on Guance Cloud in the next step.

Step 3: Use Terraform to create a dashboard
-------------------------------------------

With this, you can utilize it to generate actualrces on the Guance Cloud platform. dashboard resou

```shell
cd ./out
terraform init
terraform apply
```

Step 4: Run the DataKit to collect metrics from Prometheus
----------------------------------------------------------

The next step is to use the DataKit agent to upload real-world metrics data to Guance Cloud and display it on the dashboard that was just created.

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

Step 5: Completed!
------------------

You have successfully imported a dashboard into Grafana. You can now view the dashboard on [Dashboard List](https://console.guance.com/scene/dashboard/list).

Here is a screenshot for the succeed result.

![Preview](images/preview.png)

If you want to learn more, there are some materials may be helpful.

1.	[All the examples about Grafana Importer](https://github.com/GuanceCloud/guance-cli/tree/main/specs/iac/import/grafana)
2.	[Guance CLI Homepage](https://github.com/GuanceCloud/guance-cli)
