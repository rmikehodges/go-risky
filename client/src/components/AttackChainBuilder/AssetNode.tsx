import React, { memo } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';
import  Asset  from '../Assets/Asset';

const handleStyle = { left: 10 };

interface AssetNodeData {
  label: string
  asset: Asset | null
}

interface AssetNodeProps extends NodeProps {
      data: AssetNodeData
  };

const AssetNode = ({ data, isConnectable }: AssetNodeProps) => {
  const assetName = data.asset?.name;
  return (
    <>
      <Handle type="target" position={Position.Top} />
      <div>
        <label htmlFor="text">{assetName}</label><br></br>
      </div>
      <Handle type="source" position={Position.Bottom} id="a" />
    </>
  );
}

export default AssetNode;