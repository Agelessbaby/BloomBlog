package main

import (
	favorite "github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite/favoritesrv"
	"log"
)

func main() {
	svr := favorite.NewServer(new(FavoriteSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
