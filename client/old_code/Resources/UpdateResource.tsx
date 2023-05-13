import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import ReactDOM from 'react-dom';
import {ResourceInput, UpdateResourceInput, UpdateResourceDocument, ResourcePatch} from '../../gql/graphql'
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





const UpdateResourceMutation = graphql(`
  mutation UpdateResource($input: UpdateResourceInput!)  {
  updateResource(input: $input) {
    clientMutationId
    resource {
        id
    }
  }
}`)


const UpdateResource = (resourceInput:UpdateResourceInput) => {
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
   

    const [result, executeMutation] = useMutation(UpdateResourceDocument);
    const initialValues: ResourcePatch = resourceInput.patch;

    return (
      <div id='update-resource'>
        <button onClick={openModal}>
        Update Resource
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Resource'
      >
        <h2>Update Resource</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: ResourcePatch,
            { setSubmitting }: FormikHelpers<ResourcePatch> 
          ) => {
              const submittedValues: UpdateResourceInput = {id: resourceInput.id, clientMutationId: "tester", patch: values }
              executeMutation({input: submittedValues})
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={initialValues.name} />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" placeholder={initialValues.businessId} />

          <label htmlFor="cost">Cost</label>
          <Field id="cost" name="cost" placeholder={initialValues.cost} />

          <label htmlFor="unit">Unit</label>
          <Field id="unit" name="unit" placeholder={initialValues.unit} />

          <label htmlFor="resourceType">ResourceType</label>
          <Field as="select" name="resourceType" defaultValue={initialValues.resourceType}>
          <option  value="DEV">Extreme</option>
           </Field>
           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateResource