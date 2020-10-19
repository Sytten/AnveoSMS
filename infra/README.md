# Infrastructure

This is the infrastructure part of AnveoSMS. The infrastructure is hosted on GCP Cloud Run.
The monthly cost should not be more than a few cents per month.

## ⚠️ Requirements

Each requirement has a link to help you, feel free to open an issue if one requirement is not clearly explained.

- Golang >1.15: [Install](https://golang.org/doc/install)
- A GCP Project: [Tutoriel](https://www.techrepublic.com/article/how-to-create-your-first-project-on-google-cloud-platform/)
- The `gcloud` CLI: [Install](https://cloud.google.com/sdk/docs/install), [Configure](https://cloud.google.com/sdk/docs/initializing)
- The Pulumi CLI: [Install & Configure](https://www.pulumi.com/docs/get-started/install/)
- Docker: [Install](https://docs.docker.com/desktop/)
- A Sengrid API key: [Guide](https://github.com/Sytten/AnveoSMS/blob/main/docs/SENDGRID.md)

## 🏎️ Getting started

**Make sure you are in the `infra` folder for the next steps**

1. (Recommended) Checkout a tag: `git checkout vX.X.X`
2. Enable the following GCP APIs:
   - `gcloud services enable secretmanager.googleapis.com`
   - `gcloud services enable run.googleapis.com`
3. Build the configuration:
   - `pulumi config set gcp:project <YOUR PROJECT ID>`: Be careful to use the project ID here and not the project name
   - `pulumi config set gcp:region <YOUR DEPLOYMENT REGION>`: You can find all allowed regions [here](https://cloud.google.com/run/docs/locations), I suggest `us-east1`
   - `pulumi config set email:from`: The sender of the email, this should be the same as the email you used to register on Sendgrid
   - `pulumi config set email:to`: The receiver of the emails
   - `pulumi config set --secret email:sendgridApiKey`: The API key for Sendgrid
4. Deploy the whole thing: `pulumi up`
5. You should get as output an URL to use with Anveo as webhook:
   - If you do not, retry the previous command after a minute
   - [Follow this guide ](https://github.com/Sytten/AnveoSMS/blob/main/docs/ANVEO.md) to setup Anveo

## ⌛ Updating

1. Update the repo: `git pull`
2. Checkout a newer tag: `git checkout vX.X.X`
3. Deploy the new app: `pulumi up`
