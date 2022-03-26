package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type application struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	db         *sqlx.DB
	matchMaker matchMaker
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	db, err := sqlx.Open("mysql", "root:ferrarii4@/augments")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		db:       db,
	}

	// test
	level := 1
	host := lobbyUser{id: 99, staffSlot: 1, canceled: false}
	lobby := lobby{level: level, host: host}
	app.matchMaker.lobbies = append(app.matchMaker.lobbies, lobby)

	// app.matchManager.matches = make(map[uint64]match)

	// app.matchFinder.addUnmatched(2, 2)
	// app.matchFinder.addUnmatched(3, 3)
	// app.matchFinder.addUnmatched(4, 4)
	// app.matchFinder.addUnmatched(5, 5)
	// app.matchFinder.addUnmatched(6, 6)

	srv := &http.Server{
		Addr:     ":8080",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", srv.Addr)
	errorLog.Fatal(srv.ListenAndServe())
}
