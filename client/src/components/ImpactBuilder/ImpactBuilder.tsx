import axios from "axios";
import { useEffect, useState } from "react";
import { UUID } from 'crypto';
import BusinessInterruptionTable from "./BusinessInterruptionTable";
import ExplicitLiabilityTable from "./ExplicitLiabilityTable";
import TotalLiabilityTable from "./TotalLiabilityTable";



const ImpactBuilder = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    // const businessId: UUID = "23628819-59dd-45f3-8395-aceeca86bc9c"
    // const [impactId, setImpactId] = useState<string | null>(null);

    // setImpactId(queryParameters.get("impactId"));



  return (
    <div className='ImpactBuilder'>
      <div className='ImpactBuilder__header'>
        <h1>Impact Builder</h1>
      </div>
      <div className='ImpactBuilder__body'>
      <div className='ImpactBuilder__threat'>
        <h2>Threat</h2>
        </div>
        <div className='ImpactBuilder__body__left'>
          <div className='ImpactBuilder__body__left__businessInterruption'>
            <BusinessInterruptionTable />
          </div>
          <br></br>
          <div className='ImpactBuilder__body__left__explicitLiability'>
            <ExplicitLiabilityTable /> 
          </div>
          <div className='ImpactBuilder__body__left__totalLiability'>
            <TotalLiabilityTable />
            
          </div>
        </div>
        </div>




      </div>


  );
};

export default ImpactBuilder;