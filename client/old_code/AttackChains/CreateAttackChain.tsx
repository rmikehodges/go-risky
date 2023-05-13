import { useMutation } from 'urql';
import { graphql } from '../../gql';
import React from 'react';
import Modal from 'react-modal';
import {CreateAttackChainDocument, ActionMapInput} from '../../gql/graphql'
import { Formik, Field, Form, FormikHelpers, FieldArray, ErrorMessage } from 'formik';
import { UUID } from 'crypto';


interface CreateAttackChainInput {
  actions: ActionMapInput[]
  businessId?: UUID
  clientMutationId?: string
  name: string
}


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


const CreateAttackChainMutation = graphql(`
mutation CreateAttackChain($input: CreateAttackChainInput!) {
  createAttackChain(input: $input) {
    clientMutationId
    attackChain {
      id
    }
  }
  
  }
`);



const CreateAttackChain = () => {
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
    const [result, executeMutation] = useMutation(CreateAttackChainDocument);
    const initialValues: CreateAttackChainInput = {name: "test",actions: [{}]}

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
          initialValues={initialValues}
          onSubmit={(
            values: CreateAttackChainInput,
            { setSubmitting }: FormikHelpers<CreateAttackChainInput> 
          ) => {
              values.clientMutationId = "CreateAttackChain"
              executeMutation({input: (values)}).then(() => {
                console.log(result.data)
              })

              closeModal()
          }}
        >
          {({ values }) => (
          <Form>
          <label htmlFor="name">Name</label>
          <Field id="name" name="name" />

          <label htmlFor="businessId">businessId</label>
          <Field id="businessId" name="businessId" />


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

          

           <button type="submit">Create</button>
           <button onClick={closeModal}>Cancel</button>
          </Form>)}
        </Formik>

      </Modal>
      </div>


    )
};

export default CreateAttackChain      