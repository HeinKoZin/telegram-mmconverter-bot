package main

import (
	"encoding/json"
	"fmt"
	"github.com/tucnak/telebot"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

type Msg struct {
	ConvertedText string `json:"convertedText"`
	OriginalText  string `json:"originalText"`
}

var mmConverterApiKey = os.Getenv("HEXCORE_MMCONVERTER_TOKEN")
var zg2uniApi = "http://mmconverter.hexcores.com/api/v1/zg2uni"
var uni2zgApi = "http://mmconverter.hexcores.com/api/v1/uni2zg"

func main() {

	bot, err := telebot.NewBot(os.Getenv("TELEGRAM_MMCONVERTER_TOKEN"))
	panicIf(err)

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		zgInput, _ := regexp.MatchString("/z *", message.Text)
		uniInput, _ := regexp.MatchString("/u *", message.Text)
		if message.Text == "/hi" {
			bot.SendMessage(message.Chat,
				"Hello, "+message.Sender.FirstName+"!", nil)
		} else if zgInput {
			convertedUniString := convert(message.Text, true)
			bot.SendMessage(message.Chat, convertedUniString, nil)
		} else if uniInput {
			convertedUniString := convert(message.Text, false)
			bot.SendMessage(message.Chat, convertedUniString, nil)
		} else {
			convertedUniString := convert(message.Text, true)
			bot.SendMessage(message.Chat, convertedUniString, nil)
		}
	}
}

func convert(input string, isZg bool) string {
	var resp *http.Response
	var err error

	if isZg {
		input = strings.Replace(input, "/z ", "", 1)
		resp, err = http.PostForm(zg2uniApi, url.Values{"q": {input}, "key": {mmConverterApiKey}})
	} else {
		input = strings.Replace(input, "/u ", "", 1)
		resp, err = http.PostForm(uni2zgApi, url.Values{"q": {input}, "key": {mmConverterApiKey}})
	}

	defer resp.Body.Close()
	panicIf(err)
	fmt.Println("status : ", resp.Status)

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
	panicIf(err)

	var msg Msg
	err = json.Unmarshal([]byte(jsonDataFromHttp), &msg)
	panicIf(err)

	return msg.ConvertedText
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
