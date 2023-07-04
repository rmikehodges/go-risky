import React from 'react';

interface TotalLiabilityTableProps {
    total: number;
    remediationTotal: number;
  }


const TotalLiabilityTable: React.FC<TotalLiabilityTableProps> = ({total,remediationTotal}) => {
    
    return (
        <div className='TotalLiabilityTable'>
        <table>
        <thead> 
            <tr>
                <th>Total Liabilities</th>
                <td>
                {total}
            </td>
            <td>
                {remediationTotal}
            </td>
                </tr>    
        </thead>
        <tbody>

            </tbody>
        </table>
        </div>
    );
};


export default TotalLiabilityTable;