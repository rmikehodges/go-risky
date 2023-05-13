import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { useParams } from "react-router";
import { BusinessDocument } from '../../gql/graphql';

const BusinessQuery = graphql(`
query Business($id: UUID!) {
    business(id: $id) {
        id
    }
  }
  
`);

const Business = () => {
    let { id } = useParams();
    const [{ data }] = useQuery({query: BusinessDocument, variables: {id: id}});


  return (

    <div>
        data.business.id
    </div>
  );
};

export default Business