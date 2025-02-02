import React, { useEffect, useState } from 'react';
import { DeploymentDetailProps, DeploymentDetails } from '../types/types';
import { Details } from '../utils/http';

const DetailInformation: React.FC<DeploymentDetailProps> = ({ type, name }) => {
  const [details, setDetails] = useState<DeploymentDetails | null>(null);

  useEffect(() => {
    const fetchDetails = async () => {
      try {
        const data: DeploymentDetails = await Details(type, name);
        setDetails(data);
      } catch (error) {
        console.error('Error fetching deployment details:', error);
      }
    };

    fetchDetails();
  }, [type, name]);

  if (!details) {
    return <div>Loading...</div>;
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const renderDetails = (obj: Record<string, any>, parentKey: string = ''): JSX.Element[] => {
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
      <h1>{name}</h1>
      {renderDetails(details)}
    </div>
  );
};

export default DetailInformation;
