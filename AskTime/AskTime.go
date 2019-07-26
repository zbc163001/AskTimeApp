package main

import (
	"github.com/TarsCloud/TarsGo/tars"
    "net/http"
	"AskTimeApp"
    //新增日志组件
    "github.com/TarsCloud/TarsGo/tars/util/rogger"
)

//本地日志
var LocalLogFile *rogger.Logger

//按大小滚动日志
var LocalRollLog *rogger.Logger

//按天远程日志
var RemoteDayLog *rogger.Logger

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Method, " ", r.RequestURI, " ", r.Header, " ", r.Body)
	w.Write([]byte("Hello Tars Http"))
}

func main() { //Init servant
	
	
	imp := new(AskImp)                                    //New Imp
	app := new(AskTimeApp.Ask)                                 //New init the A Tars
	cfg := tars.GetServerConfig()                               //Get Config File Object
	
	app.AddServant(imp, cfg.App+"."+cfg.Server+".AskObj") //Register Servant
	mux := &tars.TarsHttpMux{}
	mux.HandleFunc("/", HandleRequest)
	tars.AddHttpServant(mux, cfg.App+"."+cfg.Server+".HttpServerObj")
	LocalRollLog = tars.GetLogger("")                     //初始化日志组件
	LocalLogFile = tars.GetDayLogger("daylog", 1)
	RemoteDayLog = tars.GetRemoteLogger("daylog")
	tars.Run()
}
