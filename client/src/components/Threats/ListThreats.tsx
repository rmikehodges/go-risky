import axios from "axios";
import { useEffect, useState } from "react";
import CreateThreat from "./CreateThreat";
import UpdateThreat from "./UpdateThreat";
import DeleteThreat from "./DeleteThreat";
import  Threat  from "./Threat";
import { UUID } from 'crypto';


const ListThreats = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [threats, setThreats] = useState<Threat[] | null>(null);

    useEffect(() => {
      axios.get<Threat[]>(`http://localhost:8081/v0/threats?businessId=${businessId}`)
        .then(res => {
        const threatsResp = res.data;
       setThreats(threatsResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Threats'>
      <div>
        <CreateThreat />
      </div>
    <table>
      <thead><tr>
        <th>#</th>
        <th>Name</th>
        <th>BusinessId</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {threats?.map(threat => {
        counter = counter + 1;    
        const updateThreatInput: Threat = {id: threat.id, description: threat.description ,name: threat.name, businessId: threat.businessId, createdAt: threat.createdAt}
        const deleteThreatInput = {id: threat.id}
        return (
          <tr key={threat.id}>
            <td><a href={`http://localhost:3000/threat?id=${threat.id}`}>{counter}</a></td>
            <td>{threat.name}</td>
            <td>{threat.description}</td>
            <td>{threat.businessId}</td>
            <td>{threat.createdAt?.toString()}</td>
            <td><UpdateThreat {...updateThreatInput} /></td>
            <td><DeleteThreat {...deleteThreatInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListThreats