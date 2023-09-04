## Configure Alert Policy

> All the code is in `~/lab/src/alertpolicy.tf`, you can open it on the right side.

### Create member group by email

In this example, the alert policy will send alerts to a member group. So we need to get the member id by email. And create a new member group from them.

```terraform
data "guance_members" "demo" {
  filters = [
    {
      name   = "email"
      values = [var.email] # you can also add more emails here
    }
  ]
}

resource "guance_membergroup" "demo" {
  name       = var.name
  member_ids = data.guance_members.demo.items[*].id
}
```

### Create alert policy

The alert policy will send alerts to the member group we created above.

```terraform
resource "guance_alertpolicy" "demo" {
  name           = var.name
  silent_timeout = "1h"

  statuses = [
    "critical",
    "error",
    "warning",
    "info",
    "ok",
    "nodata",
    "nodata_ok",
    "nodata_as_ok",
  ]

  alert_targets = [
    {
      type = "member_group"
      member_group = {
        id = guance_membergroup.demo.id
      }
    }
  ]
}
```
