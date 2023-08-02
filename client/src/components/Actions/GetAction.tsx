import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Action  from "./Action";


const GetAction = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [action, setAction] = useState<Action | null>(null);

    useEffect(() => {
      axios.get<Action>(`http://localhost:8081/v0/action?id=${id}`)
        .then(res => {
        const actionRes = res.data;
       setAction(actionRes);})
      }, [id]);
  return (

    <div>
        {action?.id}
    </div>
  );
};

export default GetAction