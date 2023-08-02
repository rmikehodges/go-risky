import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Resource  from "./Resource";


const GetResource = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [resource, setResource] = useState<Resource | null>(null);

    useEffect(() => {
      axios.get<Resource>(`http://localhost:8081/v0/resource?id=${id}`)
        .then(res => {
        const resourceRes = res.data;
       setResource(resourceRes);})
      }, [id]);
  return (

    <div>
        {resource?.id}
    </div>
  );
};

export default GetResource