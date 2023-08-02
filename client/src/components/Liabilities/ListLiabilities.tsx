import axios from "axios";
import { useEffect, useState } from "react";
import CreateLiability from "./CreateLiability";
import UpdateLiability from "./UpdateLiability";
import DeleteLiability from "./DeleteLiability";
import  Liability  from "./Liability";


const ListLiabilities = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [liabilities, setLiabilities] = useState<Liability[] | null>(null);

    useEffect(() => {
      axios.get<Liability[]>(`http://localhost:8081/v0/liabilities?businessId=${businessId}`)
        .then(res => {
        const liabilitiesResp = res.data;
       setLiabilities(liabilitiesResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Liabilities'>
      <div>
        <CreateLiability />
      </div>
    <table>
      <thead><tr>
      <th>id</th>
      <th>name</th>
      <th>description</th>
      <th>quantity</th>
      <th>cost</th>
      <th>type</th>
      <th>resourceType</th>
      <th>businessId</th>
      <th>detectionId</th>
      <th>mitigationId</th>
      <th>resourceId</th>
      <th>threatId</th>
      <th>impactId</th>
      <th>createdAt</th>
      </tr></thead>
      <tbody>
      {liabilities?.map(liability => {
        counter = counter + 1;    
        const updateLiabilityInput: Liability = {
          id: liability.id,
          name: liability.name,
          description: liability.description,
          quantity: liability.quantity, 
          cost:  liability.cost, 
          type:  liability.type, 
          resourceType:  liability.resourceType,
          businessId:  liability.businessId, 
          detectionId:  liability.detectionId,
          mitigationId:  liability.mitigationId,
          resourceId:  liability.resourceId, 
          threatId:  liability.threatId,
          impactId:  liability.impactId, 
          createdAt:  liability.createdAt 
        };        const deleteLiabilityInput = {id: liability.id}
        return (
          <tr key={liability.id}>
            <td><a href={`http://localhost:3000/liability?id=${liability.id}`}>{counter}</a></td>
            <td>{liability.id}</td>
            <td>{liability.name}</td>
            <td>{liability.description}</td>
            <td>{liability.quantity}</td>
            <td>{liability.cost}</td>
            <td>{liability.type}</td>
            <td>{liability.resourceType}</td>
            <td>{liability.businessId}</td>
            <td>{liability.detectionId}</td>
            <td>{liability.mitigationId}</td>
            <td>{liability.resourceId}</td>
            <td>{liability.threatId}</td>
            <td>{liability.impactId}</td>
            <td>{liability.createdAt?.toString()}</td>
            <td><UpdateLiability {...updateLiabilityInput} /></td>
            <td><DeleteLiability {...deleteLiabilityInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListLiabilities