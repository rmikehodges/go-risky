import { useMutation } from 'urql';
import { graphql } from '../../gql';
import {DeleteCapabilityDocument,DeleteCapabilityInput} from '../../gql/graphql'
import { UUID } from 'crypto';



const DeleteCapabilityMutation = graphql(`
mutation DeleteCapability ($input: DeleteCapabilityInput!) {
  deleteCapability(input: $input) {
    capability {
        id
    }
  }
}`)


const DeleteCapability = (deleteCapabilityInput:DeleteCapabilityInput) => {
    const [result, executeMutation] = useMutation(DeleteCapabilityDocument);
    const submitDeleteCapabilityInput: DeleteCapabilityInput = {clientMutationId: "tester", id: deleteCapabilityInput.id}

    return (
      <button onClick={() => executeMutation({input: submitDeleteCapabilityInput})}>
        Delete
      </button>
    )
};

export default DeleteCapability