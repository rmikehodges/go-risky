import axios from "axios";
import { useEffect, useState } from "react";
import CreateMitigation from "./CreateMitigation";
import UpdateMitigation from "./UpdateMitigation";
import DeleteMitigation from "./DeleteMitigation";
import  Mitigation  from "./Mitigation";
import { UUID } from 'crypto';


const ListMitigations = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [mitigations, setMitigations] = useState<Mitigation[] | null>(null);

    useEffect(() => {
      axios.get<Mitigation[]>(`http://localhost:8081/mitigations?businessId=${businessId}`)
        .then(res => {
        const mitigationsResp = res.data;
       setMitigations(mitigationsResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Mitigations'>
      <div>
        <CreateMitigation />
      </div>
    <table>
      <thead><tr>
        <th>#</th>
        <th>Name</th>
        <th>Description</th>
        <th>BusinessId</th>
        <th>ActionId</th>
        <th>Implemented</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {mitigations?.map(mitigation => {
        counter = counter + 1;    
        const updateMitigationInput: Mitigation = {id: mitigation.id, description: mitigation.description ,
          name: mitigation.name, businessId: mitigation.businessId, actionId: mitigation.actionId, implemented: mitigation.implemented,
          createdAt: mitigation.createdAt}
        const deleteMitigationInput = {id: mitigation.id}
        return (
          <tr key={mitigation.id}>
            <td><a href={`http://localhost:3000/mitigation?id=${mitigation.id}`}>{counter}</a></td>
            <td>{mitigation.name}</td>
            <td>{mitigation.description}</td>
            <td>{mitigation.businessId}</td>
            <td>{mitigation.actionId}</td>
            <td>{mitigation.implemented}</td>
            <td>{mitigation.createdAt?.toString()}</td>
            <td><UpdateMitigation {...updateMitigationInput} /></td>
            <td><DeleteMitigation {...deleteMitigationInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListMitigations