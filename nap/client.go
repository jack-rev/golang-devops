package nap

type Client struct {
    Client *http.Client
    AuthInfo Authentication
}
