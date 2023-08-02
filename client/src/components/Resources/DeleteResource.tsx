import axios from "axios";
import { UUID } from 'crypto';
import { useRef, useEffect, useState, useCallback } from "react";


interface DeleteResourceInput {
  id: UUID | null
}

const DeleteResource = (props:DeleteResourceInput) => {
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
    await axios.delete(`http://localhost:8081/v0/resource?id=${props.id}`)
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

export default DeleteResource
