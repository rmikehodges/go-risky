import { useMutation } from 'urql';
import { graphql } from '../../gql';
import {DeleteResourceInput, DeleteResourceDocument} from '../../gql/graphql'
import { UUID } from 'crypto';



const DeleteResourceMutation = graphql(`
mutation DeleteResource ($input: DeleteResourceInput!) {
  deleteResource(input: $input) {
    resource {
        id
    }
  }
}`)


const DeleteResource = (deleteResourceInput:DeleteResourceInput) => {
    const [result, executeMutation] = useMutation(DeleteResourceDocument);
    const submitDeleteResourceInput: DeleteResourceInput = {clientMutationId: "tester", id: deleteResourceInput.id}

    return (
      <button onClick={() => executeMutation({input: submitDeleteResourceInput})}>
        Delete
      </button>
    )
};

export default DeleteResource