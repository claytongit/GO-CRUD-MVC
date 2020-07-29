package connection

import(
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Db() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/donation")

	if err != nil {

		log.Fatal(err.Error())

	} 

	erroPing := db.Ping()

	if erroPing != nil {

		log.Fatal(erroPing.Error())
	}

	return db
	
}