import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID } from 'crypto';
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



const CreateDetection = () => {
  const [modalIsOpen, setIsOpen] = React.useState(false);
  // const [detectionInput, setDetectionInput] = useState<Detection>();
  const detectionInput: Detection = {id: null, name: "", description : "", businessId: null, implemented: false, createdAt: null}

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
      <div id='create-detection'>
        <button onClick={openModal}>
        Create Detection
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Detection'
      >
        <h2>Create Detection</h2>
        <Formik
          initialValues={detectionInput}
          onSubmit={(
            values: Detection,
            { setSubmitting }: FormikHelpers<Detection> 
          ) => {
              axios.post<Detection>(`http://localhost:8081/detection`, values).catch(err => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />

          <label htmlFor="description">Description</label>
          <Field id="description" name="description"/>

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" />
          
          <label htmlFor="implemented">implemented</label>
          <Field as="checkbox" name="implemented" defaultValue={false}/>
           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateDetection