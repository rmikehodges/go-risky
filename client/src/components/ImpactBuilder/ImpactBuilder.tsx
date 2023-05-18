import axios from "axios";
import { useEffect, useState } from "react";
import { UUID } from 'crypto';



const ImpactBuilder = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId: UUID = "23628819-59dd-45f3-8395-aceeca86bc9c"
    const [impactId, setImpactId] = useState<string | null>(null);

    setImpactId(queryParameters.get("impactId"));



  return (
    <div className='ImpactBuilder'>
      <div>
      </div>
    <table>
      <thead><tr>
        <th></th>
        <th>Exposure</th>
        <th>Remediation</th>
      </tr></thead>
      <tbody>
    </tbody>
    </table>
    </div>

  );
};

export default ImpactBuilder;