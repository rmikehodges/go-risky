import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { UpdateVulnerabilityInput, VulnerabilitiesDocument } from '../../gql/graphql';
import CreateVulnerability from './CreateVulnerability';
import DeleteVulnerability from './DeleteVulnerability';
import UpdateVulnerability from './UpdateVulnerability';

const VulnerabilitiesQuery = graphql(`
query Vulnerabilities {
    vulnerabilities {
      edges {
        node {
          id
          name
        }
      }
    }
  }
  
`);

const Vulnerabilities = () => {
  const [{ data }] = useQuery({ query: VulnerabilitiesDocument });

  return (
    <div className='Vulnerabilities'>
      <div>
        <CreateVulnerability />
      </div>
    <table>
      <thead><tr>
        <th>ID</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {data?.vulnerabilities?.edges.map(vulnerability => {
        const updateVulnerabilityInput: UpdateVulnerabilityInput = {id: vulnerability.node?.id, patch: {name: vulnerability.node?.name }}
        const deleteVulnerabilityInput = {id: vulnerability.node?.id}
        return (
          <tr key={vulnerability.node?.id}>
            <td>{vulnerability.node?.id}</td>
            <td><UpdateVulnerability {...updateVulnerabilityInput} /></td>
            <td><DeleteVulnerability {...deleteVulnerabilityInput} /></td>
          </tr>)
      })}
    </tbody>
    </table>
    </div>
  );
};
export default Vulnerabilities