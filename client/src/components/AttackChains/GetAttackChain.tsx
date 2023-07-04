import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  AttackChain  from "./AttackChain";


const GetAttackChain = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [attackChain, setAttackChain] = useState<AttackChain | null>(null);

    useEffect(() => {
      axios.get<AttackChain>(`http://localhost:8081/attackChain?id=${id}`)
        .then(res => {
        const attackChainRes = res.data;
       setAttackChain(attackChainRes);})
      }, [id]);
  return (

    <div>
        {attackChain?.id}
    </div>
  );
};

export default GetAttackChain