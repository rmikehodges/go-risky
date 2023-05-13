import React, { memo } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';
import { ActionOutput } from '../Actions/Actions';

const handleStyle = { left: 10 };

interface ActionNodeData {
  label: string
  action: ActionOutput | null
}

interface ActionNodeProps extends NodeProps {
      data: ActionNodeData
  };

const ActionNode = ({ data, isConnectable }: ActionNodeProps) => {
  const actionName = data.action?.name;
  return (
    <>
      <Handle type="target" position={Position.Top} />
      <div>
        <label htmlFor="text">{actionName}</label><br></br>
      </div>
      <Handle type="source" position={Position.Bottom} id="a" />
    </>
  );
}

export default ActionNode;