package types

type Resource struct {
	Name string `json:"name"`
}

type ResourceDetails struct {
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	Priority       string `json:"priority"`
	ServiceAccount string `json:"serviceAccount"`
	Node           string `json:"node"`
	StartTime      string `json:"startTime"`
	Labels         string `json:"labels"`
	Annotations    string `json:"annotations"`
	Status         string `json:"status"`
	IP             string `json:"ip"`
	ControlledBy   string `json:"controlledBy"`
	ContainerID    string `json:"containerID"`
	Image          string `json:"image"`
	ImageID        string `json:"imageID"`
	Port           string `json:"port"`
	HostPort       string `json:"hostPort"`
	State          string `json:"state"`
	Started        string `json:"started"`
	Ready          string `json:"ready"`
	RestartCount   string `json:"restartCount"`
	Environment    string `json:"environment"`
	Mounts         string `json:"mounts"`
	Conditions     string `json:"conditions"`
	Volumes        string `json:"volumes"`
	QoSClass       string `json:"qosClass"`
	NodeSelectors  string `json:"nodeSelectors"`
	Tolerations    string `json:"tolerations"`
	Events         string `json:"events"`
	Logs           string `json:"logs"`
}

type EndpointDetails struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Labels      string `json:"labels"`
	Annotations string `json:"annotations"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
}
