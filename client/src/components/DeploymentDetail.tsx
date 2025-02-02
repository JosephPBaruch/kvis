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
    return <div>Loading...</div>;
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const renderDetails = (obj: any, parentKey: string = ''): JSX.Element[] => {
    return Object.keys(obj).map((key) => {
      const value = obj[key];
      const displayKey = parentKey ? `${parentKey}.${key}` : key;

      if (typeof value === 'object' && value !== null) {
        return (
          <div key={displayKey}>
            <h3>{displayKey}</h3>
            {renderDetails(value, displayKey)}
          </div>
        );
      }

      return (
        <p key={displayKey}>
          <strong>{displayKey}:</strong> {value}
        </p>
      );
    });
  };

  return (
    <div>
      <h2>{name}</h2>
      {renderDetails(details)}
    </div>
  );
};

export default DeploymentDetail;
