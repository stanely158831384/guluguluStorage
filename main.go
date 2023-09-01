package api

import (
	"database/sql"
	"log"

	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)



func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := NewServer2(store)

	err = server.Start(serverAddress)
}