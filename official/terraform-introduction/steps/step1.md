# Step1: Configure Monitor

> All the code is in `~/lab/src/monitor.tf`, you can open it on the right side.

Monitor is the core of Guance. It is a JSON file that describes the metrics to be collected and the alert policies to be applied.

## Create a monitor

See an example file named `monitor.json` with the following content:

```shell
cat ~/lab/src/monitor.json
```

Then create a terraform file to define the monitor:

```terraform
resource "guance_monitor" "demo" {
  manifest     = file("${path.module}/monitor.json")
  alert_policy = {
    id = guance_alertpolicy.demo.id
  }
}
```
