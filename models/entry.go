package models

import "log"

// Entry represents a database record
type Entry struct {
	ID int64 `json:"id"`
	URL string `json:"url"`
	Encoded string `json:"encoded"`
}

// Save creates a new  record
func (e Entry) Save() (int64, error) {
	stmt, err := db.Prepare("INSERT url SET url=?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(e.URL)
	if err != nil {
		log.Fatal(err)
	}

	return res.LastInsertId()
}

// SaveEncoded persist the hash
func (e Entry) SaveEncoded() (int64, error) {
	smtp, err := db.Prepare("update url set encoded=? where id=?")
	if err != nil {
		log.Fatal(err)	
	}

	res, err := smtp.Exec(e.Encoded, e.ID)
	if err != nil {
		log.Fatal(err)
	}

	return res.RowsAffected()
}

// Verify checks if an url has already been saved
func (e Entry) Verify() (Entry, error) {
	row := db.QueryRow("SELECT * FROM url WHERE url = ?", e.URL)

	err := row.Scan(&e.ID, &e.URL, &e.Encoded)
	if err != nil {
		return e, err
	}

	return e, nil
}

// FindByHash returns a single register or an error
func (e Entry) FindByHash(code string) (Entry, error) {
	row := db.QueryRow("SELECT * FROM url WHERE encoded = ?", code)

	err := row.Scan(&e.ID, &e.URL, &e.Encoded)
	if err != nil {
		return e, err
	}

	return e, nil
}