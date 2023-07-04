import axios from "axios";
import { useEffect, useState } from "react";
import { UUID } from 'crypto';
import BusinessInterruptionTable from "./BusinessInterruptionTable";
import ExplicitLiabilityTable from "./ExplicitLiabilityTable";
import TotalLiabilityTable from "./TotalLiabilityTable";
import ThreatDropdown from "./ThreatDropdown";
import  Threat from "../Threats/Threat";
import  Liability from "../Liabilities/Liability";



const ImpactBuilder = () => {
    const businessId: UUID = "23628819-59dd-45f3-8395-aceeca86bc9c"
    const [threats, setThreats] = useState<Threat[] | null>(null);
    const [liabilities, setLiabilities] = useState<Liability[] | null>(null);
    const [explicitLiabilities, setExplicitLiabilities] = useState<Liability[] | null>(null);
    const [businessInterruptionLiabilities, setBusinessInterruptionLiabilities] = useState<Liability[] | null>(null);
    const [selectedThreat, setSelectedThreat] = useState<string>('');
    const [total, setTotal] = useState<number>(0);
    const [remediationTotal, setRemediationTotal] = useState<number>(0);


    useEffect(() => {

      axios.get<Threat[]>(`http://localhost:8081/threats?businessId=${businessId}`)
      .then(res => {
     setThreats(res.data)});
     axios.get<Liability[]>(`http://localhost:8081/liabilities?businessId=${businessId}`).then(res => {
        setLiabilities(res.data)}
        );
      } , [businessId]);


    const handleSelectThreat = (option: string) => {
      let localRemediationTotal = 0;
      let localTotal = 0;
        let explicit: Liability[] = [];
        let businessInterruption: Liability[] = [];
        if (liabilities != null) {
          for(let i=0; i<liabilities!.length; i++) {
            if (liabilities![i].threatId == option) {
              if (liabilities![i].mitigationId != null || liabilities![i].detectionId != null) {
                localRemediationTotal += liabilities![i].cost;
              } 
              else {
                localTotal += liabilities![i].cost;
              }

              if (liabilities![i].type === "EXPLICIT") {
                explicit.push(liabilities![i]);
              }
              else if (liabilities![i].type === "BUSINESS INTERRUPTION LOSS") {
                businessInterruption.push(liabilities![i]);
              }
            }
          }
        }
        
          setRemediationTotal(localRemediationTotal);
          setTotal(localTotal);
          setSelectedThreat(option)
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
        <ThreatDropdown options={threats} selectedThreat={selectedThreat} onSelectOption={handleSelectThreat}/>
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
            <TotalLiabilityTable total={total} remediationTotal={remediationTotal}/>
            
          </div>
        </div>
        </div>




      </div>


  );
};

export default ImpactBuilder;