# AnveoSMS

![Golang version](https://img.shields.io/github/go-mod/go-version/sytten/anveosms?filename=app%2Fgo.mod)
[![Release](https://github.com/sytten/anveosms/workflows/Release/badge.svg)](https://github.com/Sytten/AnveoSMS/releases)
[![Version](https://img.shields.io/github/v/tag/sytten/anveosms)](https://github.com/Sytten/AnveoSMS/packages/445142)

<p align="center">
  <img width="460" src="https://raw.githubusercontent.com/Sytten/AnveoSMS/main/.github/assets/logo.png">
</p>

AnveoSMS is a small program to receive your SMS from [Anveo](http://www.anveo.com) and forward them to an email address of your choice. [Anveo](http://www.anveo.com) is a great provider is you need a phone number to receive 2FA since it supports **short codes**. If you are willing to pay for a premium plan, [Anveo](http://www.anveo.com) does offer the same feature but if you are a thinkerer like me I invite you to use this application!

## ✅ Current features

- Receive SMS from [Anveo](http://www.anveo.com)
- Send them to any email using [Sengrid](https://sendgrid.com/)
- Deployment infrastructure on GCP

## 🏗️ How to deploy?

- The suggested way of deploying it is using the provided `Infrastructure as Code` in the [infra](https://github.com/Sytten/AnveoSMS/tree/main/infra) folder. This will deploy it on GCP Cloud Run which is very cheap. As I build more and more features, I plan to rely heavily on GCP features so this will eventually be the only way of deploying.
- If you want to host it yourself, you can use the [Docker images](https://github.com/Sytten/AnveoSMS/packages/445142) I build for every release. See the configuration setup in the [app](https://github.com/Sytten/AnveoSMS/tree/main/app) folder.

## 💡 Ideas for future features

- Store them in a database
- An API for access
- Web portal
- Mobile application
