import React, { memo } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';
import  Action  from "../Actions/Action";

const handleStyle = { left: 10 };

interface ActionNodeData {
  label: string
  action: Action | null
}

interface ActionNodeProps extends NodeProps {
      data: ActionNodeData
  };

const ActionNode = ({ data, isConnectable }: ActionNodeProps) => {
  const actionName = data.action?.name;
  return (
    <>
     <Handle type="source" position={Position.Left} />
     <Handle type="target" position={Position.Right} />
      <Handle type="target" position={Position.Top} />
      <div>
       Name: {actionName}
       Description: {data.action?.description}
        Complexity: {data.action?.complexity}
        
      </div>
      <Handle type="source" position={Position.Bottom}/>

    </>
  );
}

export default ActionNode;