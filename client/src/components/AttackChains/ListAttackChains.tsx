import axios from "axios";
import { useEffect, useState } from "react";
import CreateAttackChain from "./CreateAttackChain";
import UpdateAttackChain from "./UpdateAttackChain";
import DeleteAttackChain from "./DeleteAttackChain";
import  AttackChain  from "./AttackChain";


const ListAttackChains = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [attackChains, setAttackChains] = useState<AttackChain[] | null>(null);

    useEffect(() => {
      axios.get<AttackChain[]>(`http://localhost:8081/v0/attackChains?businessId=${businessId}`)
        .then(res => {
        const attackChainsResp = res.data;
       setAttackChains(attackChainsResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='AttackChains'>
      <div>
        <CreateAttackChain />
      </div>
    <table>
      <thead><tr>
        <th>#</th>
        <th>Name</th>
        <th>Description</th>
        <th>BusinessId</th>
        <th>ThreatId</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {attackChains?.map(attackChain => {
        counter = counter + 1;    
        const updateAttackChainInput: AttackChain = {id: attackChain.id, description: attackChain.description ,
          name: attackChain.name, businessId: attackChain.businessId, threatId: attackChain.threatId,
          createdAt: attackChain.createdAt}
        const deleteAttackChainInput = {id: attackChain.id}
        return (
          <tr key={attackChain.id}>
            <td><a href={`http://localhost:3000/attackChain?id=${attackChain.id}`}>{counter}</a></td>
            <td>{attackChain.name}</td>
            <td>{attackChain.description}</td>
            <td>{attackChain.businessId}</td>
            <td>{attackChain.threatId}</td>
            <td>{attackChain.createdAt?.toString()}</td>
            <td><UpdateAttackChain {...updateAttackChainInput} /></td>
            <td><DeleteAttackChain {...deleteAttackChainInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListAttackChains