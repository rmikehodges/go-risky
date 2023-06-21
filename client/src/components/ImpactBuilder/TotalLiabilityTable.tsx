
const TotalLiabilityTable = () => {
    const queryParameters = new URLSearchParams(window.location.search)
    
    return (
        <div className='TotalLiabilityTable'>
        <table>
        <thead> 
            <tr>
                <th>Total Liabilities</th>
                <td>
                0
            </td>
            <td>
                0
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