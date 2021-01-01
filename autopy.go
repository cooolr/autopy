package autopy

import "fmt"
import "time"
import "strconv"
import "net/url"
import "net/http"
import "io/ioutil"
import "encoding/json"

var Server_url = "http://127.0.0.1:8020/?code="

type FWdata struct {
    Id string `json:"id"`
    Type string `json:"type"`
    Color []int `json:"color"`
    Text string `json:"text"`
    Width string `json:"Width"`
    Height string `json:"Height"`
    Size string `json:"size"`
    X string `json:"X"`
    Y string `json:"Y"`
    Operation string `json:"operation"`
}

func log(level,message string) {
    local,_ := time.LoadLocation("Asia/Shanghai")
    now := time.Now().In(local).Format("2006-01-02 15:04:05")
    _message := now+" ["+level+"] "+message
    fmt.Println(_message)
}

func urlopen(uri,action string) {
    _,err := http.Get(Server_url+uri)
    if err != nil{
	log("ERROR",action+"执行失败\n"+err.Error())
	return
    }
    log("INFO",action+"执行成功")
}

func urlget(uri,action string) (string,error){
    resp,err := http.Get(Server_url+uri)
    if err != nil{
	log("ERROR",action+"执行失败\n"+err.Error())
	return "",err
    }
    defer resp.Body.Close()
    body,_ := ioutil.ReadAll(resp.Body)
    log("INFO",action+"执行成功")
    return string(body),nil
}

func str(number int) string {
    return strconv.Itoa(number)
}

func sleep(num float64) {
    time.Sleep(time.Duration(num) * time.Second)
}

func Sleep(num float64) {
    time.Sleep(time.Duration(num) * time.Second)
    local,_ := time.LoadLocation("Asia/Shanghai")
    now := time.Now().In(local).Format("2006-01-02 15:04:05")
    message := now+" [INFO] Sleep:"
    fmt.Println(message,num)
}

func Click(x,y int) {
    urlopen("0,"+str(x)+","+str(y), "点击")
    sleep(0.1)
}

func Swipe(x1,y1,x2,y2,t int) {
    urlopen("1,"+str(x1)+","+str(y1)+","+str(x2)+","+str(y2)+","+str(t), "滑动")
    sleep(0.0023*float64(t))
}

func HOME() {
    urlopen("HOME,", "主页键")
}

func BACK() {
    urlopen("BACK,", "返回键")
}

func MENU() {
    urlopen("RECENTS,", "菜单键")
}

func Capturer() {
    urlopen("2,", "截图")
}

func GetView() (string,error){
    body,err := urlget("getView,", "获取视图")
    return body,err
}

func ClickById(text string) {
    urlopen("getID,"+text, "点击ID")
}

func ClickByText(text string) {
    urlopen("getText,"+url.QueryEscape(text), "点击文字")
}

func MakeToast(text string) {
    data,_ := json.Marshal(map[string]string{"type":"Toast", "text":text})
    urlopen(url.QueryEscape(string(data)), "makeToast")
}

func SetFloatWindow(data FWdata) {
    urlopen(string(json.Marshal(data)))
}

