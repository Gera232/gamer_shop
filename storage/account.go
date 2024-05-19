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
	defer rows.Close()

	accs := make(model.Accounts, 0)
	for rows.Next() {
		acc := &model.Account{}
		err := rows.Scan(
			&acc.ID,
			&acc.Name,
			&acc.Surname,
			&acc.Email,
			&acc.Password,
			&acc.Role,
			&acc.AddressID,
			&acc.CardID,
		)
		if err != nil {
			return model.Accounts{}, err
		}
		accs = append(accs, acc)
	}

	return accs, nil
}

func GetAccountByID(id uint32) (model.Account, error) {
	stmt, err := db.Prepare(sentenceGetAccountByID)
	if err != nil {
		return model.Account{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	acc, err := scanRow(row)
	if err != nil {
		return model.Account{}, err
	}

	return acc, nil
}

func GetAccountBySurname(surname string) (model.Account, error) {
	stmt, err := db.Prepare(sentenceGetAccountBySurname)
	if err != nil {
		return model.Account{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(surname)

	acc, err := scanRow(row)
	if err != nil {
		return model.Account{}, err
	}

	return acc, nil
}

func ExistAccountSurname(surname string) bool {
	acc, _ := GetAccountBySurname(surname)
	return acc.Surname == surname
}

func ExistAccountID(id uint32) bool {
	acc, _ := GetAccountByID(id)
	return acc.ID == id
}

func Logging(surname string, passwd string) (string, bool, error) {
	acc, err := GetAccountBySurname(surname)
	if err != nil {
		return "", false, err
	}

	DecodePass, err := security.Decode(acc.Password)
	if err != nil {
		return "", false, err
	}

	if DecodePass != passwd {
		return "", false, nil
	}

	return string(acc.Role), true, nil
}

func scanRow(row *sql.Row) (model.Account, error) {
	acc := model.Account{}
	err := row.Scan(
		&acc.ID,
		&acc.Name,
		&acc.Surname,
		&acc.Email,
		&acc.Password,
		&acc.Role,
		&acc.AddressID,
		&acc.CardID,
	)
	if err != nil {
		return model.Account{}, err
	}
	return acc, nil
}
