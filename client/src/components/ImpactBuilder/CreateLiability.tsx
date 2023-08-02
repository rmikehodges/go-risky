import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";

import Liability from '../Liabilities/Liability';

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



const CreateLiability = () => {
  const [modalIsOpen, setIsOpen] = React.useState(false);
  // const [liabilityInput, setLiabilityInput] = useState<Liability>();
  const liabilityInput: Liability = {id: null, name: "", description : "", quantity: 0, cost: 0, type: "", resourceType: "", 
  businessId: null, detectionId: null, mitigationId: null, resourceId: null, threatId: null,impactId: null, createdAt: null}

   const openModal = () => {
    setIsOpen(true)
   }

   const afterOpenModal = () => {
    // references are now sync'd and can be accessed.
  }

  const closeModal = () => {
    setIsOpen(false);
  }
   return (
      <>
        <button onClick={openModal}>
        +
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create Liability'
      >
        <h2>Create Liability</h2>
        <Formik
          initialValues={liabilityInput}
          onSubmit={(
            values: Liability,
            { setSubmitting }: FormikHelpers<Liability> 
          ) => {
              axios.post<Liability>(`http://localhost:8081/v0/liability`, values).catch(err => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />

          <label htmlFor="description">description</label>
          <Field id="description" name="description" />

          <label htmlFor="quantity">quantity</label>
          <Field id="quantity" name="quantity" />

          <label htmlFor="type">Type</label>
          <Field as="select" name="type">
          <option  value="EXPLICIT">Explicit</option>
          <option  value="BUSINESS INTERRUPTION LOSS">Business Interruption Loss</option>
           </Field>

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" />

          <label htmlFor="detectionId">detectionId</label>
          <Field id="detectionId" name="detectionId" />

          <label htmlFor="mitigationId">mitigationId</label>
          <Field id="mitigationId" name="mitigationId" />

          <label htmlFor="resourceId">resourceId</label>
          <Field id="resourceId" name="resourceId"/>

          <label htmlFor="threatId">threatId</label>
          <Field id="threatId" name="threatId"/>

          <label htmlFor="impactId">impactId</label>
          <Field id="impactId" name="impactId"/>


           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </>


    )
};

export default CreateLiability