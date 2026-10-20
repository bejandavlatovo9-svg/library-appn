package main
import (
	"database/sql"
	"fmt"
	"log"
	
	_"github.com/lib/pq"
)

func main(){
	db, err := sql.Open("postgres", 
"host=localhost port = 5433 user = postgres password= BEZH0909 dbname = postgres sslmode=disable")
if err != nil{
	log.Fatal(err)
}
defer db.Close()
if err := db.Ping(); err != nil{
	log.Fatal(err)
}
var title string
err = db.QueryRow(
	"SELECT title FROM noteooks ORDER BY id LIMIT 1",
).Scan(&title)
if err != nil{
	log.Fatal(err)
}
fmt.Println("Book:", title)
}