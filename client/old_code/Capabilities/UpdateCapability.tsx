import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import ReactDOM from 'react-dom';
import {UpdateCapabilityDocument, UpdateCapabilityInput, CapabilityPatch} from '../../gql/graphql'
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID, randomUUID } from 'crypto';

const customStyles = {
  content: {
    top: '50%',
    left: '50%',
    right: 'auto',
    bottom: 'auto',
    marginRight: '-50%',
    transform: 'translate(-50%, -50%)',
  },
};


Modal.setAppElement('#root');





const UpdateCapabilityMutation = graphql(`
  mutation UpdateCapability($input: UpdateCapabilityInput!)  {
  updateCapability(input: $input) {
    clientMutationId
    capability {
        id
    }
  }
}`)


const UpdateCapability = (updateCapabilityInput:UpdateCapabilityInput) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);

   const openModal = () => {
    setIsOpen(true)
   }

   const afterOpenModal = () => {
    // references are now sync'd and can be accessed.
  }

  const closeModal = () => {
    setIsOpen(false);
  }
   

    const [result, executeMutation] = useMutation(UpdateCapabilityDocument);
    const initialValues: CapabilityPatch = updateCapabilityInput.patch;

    return (
      <div id='update-capability'>
        <button onClick={openModal}>
        Update Capability
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Capability'
      >
        <h2>Update Capability</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: CapabilityPatch,
            { setSubmitting }: FormikHelpers<CapabilityPatch> 
          ) => {
              const submittedValues: UpdateCapabilityInput = {id: updateCapabilityInput.id, clientMutationId: "tester", patch: values }
              executeMutation({input: submittedValues})
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateCapabilityInput.patch.name} />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" placeholder={updateCapabilityInput.patch.businessId} />

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateCapability