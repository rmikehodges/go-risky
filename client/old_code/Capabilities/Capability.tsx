import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { useParams } from "react-router";
import { CapabilityDocument } from '../../gql/graphql';

const CapabilityQuery = graphql(`
query Capability($id: UUID!) {
    capability(id: $id) {
        id
    }
  }
  
`);

const Capability = () => {
    let { id } = useParams();
    const [{ data }] = useQuery({query: CapabilityDocument, variables: {id: id}});


  return (

    <div>
        data.capability.id
    </div>
  );
};

export default Capability