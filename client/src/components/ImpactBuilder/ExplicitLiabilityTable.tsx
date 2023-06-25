import React from 'react';
import { LiabilityOutput } from '../Liabilities/Liability';
import { useEffect, useState } from "react";


interface ExplicitLiabilityTableProps {
    explicitLiabilities: LiabilityOutput[] | null;
  }
  
const ExplicitLiabilityTable: React.FC<ExplicitLiabilityTableProps> = ({ explicitLiabilities }) => {
    const [complianceTotal, setComplianceTotal] = useState<number>(0);
    const [complianceRemediationTotal, setComplianceRemediationTotal] = useState<number>(0);
    const [legalTotal, setLegalTotal] = useState<number>(0);
    const [legalRemediationTotal, setLegalRemediationTotal] = useState<number>(0);
    const [headcountTotal, setHeadcountTotal] = useState<number>(0);
    const [headcountRemediationTotal, setHeadcountRemediationTotal] = useState<number>(0);
    const [consultingTotal, setConsultingTotal] = useState<number>(0);
    const [consultingRemediationTotal, setConsultingRemediationTotal] = useState<number>(0);
    const [cashTotal, setCashTotal] = useState<number>(0);
    const [cashRemediationTotal, setCashRemediationTotal] = useState<number>(0);
    const [otherTotal, setOtherTotal] = useState<number>(0);
    const [otherRemediationTotal, setOtherRemediationTotal] = useState<number>(0);

    useEffect(() => {
        if (explicitLiabilities != null) {
        for (let i = 0; i < explicitLiabilities!.length; i++) {
            if (explicitLiabilities![i].mitigationId != null || explicitLiabilities![i].detectionId != null) {
                if (explicitLiabilities![i].resourceType === 'COMPLIANCE') {
                    setComplianceRemediationTotal(complianceRemediationTotal + explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'LEGAL') {
                    setLegalRemediationTotal(legalRemediationTotal + explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'EMPLOYEE') {
                    setHeadcountRemediationTotal(headcountRemediationTotal + explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'CONSULTING') {
                    setConsultingRemediationTotal(consultingRemediationTotal + explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'CASH') {
                    setCashRemediationTotal(cashRemediationTotal + explicitLiabilities![i].cost);
                } else {
                    setOtherRemediationTotal(otherRemediationTotal + explicitLiabilities![i].cost);
                }

            } else {
                if (explicitLiabilities![i].resourceType === 'COMPLIANCE') {
                    setComplianceTotal(complianceTotal + explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'LEGAL') {
                    setLegalTotal(legalTotal + explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'EMPLOYEE') {
                    setHeadcountTotal(headcountTotal + explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'CONSULTING') {
                    setConsultingTotal(consultingTotal + explicitLiabilities![i].cost);
                } else if (explicitLiabilities![i].resourceType === 'CASH') {
                    setCashTotal(cashTotal + explicitLiabilities![i].cost);
                } else {
                    setOtherTotal(otherTotal + explicitLiabilities![i].cost);
                }
            }

        }
    }
    });

    
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
            <td>
                <tr>
                    <td>Compliance</td>
                </tr>
                <tr>
                    <td>Legal</td>
                </tr>
                <tr>
                    <td>Headcount</td>
                </tr>
                <tr>
                    <td>Consulting</td>
                </tr>
                <tr>
                    <td>Cash</td>
                </tr>
                                <tr>
                    <td>Other</td>
                </tr>
            </td>
            <td>
                <tr>
                    <td>{complianceTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{legalTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{headcountTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{consultingTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{cashTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{otherTotal}</td>
                    <td>+</td>
                </tr>
            </td>
            <td>
                <tr>
                    <td>{complianceRemediationTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{legalRemediationTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{headcountRemediationTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{consultingRemediationTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{cashRemediationTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{otherRemediationTotal}</td>
                    <td>+</td>
                </tr>
            </td>
            </tbody>
        </table>
        </div>
    );
};


export default ExplicitLiabilityTable;