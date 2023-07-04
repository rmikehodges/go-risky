import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";
import { useEffect, useState } from "react";
import Resource from './Resource';

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


const UpdateResource = (updateResourceInput:Resource) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [resourceInput, setResourceInput] = useState(updateResourceInput);

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
      <div id='update-resource'>
        <button onClick={openModal}>
        Update Resource
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Resource'
      >
        <h2>Update Resource</h2>
        <Formik
          initialValues={resourceInput}
          onSubmit={(
            values: Resource,
            { setSubmitting }: FormikHelpers<Resource> 
          ) => {
            axios.patch<Resource>(`http://localhost:8081/resource`, values).catch((err) => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <div>
          <label htmlFor="name">Name</label>
          <Field type="text" id="name" name="name" />
        </div>
        <div>
          <label htmlFor="description">Description</label>
          <Field type="text" id="description" name="description" />
        </div>
        <div>
          <label htmlFor="cost">Cost</label>
          <Field type="number" id="cost" name="cost" />
        </div>
        <div>
          <label htmlFor="unit">Unit</label>
          <Field type="text" id="unit" name="unit" />
        </div>
        <div>
          <label htmlFor="total">Total</label>
          <Field type="number" id="total" name="total" />
        </div>
        <div>
          <label htmlFor="resourceType">Resource Type</label>
          <Field type="text" id="resourceType" name="resourceType" />
        </div>
        <div>
          <label htmlFor="businessId">Business ID</label>
          <Field type="text" id="businessId" name="businessId" />
        </div>
           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateResource