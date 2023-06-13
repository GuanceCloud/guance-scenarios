#!/bin/sh

# Install Terraform
curl -O https://releases.hashicorp.com/terraform/"$TF_VERSION"/terraform_"$TF_VERSION"_linux_amd64.zip
unzip terraform_"$TF_VERSION"_linux_amd64.zip -d /usr/local/bin/
rm terraform_"$TF_VERSION"_linux_amd64.zip

# Install Guance CLI
echo "deb [trusted=yes] https://releases.guance.io/apt/ /" | sudo tee /etc/apt/sources.list.d/guance.list
sudo apt update
sudo apt install guance
