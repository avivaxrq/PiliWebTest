package main

import (
  "html/template"
  "net/http"
  // "strings"
  "fmt"
  // "log"
  // "io/ioutil"
  // "encoding/json"
  // "github.com/qiniu/api/rs"
  // . "github.com/qiniu/api/conf"
)


// const (
//   BUCKET = "avivaxrq"
//   DOMAIN = "7u2o6v.com2.z0.glb.qiniucdn.com"
//   AKEY = "3GjLFqbAzSuZ8cwwP05fBj-RQHaFymC8FWStn0vd"
//   SKEY = "Htj0FvLqylb9QKA5KUxwdDf0wM3vNsEbncbva2EF"
// )

// func init() {
//   ACCESS_KEY = AKEY
//   SECRET_KEY = SKEY
// }

// type UploadRet struct {
//   Key string `json:"key"`
// }

// var getInfoTmpl = template.Must(template.New("getinfo").ParseFiles("getinfo.html"))


// func getInfo(w http.ResponseWriter, r *http.Request) {
//     userInfo :=make(map[string]string)

//     policy := rs.PutPolicy{
//       Scope: BUCKET,
//       EndUser: "userId",
//       SaveKey: "$(sha1)",
//     }
//     token := policy.Token(nil)
//     fmt.Println(token)
//     fmt.Printf("%T", token)
//     userInfo["token"] = token

//     s := strings.Split(r.RemoteAddr, ":")
//       s[0] = "58.20.0.17"
//     userInfo["ip"] = s[0]

//     resp,_ := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip="+s[0])
//     defer resp.Body.Close()
//     by,_ := ioutil.ReadAll(resp.Body)
//     var ret map[string]interface{}
//     json.Unmarshal(by,&ret)
//     var data = ret["data"];
//     dataMap, ok := data.(map[string]interface{})
//     if(ok){
//         var country = dataMap["country"]
//         countryInfo, ok :=country.(string)
//         if(ok){
//           userInfo["country"] = countryInfo
//         }
//         var region = dataMap["region"]
//         regionInfo, ok :=region.(string)
//         if(ok){
//           userInfo["region"] = regionInfo
//         }
//         var city = dataMap["city"]
//         cityInfo, ok :=city.(string)
//         if(ok){
//           userInfo["city"] = cityInfo
//         }
//         var isp = dataMap["isp"]
//         ispInfo, ok :=isp.(string)
//         if(ok){
//           userInfo["isp"] = ispInfo
//         }
//       }
//     for _, v := range r.Header["User-Agent"]{
//       userInfo["useragent"] = v
//     }

//     if r.Method !="GET" {
//         r.ParseForm()
//         email := r.Form["email"][0]
//         phone := r.Form["phone"][0]
//         description := r.Form["description"][0]
//         userInfo["email"] = email
//         userInfo["phone"] = phone
//         userInfo["description"] = description

//         // postsupportvalue := new {
//         //   "Title": {"PiliWebTest"},
//         //   "Type": {"Pili"},
//         //   "Content": {userInfo["ip"] + userInfo["country"] + userInfo["region"] + userInfo["isp"] + userInfo["useragent"] + userInfo["phone"] + userInfo["description"]},
//         //   "Email": {userInfo["email"]},
//         //   "Source": {"Pili"},
//         //   "Service":{"none"},
//         //   "Attachment": {["none"]},
//         //   "TitlePrefix": {"Pili"}
//         // }
//     // resp, err := http.PostForm("https://portal.qiniu.com/support/new",postsupportvalue)



//   }
//   getInfoTmpl.ExecuteTemplate(w, "getinfo", userInfo)
// }

var piliwebadminTmpl = template.Must(template.New("piliwebadmin").ParseFiles("piliwebadmin.html"))

func piliwebadmin(w http.ResponseWriter, r *http.Request) {
  // domain := "xxx.com"
  // url = ""
  // for i := 0; ; i++ {
  //   url = domain + string(i)
  // }

  if r.Method !="GET" {
    r.ParseForm()

  }

  piliwebadminTmpl.ExecuteTemplate(w, "piliwebadmin", piliwebgenerate)

}


func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  http.HandleFunc("/", piliwebadmin) //设置访问的路由
  err := http.ListenAndServe(":9000", nil) //设置监听的端口
  if err != nil {
      log.Fatal("ListenAndServe: ", err)
  }

}






