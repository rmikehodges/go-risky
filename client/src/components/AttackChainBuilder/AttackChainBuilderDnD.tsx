import React, { useState, useRef, useCallback, useEffect } from 'react';
import ReactFlow, {
  ReactFlowProvider,
  addEdge,
  useNodesState,
  useEdgesState,
  Controls,
  ReactFlowInstance,
  ConnectionLineType,
  getIncomers,
  getOutgoers,
  getConnectedEdges,
  Node,
  Background,
  BackgroundVariant,
} from 'reactflow';
import 'reactflow/dist/style.css';
import { useMemo } from 'react';

import Sidebar from './Sidebar';
import ActionNode from './ActionNode';
import '../../index.css';
import  Action  from '../Actions/Action';
import  Threat  from '../Threats/Threat';
import  Asset  from '../Assets/Asset';
import AssetNode from './AssetNode';
import AttackChain from '../AttackChains/AttackChain';
import ThreatNode from './ThreatNode';
import AttackChainStepNode from './AttackChainStepNode';
import ImpactBuilder from '../ImpactBuilder/ImpactBuilder';
import { on } from 'events';
import axios from 'axios';
import AttackChainStep from '../AttackChainSteps/AttackChainStep';
import { UUID } from 'crypto';
import AttackChainDropdown from './AttackChainDropdown';

const initialNodes: Node[] = [];

let id = 0;
const getId = () => `dndnode_${id++}`;

// type draggedObjectType = Asset | Threat | Action


// const isThreat = (object: any) => {
//     return object && object.type === 'Threat'
// }
// const isAsset = (object: any) => {
//     return object && object.type === 'Asset'
// }
// const isAction = (object: any) => {
//     return object && object.type === 'Action'
// }


