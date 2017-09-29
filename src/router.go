package src

import (
	"github.com/fuzi-goTeam/go-image-website/src/controller"
	"github.com/teambition/gear"
)

// NewRouter create new index Router
func NewRouter() (router *gear.Router) {
	router = gear.NewRouter()
	home := controller.ControllerHome{}
	router.Get("/", home.Index)
	return
}
