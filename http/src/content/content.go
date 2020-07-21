package content
import(
    "net/http"
    "io/ioutil"
    "encoding/json"
    "log"
    "fmt"
)

type GetResponse struct {
    Origin string               `json:"origin"`
    URL string                  `json:"url"`
    Headers map[string]string   `json:"headers"`
}

func (r *GetResponse) ToString() string {
    s := fmt.Sprintf("GET Response\nOrigin IP: %s\nRequest URL: %s", r.Origin, r.URL)
    for key, value := range r.Headers {
        s += fmt.Sprintf("Header: %s = %s\n", key, value)
    }
    return s
}

func ProcessContent() {
    resp, err := http.Get("https://httpbin.org/get"); if err != nil {log.Fatalln("unable to read response")}
    defer resp.Body.Close()

    content, err := ioutil.ReadAll(resp.Body); if err != nil {log.Fatalln("unable to read response")}

    respContent := &GetResponse{}
    json.Unmarshal(content, respContent)
    fmt.Println(respContent.ToString())
}
