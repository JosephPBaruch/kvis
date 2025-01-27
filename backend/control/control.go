package control

import (
	"backend/types"
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"
)

func Get(option string) ([]types.Resource, error) {
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

	var resources []types.Resource
	for _, item := range result.Items {
		resources = append(resources, types.Resource{Name: item.Metadata.Name})
	}

	return resources, nil
}

func GetPodDetails(podName string) (*types.ResourceDetails, error) {
	details := &types.ResourceDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "pod", podName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Namespace:") {
			details.Namespace = strings.TrimSpace(strings.TrimPrefix(line, "Namespace:"))
		} else if strings.HasPrefix(line, "Priority:") {
			details.Priority = strings.TrimSpace(strings.TrimPrefix(line, "Priority:"))
		} else if strings.HasPrefix(line, "Service Account:") {
			details.ServiceAccount = strings.TrimSpace(strings.TrimPrefix(line, "Service Account:"))
		} else if strings.HasPrefix(line, "Node:") {
			details.Node = strings.TrimSpace(strings.TrimPrefix(line, "Node:"))
		} else if strings.HasPrefix(line, "Start Time:") {
			details.StartTime = strings.TrimSpace(strings.TrimPrefix(line, "Start Time:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Status:") {
			details.Status = strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
		} else if strings.HasPrefix(line, "IP:") {
			details.IP = strings.TrimSpace(strings.TrimPrefix(line, "IP:"))
		} else if strings.HasPrefix(line, "Controlled By:") {
			details.ControlledBy = strings.TrimSpace(strings.TrimPrefix(line, "Controlled By:"))
		} else if strings.HasPrefix(line, "Container ID:") {
			details.ContainerID = strings.TrimSpace(strings.TrimPrefix(line, "Container ID:"))
		} else if strings.HasPrefix(line, "Image:") {
			details.Image = strings.TrimSpace(strings.TrimPrefix(line, "Image:"))
		} else if strings.HasPrefix(line, "Image ID:") {
			details.ImageID = strings.TrimSpace(strings.TrimPrefix(line, "Image ID:"))
		} else if strings.HasPrefix(line, "Port:") {
			details.Port = strings.TrimSpace(strings.TrimPrefix(line, "Port:"))
		} else if strings.HasPrefix(line, "Host Port:") {
			details.HostPort = strings.TrimSpace(strings.TrimPrefix(line, "Host Port:"))
		} else if strings.HasPrefix(line, "State:") {
			details.State = strings.TrimSpace(strings.TrimPrefix(line, "State:"))
		} else if strings.HasPrefix(line, "Started:") {
			details.Started = strings.TrimSpace(strings.TrimPrefix(line, "Started:"))
		} else if strings.HasPrefix(line, "Ready:") {
			details.Ready = strings.TrimSpace(strings.TrimPrefix(line, "Ready:"))
		} else if strings.HasPrefix(line, "Restart Count:") {
			details.RestartCount = strings.TrimSpace(strings.TrimPrefix(line, "Restart Count:"))
		} else if strings.HasPrefix(line, "Environment:") {
			details.Environment = strings.TrimSpace(strings.TrimPrefix(line, "Environment:"))
		} else if strings.HasPrefix(line, "Mounts:") {
			details.Mounts = strings.TrimSpace(strings.TrimPrefix(line, "Mounts:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			details.Conditions = strings.TrimSpace(strings.TrimPrefix(line, "Conditions:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "QoS Class:") {
			details.QoSClass = strings.TrimSpace(strings.TrimPrefix(line, "QoS Class:"))
		} else if strings.HasPrefix(line, "Node-Selectors:") {
			details.NodeSelectors = strings.TrimSpace(strings.TrimPrefix(line, "Node-Selectors:"))
		} else if strings.HasPrefix(line, "Tolerations:") {
			details.Tolerations = strings.TrimSpace(strings.TrimPrefix(line, "Tolerations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	// Get logs
	logsCmd := exec.Command("kubectl", "logs", podName)
	var logsOut bytes.Buffer
	logsCmd.Stdout = &logsOut
	err = logsCmd.Run()
	if err != nil {
		return nil, err
	}
	details.Logs = logsOut.String()

	return details, nil
}

func GetNodeDetails(nodeName string) (*types.ResourceDetails, error) {
	details := &types.ResourceDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "node", nodeName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Namespace:") {
			details.Namespace = strings.TrimSpace(strings.TrimPrefix(line, "Namespace:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Status:") {
			details.Status = strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
		} else if strings.HasPrefix(line, "IP:") {
			details.IP = strings.TrimSpace(strings.TrimPrefix(line, "IP:"))
		} else if strings.HasPrefix(line, "Controlled By:") {
			details.ControlledBy = strings.TrimSpace(strings.TrimPrefix(line, "Controlled By:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			details.Conditions = strings.TrimSpace(strings.TrimPrefix(line, "Conditions:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "QoS Class:") {
			details.QoSClass = strings.TrimSpace(strings.TrimPrefix(line, "QoS Class:"))
		} else if strings.HasPrefix(line, "Node-Selectors:") {
			details.NodeSelectors = strings.TrimSpace(strings.TrimPrefix(line, "Node-Selectors:"))
		} else if strings.HasPrefix(line, "Tolerations:") {
			details.Tolerations = strings.TrimSpace(strings.TrimPrefix(line, "Tolerations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	return details, nil
}

func GetDeploymentDetails(deploymentName string) (*types.ResourceDetails, error) {
	details := &types.ResourceDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "deployment", deploymentName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Namespace:") {
			details.Namespace = strings.TrimSpace(strings.TrimPrefix(line, "Namespace:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Status:") {
			details.Status = strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
		} else if strings.HasPrefix(line, "Controlled By:") {
			details.ControlledBy = strings.TrimSpace(strings.TrimPrefix(line, "Controlled By:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			details.Conditions = strings.TrimSpace(strings.TrimPrefix(line, "Conditions:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "QoS Class:") {
			details.QoSClass = strings.TrimSpace(strings.TrimPrefix(line, "QoS Class:"))
		} else if strings.HasPrefix(line, "Node-Selectors:") {
			details.NodeSelectors = strings.TrimSpace(strings.TrimPrefix(line, "Node-Selectors:"))
		} else if strings.HasPrefix(line, "Tolerations:") {
			details.Tolerations = strings.TrimSpace(strings.TrimPrefix(line, "Tolerations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	return details, nil
}

func GetServiceDetails(serviceName string) (*types.ResourceDetails, error) {
	details := &types.ResourceDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "service", serviceName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Namespace:") {
			details.Namespace = strings.TrimSpace(strings.TrimPrefix(line, "Namespace:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Status:") {
			details.Status = strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
		} else if strings.HasPrefix(line, "Controlled By:") {
			details.ControlledBy = strings.TrimSpace(strings.TrimPrefix(line, "Controlled By:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			details.Conditions = strings.TrimSpace(strings.TrimPrefix(line, "Conditions:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "QoS Class:") {
			details.QoSClass = strings.TrimSpace(strings.TrimPrefix(line, "QoS Class:"))
		} else if strings.HasPrefix(line, "Node-Selectors:") {
			details.NodeSelectors = strings.TrimSpace(strings.TrimPrefix(line, "Node-Selectors:"))
		} else if strings.HasPrefix(line, "Tolerations:") {
			details.Tolerations = strings.TrimSpace(strings.TrimPrefix(line, "Tolerations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	return details, nil
}

func GetNamespaceDetails(namespaceName string) (*types.ResourceDetails, error) {
	details := &types.ResourceDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "namespace", namespaceName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Status:") {
			details.Status = strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
		} else if strings.HasPrefix(line, "Controlled By:") {
			details.ControlledBy = strings.TrimSpace(strings.TrimPrefix(line, "Controlled By:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			details.Conditions = strings.TrimSpace(strings.TrimPrefix(line, "Conditions:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "QoS Class:") {
			details.QoSClass = strings.TrimSpace(strings.TrimPrefix(line, "QoS Class:"))
		} else if strings.HasPrefix(line, "Node-Selectors:") {
			details.NodeSelectors = strings.TrimSpace(strings.TrimPrefix(line, "Node-Selectors:"))
		} else if strings.HasPrefix(line, "Tolerations:") {
			details.Tolerations = strings.TrimSpace(strings.TrimPrefix(line, "Tolerations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	return details, nil
}

func GetConfigMapDetails(configMapName string) (*types.ResourceDetails, error) {
	details := &types.ResourceDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "configmap", configMapName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Namespace:") {
			details.Namespace = strings.TrimSpace(strings.TrimPrefix(line, "Namespace:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Status:") {
			details.Status = strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
		} else if strings.HasPrefix(line, "Controlled By:") {
			details.ControlledBy = strings.TrimSpace(strings.TrimPrefix(line, "Controlled By:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			details.Conditions = strings.TrimSpace(strings.TrimPrefix(line, "Conditions:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "QoS Class:") {
			details.QoSClass = strings.TrimSpace(strings.TrimPrefix(line, "QoS Class:"))
		} else if strings.HasPrefix(line, "Node-Selectors:") {
			details.NodeSelectors = strings.TrimSpace(strings.TrimPrefix(line, "Node-Selectors:"))
		} else if strings.HasPrefix(line, "Tolerations:") {
			details.Tolerations = strings.TrimSpace(strings.TrimPrefix(line, "Tolerations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	return details, nil
}

func GetPVCDetails(pvcName string) (*types.ResourceDetails, error) {
	details := &types.ResourceDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "pvc", pvcName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Namespace:") {
			details.Namespace = strings.TrimSpace(strings.TrimPrefix(line, "Namespace:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Status:") {
			details.Status = strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
		} else if strings.HasPrefix(line, "Controlled By:") {
			details.ControlledBy = strings.TrimSpace(strings.TrimPrefix(line, "Controlled By:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			details.Conditions = strings.TrimSpace(strings.TrimPrefix(line, "Conditions:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "QoS Class:") {
			details.QoSClass = strings.TrimSpace(strings.TrimPrefix(line, "QoS Class:"))
		} else if strings.HasPrefix(line, "Node-Selectors:") {
			details.NodeSelectors = strings.TrimSpace(strings.TrimPrefix(line, "Node-Selectors:"))
		} else if strings.HasPrefix(line, "Tolerations:") {
			details.Tolerations = strings.TrimSpace(strings.TrimPrefix(line, "Tolerations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	return details, nil
}

func GetEndpointDetails(endpointName string) (*types.EndpointDetails, error) {
	details := &types.EndpointDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "endpoints", endpointName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Namespace:") {
			details.Namespace = strings.TrimSpace(strings.TrimPrefix(line, "Namespace:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Addresses:") {
			details.IP = strings.TrimSpace(strings.TrimPrefix(line, "Addresses:"))
		} else if strings.HasPrefix(line, "Ports:") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "https") {
			details.Port = strings.TrimSpace(strings.Fields(line)[1])
		}
	}

	return details, nil
}

func GetIngressDetails(ingressName string) (*types.ResourceDetails, error) {
	details := &types.ResourceDetails{}
	// Get describe
	describeCmd := exec.Command("kubectl", "describe", "ingress", ingressName)
	var describeOut bytes.Buffer
	describeCmd.Stdout = &describeOut
	err := describeCmd.Run()
	if err != nil {
		return nil, err
	}
	describeOutput := describeOut.String()

	// Parse describe output
	lines := strings.Split(describeOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Name:") {
			details.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		} else if strings.HasPrefix(line, "Namespace:") {
			details.Namespace = strings.TrimSpace(strings.TrimPrefix(line, "Namespace:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Status:") {
			details.Status = strings.TrimSpace(strings.TrimPrefix(line, "Status:"))
		} else if strings.HasPrefix(line, "Controlled By:") {
			details.ControlledBy = strings.TrimSpace(strings.TrimPrefix(line, "Controlled By:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			details.Conditions = strings.TrimSpace(strings.TrimPrefix(line, "Conditions:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "QoS Class:") {
			details.QoSClass = strings.TrimSpace(strings.TrimPrefix(line, "QoS Class:"))
		} else if strings.HasPrefix(line, "Node-Selectors:") {
			details.NodeSelectors = strings.TrimSpace(strings.TrimPrefix(line, "Node-Selectors:"))
		} else if strings.HasPrefix(line, "Tolerations:") {
			details.Tolerations = strings.TrimSpace(strings.TrimPrefix(line, "Tolerations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	return details, nil
}
