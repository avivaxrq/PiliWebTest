package main

import (
  "html/template"
  "net/http"
  "fmt"
  "strings"
  "log"
  "io/ioutil"
  "encoding/json"
)


// func hogeHandler(w http.ResponseWriter, r *http.Request) {
//   a := make(map[string]string)
//   a["a"] = "abc";
//   a["b"] = "b";
//   hogeTmpl.ExecuteTemplate(w, "base", a)
// }

// // var piyoTmpl = template.Must(template.New("piyo").ParseFiles("base.html", "piyo.html"))

// func piyoHandler(w http.ResponseWriter, r *http.Request) {
//   piyoTmpl.ExecuteTemplate(w, "base", "Piyo")
// }

var getInfoTmpl = template.Must(template.New("getinfo").ParseFiles("getinfo.html"))
// var loginTmpl = template.Must(template.New("login").ParseFiles("login.gtpl"))


func getInfo(w http.ResponseWriter, r *http.Request) {
    userInfo :=make(map[string]string)
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

            // r.ParseForm()
        // fmt.Printf("%T",r.Form["email"])
        //         fmt.Printf("%T",r.Form["email"][0])

        // email := r.Form["email"][0]
        // userInfo["email"] = email
    getInfoTmpl.ExecuteTemplate(w, "getinfo", userInfo)
}


    // var data = ret["data"];

    // dataMap, ok := data.(map[string]interface{})

    // if(ok){


    // if r.Method == "GET" {
    //     t, _ := template.ParseFiles("login.gtpl")
    //     t.Execute(w, nil)
    // } else {
        //请求的是登陆数据，那么执行登陆的逻辑判断
    // r.ParseForm()
      // emailmap := r.Form["email"]
//
      // emailstring, ok := email.(string)
      // if (ok){
        // userInfo["email"] = emailstring
      // }
    //   var
    //   emailInfo, ok := email.([]string)
    //   if(ok) {
    // // userInfo["email"] = r.Form["email"]
    //     emailInfomation := emailInfo[0]
    //   }
    // userInfo["phone"] = r.Form["phone"]
    // fmt.Println(w, email)
    // fmt.Println(w, phone)
    // fmt.Fprintf(w, "%T", email[0])

        // r.ParseForm()
        // // fmt.Printf("%T",r.Form["email"])
        // //         fmt.Printf("%T",r.Form["email"][0])

        // email := r.Form["email"][0]
        // userInfo["email"] = email
        // emailstring,ok := email.(string)
        // if(ok){
        //   userInfo["email"] = emailstring
        // }
        // phone := r.Form["phone"]
//         fmt.Fprintf(w, "%T", email[0])


func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        //请求的是登陆数据，那么执行登陆的逻辑判断
        r.ParseForm()
        email := r.Form["email"]
        phone := r.Form["phone"]
        fmt.Fprintln(w, email)
        fmt.Fprintln(w, phone)
        fmt.Fprintln(w, email)
        fmt.Fprintf(w, "%T", email)
        fmt.Fprintf(w, "%T", email[0])
    }
}



func main() {
  // http.HandleFunc("/", hogeHandler)
  // http.HandleFunc("/hoge", hogeHandler)
  // http.ListenAndServe(":8080", nil)

  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  http.HandleFunc("/", getInfo) //设置访问的路由
  http.HandleFunc("/login", login) //设置访问的路由
  err := http.ListenAndServe(":9090", nil) //设置监听的端口
  if err != nil {
      log.Fatal("ListenAndServe: ", err)
  }

}






