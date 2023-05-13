import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { useParams } from "react-router";
import { ExposureDocument } from '../../gql/graphql';

const ExposureQuery = graphql(`
query Exposure($id: UUID!) {
    exposure(id: $id) {
        id
    }
  }
  
`);

const Exposure = () => {
    let { id } = useParams();
    const [{ data }] = useQuery({query: ExposureDocument, variables: {id: id}});


  return (

    <div>
        data.exposure.id
    </div>
  );
};

export default Exposure