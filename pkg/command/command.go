package command

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type CommandType int

const (
	Exec CommandType = iota
	Http             = iota
)

type HttpMethod string

const (
	Get  HttpMethod = "get"
	Post            = "post"
)

type Command struct {
	Type CommandType
	Arg  string
}

type ExecCommandArg struct {
	Command string `json:"command"`
}

type HttpCommandArg struct {
	Path        string     `json:"path"`
	Method      HttpMethod `json:"method"`
	ContentType string     `json:"contentType,omitempty"`
	Body        string     `json:"body,omitempty"`
}

func (c *Command) Execute() {
	switch c.Type {
	case Exec:
		ExecuteExec(c)
	case Http:
		ExecuteHttp(c)
	}
}

func ExecuteExec(c *Command) {
	var arg ExecCommandArg
	err := json.Unmarshal([]byte(c.Arg), &arg)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(arg.Command)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func ExecuteHttp(c *Command) {
	var arg HttpCommandArg
	err := json.Unmarshal([]byte(c.Arg), &arg)
	if err != nil {
		log.Fatal(err)
	}
	switch arg.Method {
	case Get:
		resp, err := http.Get(arg.Path)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
	case Post:
		body := strings.NewReader(arg.Body)
		resp, err := http.Post(arg.Path, arg.ContentType, body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
	}
}
