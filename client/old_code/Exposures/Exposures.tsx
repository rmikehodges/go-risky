import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { ExposuresDocument } from '../../gql/graphql';
import CreateExposure from './CreateExposure';
import DeleteExposure from './DeleteExposure';
import UpdateExposure from './UpdateExposure';

const ExposuresQuery = graphql(`
query Exposures {
    exposures {
      edges {
        node {
          id
    
        }
      }
    }
  }
  
`);

const Exposures = () => {
  const [{ data }] = useQuery({ query: ExposuresDocument });

  return (
    <div className='Exposures'>
      <div>
        <CreateExposure />
      </div>
    <table>
      <thead><tr>
        <th>ID</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {data?.exposures?.edges.map(exposure => {
        const updateExposureInput = {id: exposure.node?.id, patch:{}}
        const deleteExposureInput = {id: exposure.node?.id}
        return (
          <tr key={exposure.node?.id}>
            <td>{exposure.node?.id}</td>
            <td><UpdateExposure {...updateExposureInput} /></td>
            <td><DeleteExposure {...deleteExposureInput} /></td>
          </tr>)
      })}
    </tbody>
    </table>
    </div>
  );
};

export default Exposures