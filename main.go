package main

import (
	"github.com/shiki-tak/YamaWikipedia/infrastructure"
)

func main() {
	infrastructure.Router.Start(":1313")
}
