import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import {CreateBusinessDocument, BusinessInput, CreateBusinessInput} from '../../gql/graphql'
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

const CreateBusinessMutation = graphql(`
mutation CreateBusiness($input: CreateBusinessInput!) {
  createBusiness(input: $input) {
    clientMutationId
    business {
      id
    }
  }
  
  }
`);



const CreateBusiness = () => {
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
    const [result, executeMutation] = useMutation(CreateBusinessDocument);
    const initialValues = {name: "", revenue: 67896789.99}

   return (
      <div id='create-business'>
        <button onClick={openModal}>
        Create Business
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Business'
      >
        <h2>Create Business</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: BusinessInput,
            { setSubmitting }: FormikHelpers<BusinessInput> 
          ) => {
              const submittedValues: CreateBusinessInput = {clientMutationId: "CreateBusiness", business: values }
              executeMutation({input: submittedValues}).then(() => {
                console.log(result.data?.createBusiness)
              })

              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />

          <label htmlFor="revenue">revenue</label>
          <Field id="revenue" name="revenue" />
           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateBusiness