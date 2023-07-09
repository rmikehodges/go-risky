import React, { memo } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';
import  AttackChainStep  from "../AttackChainSteps/AttackChainStep";
import Asset from '../Assets/Asset';
import AttackChain from '../AttackChains/AttackChain';
import Detection from '../Detections/Detection';
import Mitigation from '../Mitigations/Mitigation';
import Action from '../Actions/Action';
import { UUID } from 'crypto';

const handleStyle = { left: 10 };

interface AttackChainStepNodeData {
  label: string
  id: UUID | null
  action: Action
  asset: Asset | null
  attackChain: AttackChain | null
  detection: Detection | null
  mitigation: Mitigation | null
  position: number
}

interface AttackChainStepNodeProps extends NodeProps {
      data: AttackChainStepNodeData
  };

const AttackChainStepNode = ({ data, isConnectable }: AttackChainStepNodeProps) => {
  const actionName = data.action?.name;
  return (
    <>
     <Handle type="source" position={Position.Left} />
     <Handle type="target" position={Position.Right} />
      <Handle type="target" position={Position.Top} />
      <div>
       Name: {actionName}<br></br>
       Description: {data.action?.description}<br></br>
        Complexity: {data.action?.complexity}<br></br>
        
      </div>
      <Handle type="source" position={Position.Bottom}/>

    </>
  );
}

export default AttackChainStepNode;