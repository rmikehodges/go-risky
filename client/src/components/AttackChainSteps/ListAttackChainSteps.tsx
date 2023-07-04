import axios from "axios";
import { useEffect, useState } from "react";
import CreateAttackChainStep from "./CreateAttackChainStep";
import UpdateAttackChainStep from "./UpdateAttackChainStep";
import DeleteAttackChainStep from "./DeleteAttackChainStep";
import  AttackChainStep  from "./AttackChainStep";


const ListAttackChainSteps = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [attackChainSteps, setAttackChainSteps] = useState<AttackChainStep[] | null>(null);

    useEffect(() => {
      axios.get<AttackChainStep[]>(`http://localhost:8081/attackChainSteps?businessId=${businessId}`)
        .then(res => {
        const attackChainStepsResp = res.data;
       setAttackChainSteps(attackChainStepsResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='AttackChainSteps'>
      <div>
        <CreateAttackChainStep />
      </div>
    <table>
      <thead><tr>
        <th>#</th>
        <th>actionId</th>
        <th>assetId</th>
        <th>attackChainId</th>
        <th>Position</th>
        <th>BusinessId</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {attackChainSteps?.map(attackChainStep => {
        counter = counter + 1;    
        const updateAttackChainStepInput: AttackChainStep = {id: attackChainStep.id, actionId: attackChainStep.actionId, assetId: attackChainStep.assetId, attackChainId: attackChainStep.attackChainId, 
          position: attackChainStep.position, businessId: attackChainStep.businessId,
          createdAt: attackChainStep.createdAt}
        const deleteAttackChainStepInput = {id: attackChainStep.id}
        return (
          <tr key={attackChainStep.id}>
            <td><a href={`http://localhost:3000/attackChainStep?id=${attackChainStep.id}`}>{counter}</a></td>
            <td>{attackChainStep.actionId}</td>
            <td>{attackChainStep.assetId}</td>
            <td>{attackChainStep.attackChainId}</td>
            <td>{attackChainStep.position}</td>
            <td>{attackChainStep.businessId}</td>
            <td>{attackChainStep.createdAt?.toString()}</td>
            <td><UpdateAttackChainStep {...updateAttackChainStepInput} /></td>
            <td><DeleteAttackChainStep {...deleteAttackChainStepInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListAttackChainSteps