package storage

import (
	model "back-end/model/account"
	"back-end/security"
	"database/sql"
)

var (
	sentenceCreateAccount       = "INSERT INTO account (name, surname, email, password, role) VALUES (?, ?, ?, ?, ?);"
	sentenceGetAccountByID      = "SELECT account_id, name, surname, email, password, role, COALESCE(address_id, 0) AS address_id, COALESCE(card_id, 0) AS card_id FROM account WHERE account_id = ?;"
	sentenceGetAccountBySurname = "SELECT account_id, name, surname, email, password, role, COALESCE(address_id, 0) AS address_id, COALESCE(card_id, 0) AS card_id FROM account WHERE surname = ?;"
)

func CreateAccount(m *model.Account) error {
	stmt, err := db.Prepare(sentenceCreateAccount)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Name,
		m.Surname,
		m.Email,
		security.Encode(m.Password),
		m.Role,
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAccount() {}

func DeleteAccount() {}

func GetAccounts() {}

func GetAccountByID(id uint32) (model.Account, error) {
	stmt, err := db.Prepare(sentenceGetAccountByID)
	if err != nil {
		return model.Account{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	account, err := scanRow(row)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func GetAccountBySurname(surname string) (model.Account, error) {
	stmt, err := db.Prepare(sentenceGetAccountBySurname)
	if err != nil {
		return model.Account{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(surname)

	account, err := scanRow(row)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func ExistAccount(surname string) bool {
	account, _ := GetAccountBySurname(surname)
	return account.Surname == surname
}

func Logging(surname string, password string) (string, bool, error) {
	account, err := GetAccountBySurname(surname)
	if err != nil {
		return "", false, err
	}

	DecodePass, err := security.Decode(account.Password)
	if err != nil {
		return "", false, err
	}

	if DecodePass != password {
		return "", false, nil
	}

	return string(account.Role), true, nil
}

func scanRow(row *sql.Row) (model.Account, error) {
	account := model.Account{}
	err := row.Scan(
		&account.ID,
		&account.Name,
		&account.Surname,
		&account.Email,
		&account.Password,
		&account.Role,
		&account.AddressID,
		&account.CardID,
	)
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}
