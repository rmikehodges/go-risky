import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID } from 'crypto';
import axios from "axios";
import { useEffect, useState } from "react";

import Action from './Action';

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



const CreateAction = () => {
  const [modalIsOpen, setIsOpen] = React.useState(false);
  // const [actionInput, setActionInput] = useState<Action>();
  const actionInput: Action = {id: null, name: "", description : "", capabilityId: null, vulnerabilityId: null, businessId: null, complexity: "", createdAt: null}

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
      <div id='create-action'>
        <button onClick={openModal}>
        Create Action
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Action'
      >
        <h2>Create Action</h2>
        <Formik
          initialValues={actionInput}
          onSubmit={(
            values: Action,
            { setSubmitting }: FormikHelpers<Action> 
          ) => {
              axios.post<Action>(`http://localhost:8081/v0/action`, values).catch(err => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" />

          <label htmlFor="capabilityId">CapabilityId</label>
          <Field id="capabilityId" name="capabilityId"/>

          <label htmlFor="vulnerabilityId">VulnerabilityId</label>
          <Field id="vulnerabilityId" name="vulnerabilityId"/>

          <label htmlFor="complexity">Complexity</label>
          <Field as="select" name="complexity">
          <option  value="EXTREME">Extreme</option>
          <option  value="HIGH">High</option>
          <option  value="MEDIUM">Medium</option>
          <option  value="LOW">Low</option>
           </Field>
           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateAction