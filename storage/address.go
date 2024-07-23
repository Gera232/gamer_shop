package storage

import (
	"back-end/types"
	"os"
)

func CreateAddress(addrs *types.Address) error {
	sentence := os.Getenv("SENTENCE_CREATE_ADDRESS")

	stmt, err := db.Prepare(sentence)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		addrs.Place,
		addrs.Street,
		addrs.Height,
		addrs.Floor,
		addrs.Department,
		addrs.BetweenStreets,
		addrs.Observations,
		addrs.Shipment,
		addrs.Location_ID,
		addrs.Account_ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAddress(id uint32) error {
	sentence := os.Getenv("SENTENCE_DELETE_ADDRESS")

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

func GetAddresses(id uint32) (*types.Addresses, error) {
	sentence := os.Getenv("SENTENCE_GET_ADDRESSES")

	stmt, err := db.Prepare(sentence)
	if err != nil {
		return &types.Addresses{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return &types.Addresses{}, err
	}
	defer rows.Close()

	addrses := make(types.Addresses, 0)
	for rows.Next() {
		addrs := &types.Address{}
		err := rows.Scan(
			&addrs.ID,
			&addrs.Place,
			&addrs.Street,
			&addrs.Height,
			&addrs.Floor,
			&addrs.Department,
			&addrs.BetweenStreets,
			&addrs.Observations,
			&addrs.Shipment,
			&addrs.Location_ID,
			&addrs.Account_ID,
		)
		if err != nil {
			return &types.Addresses{}, err
		}
		addrses = append(addrses, addrs)
	}

	return &addrses, nil
}
