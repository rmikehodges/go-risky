import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import ReactDOM from 'react-dom';
import {UpdateAttackChainDocument,AttackChainPatchRecordInput , UpdateAttackChainInput, ActionMapInput, InputMaybe} from '../../gql/graphql'
import { Formik, Field, Form, FormikHelpers, FieldArray, ErrorMessage } from 'formik';
import { UUID, randomUUID } from 'crypto';

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


interface AttackChainPatch {
  actions: InputMaybe<ActionMapInput>[]
  businessId?: string
  name?: InputMaybe<string>
}


Modal.setAppElement('#root');

const UpdateAttackChainMutation = graphql(`
  mutation UpdateAttackChain($input: UpdateAttackChainInput!)  {
  updateAttackChain(input: $input) {
    clientMutationId
    attackChain {
        id
    }
  }
}`)


const UpdateAttackChain = (updateAttackChainInput:UpdateAttackChainInput) => {
   const [modalIsOpen, setIsOpen] = React.useState(false);

   const openModal = () => {
    setIsOpen(true)
   }

   const afterOpenModal = () => {
    // references are now sync'd and can be accessed.
  }

  const closeModal = () => {
    setIsOpen(false);
  }
   

    const [result, executeMutation] = useMutation(UpdateAttackChainDocument);
    let tempActions;

    if (updateAttackChainInput.patch.actions != null) {
      tempActions = updateAttackChainInput.patch.actions
    } else {
      tempActions = [{}];
    }

    const initialValues: AttackChainPatch = {name: updateAttackChainInput.patch.name, businessId: updateAttackChainInput.patch.businessId, actions: tempActions}

    return (
      <div id='update-attackChain'>
        <button onClick={openModal}>
        Update AttackChain
      </button>
      <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel='Update AttackChain'
      >
        <h2>Update AttackChain</h2>
        <Formik
          initialValues={initialValues}
          onSubmit={(
            values: AttackChainPatch,
            { setSubmitting }: FormikHelpers<AttackChainPatch> 
          ) => {
              const submittedValues: UpdateAttackChainInput = {id: updateAttackChainInput.id, clientMutationId: "UpdateAttackChain", patch: values }
              executeMutation({input: submittedValues})
              closeModal()
          }}
        >
          {({ values }) => (
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" placeholder={updateAttackChainInput.patch?.name} />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" placeholder={updateAttackChainInput.patch?.businessId} />

          <FieldArray name="actions">
          {({ insert, remove, push }) => (
              <div>
                {values.actions.length > 0 &&
                  values.actions.map((action, index) => (
                    <div className="row" key={index}>
                      <div className="col">
                        <label htmlFor={`friends.${index}.id`}>ID</label>
                        <Field
                          name={`actions.${index}.id`}
                          placeholder="id"
                          type="text"
                        />
                        <ErrorMessage
                          name={`actions.${index}.id`}
                          component="div"
                          className="field-error"
                        />
                      </div>
                      <div className="col">
                        <label htmlFor={`actions.${index}.position`}>Position</label>
                        <Field
                          name={`actions.${index}.position`}
                          placeholder="0"
                          type="number"
                        />
                        <ErrorMessage
                          name={`actions.${index}.position`}
                          component="div"
                          className="field-error"
                        />
                      </div>
                      <div className="col">
                        <button
                          type="button"
                          className="secondary"
                          onClick={() => remove(index)}
                        >
                          X
                        </button>
                      </div>
                    </div>
                  ))}
                <button
                  type="button"
                  className="secondary"
                  onClick={() => push({ id: '', position: '' })}
                >
                  Add Action
                </button>
              </div>
            )}
          </FieldArray>

           <button type="submit">Update</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>)}
        </Formik>

      </Modal>
      </div>


    )
};

export default UpdateAttackChain