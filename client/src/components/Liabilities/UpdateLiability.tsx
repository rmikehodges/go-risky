import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";
import { useState } from "react";
import Liability from './Liability';

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


const UpdateLiability = (updateLiabilityInput:Liability) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [liabilityInput, setLiabilityInput] = useState(updateLiabilityInput);

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
      <div id='update-liability'>
        <button onClick={openModal}>
        Update Liability
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Liability'
      >
        <h2>Update Liability</h2>
        <Formik
          initialValues={liabilityInput}
          onSubmit={(
            values: Liability,
            { setSubmitting }: FormikHelpers<Liability> 
          ) => {
            axios.patch<Liability>(`http://localhost:8081/liability`, values).catch((err) => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
        <Field id="name" name="name" placeholder={updateLiabilityInput.name} />

        <label htmlFor="description">Description</label>
        <Field id="description" name="description" placeholder={updateLiabilityInput.description} />

        <label htmlFor="quantity">Quantity</label>
        <Field id="quantity" name="quantity" placeholder={updateLiabilityInput.quantity} />
        
        <label htmlFor="type">Type</label>
        <Field id="type" name="type" placeholder={updateLiabilityInput.type} />

        <label htmlFor="businessId">Business ID</label>
        <Field id="businessId" name="businessId" placeholder={updateLiabilityInput.businessId} />

        <label htmlFor="detectionId">Detection ID</label>
        <Field id="detectionId" name="detectionId" placeholder={updateLiabilityInput.detectionId} />

        <label htmlFor="mitigationId">Mitigation ID</label>
        <Field id="mitigationId" name="mitigationId" placeholder={updateLiabilityInput.mitigationId} />

        <label htmlFor="resourceId">Resource ID</label>
        <Field id="resourceId" name="resourceId" placeholder={updateLiabilityInput.resourceId} />

        <label htmlFor="threatId">Threat ID</label>
        <Field id="threatId" name="threatId" placeholder={updateLiabilityInput.threatId} />

        <label htmlFor="impactId">Impact ID</label>
        <Field id="impactId" name="impactId" placeholder={updateLiabilityInput.impactId} />

        <label htmlFor="createdAt">Created At</label>
        <Field id="createdAt" name="createdAt" placeholder={updateLiabilityInput.createdAt} />
           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateLiability