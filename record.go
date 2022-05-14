package record

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
	Ips     map[string]string `json:"ips,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
		Ips:     make(map[string]string),
	}
}

// Record plugin.
type Record struct {
	next    http.Handler
	headers map[string]string
	ips     map[string]string
	name    string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Headers) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}
	return &Record{
		headers: config.Headers,
		ips:     config.Headers,
		next:    next,
		name:    name,
	}, nil
}

func (a *Record) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	log.Println("Record RemoteAddr", req.RemoteAddr)
	log.Println("Record ips", a.ips)

	exit := req.Header.Get("exit")
	if exit != "" {
		_, _ = rw.Write([]byte("你被自己要求退出"))
		log.Println("exit= ", exit)
		return
	}

	for key, value := range a.headers {
		log.Println("Record key-value= ", key, value)
		//把头加入
		req.Header.Set(key, value)
	}

	a.next.ServeHTTP(rw, req)
}
