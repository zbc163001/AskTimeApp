package main
import (
"AskTimeApp"
// "github.com/TarsCloud/TarsGo/tars"
// "fmt"
// "net/http"
// "time"
)

type AskServerImp struct {
}

func (imp *AskServerImp) Adde(a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *AskServerImp) Sube(a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *AskServerImp) AskTimee(Info *AskTimeApp.SayInfo, HiInfo *AskTimeApp.GreetInfo) (int32, error) {
	// comm := tars.NewCommunicator()
	// obj := "AskTimeApp.AskTime.AskObj@tcp -h 127.0.0.1 -p 10015 -t 60000"
	// app := new(AskTimeApp.Ask)
	// comm.StringToProxy(obj, app)
	// for {
	// 	var ret int32
	// 	var err error
    //             //对服务端发起tars协议的rpc调用
	// 	//入参定义
	// 	info := new(AskTimeApp.SayInfo)
	// 	info.Name = "tcb"
	// 	info.Age = 1
	// 	info.Address = "tencent"
	// 	info.Phonenum = "1000000"
	// 	//出参定义
	// 	hiinfo := new(AskTimeApp.GreetInfo)
    //             //对服务发起tars协议的调用
	// 	ret, err = app.AskTime(info, hiinfo)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	retHttp,errHttp := http.Get("http://127.0.0.1/")
	// 	if errHttp != nil {
	// 		fmt.Println(err)
	// 	}
	// 	//defer retHttp.Body.Close()
	// 	//body,err := ioutil.ReadAll(retHttp.Body)
	// 	fmt.Println(retHttp)
	// 	fmt.Println("ret: ", ret, "info: ", info, "|hiinfo:", hiinfo)
	// 	time.Sleep(2000 * time.Millisecond)
	// }
	//AskTimeFunc();
	return 0, nil
}
