package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luispfcanales/goweb/internal/core/services"
	"github.com/luispfcanales/goweb/internal/handlers/userhdl"
	"github.com/luispfcanales/goweb/internal/repository"
	"github.com/luispfcanales/goweb/pkg/gotemplate"
)

//Run initial web server
func Run(port string) error {
	gotpl, err := gotemplate.New("templates/*.gotpl")
	if err != nil {
		return fmt.Errorf("error to parse template %v", err)
	}
	storage := repository.New()
	usersrv := services.NewUserService(storage)
	userHandle := userhdl.New(usersrv, gotpl)

	http.HandleFunc("/", userHandle.Home)

	log.Println("listen on port :", port)
	return http.ListenAndServe(":"+port, nil)
}
