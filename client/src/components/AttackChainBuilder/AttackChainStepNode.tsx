import React, { memo, useState } from 'react';
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
  asset: Asset
  position: number
}

interface AttackChainStepNodeProps extends NodeProps {
      data: AttackChainStepNodeData
  };

const AttackChainStepNode = ({ data, isConnectable }: AttackChainStepNodeProps) => {
  const actionName = data.action?.name;
  const [toolbarVisible, setToolbarVisible] = useState<boolean>(false);
  const [toolbarPosition, setToolbarPosition] = useState<Position>(Position.Top);

  const onDoubleClick = () => {
    if (toolbarVisible) {
      setToolbarVisible(false);
    }
    else {
      setToolbarVisible(true);
    }
  }


  return (
    <div onDoubleClick={onDoubleClick}>
      <NodeToolbar isVisible={toolbarVisible} position={toolbarPosition}>
      Action Name: {data.action.name}<br></br>
      Description: {data.action.description}<br></br>
      Complexity: {data.action.complexity}<br></br>
      Capability: {data.action.capabilityId}<br></br>
      Vulnerability: {data.action.vulnerabilityId}<br></br>
      Asset: <br></br>
      Detection: <br></br>
      Mitigation: <br></br>
      </NodeToolbar>

     <Handle type="target" position={Position.Left} />
     <Handle type="source" position={Position.Right} />
      <Handle type="target" position={Position.Top} />
      <div>
        <label htmlFor="text">{actionName}</label>
      </div>
      <Handle type="source" position={Position.Bottom}/>

    </div>
  );
}

export default AttackChainStepNode;