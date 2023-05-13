import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { ThreatsDocument, UpdateThreatInput } from '../../gql/graphql';
import CreateThreat from './CreateThreat';
import DeleteThreat from './DeleteThreat';
import UpdateThreat from './UpdateThreat';

const ThreatsQuery = graphql(`
query Threats {
    threats {
      edges {
        node {
          id
          name
        }
      }
    }
  }
  
`);

const Threats = () => {
  const [{ data }] = useQuery({ query: ThreatsDocument });

  return (
    <div className='Threats'>
      <div>
        <CreateThreat />
      </div>
    <table>
      <thead><tr>
        <th>ID</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {data?.threats?.edges.map(threat => {
        const updateThreatInput: UpdateThreatInput = {id: threat.node?.id, patch: {name: threat.node?.name }}
        const deleteThreatInput = {id: threat.node?.id}
        return (
          <tr key={threat.node?.id}>
            <td>{threat.node?.id}</td>
            <td><UpdateThreat {...updateThreatInput} /></td>
            <td><DeleteThreat {...deleteThreatInput} /></td>
          </tr>)
      })}
    </tbody>
    </table>
    </div>
  );
};

export default Threats