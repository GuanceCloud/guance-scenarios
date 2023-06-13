# Terraform Introduction

Terraform is an open-source infrastructure as code software tool created by HashiCorp. It enables users to define and provision an infrastructure using a high-level configuration language known as HashiCorp Configuration Language (HCL), or optionally JSON.

You can see the follow resources for details:

1. [Terraform Website](https://www.terraform.io/)
2. [Terraform Guance Cloud Provider](https://registry.terraform.io/providers/GuanceCloud/guance)

![Intro](./images/intro.svg)

In this lab, we will implement a typical monitoring and alerting scenario together. Configure a monitor and an alert policy to notify a member group or notification object, such as an IM group, whenever a condition is triggered. At the same time, configure an additional silent rule to silence the alert message when the monthly scheduled downtime arrives for system maintenance.
