import React from 'react';
import { DeploymentDetailProps } from '../types/types';


const DeploymentDetail: React.FC<DeploymentDetailProps> = ({ name }) => {
  return (
    <div>
      <h2>Deployment Details</h2>
      <p><strong>Name:</strong> {name}</p>
      {/* Add more details about the deployment here */}
    </div>
  );
};

export default DeploymentDetail;
