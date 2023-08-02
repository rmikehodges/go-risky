import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Threat  from "./Threat";


const GetThreat = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [threat, setThreat] = useState<Threat | null>(null);

    useEffect(() => {
      axios.get<Threat>(`http://localhost:8081/v0/threat?id=${id}`)
        .then(res => {
        const threatRes = res.data;
       setThreat(threatRes);})
      }, [id]);
  return (

    <div>
        {threat?.id}
    </div>
  );
};

export default GetThreat