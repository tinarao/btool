package main

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tinarao/btool/internal/config"
	"github.com/tinarao/btool/internal/tg"
	"github.com/urfave/cli/v3"
)

func main() {
	config.Load()

	c := &cli.Command{
		Name:  "btool",
		Usage: "simple backup tool integrated with Telegram",
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					b, err := tg.New()
					if err != nil {
						fmt.Printf("failed to start a bot: %s", err.Error())
						os.Exit(1)
					}

					CreateBackup(b)

					return nil
				},
			},
			{
				Name:    "last_backup",
				Aliases: []string{"lb"},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(config.Cfg.LastBackupDate)
					return nil
				},
			},
		},
	}

	ctx := context.Background()
	c.Run(ctx, os.Args)
}

func CreateBackup(bot *tg.TelegramBot) {
	for _, path := range config.Cfg.Paths {
		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("failed to get user home directory: %s\n", err.Error())
			os.Exit(1)
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("warning: directory %s does not exist, skipping\n", path)
			continue
		}

		timestamp := GetTodayIsoDate()

		s := strings.Split(path, "/")
		lastEntry := s[len(s)-1]

		targetDir := filepath.Join(homedir, config.Cfg.TargetDir)
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			os.Mkdir(targetDir, 0755)
		}

		filename := fmt.Sprintf("%s-%s.tar.gz", lastEntry, timestamp)
		target := filepath.Join(targetDir, filename)

		fmt.Printf("archiving %s to %s\n", path, target)

		if err := TarDirectory(path, target); err != nil {
			fmt.Printf("failed to tar: %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("successfully archived: %s\n", target)

		if err := bot.SendFile(context.Background(), target); err != nil {
			fmt.Printf("failed to send a file via Telegram: %s\n", err.Error())
			os.Exit(1)
		}

		config.Cfg.SetLastBackupTime(timestamp)
	}
}

func TarDirectory(source, target string) error {
	file, err := os.Create(target)
	if err != nil {
		return err
	}

	defer file.Close()

	var writer io.Writer = file
	if filepath.Ext(target) == ".gz" {
		gzipWriter := gzip.NewWriter(file)
		defer gzipWriter.Close()
		writer = gzipWriter
	}

	tarWriter := tar.NewWriter(writer)
	defer tarWriter.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(tarWriter, file)
		return err
	})
}

func GetTodayIsoDate() string {
	return time.Now().Format(time.RFC3339)
}
