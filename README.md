
# telegram-pusher-hub

1. Create new bot using [@BotFather](//t.me/BotFather)
2. Create new channel
3. Create bot to channel admins
4. Start **telegram-pusher-hub**: `docker run --rm --net=host ivan1993spb/telegram-pusher-hub --channel_username=@<channel-username> --token=<telegram-bot-token>`

Tests:

```
curl -X POST -d "text=text <b>bold</b> ok" -d parse_mode=HTML http://localhost:8080/api/v1/message
curl -X POST -d "text=text *bold* MD" -d parse_mode=Markdown http://localhost:8080/api/v1/message
```

For details see `docker run --rm ivan1993spb/telegram-pusher-hub --help`
