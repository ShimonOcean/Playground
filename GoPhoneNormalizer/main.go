package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "shimon"
	password = "your-password"
	dbname   = "phone_normalize"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	// db, err := sql.Open("postgres", psqlInfo)
	// must(err)

	// err = createDB(db, dbname)
	// must(err)
	// db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(createPNTable(db))
	id, err := insertPhone(db, "1234567890")
	must(err)
	fmt.Println("id=", id)
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createPNTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)`
	_, err := db.Exec(statement)
	return err
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

// Normalizes phone numbers
// Iterate through string, catching all numerical characters and disposing of extra
func normalize(phone string) string {
	// bytes buffer is more efficient than string concatenation with +
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			// WriteRune: appends UTF-8 of input to buffer
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

func normalizeRegex(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}
