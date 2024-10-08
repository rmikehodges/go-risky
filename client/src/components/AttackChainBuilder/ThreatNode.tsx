import React, { memo } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';
import  Threat  from '../Threats/Threat';

const handleStyle = { left: 10 };

interface ThreatNodeData {
  label: string
  threat: Threat | null
}

interface ThreatNodeProps extends NodeProps {
      data: ThreatNodeData
  };

const ThreatNode = ({ data, isConnectable }: ThreatNodeProps) => {
  const threatName = data.threat?.name;
  return (
    <>
      <Handle type="target" position={Position.Top} />
      <div>
        <label htmlFor="text">{threatName}</label><br></br>
      </div>
      <Handle type="source" position={Position.Bottom} id="a" />
    </>
  );
}

export default ThreatNode;