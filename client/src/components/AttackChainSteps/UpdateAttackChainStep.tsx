import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";
import {useState } from "react";
import AttackChainStep from './AttackChainStep';

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


const UpdateAttackChainStep = (updateAttackChainStepInput:AttackChainStep) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [mitigationInput, setAttackChainStepInput] = useState(updateAttackChainStepInput);

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
        Update AttackChainStep
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update AttackChainStep'
      >
        <h2>Update AttackChainStep</h2>
        <Formik
          initialValues={mitigationInput}
          onSubmit={(
            values: AttackChainStep,
            { setSubmitting }: FormikHelpers<AttackChainStep> 
          ) => {
            axios.patch<AttackChainStep>(`http://localhost:8081/mitigation`, values).catch((err) => console.log(err))
              closeModal()
          }}
        >
          <Form>

          <label htmlFor="actionId">actionId</label>
          <Field id="actionId" name="actionId" />

          <label htmlFor="assetId">assetId</label>
          <Field id="assetId" name="assetId" />

          <label htmlFor="attackChainId">attackChainId</label>
          <Field id="attackChainId" name="attackChainId" />

          <label htmlFor="nextStep">nextStep</label>
          <Field id="nextStep" name="nextStep" />

          <label htmlFor="previousStep">previousStep</label>
          <Field id="previousStep" name="previousStep" />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" />

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateAttackChainStep