import { Deployment, DeploymentDetails } from '../types/types';

const baseURL = 'http://localhost:8090';

const ListObjects = async (type: string): Promise<Deployment[]> => {
  const response = await fetch(`${baseURL}/${type}`, {
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

const Details = async (subpath: string | undefined, name: string | undefined): Promise<DeploymentDetails> => {
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
  ListObjects,
  Details,
};