interface DeploymentDetails {
  name: string;
  namespace: string;
  creationTimestamp: string;
  labels: string;
  annotations: string;
  selector: string;
  replicas: string;
  strategyType: string;
  minReadySeconds: string;
  rollingUpdateStrategy: string;
  podTemplate: {
    labels: string;
    containers: {
      image: string;
      port: string;
      hostPort: string;
      environment: string;
      mounts: string;
    };
    volumes: string;
  };
  conditions: {
    available: string;
    progressing: string;
  };
  oldReplicaSets: string;
  newReplicaSet: string;
  events: string;
}

interface PodDetails {
  name: string;
  namespace: string;
  priority: string;
  serviceAccount: string;
  node: string;
  startTime: string;
  labels: string;
  annotations: string;
  status: string;
  ip: string;
  controlledBy: string;
  containerID: string;
  image: string;
  imageID: string;
  port: string;
  hostPort: string;
  state: string;
  started: string;
  ready: string;
  restartCount: string;
  environment: string;
  mounts: string;
  conditions: string;
  volumes: string;
  qosClass: string;
  nodeSelectors: string;
  tolerations: string;
  events: string;
  logs: string;
}

interface ServiceDetails {
  name: string;
  namespace: string;
  labels: string;
  annotations: string;
  selector: string;
  type: string;
  ipFamilyPolicy: string;
  ipFamilies: string;
  ip: string;
  ips: string;
  port: string;
  targetPort: string;
  endpoints: string;
  sessionAffinity: string;
  events: string;
}

interface ConfigMapDetails {
  name: string;
  namespace: string;
  labels: string;
  annotations: string;
  caCert: string;
  events: string;
}

interface IngressRule {
  host: string;
  path: string;
  backend: string;
}

interface IngressDetails {
  name: string;
  namespace: string;
  address: string;
  ingressClass: string;
  defaultBackend: string;
  rules: IngressRule[];
  annotations: string;
  events: string;
}

interface PVCDetails {
  name: string;
  namespace: string;
  storageClass: string;
  status: string;
  volume: string;
  labels: string;
  annotations: string;
  finalizers: string;
  capacity: string;
  accessModes: string;
  volumeMode: string;
  usedBy: string;
  events: string;
}

interface EndpointDetails {
  name: string;
  namespace: string;
  labels: string;
  annotations: string;
  ip: string;
  port: string;
}

interface NodeDetails {
  name: string;
  roles: string;
  labels: string;
  annotations: string;
  creationTimestamp: string;
  taints: string;
  unschedulable: string;
  internalIP: string;
  hostname: string;
  cpu: string;
  ephemeralStorage: string;
  memory: string;
  pods: string;
  machineID: string;
  systemUUID: string;
  bootID: string;
  kernelVersion: string;
  osImage: string;
  operatingSystem: string;
  architecture: string;
  containerRuntimeVersion: string;
  kubeletVersion: string;
  kubeProxyVersion: string;
  podCIDR: string;
  podCIDRs: string;
  providerID: string;
}

interface NamespaceDetails {
  name: string;
  labels: string;
  annotations: string;
  status: string;
  resourceQuota: string;
  limitRange: string;
}

export interface KubePageProps {
  name?: string | undefined;
  typeOption: string | undefined;
}

export interface ListObject {
  name: string | undefined;
}

export type ReturnPromiseDetails = Promise<DeploymentDetails | PodDetails | ServiceDetails | ConfigMapDetails | IngressDetails | PVCDetails | EndpointDetails | NodeDetails | NamespaceDetails | null>;

export type ReturnPromiseList = Promise<ListObject[]>;

export type DetailInformationType = DeploymentDetails | PodDetails | ServiceDetails | ConfigMapDetails | IngressDetails | PVCDetails | EndpointDetails | NodeDetails | NamespaceDetails | null;

export type ListInfoType = ListObject[] | null;
