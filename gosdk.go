package main

import (
  "log"
  "fmt"
  "net/http"
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

var uploadFormFmt = `
 <html>
  <body>
   <form method="post" action="http://up.qiniu.com/" enctype="multipart/form-data">
    <input name="token" type="hidden" value="%s">
    Album belonged to: <input name="x:album" value="albumId"><br>
    Image to upload: <input name="file" type="file"/><br>
    <input type="submit" value="Upload">
   </form>
  </body>
 </html>
 `
//
type UploadRet struct {
  Key string `json:"key"`
}

func handleUpload(w http.ResponseWriter, req *http.Request) {
  policy := rs.PutPolicy{
    Scope: BUCKET,
    EndUser: "userId",
    SaveKey: "$(sha1)",
  }
  token := policy.Token(nil)
  log.Println("token:", token)
  uploadForm := fmt.Sprintf(uploadFormFmt, token)
  w.Write([]byte(uploadForm))
}

func main() {
  http.HandleFunc("/upload", handleUpload)
  log.Fatal(http.ListenAndServe(":8765", nil))
}

// --------------------------------------------------------------------------------

