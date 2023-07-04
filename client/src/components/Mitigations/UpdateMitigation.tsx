import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";
import { useEffect, useState } from "react";
import Mitigation from './Mitigation';

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


const UpdateMitigation = (updateMitigationInput:Mitigation) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [mitigationInput, setMitigationInput] = useState(updateMitigationInput);

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
      <div id='update-mitigation'>
        <button onClick={openModal}>
        Update Mitigation
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Mitigation'
      >
        <h2>Update Mitigation</h2>
        <Formik
          initialValues={mitigationInput}
          onSubmit={(
            values: Mitigation,
            { setSubmitting }: FormikHelpers<Mitigation> 
          ) => {
            axios.patch<Mitigation>(`http://localhost:8081/mitigation`, values).catch((err) => console.log(err))
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

          <label htmlFor="actionId">actionId</label>
          <Field id="actionId" name="actionId" />

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

export default UpdateMitigation