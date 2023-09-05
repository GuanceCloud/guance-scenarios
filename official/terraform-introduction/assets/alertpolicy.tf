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
