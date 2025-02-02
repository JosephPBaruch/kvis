import React, { useEffect, useState } from 'react';
import { DeploymentDetailProps, DeploymentDetails } from '../types/types';
import { deploymentDetails } from '../utils/http';

const DeploymentDetail: React.FC<DeploymentDetailProps> = ({ name }) => {
  const [details, setDetails] = useState<DeploymentDetails | null>(null);

  useEffect(() => {
    const fetchDetails = async () => {
      try {
        const data = await deploymentDetails(name!);
        setDetails(data);
      } catch (error) {
        console.error('Error fetching deployment details:', error);
      }
    };

    fetchDetails();
  }, [name]);

  if (!details) {
    return <div>error: a resource cannot be retrieved by name across all namespaces</div>;
  }

  return (
    <div>
      <h1>{name}</h1>
      <p><strong>Name:</strong> {details.name}</p>
      <p><strong>Namespace:</strong> {details.namespace}</p>
      <p><strong>Creation Timestamp:</strong> {details.creationTimestamp}</p>
      <p><strong>Labels:</strong> {details.labels}</p>
      <p><strong>Annotations:</strong> {details.annotations}</p>
      <p><strong>Selector:</strong> {details.selector}</p>
      <p><strong>Replicas:</strong> {details.replicas}</p>
      <p><strong>Strategy Type:</strong> {details.strategyType}</p>
      <p><strong>Min Ready Seconds:</strong> {details.minReadySeconds}</p>
      <p><strong>Rolling Update Strategy:</strong> {details.rollingUpdateStrategy}</p>
      <div>
        <h3>Pod Template</h3>
        <p><strong>Labels:</strong> {details.podTemplate.labels}</p>
        <p><strong>Containers:</strong></p>
        <ul>
          <li><strong>Image:</strong> {details.podTemplate.containers.image}</li>
          <li><strong>Port:</strong> {details.podTemplate.containers.port}</li>
          <li><strong>Host Port:</strong> {details.podTemplate.containers.hostPort}</li>
          <li><strong>Environment:</strong> {details.podTemplate.containers.environment}</li>
          <li><strong>Mounts:</strong> {details.podTemplate.containers.mounts}</li>
        </ul>
        <p><strong>Volumes:</strong> {details.podTemplate.volumes}</p>
      </div>
      <div>
        <h3>Conditions</h3>
        <p><strong>Available:</strong> {details.conditions.available}</p>
        <p><strong>Progressing:</strong> {details.conditions.progressing}</p>
      </div>
      <p><strong>Old Replica Sets:</strong> {details.oldReplicaSets}</p>
      <p><strong>New Replica Set:</strong> {details.newReplicaSet}</p>
      <p><strong>Events:</strong> {details.events}</p>
    </div>
  );
};

export default DeploymentDetail;
