package main

import (
	"github.com/gocraft/web"
	"github.com/mlctrez/gowebsocket/chat"
	"golang.org/x/net/websocket"
	"net/http"
)

type Context struct {
}

func main() {
	server := chat.NewServer()
	go server.Listen()

	router := web.New(Context{})

	router.Middleware(web.StaticMiddleware("webroot", web.StaticOption{IndexFile:"index.html"}))

	router.Get("/entry", func(rw web.ResponseWriter, req *web.Request) {
		websocket.Handler(server.OnConnected).ServeHTTP(rw, req.Request)
	})

	http.ListenAndServe("localhost:8080", router)

}
