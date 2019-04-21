package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}

func (c App) AddMaker() revel.Result {
	return c.Render()
}

func (c App) Search() revel.Result {
	return c.Render()
}
