package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/to4to/go-stock-api/model"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

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

func CreateStock(w http.ResponseWriter, r *http.Request) {

	var stock model.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable To Decode")
	}

	insertID := insertStock(stock)

	res := response{
		ID:      insertID,
		Message: "Stock Created Successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	stock, err := getStock(int64(id))
	if err != nil {

		log.Fatalf("Unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {

	stocks, err := getAllStock()
	if err != nil {
		log.Fatal("Unable To Get Stocks")
	}
	json.NewEncoder(w).Encode(stocks)

}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var stock model.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteStock(int64(id))

	// format the message string
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

///////////////////////////////////////
////////Handlers
//////////////////////////////////////

func getStock(id int64) (model.Stock, error) {
	db := createConnection()

	defer db.Close()

	var stock model.Stock
	sqlStmt := `SELECT * FROM stocks WHERE stockid=$1`

	row := db.QueryRow(sqlStmt, id)

	err := row.Scan(&stock.Id, &stock.Company, &stock.Name, &stock.Price)

	switch err {

	}

}

func insertStock(stock model.Stock) int64 {

	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int

}

func getAllStock() {
	db := createConnection()

}



func deleteStock(){}
