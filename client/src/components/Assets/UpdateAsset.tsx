import React from 'react';
import Modal from 'react-modal';
import { Formik, Field, Form, FormikHelpers } from 'formik';
import { UUID } from 'crypto';
import axios from "axios";
import { useEffect, useState } from "react";
import Asset from './Asset';

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


const UpdateAsset = (updateAssetInput:Asset) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);
   const [assetInput, setAssetInput] = useState(updateAssetInput);

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
      <div id='update-asset'>
        <button onClick={openModal}>
        Update Asset
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update Asset'
      >
        <h2>Update Asset</h2>
        <Formik
          initialValues={assetInput}
          onSubmit={(
            values: Asset,
            { setSubmitting }: FormikHelpers<Asset> 
          ) => {
            axios.patch<Asset>(`http://localhost:8081/asset`, values).catch((err) => console.log(err))
              closeModal()
          }}
        >
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateAssetInput.name} />

          <label htmlFor="description">Description</label>
          <Field id="description" name="description" placeholder={updateAssetInput.description} />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" placeholder={updateAssetInput.businessId} />

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateAsset