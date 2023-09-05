resource "guance_monitor" "demo" {
  manifest     = file("${path.module}/monitor.json")
  alert_policy = {
    id = guance_alertpolicy.demo.id
  }
}
