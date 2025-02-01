package main

import (
	"fmt"
	"github.com/user/urlProject/internal/config"
)

func main() {

	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init config: cleanenv
	// TODO: init logger: slog
	// TODO: init storage: sqlite
	// TODO: init router: chi, "chi render"
	// TODO: run server
}
