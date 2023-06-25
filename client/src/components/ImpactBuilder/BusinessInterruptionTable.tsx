import React from 'react';
import { useEffect, useState } from "react";
import { LiabilityOutput } from '../Liabilities/Liability';

interface BusinessInterruptionTableProps {
    businessInterruptionLiabilities: LiabilityOutput[] | null;
  }

const BusinessInterruptionTable: React.FC<BusinessInterruptionTableProps> = ({ businessInterruptionLiabilities}) => {
    const [employeeTotal, setEmployeeTotal] = useState<number>(0);
    const [employeeRemediationTotal, setEmployeeRemediationTotal] = useState<number>(0);
    const [overtimeTotal, setOvertimeTotal] = useState<number>(0);
    const [overtimeRemediationTotal, setOvertimeRemediationTotal] = useState<number>(0);
    const [revenueTotal, setRevenueTotal] = useState<number>(0);
    const [revenueRemediationTotal, setRevenueRemediationTotal] = useState<number>(0);
    const [otherTotal, setOtherTotal] = useState<number>(0);
    const [otherRemediationTotal, setOtherRemediationTotal] = useState<number>(0);


    useEffect(() => {
        if (businessInterruptionLiabilities != null) {
        for (let i = 0; i < businessInterruptionLiabilities!.length; i++) {
            if (businessInterruptionLiabilities![i].mitigationId != null || businessInterruptionLiabilities![i].detectionId != null) {
                if (businessInterruptionLiabilities![i].resourceType === 'EMPLOYEE') {
                    setEmployeeRemediationTotal(employeeRemediationTotal + businessInterruptionLiabilities![i].cost);
                }
                else if (businessInterruptionLiabilities![i].resourceType === 'OVERTIME') {
                    setOvertimeRemediationTotal(overtimeRemediationTotal + businessInterruptionLiabilities![i].cost);
                }
                else if (businessInterruptionLiabilities![i].resourceType === 'REVENUE') {
                    setRevenueRemediationTotal(revenueRemediationTotal + businessInterruptionLiabilities![i].cost);
                } else {
                    setOtherRemediationTotal(otherRemediationTotal + businessInterruptionLiabilities![i].cost);
                }
            } else {
                if (businessInterruptionLiabilities![i].resourceType === 'EMPLOYEE') {
                    setEmployeeTotal(employeeTotal + businessInterruptionLiabilities![i].cost);
                }
                else if (businessInterruptionLiabilities![i].resourceType === 'OVERTIME') {
                    setOvertimeTotal(overtimeTotal + businessInterruptionLiabilities![i].cost);
                }
                else if (businessInterruptionLiabilities![i].resourceType === 'REVENUE') {
                    setRevenueTotal(revenueTotal + businessInterruptionLiabilities![i].cost);
                } else {
                    setOtherTotal(otherTotal + businessInterruptionLiabilities![i].cost);
                }
            }
        }
    } 
    }, [businessInterruptionLiabilities])

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