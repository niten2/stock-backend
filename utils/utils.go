package utils

// import (
//   "net/http"
// 	"io/ioutil"
//   "log"
//   "encoding/json"
// )

// func LoggerMiddleware(h http.Handler) http.Handler {
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     // formatRequest(r)

//     h.ServeHTTP(w, r)
//   })
// }

// func FormatRequest(r *http.Request) {
//   var data map[string] interface{}

//   body, _ := ioutil.ReadAll(r.Body)

//   json.Unmarshal(body, &data)

//   log.Println(data["query"])
//   log.Println(data["variables"])
// }
