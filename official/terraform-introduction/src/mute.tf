resource "guance_mute" "demo" {
  // mute ranges
  mute_ranges = [
    {
      type = "alert_policy"

      alert_policy = {
        id = guance_alertpolicy.demo.id
      }
    }
  ]

  // notify options
  notify = {
    message = <<EOF
      Muted
    EOF

    before_time = "15m"
  }

  notify_targets = [
    {
      type = "member_group"

      member_group = {
        id = guance_membergroup.demo.id
      }
    }
  ]

  // cron options
  repeat = {
    crontab_duration = "30s"
    start            = "05:00:00"
    end              = "10:00:00"
    expire           = "2023-12-31T12:00:00Z"
    crontab          = {
      min   = "0"
      hour  = "0"
      day   = "*"
      month = "*"
      week  = "*"
    }
  }

  mute_tags = [
    {
      key   = "host"
      value = "*"
    }
  ]
}
