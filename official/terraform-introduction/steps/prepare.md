## Prepare the environment

### Lab environment

In this course we will use the [Terraform](https://www.terraform.io/) to create the observability stack. All of the related tools is installed in the lab environment. You can open the integrated-terminal of IDE on the right side.

### Login Guance

The first, you should login Guance with your account.

1. Open the [Guance Console](https://console.guance.io/) in your browser.
2. Login with your account.
3. Open the [API Key Management](https://console.guance.com/workspace/apiManage) page, create a new key, and copy the `Key ID` content.
4. Open the integrated-terminal of IDE on the right side, and run the following command to login Guance:

```shell
export GUANCE_ACCESS_TOKEN=<your key id>
```

### Understand the lab environment

In the right side, you can see the lab environment. It contains the following files and directories:

```shell
src/ # The directory contains all the terraform files.
```
