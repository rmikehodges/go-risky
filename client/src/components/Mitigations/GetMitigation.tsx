import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Mitigation  from "./Mitigation";


const GetMitigation = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [mitigation, setMitigation] = useState<Mitigation | null>(null);

    useEffect(() => {
      axios.get<Mitigation>(`http://localhost:8081/v0/mitigation?id=${id}`)
        .then(res => {
        const mitigationRes = res.data;
       setMitigation(mitigationRes);})
      }, [id]);
  return (

    <div>
        {mitigation?.id}
    </div>
  );
};

export default GetMitigation