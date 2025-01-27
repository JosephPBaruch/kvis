package control

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"
)

type Resource struct {
	Name string `json:"name"`
}

func Get(option string) ([]Resource, error) {
	cmd := exec.Command("kubectl", "get", option, "--all-namespaces", "-o", "json")
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

	var resources []Resource
	for _, item := range result.Items {
		resources = append(resources, Resource{Name: item.Metadata.Name})
	}

	return resources, nil
}

type ResourceDetails struct {
	Describe string `json:"describe"`
	Logs     string `json:"logs"`
	Events   string `json:"events"`
}

func GetResourceDetails(resourceType, resourceName string) (*ResourceDetails, error) {
	details := &ResourceDetails{}

	// Get describe
	describeCmd := exec.Command("kubectl", "describe", resourceType, resourceName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	details.Describe = strings.ReplaceAll(strings.ReplaceAll(describeOut.String(), "\t", ""), "\n", "")

	// Get logs (only for pods)
	if resourceType == "pod" {
		logsCmd := exec.Command("kubectl", "logs", resourceName)
		var logsOut bytes.Buffer
		logsCmd.Stdout = &logsOut
		err := logsCmd.Run()
		if err != nil {
			return nil, err
		}
		details.Logs = strings.ReplaceAll(strings.ReplaceAll(logsOut.String(), "\t", ""), "\n", "")
	}

	// Get events
	eventsCmd := exec.Command("kubectl", "get", "events", "--field-selector", "involvedObject.name="+resourceName, "-o", "json")
	var eventsOut bytes.Buffer
	eventsCmd.Stdout = &eventsOut
	err = eventsCmd.Run()
	if err != nil {
		return nil, err
	}

	var eventsResult struct {
		Items []struct {
			Type    string `json:"type"`
			Reason  string `json:"reason"`
			Message string `json:"message"`
			Source  struct {
				Component string `json:"component"`
			} `json:"source"`
			FirstTimestamp string `json:"firstTimestamp"`
			LastTimestamp  string `json:"lastTimestamp"`
		} `json:"items"`
	}

	err = json.Unmarshal(eventsOut.Bytes(), &eventsResult)
	if err != nil {
		return nil, err
	}

	eventsFormatted := ""
	for _, event := range eventsResult.Items {
		eventsFormatted += strings.ReplaceAll(strings.ReplaceAll(event.Type+" "+event.Reason+" "+event.FirstTimestamp+" "+event.LastTimestamp+" "+event.Source.Component+" "+event.Message, "\t", ""), "\n", "") + "\n"
	}
	details.Events = eventsFormatted

	return details, nil
}
