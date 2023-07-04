import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";
import { useEffect, useState } from "react";
import Detection from './Detection';

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


const UpdateDetection = (updateDetectionInput:Detection) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [detectionInput, setDetectionInput] = useState(updateDetectionInput);

   const openModal = () => {
    setIsOpen(true)
   }

   const afterOpenModal = () => {
    // references are now sync'd and can be accessed.
  }

  const closeModal = () => {
    setIsOpen(false);
    window.location.reload();
  }

    return (
      <div id='update-detection'>
        <button onClick={openModal}>
        Update Detection
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Detection'
      >
        <h2>Update Detection</h2>
        <Formik
          initialValues={detectionInput}
          onSubmit={(
            values: Detection,
            { setSubmitting }: FormikHelpers<Detection> 
          ) => {
            axios.patch<Detection>(`http://localhost:8081/detection`, values).catch((err) => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateDetectionInput.name} />

          <label htmlFor="description">Description</label>
          <Field id="description" name="description" placeholder={updateDetectionInput.description} />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" placeholder={updateDetectionInput.businessId} />


          <label htmlFor="implemented">implemented</label>
          <Field as="checkbox" name="implemented" defaultValue={false}/>

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateDetection