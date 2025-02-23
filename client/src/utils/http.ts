import { ReturnPromiseList, ReturnPromiseDetails } from '../types/types';

const baseURL = 'http://192.168.254.45/backend';

const ListObjects = async (type: string | undefined): ReturnPromiseList => {
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

const Details = async (subpath: string | undefined, name: string | undefined): ReturnPromiseDetails => {
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