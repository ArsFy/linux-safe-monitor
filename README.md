## Linux Safe Monitor

> This project is in development

Login log monitoring, process monitoring, push notification on status changes, and support for automatically terminating processes outside the whitelist.

### What can it do?

- Monitor all processes on the server and push notifications when new processes appear or when processes terminate.

- In development ...

### Support Notification methods

- Telegram
- In development ...

### Manual installation

#### 1. Download

Download the compressed package in [releases](releases) or [Build](#build)

```
linux-safe
    - safe-monitor
    - white-list.json
    - config.json
```

#### 2. Edit Config

- `lang` : Language ( `en` / `cn` ) 
- `kill_mode`:
    - `-1` : No remind and kill
    - `0` : Only remind
	- `1` : Remind and kill (White list)
	- `2` : Remind and kill (All)
- `check_time` : Interval time (ms)
- `enable_telegram` : `true` / `false`
- `telegram_token` : [@BotFather](https://t.me/BotFather)
- `telegram_chat_id` : UserID or ChatID (Int64)

```json
{
    "lang": "en",
    "kill_mode": 0,
    "check_time": 5000,
    "enable_telegram": false,
    "telegram_token": "",
    "telegram_chat_id": 0
}
```

#### 3. Run

```
./safe-monitor
```

or add linux service

[<img src="https://opengraph.githubassets.com/0ce367d2a8cee652c1242cb4a99af11939ad2161e47eac849791a8695027a549/ArsFy/add_service" width="40%" />](https://github.com/ArsFy/add_service)

-----

#### <span id="build">Build</span>

```
git clone https://github.com/ArsFy/linux-safe-monitor.git
cd linux-safe-monitor
go mod tidy
go build .
```