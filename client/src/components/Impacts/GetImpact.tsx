import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Impact  from "./Impact";


const GetImpact = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [impact, setImpact] = useState<Impact | null>(null);

    useEffect(() => {
      axios.get<Impact>(`http://localhost:8081/impact?id=${id}`)
        .then(res => {
        const impactRes = res.data;
       setImpact(impactRes);})
      }, [id]);
  return (

    <div>
        {impact?.id}
    </div>
  );
};

export default GetImpact