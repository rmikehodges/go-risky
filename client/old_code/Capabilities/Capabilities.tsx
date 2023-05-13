import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { CapabilitiesDocument } from '../../gql/graphql';
import CreateCapability from './CreateCapability';
import DeleteCapability from './DeleteCapability';
import UpdateCapability from './UpdateCapability';

const CapabilitiesQuery = graphql(`
query Capabilities {
    capabilities {
      edges {
        node {
          id
    
        }
      }
    }
  }
  
`);

const Capabilities = () => {
  const [{ data }] = useQuery({ query: CapabilitiesDocument });

  return (
    <div className='Capabilitys'>
      <div>
        <CreateCapability />
      </div>
    <table>
      <thead><tr>
        <th>ID</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {data?.capabilities?.edges.map(capability => {
        const updateCapabilityInput = {id: capability.node?.id, patch:{}}
        const deleteCapabilityInput = {id: capability.node?.id}
        return (
          <tr key={capability.node?.id}>
            <td>{capability.node?.id}</td>
            <td><UpdateCapability {...updateCapabilityInput} /></td>
            <td><DeleteCapability {...deleteCapabilityInput} /></td>
          </tr>)
      })}
    </tbody>
    </table>
    </div>
  );
};

export default Capabilities