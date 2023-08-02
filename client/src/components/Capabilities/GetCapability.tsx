import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Capability  from "./Capability";


const GetCapability = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [capability, setCapability] = useState<Capability | null>(null);

    useEffect(() => {
      axios.get<Capability>(`http://localhost:8081/v0/capability?id=${id}`)
        .then(res => {
        const capabilityRes = res.data;
       setCapability(capabilityRes);})
      }, [id]);
  return (

    <div>
        {capability?.id}
    </div>
  );
};

export default GetCapability