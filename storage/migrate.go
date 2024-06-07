package storage

import "os"

func Migrate() error {
	sentenceMigrate := os.Getenv("SENTENCE_MIGRATE")

	stmt, err := db.Prepare(sentenceMigrate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
