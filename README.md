# 🎵 Telegram Detune Bot (Go + FFmpeg)

Telegram-бот на **Go**, который принимает **MP3 / WAV** файлы, предлагает выбрать частоту (detune / tuning) и возвращает аудио с изменённой высотой звука **без изменения темпа**.

Бот использует **FFmpeg** для качественного pitch shifting и работает на **Windows / Linux / macOS**.

https://youtube.com/shorts/5kk_1tL5PE4

---

## ✨ Возможности

- 📥 Приём **MP3 и WAV**
- 🎛 Выбор частоты через **inline-кнопки**
- 🎵 Pitch shift без изменения темпа
- 🧹 Автоматическая очистка временных файлов

---

## 🎚 Поддерживаемые частоты

- **440 Hz** — Standard Tuning (Modern Concert Pitch)
- **432 Hz** — Verdi / Scientific Tuning
- **415 Hz** — Baroque Tuning
- **392 Hz** — French Baroque Tuning
- **465 Hz** — Classical / Viennese Tuning
- **430 Hz** — 19th Century European Tuning
- **528 Hz** — Love Frequency / DNA Repair
- **396 Hz** — Liberating Fear & Guilt
- **417 Hz** — Facilitating Change
- **639 Hz** — Connection & Relationships
- **741 Hz** — Awakening Intuition
- **852 Hz** — Spiritual Order

---

## 🛠 Требования

### 1️⃣ Go
Версия **Go 1.18+**

### 2️⃣ FFmpeg
Должен быть доступен в PATH

---

### Инструкция

### 1️⃣ Клонируй репозиторий
```git clone https://github.com/yourname/detune-telegram-bot.git```
```cd detune-telegram-bot```

### 2️⃣ Инициализируй Go-модуль
```go mod init detune-bot```

### 3️⃣ Установи зависимости
```go get github.com/go-telegram-bot-api/telegram-bot-api/v5```

### 4️⃣ Создай папку для временных файлов
```mkdir temp```

---

⚙️ Настройка

Открой main.go и вставь токен бота:

```bot, err := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")```


Получить токен можно через @BotFather в Telegram.

---

▶️ Запуск
Запуск напрямую
```go run main.go```

Сборка бинарника
```go build -o detune-bot```
./detune-bot

