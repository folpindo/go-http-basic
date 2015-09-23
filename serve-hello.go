package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func main() {
    http.HandleFunc(
        "/",
        func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type","application/json")
            w.WriteHeader(http.StatusOK)
            enc := json.NewEncoder(w)
            var data map[string]interface{}
            s := `{"status":"Ok"}`
            err := json.Unmarshal([]byte(s),&data)
            if err != nil {
              panic(err)
            }
            if err := enc.Encode(data); nil != err {
              fmt.Fprintf(w, `{"error":"%s"}`, err)
            }
            //fmt.Fprintln(w, '{"greeting":"Hello, www!"}')
        },

    )
    http.ListenAndServe(":8080", nil)
}
