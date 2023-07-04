import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Asset  from "./Asset";


const GetAsset = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [asset, setAsset] = useState<Asset | null>(null);

    useEffect(() => {
      axios.get<Asset>(`http://localhost:8081/asset?id=${id}`)
        .then(res => {
        const assetRes = res.data;
       setAsset(assetRes);})
      }, [id]);
  return (

    <div>
        {asset?.id}
    </div>
  );
};

export default GetAsset