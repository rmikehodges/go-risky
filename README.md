# Risky

## Where I Left Off
Rough unit tests are started, need actual data to test. Seeding sql statements are there, but need to be tied back to business id. Probably need to put seeding into a function to be able to use the variable of the business_id since its required.

Optional fields are handled using pointers to the original values so that they can possibly be nil. This is a little weird, but it works. I'm not sure if this is the best way to do it, but it works for now. Its the cleanest way I can think of doing this without having to create a new struct for each optional field.

5/6 Got pgTap installed on the database, need to architect it and do schema validation + unit tests for the database functions.

5/7 Going to take a break from writing tests to work on the UI flows

5/7 Action UI flow is ported over. Requires a decent bit of effort, but should be fairfly straighforward to automate.

5/9 https://reactflow.dev/docs/examples/overview/ - that will help the attack chain builder, 2 different views, tiered table and a flow chart. I think the flow chart is the best way to go, but the tiered table is a good backup.

5/13 I think going towards a balance sheet feel for impact builder. This will allow people familiar with finance to understand the impacts while getting infosec people familiar with it. However, it should be really easy to use. This is because we will have different categories.
1. Compliance Costs
2. Liability Transfer (People hours)
3. Revenue per head due to moving to non-productive task
4. DFIR Costs - lots of companies have someone on retainer for incident response
5. Asset Replacement - could be physical or digital
6. Legal Costs - most companies have someone on retainer, which has a known rate
This should mirror what a balance sheet looks like with categories like YNAB. 

There should be a balance sheet view of cost of mitigation vs cost of exploitation

In order to tie this all together, we are going to need a resource builder 


Attack Chain Builder:
THere should be a view where you can see all the attack chains that a certain action is in. This would help gauge the impact of fixes at different stages in the attack chain.

Each node should have a small box showing whats in it and allow updating upon double clicking

There should be a filter on the items

There should be the ability to add mitigations and detections and it should be indicated somehow, color, or something.


Accounting:
https://www.mdd.com/forensic-accounting-articles/accounting-for-business-interruption-loss-after-cyber-attack/ - should probably read up on forensic accounting methodologies to make impact analysis solid.

Accountants go by revenue / employee hours. This means

5/16 - Reworked the database middleware and started writing tests for the database functions.

5/29 - I figured out that I need to do pointers to uuid.UUID instead of the value since I will need represent nil values coming from the database to deal with optional relationships. Otherwise the all 0 uuid value will be rejected. 

5/30 - finished unit tests for model with the exception of updateImpact and updateAttackChainStep, the step has the issue of the combined primary key that I need to figure out. I solved this by adding in a unique primary key and also refining the logic in the impact calculation.

06/05 - finished all tests for model and updated the handler

06/10 refactored a ton of code and go all the handlers and database tests working. Looking at impaact builder. I need a few things
- creating a new attack chain flow
- connecting edges to nodes creating attackChainSteps
- a workable sidebar view with input filtering and collapsable lists, likely split out capabilities and vulnerabilities.
- 

06/24 TODO: See if automatic liability calculation is working by linking objects. I really need to make a good init script to create a sandbox env with correct object linking so frontend functionality can be tested.

6/26: Fleshed out liability relation seed script in seed.sql and relevant database changes. Got the liability lookup working and calculating each table. Need to total everything up by passing the state up to the main component or just roll it all into one. Noticed that useEffect isnt always necesary and state is only really necessary for conditional rendering and passing around values.

6/30: Looked into bootstrap themes which seems worth the cost for getting this in a decent place. Also looked at how to do RSS feeds and ChatGPT for users which doesnt seem that easy from the frontend. I would rather do it all in the frontend, but the backend could be safer for users. I need to get an opinion on that from a software dev I trust. Probably Tobias.


7/3
Consider doing this for the td tags in the table to make it dynamic to object changes:
```import React from 'react';

const ResourceTable = ({ resource }) => {
  return (
    <table>
      <tbody>
        <tr>       
          {Object.values(resource).map((value, index) => (
            <td key={index}>{value}</td>
          ))}
        </tr>
      </tbody>
    </table>
  );
};

export default ResourceTable;
```

I created all the components for all of different objects and their respective CRUD operations. Still need to test it out, but its looking pretty good so far. Probably need to go and consolidate my thoughts one of these days, but the frontend is really coming along. Made huge progress today focusing on just the bare minimum. Need to figure out the attack chain builder next and how Im going to create the attackChainSteps. Overall great work today.


7/4 - Changed the action nodes to be attackChainStep nodes and realized that I need to do a more thorough design of the attackChainStep node functionality and how I want everything to operate. Real time updates would be nice on the impact builder if
detections or mitigations are added.


## Node Actions
Node Actions:

Each node will be associated to an Action and AttackChain in the beginning. AssetId, DetectionId, and MitigationId can be added later.

OnDrop:
Create AttackChainStep and get the id back to populate the item - Done

OnDelete:
Delete AttackChainStep by Id and move the positions of all other AttackChainSteps accordingly - Done, fully connected end-to-end

Moving Edges around;
Update AttackChainSteps to move their position based on connection of edges.
- Initial value is -1
- If connected on right, but not connected on left it goes to 0, and currentPosition is updated
- Adding connection to the left and make the position currentPosition+1
- Right now the update on the database is not working how I would like as it doesn't appear to be incrementing anything
- How do I want to handle if an attackChain is disconnected somewhere in the middle?
  - Would it be good to make multiple paths in an attackChain to realize a threat? If so then if an edge is broken
    then the node without a left edge would be a position 0 and then we could update all nodes.

OnDoubleClick
Show a larger modal where users can select Assets, Detection and/or Mitigation.
Every detection and mitigation should probably create a liability associated with a resource automatically for the user

OnConnection
populate a node’s position - Done

onDeleteEdge
- update a node’s position to negative 1 

Expand Related Attack Chains on Action and Action+Asset Combo
- the former for problematic capabilities
- the latter for hot spots that need detection/hardening
- Can use DB function to list both
- Can get AttackChainSteps based on Action
- Show/Hide alternative chains in the graph since we are working on one chain at a time. But updates will be
allowed because of the nature of the Custom Node. 

7/10 - I need to remove the model, output, and input objects and consolidate on a single object definition defined in another package. Its creating too much complexity. This is done, was very easy due to lack of functionality built on top of it.

7/11 - Added functionality to select the attackChain to work on and grab the attackchain steps. However
the attackchainstep fetching isn't working correctly after creating an attackchainstep in the UI

7/12 - Changed the datamodel of attackchainstep to be a linked list so its easier to query and update and render.



## TODO Testing
2. Handlers
4. pgTap testing of more complex database logic (calculate_impact etc.)

## pgTAP Installation
1. make PERL=PATH to Perl


## TODO UI Flows
1. ImpactBuilder
- Once a user inputs a threat, the existing liabilities for that threat should populate. The plus sign will trigger a createLiability popup modal where the user can input liabilities 
- Work out flow on creating attackChainSteps on node connection
- Specify threat and then get the listing of Attack Chains based on the threat to choose from or create new. This will create an option to show all attack chains for a threat so you can identify critical assets and actions.


## TODO Backend
1. Import MITRE Attack into Actions once the builder is in working order.