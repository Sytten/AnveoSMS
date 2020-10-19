# Anveo

This guide will help you setup an Anveo account for its use with this project.

```golang
If you like this project, please consider using my referal code when subscribing: 2342832
Thanks!
```

1. Create a consumer account: https://www.anveo.com/consumer/default.asp. You will need to enter personnal information aobut your main residence but this is only for billign and legal purposes (it won't impact your ability to purchase a number from a country).
2. Once on the main screen, click on `Add funds` or navigate to https://www.anveo.com/addmoneyother.asp. Funds take some time to be process (took ~2h in my case).

<p align="center">
  <img width="300" src="https://raw.githubusercontent.com/Sytten/AnveoSMS/main/docs/assets/anveo_account.png">
</p>

3. When the funds are in, click on `Phones Numbers -> Order a New Phone Number` and select `Geographic` or `Mobile`. be sure to pick a phone number in a country where `SMS` are supported.

<p align="center">
  <img width="500" src="https://raw.githubusercontent.com/Sytten/AnveoSMS/main/docs/assets/anveo_order_new.png">
</p>

<p align="center">
  <img width="400" src="https://raw.githubusercontent.com/Sytten/AnveoSMS/main/docs/assets/anveo_order_choices.png">
</p>

4. Once you click on `Order phone numbers selected`, you must choose a plan. If this is only going to be used for SMS, I suggest going for the cheapest `per minute` plan.

<p align="center">
  <img width="300" src="https://raw.githubusercontent.com/Sytten/AnveoSMS/main/docs/assets/anveo_rate_plan.png">
</p>

5. Once the order is completed, click on `Phones Numbers -> Manage Phone Numbers` and then click on the `edit` button of your recently purchased phone number.

<p align="center">
  <img width="600" src="https://raw.githubusercontent.com/Sytten/AnveoSMS/main/docs/assets/anveo_manage.png">
</p>

6. Click on the `SMS` tab, tick the `Forward to URL` option and paste the link provided in the output of the applied infrastructure (in the [infra](https://github.com/Sytten/AnveoSMS/tree/main/infra) folder). Click `save` and you are done!

<p align="center">
  <img width="400" src="https://raw.githubusercontent.com/Sytten/AnveoSMS/main/docs/assets/anveo_phone_edit.png">
</p>
