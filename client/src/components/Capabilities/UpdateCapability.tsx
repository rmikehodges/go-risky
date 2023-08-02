import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID } from 'crypto';
import axios from "axios";
import { useEffect, useState } from "react";
import Capability from './Capability';

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


const UpdateCapability = (updateCapabilityInput:Capability) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [capabilityInput, setCapabilityInput] = useState(updateCapabilityInput);

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
      <div id='update-capability'>
        <button onClick={openModal}>
        Update Capability
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Capability'
      >
        <h2>Update Capability</h2>
        <Formik
          initialValues={capabilityInput}
          onSubmit={(
            values: Capability,
            { setSubmitting }: FormikHelpers<Capability> 
          ) => {
            axios.patch<Capability>(`http://localhost:8081/v0/capability`, values).catch((err) => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateCapabilityInput.name} />

          <label htmlFor="description">Description</label>
          <Field id="description" name="description" placeholder={updateCapabilityInput.description} />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" placeholder={updateCapabilityInput.businessId} />

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateCapability