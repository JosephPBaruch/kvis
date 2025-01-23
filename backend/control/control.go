package control

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

type Pod struct {
	Name string `json:"name"`
}

func GetPods() ([]Pod, error) {
	cmd := exec.Command("kubectl", "get", "pods", "-o", "json")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	var result struct {
		Items []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
		} `json:"items"`
	}

	err = json.Unmarshal(out.Bytes(), &result)
	if err != nil {
		return nil, err
	}

	var pods []Pod
	for _, item := range result.Items {
		pods = append(pods, Pod{Name: item.Metadata.Name})
	}

	return pods, nil
}
