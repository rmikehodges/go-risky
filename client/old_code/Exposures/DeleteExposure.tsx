import { useMutation } from 'urql';
import { graphql } from '../../gql';
import {DeleteExposureDocument, DeleteExposureInput} from '../../gql/graphql'
import { UUID } from 'crypto';



const DeleteExposureMutation = graphql(`
mutation DeleteExposure ($input: DeleteExposureInput!) {
  deleteExposure(input: $input) {
    exposure {
        id
    }
  }
}`)


const DeleteExposure = (deleteExposureInput:DeleteExposureInput) => {
    const [result, executeMutation] = useMutation(DeleteExposureDocument);
    const submitDeleteExposureInput: DeleteExposureInput = {clientMutationId: "tester", id: deleteExposureInput.id}

    return (
      <button onClick={() => executeMutation({input: submitDeleteExposureInput})}>
        Delete
      </button>
    )
};

export default DeleteExposure