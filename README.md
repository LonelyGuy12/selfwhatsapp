# selfwhatsapp

A self-hosted WhatsApp bot that runs on your own account using [whatsmeow](https://github.com/tulir/whatsmeow). No third-party servers — your session stays on your machine.

## How it works

The bot connects to WhatsApp's WebSocket as your account. It listens for incoming messages and can react to them automatically. Your session is stored in a local SQLite database so you only need to scan the QR code once.

## Project structure

```
main.go          — entry point, connects to WhatsApp and blocks until Ctrl+C
src/
  store.go       — opens the SQLite database that holds your session
  client.go      — creates the WhatsApp client and handles QR / session login
  handler.go     — receives all incoming events and acts on messages
```

## Requirements

- Go 1.21+
- GCC (needed to compile `go-sqlite3`)
  - macOS: `xcode-select --install`
  - Linux: `sudo apt install gcc`

## Setup

```bash
git clone https://github.com/LonelyGuy12/selfwhatsapp.git
cd selfwhatsapp
go mod download
```

## Running

```bash
go run main.go
```

On first run a QR code is printed in the terminal. Scan it with WhatsApp on your phone (**Linked Devices → Link a Device**). After that your session is saved in `whatsapp.db` and the QR code will not appear again.

## Auto-replies

| Trigger (anyone messages you) | Bot replies |
|-------------------------------|-------------|
| `hi` (case-insensitive)       | `hewwo senpai! :3` |

## Commands

Send a message **from your own account** starting with `*&` to trigger a command.

Example: `*&lyrics` — detected and printed to the terminal (extend `handler.go` to do something with it).

## Adding new auto-replies

Open `src/handler.go` and add another `strings.EqualFold` check inside the `!v.Info.IsFromMe` block:

```go
if strings.EqualFold(strings.TrimSpace(text), "hello") {
    Client.SendMessage(context.Background(), v.Info.Chat, &waE2E.Message{
        Conversation: proto.String("hewwo! :3"),
    })
}
```

## Session file

Your WhatsApp session is stored in `whatsapp.db`. This file is listed in `.gitignore` and should never be committed or shared.
