import { Deployment, DeploymentDetails } from '../types/types';

const baseURL = 'http://localhost:8090';

const deploymentsList = async (): Promise<Deployment[]> => {
  const response = await fetch(`${baseURL}/deployments`, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();
  return data;
};

const Details = async (subpath: string, name: string | undefined): Promise<DeploymentDetails> => {
  const response = await fetch(`${baseURL}/${subpath}/${name}`, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();
  return data;
};

export {
  deploymentsList,
  Details,
};