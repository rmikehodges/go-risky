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
  Edge,
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
  const [actions, setActions] = useState<Action[]>([]);
  const [showOverlappingAttackChains, setShowOverlappingAttackChains] = useState<boolean>(false);


  //TODO: Make the graph render the attack chain steps in the correct order

  useEffect(() => {
    axios.get('http://localhost:8081/v0/attackChains?businessId='+businessId).then((res) => {
      setAttackChains(res.data);
    });
    axios.get<Action[]>(`http://localhost:8081/v0/actions?businessId=${businessId}`)
    .then(res => {setActions(res.data)});
  }, []);

  //TODO: Update nodes once this is done, its done in the backend automatically, but needs to reflect the state
  //in the frontend.
  const onNodesDelete = useCallback(
    (deleted: Node[]) => {
      setEdges(
        deleted.reduce((acc, node) => {
          //TODO: Add Error Handling
          if (node.type === 'attackChainStep') {
            axios.delete(`http://localhost:8081/v0/attackChainStep?id=${node.data.attackChainStep.id}`)
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
            attackChainId: attackChainId,
            previousStep: previousStepId,
          }
        };
        axios.patch('http://localhost:8081/v0/attackChainStep', node.data.attackChainStep).then((res) => {
          console.log("success ")
        });
      } else if (node.id === params.source) { 
        node.data = {
          ...node.data,
          attackChainStep: {
            ...node.data.attackChainStep,
            attackChainId: attackChainId,
            nextStep: nextStepId,
          }
        };
        axios.patch('http://localhost:8081/v0/attackChainStep', node.data.attackChainStep).then((res) => {
          console.log(res)
        });
      }
      return node;
    })
  );
    // setAttackChainStepPosition(attackChainStepPosition + 1);
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
        axios.post('http://localhost:8081/v0/attackChainStep', attackChainStepData).then((res) => {
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

  //TODO: Place nodes in order of attack chain steps
  const handleSelectedAttackChain = (event: any) => {
    let selectedOption: UUID = event.target.value as UUID;
    let attackChainStepsResponse: AttackChainStep[] = [];
    let tempNodes: Node[] = [];

    let currentNode: Node;
    let currentEdge: Edge;
    let edgesToSet: Edge[] = [];
    let nodesToSet: Node[] = [];

    axios.get('http://localhost:8081/v0/attackChainSteps?attackChainId='+selectedOption+'&businessId='+businessId).then((res) => { 
      attackChainStepsResponse= res.data;
      let result = attackChainStepsResponse.map((tempAttackChainStep) => {
        let currentAction: Action = {id: null, name: "INVALID", description: null, businessId: null, createdAt: null, capabilityId: null, vulnerabilityId: null, complexity: ""};
        for (let i = 0; i < actions?.length; i++) {
          currentAction = actions[i] as Action;
          if (currentAction.id == tempAttackChainStep.actionId) {
            currentNode = {
              id: getId(),
              type: 'attackChainStep',
              position:{x:0, y:0},
              data: { label: currentAction.name, action: currentAction, attackChainStep: tempAttackChainStep },
            };
            return currentNode;
          }
        }
      });   
      tempNodes = result as Node[];

      nodesToSet = tempNodes;

      for (let i =0; i< tempNodes.length; i++) {
        let yCounter = 0;
        for (let j = 0; j < tempNodes.length; j++) {
          if (tempNodes[i].data.attackChainStep.id == tempNodes[j].data.attackChainStep.previousStep) {
            currentEdge = { id: getId(), source: tempNodes[i].id, target: tempNodes[j].id, type: "smoothstep" };
            nodesToSet[j].position.x = nodesToSet[i].position.x + 200;
            nodesToSet[j].position.y = nodesToSet[i].position.y + 100*yCounter;
            edgesToSet.push(currentEdge);
            yCounter++;
          }

        }
      }

      setNodes(tempNodes);
      setEdges(edgesToSet); 
      setAttackChainSteps(attackChainStepsResponse);
      setAttackChainId(selectedOption);

      });

    };

  //TODO Connect this to a node.
  const onShowOverlappingAttackChainsClick = (actionId: string) => {
    setShowOverlappingAttackChains(!showOverlappingAttackChains);

    if (!showOverlappingAttackChains) {
      let attackChainStepsResponse: AttackChainStep[] = [];
      let tempNodes: Node[] = [];
      let currentNode: Node;
      let currentEdge: Edge;
      


      axios.get('http://localhost:8081/v0/attackChainSteps?businessId='+businessId+"&actionId="+actionId).then((res) => {
        attackChainStepsResponse = res.data;
        setNodes(attackChainStepsResponse.map((tempAttackChainStep) => {
          let currentAction: Action = {id: null, name: "INVALID", description: null, businessId: null, createdAt: null, capabilityId: null, vulnerabilityId: null, complexity: ""};
          for (let i = 0; i < actions?.length; i++) {
            currentAction = actions[i] as Action;
            if (currentAction.id === tempAttackChainStep.actionId) {
              break;
            }
          }
          currentNode = {
            id: getId(),
            type: 'attackChainStep',
            position: { x: 0, y: 0 },
            data: { label: currentAction.name, action: currentAction, attackChainStep: tempAttackChainStep },
          }
          tempNodes.push(currentNode);
          return currentNode
        }));
        setEdges(() => {
          let edgesToSet: Edge[] = [];
          for (let i = 0; i < tempNodes.length; i++) {
            currentNode = tempNodes[i];
            let nextNodes = tempNodes.filter((node) => node.data.attackChainStep.previousStep === currentNode.data.attackChainStep.id);
            for (let j = 0; j < nextNodes.length; j++) {
              currentEdge = { id: getId(), source: currentNode.id, target: nextNodes[j].id, type: "smoothstep" };
              edgesToSet.push(currentEdge);
          }
          }
          return edgesToSet
        });
    })
  }};


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
            <select value={attackChainId} onChange={handleSelectedAttackChain}>
              {attackChains?.map((tempAttackChain, i) => (
                <option key={i} value={tempAttackChain.id?.toString()}>{tempAttackChain.id}</option>
                ))}
              </select>
            
            </div>
          </div>
          <Sidebar actions={actions}/>
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