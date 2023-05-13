import { useMutation } from 'urql';
import { graphql } from '../../gql';
import {DeleteBusinessDocument, DeleteBusinessInput} from '../../gql/graphql'
import { UUID } from 'crypto';



const DeleteBusinessMutation = graphql(`
mutation DeleteBusiness ($input: DeleteBusinessInput!) {
  deleteBusiness(input: $input) {
    business {
        id
    }
  }
}`)


const DeleteBusiness = (deleteBusinessInput:DeleteBusinessInput) => {
    const [result, executeMutation] = useMutation(DeleteBusinessDocument);
    const submitDeleteBusinessInput: DeleteBusinessInput = {clientMutationId: "tester", id: deleteBusinessInput.id}

    return (
      <button onClick={() => executeMutation({input: submitDeleteBusinessInput})}>
        Delete
      </button>
    )
};

export default DeleteBusiness