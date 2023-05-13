import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import {CreateResourceInput, ResourceInput, ResourcesType, CreateResourceDocument} from '../../gql/graphql'
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


const CreateResourceMutation = graphql(`
mutation CreateResource($input: CreateResourceInput!) {
  createResource(input: $input) {
    clientMutationId
    resource {
      id
    }
  }
  
  }
`);



const CreateResource = () => {
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
    const [result, executeMutation] = useMutation(CreateResourceDocument);
    const initialValues = {name: "", cost: 0, total: 0, unit: "sprint", resourceType: ResourcesType.Cash}

   return (
      <div id='create-resource'>
        <button onClick={openModal}>
        Create Resource
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Resource'
      >
        <h2>Create Resource</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: ResourceInput,
            { setSubmitting }: FormikHelpers<ResourceInput> 
          ) => {
              const submittedValues: CreateResourceInput = {clientMutationId: "CreateResource", resource: values }
              executeMutation({input: submittedValues}).then(() => {
                console.log(result.data?.createResource)
              })

              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />

          <label htmlFor="Total">Total</label>
          <Field id="total" total="total" />

          <label htmlFor="cost">Cost</label>
          <Field id="cost" cost="cost" />


           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateResource