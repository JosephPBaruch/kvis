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

export interface KubePageProps {
  name?: string | undefined;
  typeOption: string | undefined;
}

export interface ListObject {
  name: string | undefined;
}

export type ReturnPromiseDetails = Promise<DeploymentDetails | null>;

export type ReturnPromiseList = Promise<ListObject[]>;

export type DetailInformationType = DeploymentDetails | null;

export type ListInfoType = ListObject[] | null;
