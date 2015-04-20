package scut

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ConfigServer struct {
	config *Config
}

func NewConfigServer(config *Config) (*ConfigServer, error) {
	if config == nil {
		return nil, fmt.Errorf("config is nil")
	}

	server := ConfigServer{config: config}

	return &server, nil
}

func (server *ConfigServer) Listen(address string) error {
	return http.ListenAndServe(address, server)
}

func (server ConfigServer) ServeHTTP(
	writer http.ResponseWriter, request *http.Request,
) {
	path := strings.Split(strings.Trim(request.URL.Path, "/"), "/")

	rawBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	var body interface{}
	if len(rawBody) != 0 {
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
	}

	switch request.Method {
	case "GET":
		server.handleGET(writer, path...)
	case "PATCH":
		server.handlePATCH(writer, body, path...)
	case "PUT":
		server.handlePUT(writer, body, path...)
	default:
		writer.WriteHeader(400)
	}
}

func (server *ConfigServer) handleGET(writer http.ResponseWriter, path ...string) {
	var data interface{}

	if len(path) == 1 && path[0] == "" {
		data = server.config.GetRoot()
	} else {
		data = server.config.Get(path...)
		if data == nil {
			writer.WriteHeader(404)
			return
		}
	}

	jsonedData, _ := json.MarshalIndent(data, "", "  ")
	writer.Write(jsonedData)
}

func (server *ConfigServer) handlePATCH(
	writer http.ResponseWriter, value interface{}, path ...string,
) {
	server.config.Set(value, path...)
}

func (server *ConfigServer) handlePUT(
	writer http.ResponseWriter, value interface{}, path ...string,
) {
	switch root := value.(type) {
	case map[string]interface{}:
		server.config.SetRoot(root)
	default:
		writer.WriteHeader(400)
	}
}
