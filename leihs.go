package leihs

import (
	"net/http"
	"regexp"
)

// Leihs ....
type Leihs struct {
	token  string
	url    string
	client http.Client
}

// Config ...
type Config struct {
	Token    string
	LeihsURL string
}

// NewLeihs ...
func NewLeihs(c *Config) *Leihs {
	l := &Leihs{
		token: c.Token,
		url:   c.LeihsURL,
	}
	return l
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
