Grafana Importer Introduction
=============================

Importer Grafana dashboard by Guance CLI tookit

Introduction
------------

![Preview](./images/preview.png)

Taking Ubuntu as an example, this guide demonstrates using the command line interface to import the Grafana dashboard [Node Exporter Full](https://grafana.com/grafana/dashboards/1860-node-exporter-full/) using Guance CLI.

We will complete the following tasks:

1.	Import Grafana JSON as Guance dashboard format
2.	Save it to Guance Cloud
3.	Collect the demo metrics
4.	Succeed!

Step 0: Prepare
---------------

### Install Guance CLI

The first step is to install the Guance CLI by running the command:

```bash
echo "deb [trusted=yes] https://releases.guance.io/apt/ /" | sudo tee /etc/apt/sources.list.d/guance.list
sudo apt update
apt install guance
```

For other ways to install Guance CLI, please refer to [Guance CLI Installation](https://github.com/GuanceCloud/guance-cli).

### Pre-installed environment

In this course, we will use some other tools only for demos. All of the related tools are installed in the lab environment. So you don't need to install them again.

1.	Terraform
2.	Docker
3.	Docker Compose

### Login Guance API

First, you should log in to Guance with your account.

1.	Open the [Guance Console](https://console.guance.io/) in your browser.
2.	Login with your account.
3.	Open the [API Key Management](https://console.guance.com/workspace/apiManage) page, create a new key, and copy the `Key ID` content.
4.	Open the [Workspace Settings](https://console.guance.com/workspace/detail) page and copy the `Token` content.
5.	Open the integrated terminal of IDE on the right side, and run the following command to log in to Guance:

	```shell
	export GUANCE_ACCESS_TOKEN=*** # your key id
	export GUANCE_WORKSPACE_TOKEN=*** # your token
	```

Step 1: Download the Grafana Dashboard JSON
-------------------------------------------

Then, download the [Node Exporter Dashboard on Grafana](https://grafana.com/grafana/dashboards/1860-node-exporter-full/).

Or use `wget` to download it:

```shell
wget https://grafana.com/api/dashboards/1860/revisions/31/download -O grafana.json
```

After running the command, a Grafana dashboard JSON file named `grafana.json` will be downloaded to the local disk. We could use it to import this file on the Guance Cloud Console.

Step 2: Convert the Grafana dashboard to Guance's
-------------------------------------------------

The next step is to use Guance CLI to convert the Grafana dashboard JSON file into a Terraform module.

### Run CLI

The CLI command is:

```shell
guance iac import grafana -f ./grafana.json -t terraform-module -o ./out -m prom
```

-	**-f**: The Grafana dashboard JSON file path.
-	**-t**: The output type. The value is `terraform-module`.
-	**-o**: The output folder path.
-	**-m**: The measurement. The default value is `prom`. If set to empty string, It will auto split metric name by `_`, the first part is the measurement.

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

Step 3: Upload the dashboard
----------------------------

With this, you can use it to generate resources on the Guance Cloud platform.

There are two ways to upload the dashboard to Guance Cloud:

1.	Upload at console directly
2.	Use Terraform to create a dashboard

Any of the above methods can be used to upload the dashboard to Guance Cloud. We recommend using the Terraform method in this course because it suits large-scale automation deployment.

### Upload to console directly

To use the Guance Cloud console, please follow these steps:

1.	Open [Create Dashboard Page](https://console.guance.com/scene/dashboard/createDashboard).
2.	Click `Import Template`.
3.	Import the generated JSON file.

### Use Terraform to create a dashboard

Run the following commands to create the dashboard:

```shell
# Change the directory to the output folder
cd ./out

# Initialize the Guance provider
terraform init

# Create the dashboard
terraform apply -var 'name=Demo'
```

The output is:

```shell
Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # guance_dashboard.main will be created
  + resource "guance_dashboard" "main" {
      + created_at = (known after apply)
      + id         = (known after apply)
      + manifest   = jsonencode(
            {
              + dashboardBindSet   = []
              + dashboardExtend    = {}
              + dashboardMapping   = []
              + dashboardOwnerType = "node"
              + dashboardType      = "CUSTOM"
              + iconSet            = {}
              + main               = {
                  + charts = []
                  + groups = []
                  + type   = "template"
                  + vars   = [
                      + ...
                    ]
                }
              + summary            = "..."
              + tagInfo            = []
              + tags               = []
              + thumbnail          = ""
              + title              = "Demo"
            }
        )
      + name       = "Demo"
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + dashboard_id = (known after apply)

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: 
```

Enter `yes` to confirm the operation.

Step 4: Run the DataKit to collect metrics from Prometheus
----------------------------------------------------------

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

1.	Collect the metrics from `http://node-exporter:9100/metrics`.
2.	(Optional) Set the source and measurement name to `prom`.
3.	(Optional) Set the `job` tag to `DataKit`.

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

-	`node-exporter`: The Prometheus exporter for hardware and OS metrics.
-	`datakit`: The DataKit agent collects metrics from Prometheus and uploads them to Guance Cloud.

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

Step 5: Completed!
------------------

You have successfully imported a dashboard into Grafana. You can now view the dashboard on [Dashboard List](https://console.guance.com/scene/dashboard/list).

Here is a screenshot for the successful result.

![Preview](./images/preview.png)

If you want to learn more, some materials may be helpful.

1.	[All the examples about Grafana Importer](https://github.com/GuanceCloud/guance-cli/tree/main/specs/iac/import/grafana)
2.	[Guance CLI Homepage](https://github.com/GuanceCloud/guance-cli)
