package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

var Params map[string]interface{}

type Configuration struct {
	Params map[string]interface{}
}

func (config Configuration) setConfigFile(configFile string) {}

func (config Configuration) get(key string) interface{} {
	return config.Params[key]
}

func (config Configuration) set(key string, value interface{}) {
	config.Params[key] = value
}

type RouteConfiguration struct {
	Routes map[string]interface{}
}

func (routeConfig RouteConfiguration) setRouteConfiguration(config interface{}) {
	routeConfig.Routes = config.(map[string]interface{})
}

func (routeConfig RouteConfiguration) getRouteConfiguration(key string) interface{} {
	return routeConfig.Routes[key]
}

func (routeConfig RouteConfiguration) setConfiguration(config Configuration, key string) {
	params := config.Params
	routeConfig.setRouteConfiguration(params)
}

type Route struct {
	Configuration RouteConfiguration
}

func (route Route) setConfig(config RouteConfiguration) {
	route.Configuration = config
}
func (route Route) getConfig() RouteConfiguration {
	return route.Configuration
}

func (route Route) create() {

}

type Front struct{}

type Controller interface {
	PreDispatch()
	Dispatch()
	PostDispatch()
}

type Index struct{}

func (index *Index) init() {

}

func (index *Index) indexAction() {

}

type Registry struct {
	Params map[string]interface{}
}

func (reg *Registry) set(key string, value interface{}) {
	reg.Params[key] = value
}

func (reg *Registry) get(key string) interface{} {
	return reg.Params[key]
}

type handler func(w http.ResponseWriter, r *http.Request)

type Application struct {
	Response http.ResponseWriter
	Request  *http.Request
	Handler  handler
	Routes   map[string]handler
}

func (app *Application) Run() {
	fmt.Println("testing run")
	http.ListenAndServe(":8001", nil)
}

func (app *Application) setResponseWriter(w http.ResponseWriter) {
	app.Response = w
}

func (app *Application) setRequest(r *http.Request) {
	app.Request = r
}

//func (app *Application) Handle(path string,data map[string]interface{}) (map[string]interface{}) {
//	data["inclusion"] = "myhandler"
//	data["inclusion2"] = "myhandler2"
//	return data
//}
type CommitDetails struct {
	Added interface{}	`json:"added"`
	Author interface{}	`json:"author"`

}

type GitWebHookPayload struct {
	EventName     string      `json:"event_name"`
	UserName      string      `json:"user_name"`
	UserEmail     string      `json:"user_email`
	RefSpec       string      `json:"ref"`
	OldRev        string      `json:"before"`
	NewRev        string      `json:"after"`
	TagAnnotation string      `json:"message"`
	TotalCommits  int      `json:"total_commits_count"`
	Commits       interface{} `json:"commits"`
	CheckoutSha   string      `json:"rev"`
	Project       interface{} `json:"project"`
}

func (app *Application) rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	decoder := json.NewDecoder(r.Body)
	m := GitWebHookPayload{}
	enc := json.NewEncoder(w)
	err := decoder.Decode(&m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.EventName)
	fmt.Println(m.UserName)
	fmt.Println(m.UserEmail)
	fmt.Println(m.RefSpec)
	fmt.Println(m.OldRev)
	fmt.Println(m.NewRev)
	fmt.Println(m.TagAnnotation)
	fmt.Println(m.TotalCommits)
	fmt.Println(m.Commits)
	fmt.Println(m.CheckoutSha)
	fmt.Println(m.Project)

	d := make(map[string]string)
	d["testing"] = "ok"
	if err := enc.Encode(d); nil != err {
		fmt.Fprintf(w, `{"error":"%s"}`, err)
	}
}

func (app *Application) InitRoutes() {
	//router := mux.NewRouter()
	//router.Path("/").Name("root").Handler(handler(app.rootHandle))
	app.Routes = make(map[string]handler)
	app.Routes["/"] = handler(app.rootHandle)
	app.Route("/")
}

func (app *Application) GetHandler(path string) {
	pathHandler := app.Routes[path]
	app.Handler = handler(pathHandler)
}

func (app *Application) Route(path string) {
	app.GetHandler(path)
	http.HandleFunc(
		path, handler(app.Handler),
	)
}

func main() {
	reg := Registry{Params: make(map[string]interface{})}
	app := Application{}
	reg.set("app", app)
	myapp := reg.get("app").(Application)
	data := make(map[string]interface{})
	data["testing"] = "sample"
	app.InitRoutes()
	fmt.Println(myapp)
	myapp.Run()
}
