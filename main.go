package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Tuning struct {
	Hz   float64
	Name string
}

var tunings = map[string]Tuning{
	"440": {440, "Standard Tuning (Modern Concert Pitch)"},
	"432": {432, "Verdi / Scientific Tuning"},
	"415": {415, "Baroque Tuning"},
	"392": {392, "French Baroque Tuning"},
	"465": {465, "Classical / Viennese Tuning"},
	"430": {430, "19th Century European"},
	"528": {528, "Love Frequency / DNA Repair"},
	"396": {396, "Liberating Fear & Guilt"},
	"417": {417, "Facilitating Change"},
	"639": {639, "Connection & Relationships"},
	"741": {741, "Awakening Intuition"},
	"852": {852, "Spiritual Order"},
}

func main() {
	bot, err := tgbotapi.NewBotAPI("СЮДА ТОКЕН")
	if err != nil {
		log.Fatal(err)
	}

	os.MkdirAll("temp", 0755)
	clearTemp()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil && update.Message.IsCommand() {
			if update.Message.Command() == "start" {
				sendStart(bot, update.Message.Chat.ID)
			}
			continue
		}

		if update.Message != nil {
			if update.Message.Audio != nil || update.Message.Document != nil {
				sendKeyboard(bot, update.Message)
			}
		}

		if update.CallbackQuery != nil {
			processTuning(bot, update.CallbackQuery)
		}
	}
}

func sendStart(bot *tgbotapi.BotAPI, chatID int64) {
	text := `
🎵 Detune Bot

Отправь MP3 или WAV файл,
затем выбери нужную частоту.

Бот изменит высоту звука
без изменения темпа.

`
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}

func sendKeyboard(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	var rows [][]tgbotapi.InlineKeyboardButton

	for key, t := range tunings {
		text := fmt.Sprintf("%s Hz — %s", key, t.Name)
		btn := tgbotapi.NewInlineKeyboardButtonData(text, key)
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(btn))
	}

	reply := tgbotapi.NewMessage(msg.Chat.ID, "Выберите детюн:")
	reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	reply.ReplyToMessageID = msg.MessageID
	bot.Send(reply)
}

func processTuning(bot *tgbotapi.BotAPI, q *tgbotapi.CallbackQuery) {
	defer clearTemp()

	tuning := tunings[q.Data]

	var fileID, ext string

	if q.Message.ReplyToMessage.Audio != nil {
		fileID = q.Message.ReplyToMessage.Audio.FileID
		ext = ".mp3"
	} else if q.Message.ReplyToMessage.Document != nil {
		fileID = q.Message.ReplyToMessage.Document.FileID
		ext = filepath.Ext(q.Message.ReplyToMessage.Document.FileName)
	} else {
		return
	}

	file, _ := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	url := file.Link(bot.Token)

	input := "temp/input_" + fileID + ext
	output := "temp/output_" + q.Data + ".mp3"

	downloadFile(url, input)
	pitchShift(input, output, tuning.Hz)

	audio := tgbotapi.NewAudio(q.Message.Chat.ID, tgbotapi.FilePath(output))
	audio.Caption = fmt.Sprintf("%s Hz — %s", q.Data, tuning.Name)
	bot.Send(audio)
}

func pitchShift(in, out string, targetHz float64) {
	ratio := targetHz / 440.0

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", in,
		"-af", "asetrate=44100*"+formatFloat(ratio)+",aresample=44100",
		out,
	)

	cmd.Run()
}

func downloadFile(url, path string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	io.Copy(out, resp.Body)
}

func clearTemp() {
	files, _ := os.ReadDir("temp")
	for _, f := range files {
		os.Remove(filepath.Join("temp", f.Name()))
	}
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}
