# sing-box-daily-telegram
sing box with send configuration in the telegram channel every day.


#Persian Explain
[Persian Article](https://telegra.ph/Small-family-servers-05-17)

# Copywriting
This project is fork of [sing-REALITY-Box](https://github.com/deathline94/sing-REALITY-Box).
The main Idea is combine [sb-server-configer](https://github.com/hrostami/sb-server-configer) with bash script.
It means that implement outstanding feature [sb-server-configer] with bash script.

# How to use
Clone the Project and run the sing-REALITY-Box bash script

```
git clone https://github.com/GostOfSarina/sing-box-daily-telegram.git
```

```
bash ./sing-REALITY-Box
```



## you to fill these files with your own information.


We have three configuration options.

```/root/public_key.txt```
I add store public key in the original project folder.


```/root/bot_token.txt```

```/root/chat_id.txt```


get bot token from [BotFather](https://t.me/BotFather)
get chat id from [Find Channel id](https://gist.github.com/mraaroncruz/e76d19f7d61d59419002db54030ebe35)

public key is automatically make with sing-Realty-Box script.

# Setup the cronjob
```
bash ./cronjob.sh
```

you can change the cronjob time in the cronjob.sh file. [easy set the time](https://crontab.guru/)


see the cronjob list
```crontab -l```

# Get New Configuration
```
bash ./renew.sh
```
