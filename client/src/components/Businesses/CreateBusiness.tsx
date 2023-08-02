import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID } from 'crypto';
import axios from "axios";
import { useEffect, useState } from "react";
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



const CreateBusiness = () => {
  const [modalIsOpen, setIsOpen] = React.useState(false);
  // const [businessInput, setBusinessInput] = useState<CreateBusinessInput>();
  const businessInput: Business = {id: null, name: "", revenue: 0, createdAt: null }

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
          initialValues={businessInput}
          onSubmit={(
            values: Business,
            { setSubmitting }: FormikHelpers<Business> 
          ) => {
              axios.post<Business>(`http://localhost:8081/v0/business`, values).catch(err => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />

          <label htmlFor="revenue">Revenue</label>
          <Field id="revenue" name="revenue"/>
           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateBusiness