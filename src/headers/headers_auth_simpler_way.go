package headers
import(
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
)

func SimpleAuth() {
    req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil) //nil = no payload
    if err != nil {
        log.Fatalln("Err - unable to create request")
    }

    req.SetBasicAuth("user", "passw0rd")
    resp, err := http.DefaultClient.Do(req)

    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("Err - could not read content")
    }
    fmt.Println(string(content))
    fmt.Println(resp.StatusCode)
}
