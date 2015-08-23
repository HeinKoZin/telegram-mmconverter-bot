# Telegram MMConvter Bot

A [Telegram](https://telegram.org) bot which allows you to converter zg <-> uni.

## Usage 

- Go to https://telegram.me/MMConverterBot and install it.
- `/z <Your Zg text> for Zg to Uni conversion.
- `/u <Your Uni text>` for Uni to Zg conversion.
- The default is Zg to Uni converston.

## Development 

Please install [Go](http://www.golang.org) and set up first, if you haven't.

```bash
$ go get github.com/yelinaung/telegram-mmconverter-bot
$ cd $GOPATH/src/github.com/yelinaung/telegram-mmconverter-bot
$ export TELEGRAM_MMCONVERTER_TOKEN="your token" && export HEXCORE_MMCONVERTER_TOKEN="token" && go run app.go
```

## Credits

Thanks to [MMConverter](http://mmconverter.hexcores.com) API by [Hexcores](http://hexcores.com).

## License
MIT

