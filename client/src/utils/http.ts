import { Deployment } from '../types/types';

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

const deploymentDetails = async (name: string): Promise<Deployment[]> => {
  const response = await fetch(`${baseURL}/deployments/${name}`, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json(); 
  console.log(data);
  return data;
};

export {
  deploymentsList,
  deploymentDetails,
};