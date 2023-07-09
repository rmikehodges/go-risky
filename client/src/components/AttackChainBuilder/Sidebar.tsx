import axios from 'axios';
import React from 'react';
import { useEffect, useState } from 'react';
import  Action  from '../Actions/Action';
import  Threat  from '../Threats/Threat';
import  Asset  from '../Assets/Asset';
import ThreatDropdown from '../ImpactBuilder/ThreatDropdown';

export default (props:any) => {
    var businessId = "23628819-59dd-45f3-8395-aceeca86bc9c"
    const [actions, setActions] = useState<Action[] | null>(null);
    const [assets, setAssets] = useState<Asset[] | null>(null);

    useEffect(() => {
        axios.get<Action[]>(`http://localhost:8081/actions?businessId=${businessId}`)
          .then(res => {
          const actionsResp = res.data;
         setActions(actionsResp)});
        axios.get<Asset[]>(`http://localhost:8081/assets?businessId=${businessId}`)
        .then(res => {
        const assetsResp = res.data;
       setAssets(assetsResp)});
        }, [businessId]);


//  const onActionDragStart = (event: any, nodeType: any, action: Action) => {
//     event.dataTransfer.setData('application/reactflow', nodeType);
//     event.dataTransfer.setData('application/json', JSON.stringify(action));
//     event.dataTransfer.effectAllowed = 'move';
//   };

  //TODO: Add in attackchain to drag start
  const onAttackChainStepDragStart = (event: any, nodeType: any, action: Action) => {
    event.dataTransfer.setData('application/reactflow', nodeType);
    event.dataTransfer.setData('application/json', JSON.stringify(action));
    event.dataTransfer.effectAllowed = 'move';
  };

  // const onThreatDragStart = (event: any, nodeType: any, threat: Threat) => {
  //   event.dataTransfer.setData('application/reactflow', nodeType);
  //   event.dataTransfer.setData('application/json', JSON.stringify(threat));
  //   event.dataTransfer.effectAllowed = 'move';
  // };

  // const onAssetDragStart = (event: any, nodeType: any, asset: Asset) => {
  //   event.dataTransfer.setData('application/reactflow', nodeType);
  //   event.dataTransfer.setData('application/json', JSON.stringify(asset));
  //   event.dataTransfer.effectAllowed = 'move';
  // };

  return (
    <div>
      <div>
      <br></br>
        Attack Chain
      </div>
    <aside>
      <div className="description">Attack Chain Steps</div>
      {actions?.map(action => {
        return (
           <div  key={action.id} className="dndnode" onDragStart={(event) => onAttackChainStepDragStart(event, 'attackChainStep', action)} draggable>
            {action.name}
          </div>

        )
          
      })}

    {/* <div className="description">Threats</div>
      {threats?.map(threat => {
        return (
           <div  key={threat.id} className="dndnode" onDragStart={(event) => onThreatDragStart(event, 'threat', threat)} draggable>
            {threat.name}
          </div>

        )
      })} */}
    {/* <div className="description">Assets</div>
      {assets?.map(asset => {
        return (
           <div  key={asset.id} className="dndnode" onDragStart={(event) => onAssetDragStart(event, 'asset', asset)} draggable>
            {asset.name}
          </div>

        )
    })} */}
    </aside>
    </div>
  );
};
