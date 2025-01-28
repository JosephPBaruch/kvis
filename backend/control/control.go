package control

import (
	"backend/types"
	"bytes"
	"encoding/base64"
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

func GetNodeDetails(nodeName string) (*types.NodeDetails, error) {
	details := &types.NodeDetails{}
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
		} else if strings.HasPrefix(line, "Roles:") {
			details.Roles = strings.TrimSpace(strings.TrimPrefix(line, "Roles:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "CreationTimestamp:") {
			details.CreationTimestamp = strings.TrimSpace(strings.TrimPrefix(line, "CreationTimestamp:"))
		} else if strings.HasPrefix(line, "Taints:") {
			details.Taints = strings.TrimSpace(strings.TrimPrefix(line, "Taints:"))
		} else if strings.HasPrefix(line, "Unschedulable:") {
			details.Unschedulable = strings.TrimSpace(strings.TrimPrefix(line, "Unschedulable:"))
		} else if strings.HasPrefix(line, "Addresses:") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "InternalIP:") {
			details.InternalIP = strings.TrimSpace(strings.TrimPrefix(line, "InternalIP:"))
		} else if strings.HasPrefix(line, "Hostname:") {
			details.Hostname = strings.TrimSpace(strings.TrimPrefix(line, "Hostname:"))
		} else if strings.HasPrefix(line, "Capacity:") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "cpu:") {
			details.CPU = strings.TrimSpace(strings.TrimPrefix(line, "cpu:"))
		} else if strings.HasPrefix(line, "ephemeral-storage:") {
			details.EphemeralStorage = strings.TrimSpace(strings.TrimPrefix(line, "ephemeral-storage:"))
		} else if strings.HasPrefix(line, "memory:") {
			details.Memory = strings.TrimSpace(strings.TrimPrefix(line, "memory:"))
		} else if strings.HasPrefix(line, "pods:") {
			details.Pods = strings.TrimSpace(strings.TrimPrefix(line, "pods:"))
		} else if strings.HasPrefix(line, "System Info:") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "Machine ID:") {
			details.MachineID = strings.TrimSpace(strings.TrimPrefix(line, "Machine ID:"))
		} else if strings.HasPrefix(line, "System UUID:") {
			details.SystemUUID = strings.TrimSpace(strings.TrimPrefix(line, "System UUID:"))
		} else if strings.HasPrefix(line, "Boot ID:") {
			details.BootID = strings.TrimSpace(strings.TrimPrefix(line, "Boot ID:"))
		} else if strings.HasPrefix(line, "Kernel Version:") {
			details.KernelVersion = strings.TrimSpace(strings.TrimPrefix(line, "Kernel Version:"))
		} else if strings.HasPrefix(line, "OS Image:") {
			details.OSImage = strings.TrimSpace(strings.TrimPrefix(line, "OS Image:"))
		} else if strings.HasPrefix(line, "Operating System:") {
			details.OperatingSystem = strings.TrimSpace(strings.TrimPrefix(line, "Operating System:"))
		} else if strings.HasPrefix(line, "Architecture:") {
			details.Architecture = strings.TrimSpace(strings.TrimPrefix(line, "Architecture:"))
		} else if strings.HasPrefix(line, "Container Runtime Version:") {
			details.ContainerRuntimeVersion = strings.TrimSpace(strings.TrimPrefix(line, "Container Runtime Version:"))
		} else if strings.HasPrefix(line, "Kubelet Version:") {
			details.KubeletVersion = strings.TrimSpace(strings.TrimPrefix(line, "Kubelet Version:"))
		} else if strings.HasPrefix(line, "Kube-Proxy Version:") {
			details.KubeProxyVersion = strings.TrimSpace(strings.TrimPrefix(line, "Kube-Proxy Version:"))
		} else if strings.HasPrefix(line, "PodCIDR:") {
			details.PodCIDR = strings.TrimSpace(strings.TrimPrefix(line, "PodCIDR:"))
		} else if strings.HasPrefix(line, "PodCIDRs:") {
			details.PodCIDRs = strings.TrimSpace(strings.TrimPrefix(line, "PodCIDRs:"))
		} else if strings.HasPrefix(line, "ProviderID:") {
			details.ProviderID = strings.TrimSpace(strings.TrimPrefix(line, "ProviderID:"))
		}
	}

	return details, nil
}

