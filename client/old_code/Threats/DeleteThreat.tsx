import { useMutation } from 'urql';
import { graphql } from '../../gql';
import {DeleteThreatInput, DeleteThreatDocument} from '../../gql/graphql'
import { UUID } from 'crypto';


const DeleteThreatMutation = graphql(`
mutation DeleteThreat ($input: DeleteThreatInput!) {
  deleteThreat(input: $input) {
    threat {
        id
    }
  }
}`)


const DeleteThreat = (deleteThreatInput:DeleteThreatInput) => {
    const [result, executeMutation] = useMutation(DeleteThreatDocument);
    const submitDeleteThreatInput: DeleteThreatInput = {clientMutationId: "tester", id: deleteThreatInput.id}

    return (
      <button onClick={() => executeMutation({input: submitDeleteThreatInput})}>
        Delete
      </button>
    )
};

export default DeleteThreat