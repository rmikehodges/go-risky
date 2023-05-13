import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import {CreateCapabilityDocument, CreateCapabilityInput, CapabilityInput} from '../../gql/graphql'
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID } from 'crypto';

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


const CreateCapabilityMutation = graphql(`
mutation CreateCapability($input: CreateCapabilityInput!) {
  createCapability(input: $input) {
    clientMutationId
    capability {
      id
    }
  }
  
  }
`);



const CreateCapability = () => {
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
    const [result, executeMutation] = useMutation(CreateCapabilityDocument);
    const initialValues = {name: ""}

   return (
      <div id='create-capability'>
        <button onClick={openModal}>
        Create Capability
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Capability'
      >
        <h2>Create Capability</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: CapabilityInput,
            { setSubmitting }: FormikHelpers<CapabilityInput> 
          ) => {
              const submittedValues: CreateCapabilityInput = {clientMutationId: "CreateCapability", capability: values }
              executeMutation({input: submittedValues}).then(() => {
                console.log(result.data?.createCapability)
              })

              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" />

           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateCapability