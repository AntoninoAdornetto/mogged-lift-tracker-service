package db

import (
	"bufio"
	"context"
	"database/sql"
	"log"
	"os"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbSource = "root:secret@tcp(localhost:3306)/ismogged?parseTime=true"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	testQueries = New(testDB)

	//@todo - revist. Triggers won't run with sqlc. Using as temp solution
	file, err := os.Open("../trigger/inactive.sql")
	if err != nil {
		log.Fatal("cannot read trigger file")
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// skipping err so I don't have to check if trigger has been created already.
	// will revisit
	testDB.ExecContext(context.Background(), strings.Join(lines, " "))

	os.Exit(m.Run())
}
