package storage

import "os"

func Migrate() error {
	sentence := os.Getenv("SENTENCE_MIGRATE_ACCOUNT")

	stmt, err := db.Prepare(sentence)
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
