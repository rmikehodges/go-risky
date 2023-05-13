import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import {UpdateExposureDocument,UpdateExposureInput, ExposurePatch} from '../../gql/graphql'
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



const UpdateExposureMutation = graphql(`
  mutation UpdateExposure($input: UpdateExposureInput!)  {
  updateExposure(input: $input) {
    clientMutationId
    exposure {
        id
    }
  }
}`)


const UpdateExposure = (updateExposureInput:UpdateExposureInput) => {
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
   

    const [result, executeMutation] = useMutation(UpdateExposureDocument);
    const initialValues: ExposurePatch = updateExposureInput.patch;

    return (
      <div id='update-exposure'>
        <button onClick={openModal}>
        Update Exposure
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Exposure'
      >
        <h2>Update Exposure</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: ExposurePatch,
            { setSubmitting }: FormikHelpers<ExposurePatch> 
          ) => {
              const submittedValues: UpdateExposureInput = {id: updateExposureInput.id, clientMutationId: "tester", patch: values }
              executeMutation({input: submittedValues})
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateExposureInput.patch.name} />
           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateExposure