import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  AttackChainStep  from "./AttackChainStep";


const GetAttackChainStep = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [attackChainStepStep, setAttackChainStep] = useState<AttackChainStep | null>(null);

    useEffect(() => {
      axios.get<AttackChainStep>(`http://localhost:8081/v0/attackChainStepStep?id=${id}`)
        .then(res => {
        const attackChainStepStepRes = res.data;
       setAttackChainStep(attackChainStepStepRes);})
      }, [id]);
  return (

    <div>
        {attackChainStepStep?.id}
    </div>
  );
};

export default GetAttackChainStep