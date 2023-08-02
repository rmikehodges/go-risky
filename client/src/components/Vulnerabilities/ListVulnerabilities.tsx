import axios from "axios";
import { useEffect, useState } from "react";
import CreateVulnerability from "./CreateVulnerability";
import UpdateVulnerability from "./UpdateVulnerability";
import DeleteVulnerability from "./DeleteVulnerability";
import  Vulnerability  from "./Vulnerability";
import { UUID } from 'crypto';


const ListVulnerabilities = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [vulnerabilities, setVulnerabilities] = useState<Vulnerability[] | null>(null);

    useEffect(() => {
      axios.get<Vulnerability[]>(`http://localhost:8081/v0/vulnerabilities?businessId=${businessId}`)
        .then(res => {
        const vulnerabilitiesResp = res.data;
       setVulnerabilities(vulnerabilitiesResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Vulnerabilities'>
      <div>
        <CreateVulnerability />
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
      {vulnerabilities?.map(vulnerability => {
        counter = counter + 1;    
        const updateVulnerabilityInput: Vulnerability = {id: vulnerability.id, description: vulnerability.description ,name: vulnerability.name, businessId: vulnerability.businessId, createdAt: vulnerability.createdAt}
        const deleteVulnerabilityInput = {id: vulnerability.id}
        return (
          <tr key={vulnerability.id}>
            <td><a href={`http://localhost:3000/vulnerability?id=${vulnerability.id}`}>{counter}</a></td>
            <td>{vulnerability.name}</td>
            <td>{vulnerability.description}</td>
            <td>{vulnerability.businessId}</td>
            <td>{vulnerability.createdAt?.toString()}</td>
            <td><UpdateVulnerability {...updateVulnerabilityInput} /></td>
            <td><DeleteVulnerability {...deleteVulnerabilityInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListVulnerabilities