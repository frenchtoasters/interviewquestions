package main

import (
	"encoding/json"

	"k8s.io/klog/v2"
)

var (
	payload = `{"name":"bob", "last":"ross"}`
)

type dataPayload struct {
	Person string `json:"name"`
	Full   string `json:last,omitempty`
}

func loadData(p string) (dataPayload, error) {
	var obj dataPayload
	err := json.Unmarshal([]byte(p), &obj)
	if err != nil {
		return dataPayload{}, err
	}
	return obj, nil
}

func parseData(s dataPayload) ([]byte, error) {
	u, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return u, nil
}

func main() {
	start, err := loadData(payload)
	if err != nil {
		klog.Errorf("error loading data")
	}
	klog.Infof("Start: %s", start)

	data, err := parseData(start)
	if err != nil {
		klog.Errorf("error parsing payload")
	}

	klog.Infof("Data: %s", data)
	return
}
