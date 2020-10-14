# Infrastructure

This is the infrastructure part of AnveoSMS. The infrastructure is hosted on GCP on Cloud Run. 
The monthly cost should not be more than a few cents per month.

## Requirements
- GCP Project
- Configured `gcloud` CLI
- Pulumi CLI
- Sendgrid API with at least `Send` permission

## Setup
Before we can deploy the infrastructure, we need to enable a few GCP APIs:
- `gcloud services enable secretmanager.googleapis.com`
