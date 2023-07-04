import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID } from 'crypto';
import axios from "axios";
import { useEffect, useState } from "react";
import Impact from './Impact';


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



const CreateImpact = () => {
  const [modalIsOpen, setIsOpen] = React.useState(false);
  // const [impactInput, setImpactInput] = useState<Impact>();
  const impactInput: Impact = {id: null, name: "", description : "", businessId: null, threatId:null, exploitationCost: 0, mitigationCost: 0, createdAt: null}

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
      <div id='create-impact'>
        <button onClick={openModal}>
        Create Impact
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Impact'
      >
        <h2>Create Impact</h2>
        <Formik
          initialValues={impactInput}
          onSubmit={(
            values: Impact,
            { setSubmitting }: FormikHelpers<Impact> 
          ) => {
              axios.post<Impact>(`http://localhost:8081/impact`, values).catch(err => console.log(err))
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

          <label htmlFor="threatId">threatId</label>
          <Field id="threatId" name="threatId" />

           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateImpact