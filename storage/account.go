package storage

import (
	"back-end/security"
	"back-end/types"
	"database/sql"
	"os"
)

func CreateAccount(acc *types.Account) error {
	sentence := os.Getenv("SENTENCE_CREATE_ACCOUNT")

	stmt, err := db.Prepare(sentence)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		acc.Name,
		acc.Surname,
		acc.Email,
		security.Encode(acc.Password),
		acc.Role,
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAccount(acc *types.Account) error {
	sentence := os.Getenv("SENTENCE_UPDATE_ACCOUNT")

	stmt, err := db.Prepare(sentence)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		acc.Name,
		acc.Surname,
		acc.Email,
		acc.Surname,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAccount(id uint32) error {
	sentence := os.Getenv("SENTENCE_DELETE_ACCOUNT")

	stmt, err := db.Prepare(sentence)
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

func GetAccounts() (types.Accounts, error) {
	sentence := os.Getenv("SENTENCE_GET_ACCOUNTS")

	stmt, err := db.Prepare(sentence)
	if err != nil {
		return types.Accounts{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return types.Accounts{}, err
	}
	defer rows.Close()

	accs := make(types.Accounts, 0)
	for rows.Next() {
		acc := &types.Account{}
		err := rows.Scan(
			&acc.ID,
			&acc.Name,
			&acc.Surname,
			&acc.Email,
			&acc.Password,
			&acc.Role,
		)
		if err != nil {
			return types.Accounts{}, err
		}
		accs = append(accs, acc)
	}

	return accs, nil
}

func GetAccountByID(id uint32) (types.Account, error) {
	sentence := os.Getenv("SENTENCE_GET_ACCOUNT_BY_ID")

	stmt, err := db.Prepare(sentence)
	if err != nil {
		return types.Account{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	acc, err := scanRowAcc(row)
	if err != nil {
		return types.Account{}, err
	}

	return acc, nil
}

func GetAccountBySurname(surname string) (types.Account, error) {
	sentence := os.Getenv("SENTENCE_GET_ACCOUNT_BY_SURNAME")

	stmt, err := db.Prepare(sentence)
	if err != nil {
		return types.Account{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(surname)

	acc, err := scanRowAcc(row)
	if err != nil {
		return types.Account{}, err
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

func Logging(surname string, passwd string) (string, uint32, bool, error) {
	acc, err := GetAccountBySurname(surname)
	if err != nil {
		return "", 0, false, nil
	}

	DecodePass, err := security.Decode(acc.Password)
	if err != nil {
		return "", 0, false, err
	}

	if DecodePass != passwd {
		return "", 0, false, nil
	}

	return string(acc.Role), acc.ID, true, nil
}

func scanRowAcc(row *sql.Row) (types.Account, error) {
	acc := types.Account{}
	err := row.Scan(
		&acc.ID,
		&acc.Name,
		&acc.Surname,
		&acc.Email,
		&acc.Password,
		&acc.Role,
	)
	if err != nil {
		return types.Account{}, err
	}
	return acc, nil
}
