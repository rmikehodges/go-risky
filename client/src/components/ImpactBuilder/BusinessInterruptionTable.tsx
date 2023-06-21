

const BusinessInterruptionTable = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    
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
                    <td>Saved Expenses</td>
                </tr>
                <tr>
                    <td>Lost Profit</td>
                </tr>
            </td>
            <td>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
            </td>
            <td>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
                <tr>
                    <td>0</td>
                    <td>+</td>
                </tr>
            </td>
            </tbody>
        </table>
        </div>
    );
};


export default BusinessInterruptionTable;