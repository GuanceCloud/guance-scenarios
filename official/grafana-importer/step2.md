## Use Guance CLI to convert grafana as Terraform module

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

1. Open [Create Dashboard Page](https://console.guance.com/scene/dashboard/createDashboard).
2. Click `Import Template`.
3. Import the generated JSON file.

In this example, we will use the Terraform module to create the dashboard on Guance Cloud in the next step.
