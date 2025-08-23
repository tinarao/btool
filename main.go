package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tinarao/btool/internal/config"
	"github.com/tinarao/btool/internal/tg"
)

func main() {
	config.Load()

	tg.Start(config.Cfg.BotToken)
}

func CreateBackup() {
	for _, path := range config.Cfg.Paths {
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("failed to get user home directory: %s\n", err.Error())
			return
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Printf("warning: directory %s does not exist, skipping\n", path)
			continue
		}

		timestamp := time.Now().Format(time.RFC3339)

		s := strings.Split(path, "/")
		lastEntry := s[len(s)-1]

		targetDir := filepath.Join(homedir, config.Cfg.TargetDir)
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			os.Mkdir(targetDir, 0755)
		}

		filename := fmt.Sprintf("%s-%s.tar.gz", lastEntry, timestamp)
		target := filepath.Join(targetDir, filename)

		log.Printf("archiving %s to %s\n", path, target)

		if err := TarDirectory(path, target); err != nil {
			log.Fatalf("failed to tar: %s\n", err.Error())
		}

		log.Printf("successfully archived: %s\n", target)
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
