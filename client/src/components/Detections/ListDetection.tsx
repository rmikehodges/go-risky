import axios from "axios";
import { useEffect, useState } from "react";
import CreateDetection from "./CreateDetection";
import UpdateDetection from "./UpdateDetection";
import DeleteDetection from "./DeleteDetection";
import  Detection  from "./Detection";
import { UUID } from 'crypto';


const ListDetections = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [detections, setDetections] = useState<Detection[] | null>(null);

    useEffect(() => {
      axios.get<Detection[]>(`http://localhost:8081/v0/detections?businessId=${businessId}`)
        .then(res => {
        const detectionsResp = res.data;
       setDetections(detectionsResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Detections'>
      <div>
        <CreateDetection />
      </div>
    <table>
      <thead><tr>
        <th>#</th>
        <th>Name</th>
        <th>Description</th>
        <th>BusinessId</th>
        <th>Implemented</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {detections?.map(detection => {
        counter = counter + 1;    
        const updateDetectionInput: Detection = {id: detection.id, description: detection.description ,name: detection.name, businessId: detection.businessId, implemented: detection.implemented, createdAt: detection.createdAt}
        const deleteDetectionInput = {id: detection.id}
        return (
          <tr key={detection.id}>
            <td><a href={`http://localhost:3000/detection?id=${detection.id}`}>{counter}</a></td>
            <td>{detection.name}</td>
            <td>{detection.description}</td>
            <td>{detection.businessId}</td>
            <td>{detection.implemented}</td>
            <td>{detection.createdAt?.toString()}</td>
            <td><UpdateDetection {...updateDetectionInput} /></td>
            <td><DeleteDetection {...deleteDetectionInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListDetections