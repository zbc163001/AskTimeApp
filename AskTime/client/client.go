package main

import (
	"fmt"
    "net/http"
	"github.com/TarsCloud/TarsGo/tars"
	"time"
    "io/ioutil"
	"AskTimeApp"
)



func main(){ 
	comm := tars.NewCommunicator()
	//obj := "AskTimeApp.AskTimeServer.AskObj@tcp -h 100.67.68.253 -p 33336 -t 60000"
	//本机测试
	obj := "AskTimeApp.AskTimeServer.AskObj@tcp -h 127.0.0.1 -p 33336 -t 60000"
	app := new(AskTimeApp.Ask)
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
	//	body,err := ioutil.ReadAll(retHttp.Body)
	//	fmt.Println(retHttp)
		fmt.Println("ret: ", ret, "info: ", info, "|hiinfo:", hiinfo)
		time.Sleep(2000 * time.Millisecond)
	}
}


