package main

import (
	"blitsms/lib/howdah"
	"database/sql"
	"fmt"
	eventsource "github.com/antage/eventsource/http"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	_ "github.com/lib/pq"
	//"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	db, _ := sql.Open("postgres", "postgres://troy@localhost/blit?sslmode=disable")
	db.SetMaxIdleConns(15)
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
	}))
	m.Get("/api", func(r render.Render) {
		users := howdah.Users(db)
		r.HTML(200, "get-strands", users)
	})
	es := eventsource.New(nil, nil)
	defer es.Close()
	http.Handle("/events", es)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.URL.Path)
	})
	http.Handle("/", m)
	go func() {
		id := 1
		for {
			es.SendMessage("tick", "tick-event", strconv.Itoa(id))
			id++
			time.Sleep(2 * time.Second)
		}
	}()

	//http.ListenAndServe("127.0.0.1:3000", nil)
	m.Run()

}
