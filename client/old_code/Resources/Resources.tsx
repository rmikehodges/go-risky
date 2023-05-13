import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { ResourceInput, ResourcePatch, ResourcesDocument, UpdateResourceInput } from '../../gql/graphql';
import CreateResource from './CreateResource';
import DeleteResource from './DeleteResource';
import UpdateResource from './UpdateResource';

const ResourcesQuery = graphql(`
query Resources {
    resources {
      edges {
        node {
          id
          name
          cost
          unit
          total
          resourceType
        }
      }
    }
  }
  
`);

const Resources = () => {
  const [{ data }] = useQuery({ query: ResourcesDocument });

  return (
    <div className='Resources'>
      <div>
        <CreateResource />
      </div>
    <table>
      <thead><tr>
        <th>ID</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {data?.resources?.edges.map(resource => {
        const updateResourceInput: UpdateResourceInput = {id: resource.node?.id, patch: {name: resource.node?.name,cost: resource.node?.cost,unit: resource.node?.unit , resourceType: resource.node?.resourceType, total: resource.node?.total }}
        const deleteResourceInput = {id: resource.node?.id}
        return (
          <tr key={resource.node?.id}>
            <td>{resource.node?.id}</td>
            <td><UpdateResource {...updateResourceInput} /></td>
            <td><DeleteResource {...deleteResourceInput} /></td>
          </tr>)
      })}
    </tbody>
    </table>
    </div>
  );
};

export default Resources