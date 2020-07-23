package nap

type RouterFunc func(client *http.Client, content interface{}) {

}

type CBRouter struct {
    Routers map[int]RouterFunc
    DefaultRouter RouterFunc
}
