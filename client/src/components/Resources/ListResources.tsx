import axios from "axios";
import { useEffect, useState } from "react";
import CreateResource from "./CreateResource";
import UpdateResource from "./UpdateResource";
import DeleteResource from "./DeleteResource";
import  Resource  from "./Resource";
import { UUID } from 'crypto';


const ListResources = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [resources, setResources] = useState<Resource[] | null>(null);

    useEffect(() => {
      axios.get<Resource[]>(`http://localhost:8081/resources?businessId=${businessId}`)
        .then(res => {
        const resourcesResp = res.data;
       setResources(resourcesResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Resources'>
      <div>
        <CreateResource />
      </div>
    <table>
      <thead><tr>
      <th>id</th>
      <th>name</th>
      <th>description</th>
      <th>cost</th>
      <th>unit</th>
      <th>total</th>
      <th>resourceType</th>
      <th>businessId</th>
      <th>createdAt</th>
      </tr></thead>
      <tbody>
      {resources?.map(resource => {
        counter = counter + 1;    
        const updateResourceInput: Resource = {
          id: resource.id,
          name: resource.name,
          description: resource.description,
          cost: resource.cost,
          unit: resource.unit,
          total: resource.total,
          resourceType: resource.resourceType,
          businessId: resource.businessId,
          createdAt: resource.createdAt
        };        
        const deleteResourceInput = {id: resource.id}
        return (
          <tr key={resource.id}>
            <td><a href={`http://localhost:3000/resource?id=${resource.id}`}>{counter}</a></td>
            <td>{resource.name}</td>
            <td>{resource.description}</td>
            <td>{resource.cost}</td>
            <td>{resource.unit}</td>
            <td>{resource.total}</td>
            <td>{resource.resourceType}</td>
            <td>{resource.businessId}</td>
            <td>{resource.createdAt?.toString()}</td>
            <td><UpdateResource {...updateResourceInput} /></td>
            <td><DeleteResource {...deleteResourceInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListResources