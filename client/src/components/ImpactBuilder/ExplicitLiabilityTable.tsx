

const ExplicitLiabilityTable = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    
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
                    <td>Additional Headcount</td>
                </tr>
                <tr>
                    <td>Consulting</td>
                </tr>
                <tr>
                    <td>Consulting</td>
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


export default ExplicitLiabilityTable;