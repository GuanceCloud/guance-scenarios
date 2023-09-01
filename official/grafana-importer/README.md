Grafana Importer Introduction
=============================

Importer Grafana dashboard by Guance CLI tookit

[TOC]

## Introduction

![Preview](images/preview.png)

Step 0: Prepare
---------------

### Install Guance CLI Toolkit

```bash
echo "deb [trusted=yes] https://releases.guance.io/apt/ /" | sudo tee /etc/apt/sources.list.d/guance.list
sudo apt update
apt install guance
```

### Install the additional tools for the lab

```bash
apt install docker docker-compose terraform
```

Step 1: Download the Grafana Dashboard JSON
-------------------------------------------

Then download the [Node Exporter Dashboard on Grafana](https://grafana.com/grafana/dashboards/1860-node-exporter-full/).

Or use `wget` to download it:

```shell
wget https://grafana.com/api/dashboards/1860/revisions/31/download -O grafana.json
```

Step 2: Use Guance CLI to convert grafana as Terraform module
-------------------------------------------------------------

```shell
guance iac import grafana -f ./input.json -t terraform-module -o ./out
```

You will get a Terraform module at `./out` folder. The folder structure is:

```
.
├── main.tf
├── manifest.json
├── outputs.tf
├── variables.tf
└── versions.tf
```

The `manifest.json` is the dashboard json file for Guance Cloud. You can also use it to import the dashboard to Guance Cloud Console directly.

In this example, we will use the Terraform module to create the dashboard on Guance Cloud.

Step 3: Run the Terraform apply to create dashboard
---------------------------------------------------

So you can apply it to create the real dashboard resources on Guance Cloud.

```shell
cd ./out
terraform init
terraform apply
```

Step 4: Run the DataKit to collect metrics from Prometheus
----------------------------------------------------------

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
