import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID } from 'crypto';
import axios from "axios";
import { useEffect, useState } from "react";
import Threat from './Threat';


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



const CreateThreat = () => {
  const [modalIsOpen, setIsOpen] = React.useState(false);
  // const [threatInput, setThreatInput] = useState<Threat>();
  const threatInput: Threat = {id: null, name: "", description : "", businessId: null, createdAt: null}

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
          initialValues={threatInput}
          onSubmit={(
            values: Threat,
            { setSubmitting }: FormikHelpers<Threat> 
          ) => {
              axios.post<Threat>(`http://localhost:8081/v0/threat`, values).catch(err => console.log(err))
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
           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateThreat