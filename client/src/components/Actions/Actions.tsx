import axios from "axios";
import { useEffect, useState } from "react";
import CreateAction from "./CreateAction";
import { UpdateActionInput } from "./UpdateAction";
import UpdateAction from "./UpdateAction";
import DeleteAction from "./DeleteAction";
import { UUID } from 'crypto';

export interface ActionOutput {
  id: UUID
  name: string
  description: string
  capabilityId: UUID | null
  vulnerabilityId: UUID | null
  businessId: UUID | null
  complexity: string
  assetId: UUID | null
  createdAt: Date
  }

const Actions = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [actions, setActions] = useState<ActionOutput[] | null>(null);

    useEffect(() => {
      axios.get<ActionOutput[]>(`http://localhost:8081/actions?businessId=${businessId}`)
        .then(res => {
        const actionsResp = res.data;
       setActions(actionsResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Actions'>
      <div>
        <CreateAction />
      </div>
    <table>
      <thead><tr>
        <th>#</th>
        <th>Name</th>
        <th>BusinessId</th>
        <th>CapabilityId</th>
        <th>VulnerabilityId</th>
        <th>Complexity</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {actions?.map(action => {
        counter = counter + 1;    
        const updateActionInput: UpdateActionInput = {id: action.id, description: action.description ,name: action.name, complexity: action.complexity,businessId: action.businessId, assetId: action.assetId, capabilityId: action.capabilityId, vulnerabilityId: action.vulnerabilityId}
        const deleteActionInput = {id: action.id}
        return (
          <tr key={action.id}>
            <td><a href={`http://localhost:3000/action?id=${action.id}`}>{counter}</a></td>
            <td>{action.name}</td>
            <td>{action.businessId}</td>
            <td>{action.capabilityId}</td>
            <td>{action.vulnerabilityId}</td>
            <td>{action.complexity}</td>
            <td>{action.createdAt.toString()}</td>
            <td><UpdateAction {...updateActionInput} /></td>
            <td><DeleteAction {...deleteActionInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default Actions