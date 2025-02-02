export interface Deployment {
  name: string | undefined;
}

export interface DeploymentDetailProps {
  name: string | undefined;
}

export interface DeploymentDetails {
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

