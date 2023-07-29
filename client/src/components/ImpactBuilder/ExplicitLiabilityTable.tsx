import React from 'react';
import  Liability  from '../Liabilities/Liability';
import { useEffect, useState, useRef } from "react";
import CreateLiability from './CreateLiability';


interface ExplicitLiabilityTableProps {
    explicitLiabilities: Liability[] | null;
  }

const ExplicitLiabilityTypes = ["COMPLIANCE", "LEGAL", "EMPLOYEE", "CONSULTING", "CASH", "OTHER"]
  
const ExplicitLiabilityTable: React.FC<ExplicitLiabilityTableProps> = ({ explicitLiabilities }) => {
    let complianceTotal = 0;
    let complianceRemediationTotal = 0;
    let legalTotal = 0;
    let legalRemediationTotal = 0;
    let headcountTotal = 0;
    let headcountRemediationTotal = 0;
    let consultingTotal = 0;
    let consultingRemediationTotal = 0;
    let cashTotal = 0;
    let cashRemediationTotal = 0;
    let otherTotal = 0;
    let otherRemediationTotal = 0;


        if (explicitLiabilities != null) {
        for (let i = 0; i < explicitLiabilities!.length; i++) {
            if (explicitLiabilities![i].mitigationId != null || explicitLiabilities![i].detectionId != null) {
                if (explicitLiabilities![i].resourceType === 'COMPLIANCE') {
                    (complianceRemediationTotal += explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'LEGAL') {
                    legalRemediationTotal += explicitLiabilities![i].cost;
                } else if (explicitLiabilities![i].resourceType === 'EMPLOYEE') {
                    headcountRemediationTotal += explicitLiabilities![i].cost;
                } else if (explicitLiabilities![i].resourceType === 'CONSULTING') {
                    consultingRemediationTotal += explicitLiabilities![i].cost;
                } else if (explicitLiabilities![i].resourceType === 'CASH') {
                    cashRemediationTotal += explicitLiabilities![i].cost;
                } else {
                    otherRemediationTotal += explicitLiabilities![i].cost;
                }

            } else {
                if (explicitLiabilities![i].resourceType === 'COMPLIANCE') {
                    (complianceTotal += explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'LEGAL') {
                    (legalTotal += explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'EMPLOYEE') {
                    (headcountTotal += explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'CONSULTING') {
                    (consultingTotal += explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'CASH') {
                    (cashTotal += explicitLiabilities![i].cost);
                } else {
                    (otherTotal += explicitLiabilities![i].cost);
                }
            }

        }
    }

    
    return (
        <div className='ExplicitLiabilityTable'>
        <table>
        <thead> 
            <tr>
                <th>Explicit Liabilities</th>
                </tr>
                <tr>
                <th></th>
                <th>Exposure</th>
                <th>Remediation</th>
                </tr>           
        </thead>
        <tbody>
                <tr>
                    <td>Compliance  {complianceTotal} <CreateLiability/>
                    {explicitLiabilities?.filter(liability => liability.resourceType == "COMPLIANCE").map(liability => <li>{liability.name}</li>)}
                    </td>
                </tr>
                <tr>
                    <td>Legal     {legalTotal}         <CreateLiability/>
     
                    {explicitLiabilities?.filter(liability => liability.resourceType == "LEGAL").map(liability => <li>{liability.name}</li>)}
                    </td>
                </tr>
                <tr>
                    <td>Headcount 
                    {explicitLiabilities?.filter(liability => liability.resourceType == "EMPLOYEE").map(liability => <li>{liability.name}</li>)}
                    </td>
                </tr>
                <tr>
                    <td>Consulting {consultingTotal}
                    {explicitLiabilities?.filter(liability => liability.resourceType == "CONSULTING").map(liability => <li>{liability.name}</li>)}
                    </td>
                </tr>
                <tr>
                    <td>Cash
                    {explicitLiabilities?.filter(liability => liability.resourceType == "CASH").map(liability => <li>{liability.name}</li>)}
                    </td>
                </tr>
                                <tr>
                    <td>Other
                    {explicitLiabilities?.filter(liability => !ExplicitLiabilityTypes.includes(liability.resourceType)).map(liability => <li>{liability.name}</li>)}
                    </td>
                </tr>
            <td>
                <tr>
                    <td></td>
                    <td><CreateLiability/></td>
                </tr>
                <tr>
                    <td>{cashTotal}</td>
                    <td><CreateLiability/></td>
                </tr>
                <tr>
                    <td>{otherTotal}</td>
                    <td><CreateLiability/></td>
                </tr>
            </td>
            <td>
                <tr>
                    <td>{complianceRemediationTotal}</td>
                    <td><CreateLiability/></td>
                </tr>
                <tr>
                    <td>{legalRemediationTotal}</td>
                    <td><CreateLiability/></td>
                </tr>
                <tr>
                    <td>{headcountRemediationTotal}</td>
                    <td><CreateLiability/></td>
                </tr>
                <tr>
                    <td>{consultingRemediationTotal}</td>
                    <td><CreateLiability/></td>
                </tr>
                <tr>
                    <td>{cashRemediationTotal}</td>
                    <td><CreateLiability/></td>
                </tr>
                <tr>
                    <td>{otherRemediationTotal}</td>
                    <td><CreateLiability/></td>
                </tr>
            </td>
            </tbody>
        </table>
        </div>
    );
};


export default ExplicitLiabilityTable;