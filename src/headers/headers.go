package headers
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "bytes"
    "encoding/base64"
)

//Simple authentication using headers in Go
func Auth() {
    req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil) //nil = no payload
    if err != nil {
        log.Fatalln("Err - unable to create request")
    }

    buffer := &bytes.Buffer{}
    enc := base64.NewEncoder(base64.URLEncoding, buffer)
    enc.Write([]byte("user:passw0rd"))
    enc.Close()

    encodedCreds, err := buffer.ReadString('\n')
    if err != nil && err.Error() != "EOF" {
        log.Fatalln("Error - failed to read encoded buffer")
    }
    req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCreds))
    resp, err := http.DefaultClient.Do(req)

    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("Err - could not read content")
    }
    fmt.Println(string(content))
    fmt.Println(resp.StatusCode)

}
