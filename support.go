package main

import (
	"net"
	"time"
	"errors"
	"bytes"
	"encoding/json"
	"io"
	"fmt"
	"net/http"
)

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

func main() {
  req:=testSupport{
    Title: "xrq test" ,
    Type:"流量统计",
    Content: "i am a test",
    Email:"xurongqian@qiniu.com",
    Source:"pili",
    TitlePrefix:"Pili Web Test",
  }

  ret:=Ret{}
  err := jsonCall(req, &ret)
  fmt.Println(err)
  fmt.Println(ret)
}

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