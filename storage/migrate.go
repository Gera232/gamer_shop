package storage

import (
	"io"
	"log"
	"os"
	"strings"
)

func Migrate(file *os.File) error {
	sentence, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	sentences := strings.Split(string(sentence), ";")

	for _, sentence := range sentences {
		if strings.TrimSpace(sentence) != "" {
			_, err := db.Exec(sentence)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}

func Migrate2(file *os.File) error {
	sentence, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	sentences := strings.Split(string(sentence), ";")

	for _, sentence := range sentences {
		if strings.TrimSpace(sentence) != "" {
			_, err := db.Exec(sentence)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}
