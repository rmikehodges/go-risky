import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import axios from "axios";
import AttackChain from './AttackChain';


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



const CreateAttackChain = () => {
  const [modalIsOpen, setIsOpen] = React.useState(false);
  // const [attackChainInput, setAttackChainInput] = useState<AttackChain>();
  const attackChainInput: AttackChain = {id: null, name: "", description : "", businessId: null, threatId:null, createdAt: null}

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
      <div id='create-attackChain'>
        <button onClick={openModal}>
        Create AttackChain
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Create AttackChain'
      >
        <h2>Create AttackChain</h2>
        <Formik
          initialValues={attackChainInput}
          onSubmit={(
            values: AttackChain,
            { setSubmitting }: FormikHelpers<AttackChain> 
          ) => {
              axios.post<AttackChain>(`http://localhost:8081/attackChain`, values).catch(err => console.log(err))
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

export default CreateAttackChain