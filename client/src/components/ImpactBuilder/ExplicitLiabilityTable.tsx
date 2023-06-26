import React from 'react';
import { LiabilityOutput } from '../Liabilities/Liability';
import { useEffect, useState, useRef } from "react";


interface ExplicitLiabilityTableProps {
    explicitLiabilities: LiabilityOutput[] | null;
  }
  
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