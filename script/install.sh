#!/bin/bash
if ! command -v wget &> /dev/null; then
    if command -v apt-get &> /dev/null; then
        sudo apt-get install -y wget
    elif command -v yum &> /dev/null; then
        sudo yum install -y wget
    elif command -v dnf &> /dev/null; then
        sudo dnf install -y wget
    elif command -v pacman &> /dev/null; then
        sudo pacman -Sy --noconfirm wget
    else
        echo "Unable to install wget. Please install it manually."
        exit 1
    fi
fi

sudo mkdir -p /opt/safe-monitor
wget -O /opt/safe-monitor/safe-monitor-linux-amd64.tar.gz https://github.com/ArsFy/linux-safe-monitor/releases/download/v0.1/safe-monitor-linux-amd64.tar.gz
sudo tar -xzf /opt/safe-monitor/safe-monitor-linux-amd64.tar.gz -C /opt/safe-monitor
sudo rm /opt/safe-monitor/safe-monitor-linux-amd64.tar.gz
sudo chmod +x /opt/safe-monitor/safe-monitor

read -p $'\nEnter notification language [en, cn](default: en): ' lang
if [[ "$lang" != "en" && "$lang" != "cn" ]]; then
    echo "Use default: en"
    lang="en"
fi
echo <<EOF

-1: No remind and kill
0: Only remind
1: Remind and kill (White list)
2: Remind and kill (All)
EOF
read -p "Enter Kill mode [-1, 0, 1, 2, 3](default: 0): " kill_mode
if [[ "$kill_mode" != "-1" && "$kill_mode" != "0" && "$kill_mode" != "1" && "$kill_mode" != "2" && "$kill_mode" != "3" ]]; then
    echo "Use default: 0"
    kill_mode="0"
fi

read -p $'Enter check time (ms) (default: 5000): ' check_time
check_time=${check_time:-5000}

read -p "Use telegram bot? [false/true] (default: false): " enable_telegram
if [[ "$enable_telegram" != "false" && "$enable_telegram" != "true" ]]; then
    echo "Use default: false"
    enable_telegram="false"
fi

if [[ "$enable_telegram" == "true" ]]; then
    read -p "Enter telegram bot token: " telegram_token
    read -p "Enter telegram chat id: " telegram_chat_id
fi

sudo cat > /opt/safe-monitor/config.json <<EOF
{
    "lang": "$lang",
    "kill_mode": $kill_mode,
    "check_time": $check_time,
    "enable_telegram": $enable_telegram,
    "telegram_token": "$telegram_token",
    "telegram_chat_id": $telegram_chat_id
}
EOF

echo "Edit /opt/safe-monitor/config.json"

sudo cat > /etc/systemd/system/safe-monitor.service <<EOF
[Unit]
Description=safe-monitor

[Service]
Type=simple
WorkingDirectory=/opt/safe-monitor/
ExecStart=sudo /opt/safe-monitor/safe-monitor
Restart=always
RestartSec=5
StartLimitInterval=3
RestartPreventExitStatus=137
 
[Install]
WantedBy=multi-user.target
EOF

service safe-monitor start

echo "service safe-monitor [start|stop|restart|status]"