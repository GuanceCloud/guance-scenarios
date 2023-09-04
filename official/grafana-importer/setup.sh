#!/bin/sh

TF_VERSION=1.5.6

# Install Terraform
curl -O https://releases.hashicorp.com/terraform/"$TF_VERSION"/terraform_"$TF_VERSION"_linux_amd64.zip
unzip terraform_"$TF_VERSION"_linux_amd64.zip -d /usr/local/bin/
rm terraform_"$TF_VERSION"_linux_amd64.zip
