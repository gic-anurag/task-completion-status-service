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

// func (server *Server) UpdateHistory(journalId string, history *History) (*Response, error) {
// 	var result Response
// 	var buf bytes.Buffer
// 	historyReq := map[string]interface{}{
// 		"id":      journalId,
// 		"history": history,
// 	}
// 	err := json.NewEncoder(&buf).Encode(historyReq)

// 	if err != nil {
// 		return &result, err
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", server.Server+"/journal/update/journal_history", strings.NewReader(buf.String()))

// 	if err != nil {
// 		return &result, err
// 	}

// 	req.Header.Add("Content-Type", "application/json")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return &result, err
// 	}

// 	defer resp.Body.Close()

// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		return &result, err
// 	}

// 	return &result, nil
// }

// func (server Server) SearchFilter(searchFilter map[string]interface{}) (*Response, error) {
// 	var result Response
// 	// payload := fmt.Sprint(searchFilter)
// 	var buf bytes.Buffer

// 	err := json.NewEncoder(&buf).Encode(searchFilter)

// 	if err != nil {
// 		return &result, err
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", server.Server+"/journal/filter/journal_entry", strings.NewReader(buf.String()))
// 	if err != nil {
// 		return &result, err
// 	}
// 	req.Header.Add("Content-Type", "application/json")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return &result, err
// 	}

// 	defer resp.Body.Close()

// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		return &result, err
// 	}
// 	return &result, nil
// }

// func (server Server) SearchByJournalId(journalId *FilterById) (*Response, error) {
// 	var result Response
// 	var buf bytes.Buffer

// 	err := json.NewEncoder(&buf).Encode(journalId)

// 	if err != nil {
// 		return &result, err
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", server.Server+"/journal/find/journal_entry_by_id", strings.NewReader(buf.String()))

// 	if err != nil {
// 		return &result, err
// 	}
// 	req.Header.Add("Content-Type", "application/json")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return &result, err
// 	}

// 	defer resp.Body.Close()

// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		return &result, err
// 	}
// 	return &result, nil
// }
