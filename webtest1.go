package main

import (
  "html/template"
  "net/http"
  "strings"
  "fmt"
  "log"
  "io/ioutil"
  "encoding/json"
  "net"
  "time"
  "errors"
  "bytes"
  "io"
  "github.com/qiniu/api/rs"
  . "github.com/qiniu/api/conf"
)


const (
  BUCKET = "avivaxrq"
  DOMAIN = "7u2o6v.com2.z0.glb.qiniucdn.com"
  AKEY = "3GjLFqbAzSuZ8cwwP05fBj-RQHaFymC8FWStn0vd"
  SKEY = "Htj0FvLqylb9QKA5KUxwdDf0wM3vNsEbncbva2EF"
)

func init() {
  ACCESS_KEY = AKEY
  SECRET_KEY = SKEY
}

type UploadRet struct {
  Key string `json:"key"`
}

type testSupport struct {
  Title       string   `json:"title"`
  Type        string   `json:"type"`
  Content     string   `json:"content"`
  Email       string   `json:"email"`
  Source      string   `json:"source"`
  TitlePrefix string   `json:"title_prefix"`
}

type Ret struct {
  Code int64 `json:"code"`
  Data struct {
    TicketId int`json:"ticket_id"`
  }`json:"data"`
}

var getInfoTmpl = template.Must(template.New("getinfo").ParseFiles("getinfo.html"))
var postSupportTmpl = template.Must(template.New("postsupport").ParseFiles("postsupport.html"))

func jsonCall(args interface{}, ret interface{})(err error) {

  tr := &http.Transport{
    ResponseHeaderTimeout: 30 * time.Second,
    DisableKeepAlives:     true,
    Dial: func(network, addr string) (net.Conn, error) {
      return net.DialTimeout(network, addr, 10*time.Second)
    },
  }

  var reader io.Reader

  data,err:=json.Marshal(args)
  if err!=nil {
    return
  }

  reader = bytes.NewReader(data)
  req,_:=http.NewRequest("POST", "https://portal.qiniu.com/support/new", reader)
  req.Header.Add("Content-Type", "application/json")

  client:=http.Client{
    Transport: tr,
  }

  resp,err:=client.Do(req)
  if err!=nil {
    return
  }

  if resp.StatusCode!=http.StatusOK {
    err = errors.New("error " + resp.Status)
    return
  }

  defer resp.Body.Close()

  buf:=bytes.NewBuffer(nil)
  n,err:=io.Copy(buf,resp.Body)
  body := buf.Bytes()
  if n==0 {
    body = []byte("null")
  }
  fmt.Println(string(body))

  err = json.Unmarshal(body, ret)

return
}


func postSupport(w http.ResponseWriter, r *http.Request) {
  data :=make(map[string]string)

  if r.Method !="GET" {
      r.ParseForm()
      if len(r.Form["email"][0]) == 0 {
        //html提示错误
        return
      }

      data["userinfo"] = r.Form["userinfo"][0]
      data["email"] = r.Form["email"][0]
      data["phone"] = r.Form["phone"][0]
      data["description"] = r.Form["description"][0]
      data["ip"] = r.Form["ip"][0]
      data["country"] = r.Form["country"][0]
      data["region"] = r.Form["region"][0]
      data["city"] = r.Form["city"][0]
      data["isp"] = r.Form["isp"][0]
      data["useragent"] = r.Form["useragent"][0]
      fmt.Println(r.Form["net"])

  }

  req := testSupport{
    Title: "徐榕谦工单测试" ,
    Type:"流量统计",
    Content: data["email"] + "\n" + data["phone"] + "\n" + data["description"] + "\n" + data["ip"] + "\n" + data["country"] + "\n" + data["region"] + "\n" + data["city"] + "\n" + data["isp"] + "\n" + data["useragent"] + "\n",
    Email:data["email"],
    Source:"Pili",
    TitlePrefix:"Pili Web Test",
  }

  ret:=Ret{}

  err := jsonCall(req, &ret)
  fmt.Println(err)
  fmt.Println(ret)

  postSupportTmpl.ExecuteTemplate(w, "postsupport", data)

}


func getInfo(w http.ResponseWriter, r *http.Request) {
    userInfo :=make(map[string]string)

    policy := rs.PutPolicy{
      Scope: BUCKET,
      EndUser: "userId",
      SaveKey: "$(sha1)",
    }
    token := policy.Token(nil)
    userInfo["token"] = token

    s := strings.Split(r.RemoteAddr, ":")
      s[0] = "58.20.0.17"
    userInfo["ip"] = s[0]

    resp,_ := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip="+s[0])
    defer resp.Body.Close()
    by,_ := ioutil.ReadAll(resp.Body)
    var ret map[string]interface{}
    json.Unmarshal(by,&ret)
    var data = ret["data"];
    dataMap, ok := data.(map[string]interface{})
    if(ok){
        var country = dataMap["country"]
        countryInfo, ok :=country.(string)
        if(ok){
          userInfo["country"] = countryInfo
        }
        var region = dataMap["region"]
        regionInfo, ok :=region.(string)
        if(ok){
          userInfo["region"] = regionInfo
        }
        var city = dataMap["city"]
        cityInfo, ok :=city.(string)
        if(ok){
          userInfo["city"] = cityInfo
        }
        var isp = dataMap["isp"]
        ispInfo, ok :=isp.(string)
        if(ok){
          userInfo["isp"] = ispInfo
        }
      }
    for _, v := range r.Header["User-Agent"]{
      userInfo["useragent"] = v
    }

  getInfoTmpl.ExecuteTemplate(w, "getinfo", userInfo)
}



func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  http.HandleFunc("/", getInfo) //设置访问的路由
  http.HandleFunc("/postsupport", postSupport)
  err := http.ListenAndServe(":9090", nil) //设置监听的端口
  if err != nil {
      log.Fatal("ListenAndServe: ", err)
  }

}






