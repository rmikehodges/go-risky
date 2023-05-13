import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import ReactDOM from 'react-dom';
import {UpdateBusinessDocument, UpdateBusinessInput, BusinessPatch, BusinessInput} from '../../gql/graphql'
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



const UpdateBusinessMutation = graphql(`
  mutation UpdateBusiness($input: UpdateBusinessInput!)  {
  updateBusiness(input: $input) {
    clientMutationId
    business {
        id
        revenue
    }
  }
}`)


const UpdateBusiness = (updateBusinessInput:UpdateBusinessInput) => {
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
   

    const [result, executeMutation] = useMutation(UpdateBusinessDocument);
    const initialValues: BusinessPatch = updateBusinessInput.patch;

    return (
      <div id='update-business'>
        <button onClick={openModal}>
        Update Business
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Business'
      >
        <h2>Update Business</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: BusinessPatch,
            { setSubmitting }: FormikHelpers<BusinessPatch> 
          ) => {
              const submittedValues: UpdateBusinessInput = {id: updateBusinessInput.id, clientMutationId: "tester", patch: values }
              executeMutation({input: submittedValues})
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateBusinessInput.patch.name} />

          <label htmlFor="revenue">revenue</label>
          <Field id="revenue" name="revenue" placeholder={updateBusinessInput.patch.revenue} />

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateBusiness