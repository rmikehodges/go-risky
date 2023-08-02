import axios from "axios";
import { useEffect, useState } from "react";
import CreateCapability from "./CreateCapability";
import UpdateCapability from "./UpdateCapability";
import DeleteCapability from "./DeleteCapability";
import  Capability  from "./Capability";
import { UUID } from 'crypto';


const ListCapabilities = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [capabilities, setCapabilities] = useState<Capability[] | null>(null);

    useEffect(() => {
      axios.get<Capability[]>(`http://localhost:8081/v0/capabilities?businessId=${businessId}`)
        .then(res => {
        const capabilitiesResp = res.data;
       setCapabilities(capabilitiesResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Capabilities'>
      <div>
        <CreateCapability />
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
      {capabilities?.map(capability => {
        counter = counter + 1;    
        const updateCapabilityInput: Capability = {id: capability.id, description: capability.description ,name: capability.name, businessId: capability.businessId, createdAt: capability.createdAt}
        const deleteCapabilityInput = {id: capability.id}
        return (
          <tr key={capability.id}>
            <td><a href={`http://localhost:3000/capability?id=${capability.id}`}>{counter}</a></td>
            <td>{capability.name}</td>
            <td>{capability.description}</td>
            <td>{capability.businessId}</td>
            <td>{capability.createdAt?.toString()}</td>
            <td><UpdateCapability {...updateCapabilityInput} /></td>
            <td><DeleteCapability {...deleteCapabilityInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListCapabilities