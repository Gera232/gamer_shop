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
	sentenceDeleteAccount       = "DELETE FROM account WHERE account_id = ?;"
	sentenceUpdateAccount       = "UPDATE account SET name = ?, surname = ?, email = ? WHERE surname = ?;"
	sentenceGetAccounts         = "SELECT account_id, name, surname, email, password, role, COALESCE(address_id, 0) AS address_id, COALESCE(card_id, 0) AS card_id FROM account;"
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

func UpdateAccount(m *model.Account) error {
	stmt, err := db.Prepare(sentenceUpdateAccount)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Name,
		m.Surname,
		m.Email,
		m.Surname,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAccount(id uint32) error {
	stmt, err := db.Prepare(sentenceDeleteAccount)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func GetAccounts() (model.Accounts, error) {
	stmt, err := db.Prepare(sentenceGetAccounts)
	if err != nil {
		return model.Accounts{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return model.Accounts{}, err
	}

	accounts := make(model.Accounts, 0)
	for rows.Next() {
		account := &model.Account{}
		err := rows.Scan(
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
			return model.Accounts{}, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

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

func ExistAccountSurname(surname string) bool {
	account, _ := GetAccountBySurname(surname)
	return account.Surname == surname
}

func ExistAccountID(id uint32) bool {
	account, _ := GetAccountByID(id)
	return account.ID == id
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
