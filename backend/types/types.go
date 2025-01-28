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

type NodeDetails struct {
	Name                    string `json:"name"`
	Roles                   string `json:"roles"`
	Labels                  string `json:"labels"`
	Annotations             string `json:"annotations"`
	CreationTimestamp       string `json:"creationTimestamp"`
	Taints                  string `json:"taints"`
	Unschedulable           string `json:"unschedulable"`
	InternalIP              string `json:"internalIP"`
	Hostname                string `json:"hostname"`
	CPU                     string `json:"cpu"`
	EphemeralStorage        string `json:"ephemeralStorage"`
	Memory                  string `json:"memory"`
	Pods                    string `json:"pods"`
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OSImage                 string `json:"osImage"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	PodCIDR                 string `json:"podCIDR"`
	PodCIDRs                string `json:"podCIDRs"`
	ProviderID              string `json:"providerID"`
}

type ServiceDetails struct {
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	Labels          string `json:"labels"`
	Annotations     string `json:"annotations"`
	Selector        string `json:"selector"`
	Type            string `json:"type"`
	IPFamilyPolicy  string `json:"ipFamilyPolicy"`
	IPFamilies      string `json:"ipFamilies"`
	IP              string `json:"ip"`
	IPs             string `json:"ips"`
	Port            string `json:"port"`
	TargetPort      string `json:"targetPort"`
	Endpoints       string `json:"endpoints"`
	SessionAffinity string `json:"sessionAffinity"`
	Events          string `json:"events"`
}
