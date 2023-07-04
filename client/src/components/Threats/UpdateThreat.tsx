import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
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


const UpdateThreat = (updateThreatInput:Threat) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [threatInput, setThreatInput] = useState(updateThreatInput);

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
      <div id='update-threat'>
        <button onClick={openModal}>
        Update Threat
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Threat'
      >
        <h2>Update Threat</h2>
        <Formik
          initialValues={threatInput}
          onSubmit={(
            values: Threat,
            { setSubmitting }: FormikHelpers<Threat> 
          ) => {
            axios.patch<Threat>(`http://localhost:8081/threat`, values).catch((err) => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateThreatInput.name} />

          <label htmlFor="description">Description</label>
          <Field id="description" name="description" placeholder={updateThreatInput.description} />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" placeholder={updateThreatInput.businessId} />

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateThreat