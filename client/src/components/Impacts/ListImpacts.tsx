import axios from "axios";
import { useEffect, useState } from "react";
import CreateImpact from "./CreateImpact";
import UpdateImpact from "./UpdateImpact";
import DeleteImpact from "./DeleteImpact";
import  Impact  from "./Impact";


const ListImpacts = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [impacts, setImpacts] = useState<Impact[] | null>(null);

    useEffect(() => {
      axios.get<Impact[]>(`http://localhost:8081/impacts?businessId=${businessId}`)
        .then(res => {
        const impactsResp = res.data;
       setImpacts(impactsResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Impacts'>
      <div>
        <CreateImpact />
      </div>
    <table>
      <thead><tr>
        <th>#</th>
        <th>Name</th>
        <th>Description</th>
        <th>BusinessId</th>
        <th>ThreatId</th>
        <th>Exploitation Cost</th>
        <th>Mitigation Cost</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {impacts?.map(impact => {
        counter = counter + 1;    
        const updateImpactInput: Impact = {id: impact.id, description: impact.description ,
          name: impact.name, businessId: impact.businessId, threatId: impact.threatId, exploitationCost: impact.exploitationCost,
          mitigationCost: impact.mitigationCost,createdAt: impact.createdAt}
        const deleteImpactInput = {id: impact.id}
        return (
          <tr key={impact.id}>
            <td><a href={`http://localhost:3000/impact?id=${impact.id}`}>{counter}</a></td>
            <td>{impact.name}</td>
            <td>{impact.description}</td>
            <td>{impact.businessId}</td>
            <td>{impact.threatId}</td>
            <td>{impact.exploitationCost}</td>
            <td>{impact.mitigationCost}</td>
            <td>{impact.createdAt?.toString()}</td>
            <td><UpdateImpact {...updateImpactInput} /></td>
            <td><DeleteImpact {...deleteImpactInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListImpacts