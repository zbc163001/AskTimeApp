package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"fmt"
	"net/http"
	"AskTimeApp"
	"io/ioutil"
	"time"
)

func main() { //Init servant
	go rpcCall();
	 imp := new(AskServerImp)                                    //New Imp
	 app := new(AskTimeApp.AskServer)                                 //New init the A Tars
	 cfg := tars.GetServerConfig()                               //Get Config File Object
	 
	app.AddServant(imp, cfg.App+"."+cfg.Server+".AskServerObj") //Register Servant
	tars.Run()

}
func rpcCall(){
	comm := tars.NewCommunicator()
	//直接访问服务器
	obj := "AskTimeApp.AskTimeServer.AskObj@tcp -h 100.67.68.253 -p 33336 -t 60000"
	//通过主控路由访问
	//obj := "Test.HelloGo.SayHelloObj"
	app := new(AskTimeApp.Ask)
	//如果这里的代码是在某个服务端里，默认是不需要配置    就比如此处的AskTimeClient 服务
	//comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 9.75.176.108 -p 17890:tcp -h 9.75.176.102 -p 17890")
	//comm.SetProperty("modulename", "goclient")
	comm.StringToProxy(obj, app)
	for {
		var ret int32
		var err error
                //对服务端发起tars协议的rpc调用
		//入参定义
		info := new(AskTimeApp.SayInfo)
		info.Name = "tcb"
		info.Age = 1
		info.Address = "tencent"
		info.Phonenum = "1000000"
		//出参定义
		hiinfo := new(AskTimeApp.GreetInfo)
                //对服务发起tars协议的调用
		ret, err = app.AskTime(info, hiinfo)
		if err != nil {
			fmt.Println(err)
			return
		}
		client:= &http.Client{}
		request, _ := http.NewRequest("GET", "http://100.67.68.253:35536/", nil)
		response,_ := client.Do(request)
		if response.StatusCode == 200 {
	 	str, _ := ioutil.ReadAll(response.Body)
	 	bodystr := string(str)
		 fmt.Println(bodystr)
		}
		defer response.Body.Close()
		fmt.Println("ret: ", ret, "info: ", info, "|hiinfo:", hiinfo)
		time.Sleep(2000 * time.Millisecond)
	}
}
