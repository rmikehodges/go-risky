import axios from "axios";
import { UUID } from 'crypto';
import { useRef, useEffect, useState, useCallback } from "react";


interface DeleteActionInput {
  id: UUID | null
}

const DeleteAction = (props:DeleteActionInput) => {
  const [isSending, setIsSending] = useState(false)
  const isMounted = useRef(true)

  // set isMounted to false when we unmount the component
  useEffect(() => {
    return () => {
      isMounted.current = false
    }
  }, [])

  const deleteRequest = useCallback(async () => {
    // don't send again while we are sending
    if (isSending) return
    // update state
    setIsSending(true)
    // send the actual request
    await axios.delete(`http://localhost:8081/action?id=${props.id}`)
    .catch(err => console.log(err))
    .then(resp => window.location.reload())
    if (isMounted.current) // only update if we are still mounted
      setIsSending(false)
  }, [isSending, props.id]) // update the callback if the state changes

    return (
      <button disabled={isSending} onClick={deleteRequest}>
                Delete
      </button>
    )
};

export default DeleteAction

//Write code that will reload the page after the action is deleted.
//Hint: You can use the window.location.reload() method.
//Hint: You can use the finally() method on a promise to execute code after the promise is resolved or rejected.
//Hint: You can use the catch() method on a promise to execute code if the promise is rejected.
//Hint: You can use the then() method on a promise to execute code if the promise is resolved.
//Hint: You can use the axios.delete() method to delete an action.
//Hint: You can use the axios.delete() method to delete an action.
//Hint: You can use the axios.delete() method to delete an action.
//Hint: You can use the axios.delete() method to delete an action.
//Hint: You can use the axios.delete() method to delete an action.