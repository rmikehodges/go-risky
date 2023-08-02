import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";
import {useState } from "react";
import Business from './Business';



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


const UpdateBusiness = (updateBusinessInput:Business) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [businessInput, setBusinessInput] = useState(updateBusinessInput);

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
      <div id='update-business'>
        <button onClick={openModal}>
        Update Business
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Business'
      >
        <h2>Update Business</h2>
        <Formik
          initialValues={businessInput}
          onSubmit={(
            values: Business,
            { setSubmitting }: FormikHelpers<Business> 
          ) => {
            axios.patch<Business>(`http://localhost:8081/v0/business`, values).catch((err) => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateBusinessInput.name} />

          <label htmlFor="revenue">revenue</label>
          <Field id="revenue" name="revenue" placeholder={updateBusinessInput.revenue} />

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateBusiness