import axios from "axios";
import { useEffect, useState } from "react";
import CreateAsset from "./CreateAsset";
import UpdateAsset from "./UpdateAsset";
import DeleteAsset from "./DeleteAsset";
import  Asset  from "./Asset";
import { UUID } from 'crypto';


const ListAssets = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId = queryParameters.get("businessId")
    const [assets, setAssets] = useState<Asset[] | null>(null);

    useEffect(() => {
      axios.get<Asset[]>(`http://localhost:8081/assets?businessId=${businessId}`)
        .then(res => {
        const assetsResp = res.data;
       setAssets(assetsResp)})
      }, [businessId]);


  var counter = 0;
  return (
    <div className='Assets'>
      <div>
        <CreateAsset />
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
      {assets?.map(asset => {
        counter = counter + 1;    
        const updateAssetInput: Asset = {id: asset.id, description: asset.description ,name: asset.name, businessId: asset.businessId, createdAt: asset.createdAt}
        const deleteAssetInput = {id: asset.id}
        return (
          <tr key={asset.id}>
            <td><a href={`http://localhost:3000/asset?id=${asset.id}`}>{counter}</a></td>
            <td>{asset.name}</td>
            <td>{asset.description}</td>
            <td>{asset.businessId}</td>
            <td>{asset.createdAt?.toString()}</td>
            <td><UpdateAsset {...updateAssetInput} /></td>
            <td><DeleteAsset {...deleteAssetInput} /></td>
          </tr>)
          
      })}
    </tbody>
    </table>
    </div>

  );
};

export default ListAssets