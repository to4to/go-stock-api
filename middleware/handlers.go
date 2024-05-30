package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/to4to/go-stock-api/model"
)

func createConnection() *sql.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Env Not Loaded")

	}
	db, err := sql.Open("postgres", os.ExpandEnv("POSTGRES_URL"))

	if err != nil {
		log.Fatal("Unable To Connect To DB", err)
	}

	//Check Connection

	err = db.Ping()
	if err != nil {
		log.Fatal("Connection Unstable", err)

	}

	fmt.Printf("Successfully Connected")

	return db

}

func GetStock() {

}

///////////////////////////////////////
////////Handlers
//////////////////////////////////////

func getStock(id int) (model.Stock, error) {
	db := createConnection()

	defer db.Close()

	var stock model.Stock
	sqlStmt:=`SELECT * FROM stocks WHERE stockid=$1`

	row:=db.QueryRow(sqlStmt,id)

	err:=row.Scan(&stock.Id,&stock.Company,&stock.Name,&stock.Price)

	switch err{


		
	}



}
