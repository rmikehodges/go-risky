import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import {CreateExposureDocument,CreateExposureInput, ExposureInput} from '../../gql/graphql'
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

const CreateExposureMutation = graphql(`
mutation CreateExposure($input: CreateExposureInput!) {
  createExposure(input: $input) {
    clientMutationId
    exposure {
      id
    }
  }
  
  }
`);



const CreateExposure = () => {
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
    const [result, executeMutation] = useMutation(CreateExposureDocument);
    const initialValues = {name: ""}

   return (
      <div id='create-exposure'>
        <button onClick={openModal}>
        Create Exposure
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Exposure'
      >
        <h2>Create Exposure</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: ExposureInput,
            { setSubmitting }: FormikHelpers<ExposureInput> 
          ) => {
              const submittedValues: CreateExposureInput = {clientMutationId: "CreateExposure", exposure: values }
              executeMutation({input: submittedValues}).then(() => {
                console.log(result.data?.createExposure)
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

export default CreateExposure