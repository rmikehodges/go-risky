import UpdateAction from '../src/components/Actions/UpdateAction';
import DeleteAction from '../src/components/Actions/DeleteAction';
import CreateAction from '../src/components/Actions/CreateAction';

interface ActionOutput {
  id: string
  name: string
  description: string
  capabilityId: string
  vulnerabilityId: string
  businessId: string
  complexity: string
  assetId: string
  createdAt: Date
  }

const Actions = () => {
    const [{ data }] = useQuery({ query: ActionsDocument, variables: {count: 100} });


  return (
    <div className='Actions'>
      <div>
        <CreateAction />
      </div>
    <table>
      <thead><tr>
        <th>ID</th>
        <th>Name</th>
        <th>BusinessId</th>
        <th>CapabilityId</th>
        <th>VulnerabilityId</th>
        <th>Complexity</th>
        <th>CreatedAt</th>
        <th>Update</th>
        <th>Delete</th>
      </tr></thead>
      <tbody>
      {data?.actions?.edges.map(action => {
        const updateActionInput = {id: action.node?.id, patch:{name: action.node?.name, complexity: action.node?.complexity,businessId: action.node?.businessId, capabilityId: action.node?.capabilityId, vulnerabilityId: action.node?.vulnerabilityId}}
        const deleteActionInput = {id: action.node?.id}
        return (
          <tr key={action.node?.id}>
            <td>{action.node?.id}</td>
            <td>{action.node?.name}</td>
            <td>{action.node?.businessId}</td>
            <td>{action.node?.capabilityId}</td>
            <td>{action.node?.vulnerabilityId}</td>
            <td>{action.node?.complexity}</td>
            <td>{action.node?.createdAt}</td>
            <td><UpdateAction {...updateActionInput} /></td>
            <td><DeleteAction {...deleteActionInput} /></td>
          </tr>)
      })}
    </tbody>
    </table>
    </div>

  );
};

export default Actions