import axios from "axios";
import { useEffect, useState } from "react";
import { UUID } from 'crypto';
import BusinessInterruptionTable from "./BusinessInterruptionTable";
import ExplicitLiabilityTable from "./ExplicitLiabilityTable";
import TotalLiabilityTable from "./TotalLiabilityTable";
import ThreatDropdown from "./ThreatDropdown";
import { ThreatOutput } from "../Threats/Threats";
import { LiabilityOutput } from "../Liabilities/Liability";



const ImpactBuilder = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    const businessId: UUID = "23628819-59dd-45f3-8395-aceeca86bc9c"
    const [threats, setThreats] = useState<ThreatOutput[] | null>(null);
    const [liabilities, setLiabilities] = useState<LiabilityOutput[] | null>(null);
    const [explicitLiabilities, setExplicitLiabilities] = useState<LiabilityOutput[] | null>(null);
    const [businessInterruptionLiabilities, setBusinessInterruptionLiabilities] = useState<LiabilityOutput[] | null>(null);
    const [selectedThreat, setSelectedThreat] = useState<string>('');

    useEffect(() => {
      axios.get<ThreatOutput[]>(`http://localhost:8081/threats?businessId=${businessId}`)
      .then(res => {
     setThreats(res.data)});
      } , [businessId]);


    const handleSelectThreat = (option: string) => {
      setSelectedThreat(option);
      axios.get<LiabilityOutput[]>(`http://localhost:8081/liabilities?businessId=${businessId}&threatId=${option}`).then(res => { 
        setLiabilities(res.data);
      });

      let explicit: LiabilityOutput[] = [];
      let businessInterruption: LiabilityOutput[] = [];
      if (liabilities === null) {
        return;
      }
      for(let i=0; i<liabilities!.length; i++) {

        if (liabilities![i].type === "EXPLICIT") {
          explicit.push(liabilities![i]);
        }
        else if (liabilities![i].type === "BUSINESS INTERRUPTION LOSS") {
          businessInterruption.push(liabilities![i]);
        }
      }
      setExplicitLiabilities(explicit)
      setBusinessInterruptionLiabilities(businessInterruption)
    }



  return (
    <div className='ImpactBuilder'>
      <div className='ImpactBuilder__header'>
        <h1>Impact Builder</h1>
      </div>
      <div className='ImpactBuilder__body'>
      <div className='ImpactBuilder__threat'>
        <h2>Threat</h2>
        <div>
        <ThreatDropdown options={threats} onSelectOption={handleSelectThreat}/>
        </div>
        </div>
        <div className='ImpactBuilder__body__left'>
          <div className='ImpactBuilder__body__left__businessInterruption'>
            <BusinessInterruptionTable businessInterruptionLiabilities={businessInterruptionLiabilities} />
          </div>
          <br></br>
          <div className='ImpactBuilder__body__left__explicitLiability'>
            <ExplicitLiabilityTable explicitLiabilities={explicitLiabilities} /> 
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