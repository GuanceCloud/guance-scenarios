## Use Guance CLI to convert grafana as Terraform module

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
