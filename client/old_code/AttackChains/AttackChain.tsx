import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { useParams } from "react-router";
import { AttackChainDocument } from '../../gql/graphql';

const AttackChainQuery = graphql(`
query AttackChain($id: UUID!) {
  attackChain(id: $id) {
    attackChainActions(orderBy: POSITION_ASC) {
      edges {
        node {
          position
          action {
            id
          }
        }
      }
    }
  }
  }
  
`);

const AttackChain = () => {
    let { id } = useParams();
    const [{ data }] = useQuery({query: AttackChainDocument, variables: {id: id}});
    let actions = "";

    data?.attackChain?.attackChainActions.edges.map(attackChainAction => {
      actions.concat(attackChainAction?.node?.action?.id)
    })

  return (
    <div>
      {actions}
    </div>
  );
};

export default AttackChain