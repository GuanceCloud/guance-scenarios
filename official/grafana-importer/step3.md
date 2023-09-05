## Upload the dashboard

With this, you can use it to generate resources on the Guance Cloud platform.

There are two ways to upload the dashboard to Guance Cloud:

1. Upload at console directly
2. Use Terraform to create a dashboard

Any of the above methods can be used to upload the dashboard to Guance Cloud. We recommend using the Terraform method in this course because it suits large-scale automation deployment.

### Upload to console directly

To use the Guance Cloud console, please follow these steps:

1. Open [Create Dashboard Page](https://console.guance.com/scene/dashboard/createDashboard).
2. Click `Import Template`.
3. Import the generated JSON file.

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
