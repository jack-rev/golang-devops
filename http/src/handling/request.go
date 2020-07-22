package main
import (
    "net/http"
    "os"
    "log"
    "io/ioutil"
    "fmt"
    "handling/process"
)

var routing = router.NewRouter()

func main() {
    resp, err := http.Get(os.Args[1]); if err != nil { log.Fatalln("Error") }
    routing.Process(resp)
}

//init() gets called before main()
//map status codes to corresponding anonymous functions
func init() {
    routing.Register(200, func(r *http.Response){
        defer r.Body.Close()
        content, err := ioutil.ReadAll(r.Body); if err != nil { log.Fatalln("error") }
        fmt.Println(string(content))
    })
    routing.Register(404, func(r *http.Response) {
        log.Fatalln("Not found (404): " , r.Request.URL.String())
    })
}
