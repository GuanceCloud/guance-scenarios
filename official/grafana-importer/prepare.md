## Prepare

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
