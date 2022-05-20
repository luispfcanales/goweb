package userhdl

import (
	"net/http"

	"github.com/luispfcanales/goweb/internal/core/ports"
	"github.com/luispfcanales/goweb/pkg/gotemplate"
)

//HTTPHandlerUser is template to handlerUsers
type HTTPHandlerUser struct {
	userService ports.UserService
	gotpl       *gotemplate.GoTemplate
}

//New setup handlers
func New(userSrv ports.UserService, tpl *gotemplate.GoTemplate) *HTTPHandlerUser {
	return &HTTPHandlerUser{
		userService: userSrv,
		gotpl:       tpl,
	}
}

//Get return a User
func (hdl *HTTPHandlerUser) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("template"))
}

//Home render home template
func (hdl *HTTPHandlerUser) Home(w http.ResponseWriter, r *http.Request) {
	hdl.gotpl.Render(w, "home", nil)
}
