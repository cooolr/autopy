package autopy

import "fmt"
import "time"
import "strconv"
import "net/http"

var server_url = "http://127.0.0.1:8020/?code="

func log(level,message string) {
    local,_ := time.LoadLocation("Asia/Shanghai")
    now := time.Now().In(local).Format("2006-01-02 15:04:05")
    _message := now+" ["+level+"] "+message
    fmt.Println(_message)
}

func urlopen(url,action string) {
    _,err := http.Get(url)
    if err != nil{
	log("ERROR",action+"执行失败\n"+err.Error())
	return
    }
    log("INFO",action+"执行成功")
}

func str(number int) string {
    return strconv.Itoa(number)
}

func Sleep(num float64) {
    time.Sleep(time.Duration(num) * time.Second)
}

func Click(x,y int) {
    urlopen(server_url+"0,"+str(x)+","+str(y),"点击")
    Sleep(0.1)
}

func Swipe(x1,y1,x2,y2,t int) {
    urlopen(server_url+"1,"+str(x1)+","+str(y1)+","+str(x2)+
    ","+str(y2)+","+str(t),"滑动")
    Sleep(0.0023*float64(t))
}

func HOME() {
    urlopen(server_url + "HOME,","主页键")
}

func BACK() {
    urlopen(server_url + "BACK,","返回键")
}

func RECENTS() {
    urlopen(server_url + "RECENTS,","菜单键")
}

func Capturer() {
    urlopen(server_url + "2,","截图")
}

func ClickByText(text string) {
    urlopen(server_url+"getText,"+text,"点击文字")
}
