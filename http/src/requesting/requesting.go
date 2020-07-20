package requesting
import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "strings"
)

func make_req() {
    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        log.Fatalln("Unable to get response ")
    }
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(content))
}

func query_req() {
    resp, err := http.Get("https://httpbin.org/get?search=some_search")
    if err != nil {
        log.Fatalln("Unable to get response")
    }
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(content))
}

func make_post() {
    resp, err := http.Post("https://httpbin.org/post", "text/plain", strings.NewReader("this is the request"))
    if err != nil {
        log.Fatalln("Unable to get response")
    }
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(content))
}

func Client() {
    client := http.DefaultClient
    req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln("Unable to get response")
    }
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)

    fmt.Println(string(content))
}
