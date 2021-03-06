package main

import (
	"context"
	"dimash/snippetbox/pkg/models/postgres"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *postgres.SnippetModel
}



func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn, err := pgxpool.Connect(context.Background(),"user=postgres password=dxrk host=172.17.0.1 port=5432 dbname=snippetbox sslmode=disable pool_max_conns=10")

	if err != nil {
		log.Fatal(err)
	}

	defer dsn.Close()

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &postgres.SnippetModel{Pool: dsn},
	}

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}