import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import ReactDOM from 'react-dom';
import {UpdateThreatInput,ThreatPatch, ThreatInput, UpdateThreatDocument} from '../../gql/graphql'
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



const UpdateThreatMutation = graphql(`
  mutation UpdateThreat($input: UpdateThreatInput!)  {
  updateThreat(input: $input) {
    clientMutationId
    threat {
        id
    }
  }
}`)


const UpdateThreat = (updateThreatInput:UpdateThreatInput) => {
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
   

    const [result, executeMutation] = useMutation(UpdateThreatDocument);
    const initialValues: ThreatPatch = updateThreatInput.patch;

    return (
      <div id='update-threat'>
        <button onClick={openModal}>
        Update Threat
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Threat'
      >
        <h2>Update Threat</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: ThreatPatch,
            { setSubmitting }: FormikHelpers<ThreatPatch> 
          ) => {
              const submittedValues: UpdateThreatInput = {id: updateThreatInput.id, clientMutationId: "tester", patch: values }
              executeMutation({input: submittedValues})
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateThreatInput.patch.name} />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" placeholder={updateThreatInput.patch.businessId} />
           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateThreat