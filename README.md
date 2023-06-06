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


## TODO
2. Write tests for database models and handlers (postgresql should probably be last)
- Wrote example tests for actions on the model and determined that I need to return more data for create and update in order to test them effectively (or do I? Since I can just get them after create and update)
3. Create Seed Script for DB


## TODO Testing
2. Handlers
4. pgTap testing of more complex database logic (calculate_impact etc.)

## pgTAP Installation
1. make PERL=PATH to Perl


## UI Flows
1. Create Business
2. Create Threat
3. Create Action - check, Delete acts weird
4. Create Attack Chain
5. Create Liabilities
6. Create Impact
