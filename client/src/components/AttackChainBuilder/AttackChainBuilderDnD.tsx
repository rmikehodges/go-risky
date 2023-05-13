import React, { useState, useRef, useCallback } from 'react';
import ReactFlow, {
  ReactFlowProvider,
  addEdge,
  useNodesState,
  useEdgesState,
  Controls,
  ReactFlowInstance,
  NodeTypes,
  Node,
} from 'reactflow';
import 'reactflow/dist/style.css';
import { useMemo } from 'react';

import Sidebar from './Sidebar';
import ActionNode from './ActionNode';

import '../../index.css';
import { ActionOutput } from '../Actions/Actions';
import { ThreatOutput } from '../Threats/Threats';
import { AssetOutput } from '../Assets/Assets';
import AssetNode from './AssetNode';
import ThreatNode from './ThreatNode';

const initialNodes: Node[] = [];

let id = 0;
const getId = () => `dndnode_${id++}`;

// type draggedObjectType = AssetOutput | ThreatOutput | ActionOutput


// const isThreatOutput = (object: any) => {
//     return object && object.type === 'ThreatOutput'
// }
// const isAssetOutput = (object: any) => {
//     return object && object.type === 'AssetOutput'
// }
// const isActionOutput = (object: any) => {
//     return object && object.type === 'ActionOutput'
// }


const AttackChainBuilderDnD = () => {
  const nodeTypes = useMemo(() => ({action: ActionNode, asset: AssetNode, threat: ThreatNode}), []);
  const reactFlowWrapper = useRef<HTMLDivElement>(null);
  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);
  const [reactFlowInstance, setReactFlowInstance] = useState<ReactFlowInstance | null>(null);


  const onConnect = useCallback((params: any) => setEdges((eds) => addEdge(params, eds)), []);

  const onDragOver = useCallback((event: any) => {
    event.preventDefault();
    event.dataTransfer.dropEffect = 'move';
  }, []);

  const onDrop = useCallback(
    (event: any) => {
      event.preventDefault();

      if (!reactFlowWrapper.current || !reactFlowInstance) {
        return null;
      }

      const reactFlowBounds = reactFlowWrapper.current?.getBoundingClientRect();
      const type = event.dataTransfer.getData('application/reactflow');
      const draggedObject = JSON.parse(event.dataTransfer.getData('application/json'));



      // check if the dropped element is valid
      if (typeof type === 'undefined' || !type) {
        return;
      }

      const position = reactFlowInstance?.project({
        x: event.clientX - reactFlowBounds?.left,
        y: event.clientY - reactFlowBounds?.top,
      });

      let data;
      if (type === 'action') {
        data = { label: draggedObject.name , action: draggedObject }
      } else if (type === 'asset') {
        data = { label: draggedObject.name , asset: draggedObject }
      } else if ( type ===  'threat') {
        data = { label: draggedObject.name , threat: draggedObject }
      }

      const newNode = {
        id: getId(),
        type: type,
        position,
        data: data,
      };

      setNodes((nds) => nds.concat(newNode));
    },
    [reactFlowInstance]
  );

  return (
    <div style={{ height: 800 }} className="dndflow">
      <ReactFlowProvider>
        <div className="reactflow-wrapper" ref={reactFlowWrapper}>
          <ReactFlow
            nodes={nodes}
            nodeTypes={nodeTypes}
            edges={edges}
            onNodesChange={onNodesChange}
            onEdgesChange={onEdgesChange}
            onConnect={onConnect}
            onInit={setReactFlowInstance}
            onDrop={onDrop}
            onDragOver={onDragOver}
            fitView
          >
            <Controls />
          </ReactFlow>
        </div>
        <Sidebar />
      </ReactFlowProvider>
    </div>
  );
};

export default AttackChainBuilderDnD;