func GetDeploymentDetails(deploymentName string) (*types.DeploymentDetails, error) {
	details := &types.DeploymentDetails{}
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
		} else if strings.HasPrefix(line, "CreationTimestamp:") {
			details.CreationTimestamp = strings.TrimSpace(strings.TrimPrefix(line, "CreationTimestamp:"))
		} else if strings.HasPrefix(line, "Labels:") {
			details.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Selector:") {
			details.Selector = strings.TrimSpace(strings.TrimPrefix(line, "Selector:"))
		} else if strings.HasPrefix(line, "Replicas:") {
			details.Replicas = strings.TrimSpace(strings.TrimPrefix(line, "Replicas:"))
		} else if strings.HasPrefix(line, "StrategyType:") {
			details.StrategyType = strings.TrimSpace(strings.TrimPrefix(line, "StrategyType:"))
		} else if strings.HasPrefix(line, "MinReadySeconds:") {
			details.MinReadySeconds = strings.TrimSpace(strings.TrimPrefix(line, "MinReadySeconds:"))
		} else if strings.HasPrefix(line, "RollingUpdateStrategy:") {
			details.RollingUpdateStrategy = strings.TrimSpace(strings.TrimPrefix(line, "RollingUpdateStrategy:"))
		} else if strings.HasPrefix(line, "Pod Template:") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "Labels:") {
			details.PodTemplate.Labels = strings.TrimSpace(strings.TrimPrefix(line, "Labels:"))
		} else if strings.HasPrefix(line, "Containers:") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "Image:") {
			details.PodTemplate.Containers.Image = strings.TrimSpace(strings.TrimPrefix(line, "Image:"))
		} else if strings.HasPrefix(line, "Port:") {
			details.PodTemplate.Containers.Port = strings.TrimSpace(strings.TrimPrefix(line, "Port:"))
		} else if strings.HasPrefix(line, "Host Port:") {
			details.PodTemplate.Containers.HostPort = strings.TrimSpace(strings.TrimPrefix(line, "Host Port:"))
		} else if strings.HasPrefix(line, "Environment:") {
			details.PodTemplate.Containers.Environment = strings.TrimSpace(strings.TrimPrefix(line, "Environment:"))
		} else if strings.HasPrefix(line, "Mounts:") {
			details.PodTemplate.Containers.Mounts = strings.TrimSpace(strings.TrimPrefix(line, "Mounts:"))
		} else if strings.HasPrefix(line, "Volumes:") {
			details.PodTemplate.Volumes = strings.TrimSpace(strings.TrimPrefix(line, "Volumes:"))
		} else if strings.HasPrefix(line, "Conditions:") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "Type") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "Available") {
			details.Conditions.Available = strings.TrimSpace(strings.TrimPrefix(line, "Available"))
		} else if strings.HasPrefix(line, "Progressing") {
			details.Conditions.Progressing = strings.TrimSpace(strings.TrimPrefix(line, "Progressing"))
		} else if strings.HasPrefix(line, "OldReplicaSets:") {
			details.OldReplicaSets = strings.TrimSpace(strings.TrimPrefix(line, "OldReplicaSets:"))
		} else if strings.HasPrefix(line, "NewReplicaSet:") {
			details.NewReplicaSet = strings.TrimSpace(strings.TrimPrefix(line, "NewReplicaSet:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	return details, nil
}

func GetServiceDetails(serviceName string) (*types.ServiceDetails, error) {
	details := &types.ServiceDetails{}
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
		} else if strings.HasPrefix(line, "Selector:") {
			details.Selector = strings.TrimSpace(strings.TrimPrefix(line, "Selector:"))
		} else if strings.HasPrefix(line, "Type:") {
			details.Type = strings.TrimSpace(strings.TrimPrefix(line, "Type:"))
		} else if strings.HasPrefix(line, "IP Family Policy:") {
			details.IPFamilyPolicy = strings.TrimSpace(strings.TrimPrefix(line, "IP Family Policy:"))
		} else if strings.HasPrefix(line, "IP Families:") {
			details.IPFamilies = strings.TrimSpace(strings.TrimPrefix(line, "IP Families:"))
		} else if strings.HasPrefix(line, "IP:") {
			details.IP = strings.TrimSpace(strings.TrimPrefix(line, "IP:"))
		} else if strings.HasPrefix(line, "IPs:") {
			details.IPs = strings.TrimSpace(strings.TrimPrefix(line, "IPs:"))
		} else if strings.HasPrefix(line, "Port:") {
			details.Port = strings.TrimSpace(strings.TrimPrefix(line, "Port:"))
		} else if strings.HasPrefix(line, "TargetPort:") {
			details.TargetPort = strings.TrimSpace(strings.TrimPrefix(line, "TargetPort:"))
		} else if strings.HasPrefix(line, "Endpoints:") {
			details.Endpoints = strings.TrimSpace(strings.TrimPrefix(line, "Endpoints:"))
		} else if strings.HasPrefix(line, "Session Affinity:") {
			details.SessionAffinity = strings.TrimSpace(strings.TrimPrefix(line, "Session Affinity:"))
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

func GetConfigMapDetails(configMapName string) (*types.ConfigMapDetails, error) {
	details := &types.ConfigMapDetails{}
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
	var caCrtData []string
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
		} else if strings.HasPrefix(line, "Data") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "ca.crt:") {
			// Start collecting ca.crt data
			caCrtData = append(caCrtData, strings.TrimSpace(strings.TrimPrefix(line, "ca.crt:")))
		} else if len(caCrtData) > 0 {
			// Collect the rest of the ca.crt data
			caCrtData = append(caCrtData, line)
		} else if strings.HasPrefix(line, "BinaryData") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		}
	}

	// Encode the collected ca.crt data
	if len(caCrtData) > 0 {
		caCrt := strings.Join(caCrtData, "\n")
		encodedCACrt := base64.StdEncoding.EncodeToString([]byte(caCrt))
		details.CACert = encodedCACrt
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

func GetIngressDetails(ingressName string) (*types.IngressDetails, error) {
	details := &types.IngressDetails{}
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
		} else if strings.HasPrefix(line, "Address:") {
			details.Address = strings.TrimSpace(strings.TrimPrefix(line, "Address:"))
		} else if strings.HasPrefix(line, "Ingress Class:") {
			details.IngressClass = strings.TrimSpace(strings.TrimPrefix(line, "Ingress Class:"))
		} else if strings.HasPrefix(line, "Default backend:") {
			details.DefaultBackend = strings.TrimSpace(strings.TrimPrefix(line, "Default backend:"))
		} else if strings.HasPrefix(line, "Rules:") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "Host") {
			// Skip the header line
			continue
		} else if strings.HasPrefix(line, "Annotations:") {
			details.Annotations = strings.TrimSpace(strings.TrimPrefix(line, "Annotations:"))
		} else if strings.HasPrefix(line, "Events:") {
			details.Events = strings.TrimSpace(strings.TrimPrefix(line, "Events:"))
		} else if strings.HasPrefix(line, "*") {
			// Parse rules
			rule := types.IngressRule{}
			rule.Host = strings.TrimSpace(strings.TrimPrefix(line, "*"))
			details.Rules = append(details.Rules, rule)
		} else if strings.HasPrefix(line, "/") {
			// Parse path and backend
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				details.Rules[len(details.Rules)-1].Path = parts[0]
				details.Rules[len(details.Rules)-1].Backend = parts[2]
			}
		}
	}

	return details, nil
}
