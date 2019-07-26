package main
import ("time"
"AskTimeApp")
type AskImp struct {
}

func (imp *AskImp) Add(a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *AskImp) Sub(a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *AskImp) AskTime(Info *AskTimeApp.SayInfo, HiInfo *AskTimeApp.GreetInfo) (int32, error) {
	//按大小滚动日志
	LocalRollLog.Debug("from:", Info)
	//按天滚动日志
	LocalLogFile.Debug("from:", Info)
	//按天远程日志
	RemoteDayLog.Debug("from:", Info)
	HiInfo.Greeting = "hello from tarsgo server"
	HiInfo.FeedBackInfo.Name = "tarsgo server"
	HiInfo.FeedBackInfo.Age = 11
	HiInfo.FeedBackInfo.Address = "tencentcloud"
	HiInfo.FeedBackInfo.Phonenum = "100000"
	HiInfo.Time = time.Now().String()    //不知道为什么首字母小写全变大写
	return 0, nil
}