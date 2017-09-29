package controller

import "github.com/teambition/gear"

// ControllerHome router path "/" controller
type ControllerHome struct{}

// Index get "/"
func (c *ControllerHome) Index(ctx *gear.Context) (err error) {
	name := "home"

	return ctx.Render(200, name, map[string]string{
		"a": "hello wrold",
	})
}
