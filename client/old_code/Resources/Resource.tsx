import { useQuery } from 'urql';
import { graphql } from '../../gql';
import { useParams } from "react-router";
import { ResourceDocument } from '../../gql/graphql';

const ResourceQuery = graphql(`
query Resource($id: UUID!) {
    resource(id: $id) {
        id
    }
  }
  
`);

const Resource = () => {
    let { id } = useParams();
    const [{ data }] = useQuery({query: ResourceDocument, variables: {id: id}});


  return (

    <div>
        data.resource.id
    </div>
  );
};

export default Resource