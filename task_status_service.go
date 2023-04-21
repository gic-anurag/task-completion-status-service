package taskCompletionStatus

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type Server struct {
	Server string
}

func NewClient(server string) *Server {
	return &Server{
		Server: server,
	}
}

type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func (server Server) AddTaskCompletionStatus(taskStatusReq *TaskCompletionStatus) (*Response, error) {
	var result Response
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(taskStatusReq)

	if err != nil {
		return &result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", server.Server+"/api/add/task-status", strings.NewReader(buf.String()))

	if err != nil {
		return &result, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return &result, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

func (server Server) FilterTaskCompletionStatus(taskStatusReq *TaskCompletionStatus) (*Response, error) {
	var result Response
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(taskStatusReq)

	if err != nil {
		return &result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", server.Server+"/api/filter/task-status", strings.NewReader(buf.String()))
	if err != nil {
		return &result, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return &result, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

func (server Server) GetTaskCompletionStatusById(taskStatusId *GetTaskCompletionStatus) (*Response, error) {
	var result Response
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(taskStatusId)

	if err != nil {
		return &result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", server.Server+"/api/get-by-id/task-status", strings.NewReader(buf.String()))

	if err != nil {
		return &result, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return &result, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}
