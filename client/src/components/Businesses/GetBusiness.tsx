import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Business  from "./Business";


const GetBusiness = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [action, setBusiness] = useState<Business | null>(null);

    useEffect(() => {
      axios.get<Business>(`http://localhost:8081/action?id=${id}`)
        .then(res => {
        const actionRes = res.data;
       setBusiness(actionRes);})
      }, [id]);
  return (

    <div>
        {action?.id}
    </div>
  );
};

export default GetBusiness