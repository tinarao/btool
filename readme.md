# btool

<p align="center">
    <img src="https://skillicons.dev/icons?i=go" />
</p>

Простенькая утилита для создания автоматических бекапов. Использует Telegram как площадку для хранения файлов.

## Как это работает

btool архивирует указанные пути и сохраняет архив в две локации - в `target_dir` и в переписку с Telegram-ботом.

## Конфигурация

btool использует конфиг в yaml формате. Создайте в домашней директории файл btool.yaml со следующим содержанием:

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

```bash
go build
go install

btool
```

Для запуска используйте `btool`. При возникновении проблем с установкой или компиляцией проверьте пути ([инструкцию](https://go.dev/doc/tutorial/compile-install)).

## TODO

Проект ещё в разработке.

- [x] Архивация
- [x] Интеграция с Telegram
- [ ] История бекапов
- [ ] CLI с различным функционалом (узнать время последней синхронизации, запустить, изменить конфиги и пр.)
