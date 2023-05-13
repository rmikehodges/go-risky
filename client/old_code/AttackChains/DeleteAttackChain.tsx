import { useMutation } from 'urql';
import { graphql } from '../../gql';
import {DeleteAttackChainDocument, DeleteAttackChainInput} from '../../gql/graphql'
import { UUID } from 'crypto';




const DeleteAttackChainMutation = graphql(`
mutation DeleteAttackChain ($input: DeleteAttackChainInput!) {
  deleteAttackChain(input: $input) {
    attackChain {
        id
    }
  }
}`)


const DeleteAttackChain = (deleteAttackChainInput:DeleteAttackChainInput) => {
    const [result, executeMutation] = useMutation(DeleteAttackChainDocument);
    const submitDeleteAttackChainInput: DeleteAttackChainInput = {clientMutationId: "DeleteAttackChain", id: deleteAttackChainInput.id}

    return (
      <button onClick={() => executeMutation({input: submitDeleteAttackChainInput})}>
        Delete
      </button>
    )
};

export default DeleteAttackChain