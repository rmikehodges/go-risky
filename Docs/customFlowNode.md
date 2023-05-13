Sure, here's an example of how you can create a custom node using React Flow:

First, create a functional component for your custom node. In this example, let's call it "MyNode":

```jsx
import React from 'react';
import { Handle } from 'react-flow-renderer';

const MyNode = ({ data }) => {
  return (
    <div>
      {/* Node content */}
      <Handle type="source" position="right" />
      {/* Output port */}
      <Handle type="target" position="left" />
      {/* Input port */}
    </div>
  );
};

export default MyNode;
```

This defines a basic node with two ports: one input ("target") and one output ("source"). You can customize the appearance and behavior of the node by adding more content to the `div`.

Next, you need to register this node type with the React Flow library. To do this, use the `registerNode` function provided by `react-flow-renderer`.

```jsx
import React from 'react';
import { Handle, registerNode } from 'react-flow-renderer';

const MyNode = ({ data }) => {
  return (
    <div>
      {/* Node content */}
      <Handle type="source" position="right" />
      {/* Output port */}
      <Handle type="target" position="left" />
      {/* Input port */}
    </div>
  );
};

registerNode({
  MyNode,
  // Set the type of the node (must be unique)
  type: 'my-node',
});
```

Finally, you can use your custom node in your React Flow graph by specifying its `type` and any `data` you want to pass to it.

```jsx
import React, { useState } from 'react';
import ReactFlow, { addEdge, Elements, removeElements } from 'react-flow-renderer';
import MyNode from './MyNode';

const initialElements: Elements = [
  {
    id: '1',
    type: 'my-node',
    data: { label: 'My Node' },
    position: { x: 0, y: 0 },
  },
];

const MyFlow = () => {
  const [elements, setElements] = useState(initialElements);

  const onConnect = (params) => setElements((els) => addEdge(params, els));
  const onElementsRemove = (elementsToRemove) =>
    setElements((els) => removeElements(elementsToRemove, els));

  return <ReactFlow elements={elements} onConnect={onConnect} onElementsRemove={onElementsRemove} />;
};

export default MyFlow;
```

That's it! Now you have a custom node that you can use in your React Flow graphs. You can customize the appearance and behavior of your node by passing different `data` and adding more content to the `div` inside the `MyNode` component.