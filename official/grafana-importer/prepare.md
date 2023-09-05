## Prepare

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

1. Terraform
2. Docker
3. Docker Compose

### Login Guance API

First, you should log in to Guance with your account.

1. Open the [Guance Console](https://console.guance.io/) in your browser.
2. Login with your account.
3. Open the [API Key Management](https://console.guance.com/workspace/apiManage) page, create a new key, and copy the `Key ID` content.
3. Open the [Workspace Settings](https://console.guance.com/workspace/detail) page and copy the `Token` content.
4. Open the integrated terminal of IDE on the right side, and run the following command to log in to Guance:

```shell
export GUANCE_ACCESS_TOKEN=*** # your key id
export GUANCE_WORKSPACE_TOKEN=*** # your token
```
