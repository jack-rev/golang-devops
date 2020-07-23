package nap

type AuthToken struct {
    Token string
}

type AuthBasic struct {
    Username string
    Password string
}

interface Authentication {
    AuthorisationHeader() string // "basic <base64-encoded string>"
}
