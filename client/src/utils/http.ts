import { Deployment } from '../types/deployment';

const baseURL = 'localhost:8090';

const deploymentsList = async (): Promise<Deployment[]> => {
  const response = await fetch(`${baseURL}/deployments`, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  console.log(response.json());
  return response.json();
};

export default deploymentsList;