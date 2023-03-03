package keychain

import (
	"encoding/csv"
	"io"
)

func readLastpassCSV(in io.Reader, f func(e entry) error) error {
	r := csv.NewReader(in)

	_, err := r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		err = f(entry{
			URL:      record[0],
			Username: record[1],
			Password: record[2],
			TOTP:     record[3],
			Extra:    record[4],
			Name:     record[5],
			Grouping: record[6],
			Fav:      record[7],
		})
		if err != nil {
			return err
		}
	}
}
