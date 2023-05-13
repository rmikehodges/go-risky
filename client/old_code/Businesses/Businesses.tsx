import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { BusinessesDocument } from '../../gql/graphql';
import CreateBusiness from './CreateBusiness';
import UpdateBusiness from './UpdateBusiness';
import DeleteBusiness from './DeleteBusiness';

const BusinessesQuery = graphql(`
query Businesses {
    businesses {
      edges {
        node {
          id
    
        }
      }
    }
  }
  
`);

const Businesses = () => {
  const [{ data }] = useQuery({ query: BusinessesDocument });

  return (
    <div className='Businesss'>
      <div>
        <CreateBusiness />
      </div>
    <table>
      <thead><tr>
        <th>ID</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {data?.businesses?.edges.map(business => {
        const updateBusinessInput = {id: business.node?.id, patch:{}}
        const deleteBusinessInput = {id: business.node?.id}
        return (
          <tr key={business.node?.id}>
            <td>{business.node?.id}</td>
            <td><UpdateBusiness {...updateBusinessInput} /></td>
            <td><DeleteBusiness {...deleteBusinessInput} /></td>
          </tr>)
      })}
    </tbody>
    </table>
    </div>
  );
};


export default Businesses