# btool

<p align="center">
    <img src="https://skillicons.dev/icons?i=go" />
</p>

Простенькая CLI-утилита для создания автоматических бекапов. Сохраняет архивы в указаную локальную директорию и сохраняет их в Telegram.

## Конфигурация

btool использует конфиг в yaml формате. Создайте в домашней директории файл `btool.yaml` со следующим содержанием:

```yaml
bot_token: "<your_bot_token>"
target_dir: "~/backups/"
chat_id: "<your_chat_id>"
paths: [
	"/home/<your_os_username>/Documents/Obsidian Vault",
	"/home/<your_os_username>/important files",
	"/home/<your_os_username>/some other stuff"
]
```

- `bot_token` содержит токен созданного Вами Telegram-бота.
- `target_dir` содержит путь, в который будут сохраняться архивы.
- `paths` - массив путей, бекап которых вы хотите регулярно совершать.
- `chat_id` - ID чата, в который бот будет отправлять сообщения.

В поле `chat_id` рекомендуется вставить свой ID (перед этим бота нужно запустить). Свой ID можно узнать через [@userinfobot](https://t.me/userinfobot).

## Использование

### Требования

- Go >1.24.6

### Установка

Для установки выполните: 

```bash
git clone https://github.com/tinarao/btool.git 
cd btool
make install
```

Это запустит bash-скрипт, который добавит необходимые пути в `.zshrc` и `.bashrc` и запустит процесс установки.

### Команды

- `btool run` - запуск синхронизации указанных в конфиге путей.
- `btool last_backup` - выводит в консоль дату последней синхронизации.
- `btool help` - выводит список доступных команд.

## TODO

- [x] Архивация
- [x] Интеграция с Telegram
- [x] CLI с различным функционалом (узнать время последней синхронизации, запустить, изменить конфиги и пр.)
- [ ] Backup конфиг-файла перед операциями записи
