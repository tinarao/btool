package tg

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/tinarao/btool/internal/config"
)

type TelegramBot struct {
	bot    *bot.Bot
	chatID int64
}

func New() (*TelegramBot, error) {
	b, err := bot.New(config.Cfg.BotToken)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		bot:    b,
		chatID: config.Cfg.ChatId,
	}, nil
}

func (t *TelegramBot) SendMessage(ctx context.Context) {
	params := &bot.SendMessageParams{
		ChatID: config.Cfg.ChatId,
		Text:   "helo",
	}

	t.bot.SendMessage(ctx, params)
}

func (t *TelegramBot) SendFile(ctx context.Context, filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open the file: %s", err.Error())
	}
	defer file.Close()

	params := &bot.SendDocumentParams{
		ChatID: t.chatID,
		Document: &models.InputFileUpload{
			Filename: file.Name(),
			Data:     file,
		},
	}

	_, err = t.bot.SendDocument(ctx, params)
	return err
}

func (t *TelegramBot) SendFiles(ctx context.Context, filePaths []string) []error {
	var errors []error

	for _, path := range filePaths {
		if err := t.SendFile(ctx, path); err != nil {
			errors = append(errors, fmt.Errorf("%s: %s", path, err.Error()))
			log.Printf("failed to send a file %s: %s", path, err.Error())
			continue
		}

		log.Printf("successfully send: %s", path)
	}

	return nil
}
