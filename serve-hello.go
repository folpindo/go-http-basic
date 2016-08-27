package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Front struct {}

type Index struct {}

func (index *Index) init() {

}

func (index *Index) indexAction(){

}

type Registry struct {
	Params map[string]interface{}
}

func (reg *Registry) set(key string, value interface{}) {
	reg.Params[key] = value
}

func (reg *Registry) get(key string) (interface{}) {
	return reg.Params[key]
}

type Application struct {
	Response http.ResponseWriter
	Request *http.Request
}

func (app *Application) Run () {
	fmt.Println("testing run")
	http.ListenAndServe(":9091", nil)
}

func (app *Application) setResponseWriter(w http.ResponseWriter) {
	app.Response = w
}

func (app *Application) setRequest(r *http.Request) {
	app.Request = r
}

func (app *Application) Handle(path string,data map[string]interface{}) (map[string]interface{}) {
	data["inclusion"] = "myhandler"
	data["inclusion2"] = "myhandler2"
	return data
}

func (app *Application) Route (path string, data map[string]interface{}) {
	http.HandleFunc(
		path,
		func(w http.ResponseWriter, r *http.Request) {
			app.setResponseWriter(w)
			app.setRequest(r)
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusOK)
			enc := json.NewEncoder(w)
			d := app.Handle(path,data)
			if err := enc.Encode(d); nil != err {
				fmt.Fprintf(w, `{"error":"%s"}`, err)
			}
		},
	)
}

func main() {
	reg := Registry{Params:make(map[string]interface{})}
	app := Application{}
	reg.set("app",app)
	myapp := reg.get("app").(Application)
	data := make(map[string]interface{})
	data["testing"] = "sample"
	myapp.Route("/myroute",data)
	fmt.Println(myapp)
	myapp.Run()
}
