import axios from "axios";
import { useEffect, useState } from "react";
import CreateBusiness from "./CreateBusiness";
import UpdateBusiness from "./UpdateBusiness";
import DeleteBusiness from "./DeleteBusiness";
import  Business  from "./Business";
import { UUID } from 'crypto';


const ListBusinesses = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [businesses, setBusinesses] = useState<Business[] | null>(null);

    useEffect(() => {
      axios.get<Business[]>(`http://localhost:8081/businesses?businessId=${businessId}`)
        .then(res => {
        const businessesResp = res.data;
       setBusinesses(businessesResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Businesses'>
      <div>
        <CreateBusiness />
      </div>
    <table>
      <thead><tr>
        <th>#</th>
        <th>Name</th>
        <th>Revenue</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {businesses?.map(business => {
        counter = counter + 1;    
        const updateBusinessInput: Business = {id: business.id,name: business.name, revenue: business.revenue,  createdAt: business.createdAt}
        const deleteBusinessInput = {id: business.id}
        return (
          <tr key={business.id}>
            <td><a href={`http://localhost:3000/business?id=${business.id}`}>{counter}</a></td>
            <td>{business.id}</td>
            <td>{business.name}</td>
            <td>{business.revenue}</td>
            <td>{business.createdAt?.toString()}</td>
            <td><UpdateBusiness {...updateBusinessInput} /></td>
            <td><DeleteBusiness {...deleteBusinessInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListBusinesses