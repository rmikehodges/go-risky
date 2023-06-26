import React from 'react';
import { useEffect, useState, useRef } from "react";
import { LiabilityOutput } from '../Liabilities/Liability';

interface BusinessInterruptionTableProps {
    businessInterruptionLiabilities: LiabilityOutput[] | null;
  }

const BusinessInterruptionTable: React.FC<BusinessInterruptionTableProps> = ({ businessInterruptionLiabilities}) => {
    let employeeTotal = 0;
    let employeeRemediationTotal = 0;
    let overtimeTotal = 0;
    let overtimeRemediationTotal = 0;
    let revenueTotal = 0;
    let revenueRemediationTotal = 0;
    let otherTotal = 0;
    let otherRemediationTotal = 0;



        if (businessInterruptionLiabilities != null) {
        for (let i = 0; i < businessInterruptionLiabilities!.length; i++) {
            if (businessInterruptionLiabilities![i].mitigationId != null || businessInterruptionLiabilities![i].detectionId != null) {
                if (businessInterruptionLiabilities![i].resourceType === 'EMPLOYEE') {
                    (employeeRemediationTotal += businessInterruptionLiabilities![i].cost);
                }
                else if (businessInterruptionLiabilities![i].resourceType === 'OVERTIME') {
                    (overtimeRemediationTotal += businessInterruptionLiabilities![i].cost);
                }
                else if (businessInterruptionLiabilities![i].resourceType === 'REVENUE') {
                    (revenueRemediationTotal += businessInterruptionLiabilities![i].cost);
                } else {
                    (otherRemediationTotal += businessInterruptionLiabilities![i].cost);
                }
            } else {
                if (businessInterruptionLiabilities![i].resourceType === 'EMPLOYEE') {
                    (employeeTotal += businessInterruptionLiabilities![i].cost);
                }
                else if (businessInterruptionLiabilities![i].resourceType === 'OVERTIME') {
                    (overtimeTotal += businessInterruptionLiabilities![i].cost);
                }
                else if (businessInterruptionLiabilities![i].resourceType === 'REVENUE') {
                    (revenueTotal += businessInterruptionLiabilities![i].cost);
                } else {
                    (otherTotal += businessInterruptionLiabilities![i].cost);
                }
            }
        }
    } 

    return (
        <div className='BusinessInterruptionTable'>
        <table>
        <thead> 
            <tr>
                <th>Business Interruption Loss</th>
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
                    <td>Employee Time</td>
                </tr>
                <tr>
                    <td>Overtime</td>
                </tr>
                <tr>
                    <td>Lost Revenue</td>
                </tr>
                <tr>
                    <td>Other</td>
                </tr>
                {/* <tr>
                    <td>Lost Profit</td>
                </tr> */}
            </td>
            <td>
                <tr>
                    <td>{employeeTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{overtimeTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{revenueTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{otherTotal}</td>
                    <td>+</td>
                </tr>
                {/* <tr>
                    <td>0</td>
                    <td>+</td>
                </tr> */}
            </td>
            <td>
                <tr>
                    <td>{employeeRemediationTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{overtimeRemediationTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{revenueRemediationTotal}</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>{otherRemediationTotal}</td>
                    <td>+</td>
                </tr>
                {/* <tr>
                    <td>0</td>
                    <td>+</td>
                </tr> */}
            </td>
            </tbody>
        </table>
        </div>
    );
};


export default BusinessInterruptionTable;