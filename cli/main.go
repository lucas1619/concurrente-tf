package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type hello struct {
	app.Compo
	Ammount     string
	Corporation string
	OnClick     func() app.UI
}
type Frame struct {
	Cmd    string   `json:"cmd"`
	Sender string   `json:"sender"`
	Data   []string `json:"data"`
}

func send(remote string, frame Frame) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		defer cn.Close()
		enc := json.NewEncoder(cn)
		enc.Encode(frame)
	}
}

func (h *hello) Render() app.UI {
	//form
	return app.Form().Body(
		app.Input().
			Type("number").
			Value(h.Corporation).
			Placeholder("Empresa licitadora").
			AutoFocus(true).
			OnChange(h.ValueTo(&h.Corporation)).
			Required(true),
		app.Input().
			Type("text").
			Value(h.Ammount).
			Placeholder("Monto de licitacion").
			AutoFocus(true).
			OnChange(h.ValueTo(&h.Ammount)).
			Required(true),
		app.Button().OnClick(h.OnClick()),
	)
}

func main() {
	app.Route("/", &hello{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
