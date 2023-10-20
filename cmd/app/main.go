package main

import (
	"github.com/Oleg-OMON/gin-rest-api.git/internal/routers"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/store"

	_ "github.com/lib/pq"
)

func main() {

	s := new(store.Store)
	s.Open()
	routers.InitRouters(s)

}
