import axios from "axios";
import { useEffect, useState } from "react";
import  Liability  from "./Liability";


const GetLiability = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [liability, setLiability] = useState<Liability | null>(null);

    useEffect(() => {
      axios.get<Liability>(`http://localhost:8081/liability?id=${id}`)
        .then(res => {
        const liabilityRes = res.data;
       setLiability(liabilityRes);})
      }, [id]);
  return (

    <div>
        {liability?.id}
    </div>
  );
};

export default GetLiability