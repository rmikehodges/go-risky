import React, { memo } from 'react';
import { Handle, Position, NodeProps, NodeToolbar } from 'reactflow';
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
  action: Action
  attackChainStep: AttackChainStep
  position: number
  toolbarVisible: boolean
  toolbarPosition: Position
}

interface AttackChainStepNodeProps extends NodeProps {
      data: AttackChainStepNodeData
  };

const AttackChainStepNode = ({ data, isConnectable }: AttackChainStepNodeProps) => {
  const actionName = data.action?.name;

  return (
    <>
      <NodeToolbar isVisible={data.toolbarVisible} position={data.toolbarPosition}>
      Position: {data.attackChainStep.position}<br></br>
      </NodeToolbar>

     <Handle type="target" position={Position.Left} />
     <Handle type="source" position={Position.Right} />
      <Handle type="target" position={Position.Top} />
      <div>
       Name: {actionName}<br></br>
       Position: {data.attackChainStep.position}<br></br>
      </div>
      <Handle type="source" position={Position.Bottom}/>

    </>
  );
}

export default AttackChainStepNode;