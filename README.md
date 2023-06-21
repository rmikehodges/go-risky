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