const AttackChainBuilderDnD = () => {
  var businessId:UUID = "23628819-59dd-45f3-8395-aceeca86bc9c"
  const nodeTypes = useMemo(() => ({action: ActionNode, asset: AssetNode, threat: ThreatNode, attackChainStep: AttackChainStepNode}), []);
  const reactFlowWrapper = useRef<HTMLDivElement>(null);
  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);
  const [reactFlowInstance, setReactFlowInstance] = useState<ReactFlowInstance | null>(null);
  const [attackChainId, setAttackChainId] = useState<UUID>("20036fa3-45c6-47b2-a343-f88bcd4f5e07");
  const [attackChains, setAttackChains] = useState<AttackChain[]>([]);
  const [attackChainSteps, setAttackChainSteps] = useState<AttackChainStep[]>([]);
  const [attackChainStepPosition, setAttackChainStepPosition] = useState<number>(0);


  //TODO: Make the graph render the attack chain steps in the correct order

  useEffect(() => {
    axios.get('http://localhost:8081/attackChains?businessId='+businessId).then((res) => {
      setAttackChains(res.data);
    });
  }, []);

  //TODO: Update nodes once this is done, its done in the backend automatically, but needs to reflect the state
  //in the frontend. The delete on the database layer may not be ideal. 
  const onNodesDelete = useCallback(
    (deleted: Node[]) => {
      setEdges(
        deleted.reduce((acc, node) => {
          //TODO: Add Error Handling
          if (node.type === 'attackChainStep') {
            axios.delete(`http://localhost:8081/attackChainStep?id=${node.data.attackChainStep.id}`)
          }
          const incomers = getIncomers(node, nodes, edges);
          const outgoers = getOutgoers(node, nodes, edges);
          const connectedEdges = getConnectedEdges([node], edges);

          const remainingEdges = acc.filter((edge) => !connectedEdges.includes(edge));

          const createdEdges = incomers.flatMap(({ id: source }) =>
            outgoers.map(({ id: target }) => ({ id: `${source}->${target}`, source, target }))
          );

          return [...remainingEdges, ...createdEdges];
        }, edges)
      );
    },
    [nodes, edges]
  );

  const onConnect = useCallback((params: any) => {
    let nextStepId: UUID;
    let previousStepId: UUID; 
    nodes.forEach((node) => {
      if (params.source === node.id) {
        previousStepId = node.data.attackChainStep.id;
      } else if (params.target === node.id) {
        nextStepId = node.data.attackChainStep.id;
      }

    });
    setNodes((nds) =>
    nds.map((node) => {
      if (node.id === params.target) {
        node.data = {
          ...node.data,
          attackChainStep: {
            ...node.data.attackChainStep,
            previousStep: previousStepId,
          }
        };
        axios.patch('http://localhost:8081/attackChainStep', node.data.attackChainStep).then((res) => {
          console.log(node.data.attackChainStep)
        });
      } else if (node.id === params.source) {
        node.data = {
          ...node.data,
          attackChainStep: {
            ...node.data.attackChainStep,
            nextStep: nextStepId,
          }
        };
        axios.patch('http://localhost:8081/attackChainStep', node.data.attackChainStep).then((res) => {
          console.log(node.data.attackChainStep)
        });
      }
      return node;
    })
  );
    setAttackChainStepPosition(attackChainStepPosition + 1);
    setEdges((eds) => addEdge({...params,type:"smoothstep"}, eds))
  }, [nodes, edges, attackChainStepPosition]);

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
      let attackChainStepData: AttackChainStep;
      if (type === 'attackChainStep') { }
        attackChainStepData = { actionId: draggedObject.id, businessId: businessId, attackChainId: attackChainId, id: null, createdAt:null, assetId:null, nextStep:null, previousStep: null }
        axios.post('http://localhost:8081/attackChainStep', attackChainStepData).then((res) => {
          attackChainStepData.id = res.data;
          console.log(res.data)
          data = { label: draggedObject.name , action: draggedObject, attackChainStep: attackChainStepData }
          const newNode = {
            id: getId(),
            type: type,
            position,
            data: data,
          };
    
          setNodes((nds) => nds.concat(newNode));
        });


    },
    [reactFlowInstance]
  );

  const handleSelectedAttackChain = (option: string) => {
    let selectedOption: UUID = option as UUID;
    setAttackChainId(selectedOption);
    console.log('http://localhost:8081/attackChainSteps?attackChainId='+option+'&businessId='+businessId)
    axios.get('http://localhost:8081/attackChainSteps?attackChainId='+option+'&businessId='+businessId).then((res) => {
    setAttackChainSteps(res.data);
    });
  }

  return (
    <div className="flow-container">
    <div style={{ height: 800 }} className="dndflow">
      <ReactFlowProvider>
        <div className="reactflow-wrapper" ref={reactFlowWrapper}>
          <ReactFlow
            nodes={nodes}
            nodeTypes={nodeTypes}
            edges={edges}
            onNodesChange={onNodesChange}
            onNodesDelete={onNodesDelete}
            onEdgesChange={onEdgesChange}
            onConnect={onConnect}
            connectionLineType={ConnectionLineType.Step}
            onInit={setReactFlowInstance}
            onDrop={onDrop}
            onDragOver={onDragOver}
            fitView
          >
            <Controls />
            <Background color="#ccc" variant={BackgroundVariant.Lines} />
          </ReactFlow>
        </div>
      </ReactFlowProvider>
      </div>
      <div className="sidebar-container">
        <div className="sidebar">
          <div className="sidebar-header">
            <h2>Attack Chain Builder</h2>
            <div className="sidebar-header-select">
            <AttackChainDropdown options={attackChains} selectedAttackChain={attackChainId} onSelectOption={handleSelectedAttackChain}/>
            </div>
          </div>
          <Sidebar />
        </div>
      </div>
      <div className="impactBuilder"><ImpactBuilder /></div>
      {/* <div className="bottomtoolbar-container">
        <div className="box">Newsfeed</div>
        <div className="box">Chatbot</div>
      </div> */}


    </div>
  );
};

export default AttackChainBuilderDnD;