import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import {CreateThreatInput, ThreatInput, CreateThreatDocument} from '../../gql/graphql'
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


const CreateThreatMutation = graphql(`
mutation CreateThreat($input: CreateThreatInput!) {
  createThreat(input: $input) {
    clientMutationId
    threat {
      id
    }
  }
  
  }
`);



const CreateThreat = () => {
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
    const [result, executeMutation] = useMutation(CreateThreatDocument);
    const initialValues = {name: ""}

   return (
      <div id='create-threat'>
        <button onClick={openModal}>
        Create Threat
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Threat'
      >
        <h2>Create Threat</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: ThreatInput,
            { setSubmitting }: FormikHelpers<ThreatInput> 
          ) => {
              const submittedValues: CreateThreatInput = {clientMutationId: "CreateThreat", threat: values }
              executeMutation({input: submittedValues}).then(() => {
                console.log(result.data?.createThreat)
              })

              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />


           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateThreat