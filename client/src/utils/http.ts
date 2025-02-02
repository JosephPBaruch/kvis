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
  const data = await response.json(); // Store the result of response.json()
  console.log(data);
  return data;
};

export default deploymentsList;