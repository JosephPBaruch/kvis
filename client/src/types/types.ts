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

export interface KubePageProps {
  name?: string | undefined;
  typeOption: string | undefined;
}

export interface ListObject {
  name: string | undefined;
}

export type ReturnPromiseDetails = Promise<DeploymentDetails | PodDetails | null>;

export type ReturnPromiseList = Promise<ListObject[]>;

export type DetailInformationType = DeploymentDetails | PodDetails | null;

export type ListInfoType = ListObject[] | null;
