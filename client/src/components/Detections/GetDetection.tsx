import { useParams } from "react-router";
import axios from "axios";
import { useEffect, useState } from "react";
import  Detection  from "./Detection";


const GetDetection = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const id = queryParameters.get("id")
    const [detection, setDetection] = useState<Detection | null>(null);

    useEffect(() => {
      axios.get<Detection>(`http://localhost:8081/detection?id=${id}`)
        .then(res => {
        const detectionRes = res.data;
       setDetection(detectionRes);})
      }, [id]);
  return (

    <div>
        {detection?.id}
    </div>
  );
};

export default GetDetection