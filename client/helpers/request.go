package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RPCMethod struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

func DoRequest(url string, r *RPCMethod) (interface{}, error) {
	j, err := json.Marshal(r.Params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json-rpc", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
