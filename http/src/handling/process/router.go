package router
import (
    "net/http"
    "log"
)

type HandlerFunc func(*http.Response)

type ResponseRouter struct {
    Handlers map[int]HandlerFunc
    DefaultHandler HandlerFunc
}

func NewRouter() *ResponseRouter {
    return &ResponseRouter {
        Handlers: make(map[int]HandlerFunc),
        DefaultHandler: func(resp *http.Response) {
            log.Fatalln("Unhandled Response: ", resp.StatusCode)
        },
    }
}

func (respRouter *ResponseRouter) Register(status int, handler HandlerFunc) {
    respRouter.Handlers[status] = handler
}

func (respRouter *ResponseRouter) Process(resp *http.Response) {
    returnedFunc, ok := respRouter.Handlers[resp.StatusCode]
    if !ok {
        respRouter.DefaultHandler(resp)
        return
    }
    returnedFunc(resp)
}
