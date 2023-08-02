import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";
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



const CreateAttackChainStep = () => {
  const [modalIsOpen, setIsOpen] = React.useState(false);
  const attackChainStepInput: AttackChainStep = {id: null, actionId: null, assetId: null, attackChainId: null, businessId: null, nextStep: null, previousStep: null, createdAt: null}

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
      <div id='create-attackChainStep'>
        <button onClick={openModal}>
        Create AttackChainStep
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create AttackChainStep'
      >
        <h2>Create AttackChainStep</h2>
        <Formik
          initialValues={attackChainStepInput}
          onSubmit={(
            values: AttackChainStep,
            { setSubmitting }: FormikHelpers<AttackChainStep> 
          ) => {
              axios.post<AttackChainStep>(`http://localhost:8081/v0/attackChainStep`, values).catch(err => console.log(err))
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

          <label htmlFor="position">position</label>
          <Field id="position" name="position" />


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

export default CreateAttackChainStep