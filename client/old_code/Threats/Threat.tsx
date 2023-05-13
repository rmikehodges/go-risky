import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { useParams } from "react-router";
import { ThreatDocument } from '../../gql/graphql';

const ThreatQuery = graphql(`
query Threat($id: UUID!) {
    threat(id: $id) {
        id
    }
  }
  
`);

const Threat = () => {
    let { id } = useParams();
    const [{ data }] = useQuery({query: ThreatDocument, variables: {id: id}});


  return (

    <div>
        data.threat.id
    </div>
  );
};

export default Threat