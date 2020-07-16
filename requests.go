package main
import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        log.Fatalln("Unable to get response ")
    }
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(content))
}
