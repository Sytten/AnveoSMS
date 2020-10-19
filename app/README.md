# Application

This is the application project itself, please follow these instructions **only if** running locally or self-hosting.

## ⚠️ Requirements

- Golang >1.15
- A Sengrid API key: [Guide](https://github.com/Sytten/AnveoSMS/blob/main/docs/SENDGRID.md)
- A Public Bucket or CDN

## 🏎️ Getting started

1. Copy the content of the [assets folder of infra](https://github.com/Sytten/AnveoSMS/tree/main/infra/assets) in the bucket
2. Follow the steps in [configuration](#📚-Configuration)
3. Execute `make run` to launch the server
4. Execute `curl http://localhost:9000/webhooks/anveo?from=test&to=me&message=Hello`
5. You should receive an email if everything went well!

## 📚 Configuration

- An example configuration is provided in `config.example.yml`
- The configuration name is `config.yml`
- The configuration should be beside the binary
