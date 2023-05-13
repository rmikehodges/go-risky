import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { AttackChainsDocument } from '../../gql/graphql';
import UpdateAttackChain from './UpdateAttackChain';
import DeleteAttackChain from './DeleteAttackChain';
import CreateAttackChain from './CreateAttackChain';

const AttackChainsQuery = graphql(`
query AttackChains {
  attackChains {
    edges {
      node {
        id
        name
        attackChainActions(orderBy: POSITION_ASC) {
          edges {
            node {
              action {
                id
              }
            }
          }
        }
      }
    }
  }
}
  
`);

const AttackChains = () => {
    const [{ data }] = useQuery({ query: AttackChainsDocument });

    return (
      <div className='Actions'>
        <div>
          <CreateAttackChain />
        </div>
      <table>
        <thead><tr>
          <th>ID</th>
          <th>Actions</th>
          <th>Update</th>
          <th>Delete</th>
        </tr></thead>
        <tbody>
        {data?.attackChains?.edges.map(attackChain => {

          let actions = "";

          attackChain.node?.attackChainActions.edges.map(attackChainAction => {
            actions.concat(attackChainAction.node?.action?.id)
          })


          const updateActionInput = {id: attackChain.node?.id, patch:{}}
          const deleteActionInput = {id: attackChain.node?.id}
          return (
            <tr key={attackChain.node?.id}>
              <td>{attackChain.node?.id}</td>
              <td>{actions}</td>
              <td><UpdateAttackChain {...updateActionInput} /></td>
              <td><DeleteAttackChain {...deleteActionInput} /></td>
            </tr>)
        })}
      </tbody>
      </table>
      </div>
    );
};

export default AttackChains;