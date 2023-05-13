import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";

interface ActionOutput {
  id: string
  name: string
  description: string
  capabilityId: string
  vulnerabilityId: string
  businessId: string
  complexity: string
  assetId: string
  createdAt: Date
  }

const Action = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [action, setAction] = useState<ActionOutput | null>(null);

    useEffect(() => {
      axios.get<ActionOutput>(`http://localhost:8081/action?id=${id}`)
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

export default Action