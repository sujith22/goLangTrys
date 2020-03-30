package main

import (
	"log"

	"./routes"
	_ "github.com/lib/pq"
)

const (
	host     = "arjuna.db.elephantsql.com"
	port     = 5432
	user     = "ovptavph"
	password = "pYnPqqmXVve1ILKykcCL6nYeR3u01tYN"
	dbname   = "ovptavph"
)

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	log.Printf("*Starting server*")
	routes.RegisterRoutes()
}
