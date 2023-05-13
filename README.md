# Risky

## Where I Left Off
Rough unit tests are started, need actual data to test. Seeding sql statements are there, but need to be tied back to business id. Probably need to put seeding into a function to be able to use the variable of the business_id since its required.

Optional fields are handled using pointers to the original values so that they can possibly be nil. This is a little weird, but it works. I'm not sure if this is the best way to do it, but it works for now. Its the cleanest way I can think of doing this without having to create a new struct for each optional field.

5/6 Got pgTap installed on the database, need to architect it and do schema validation + unit tests for the database functions.

5/7 Going to take a break from writing tests to work on the UI flows

5/7 Action UI flow is ported over. Requires a decent bit of effort, but should be fairfly straighforward to automate.

5/9 https://reactflow.dev/docs/examples/overview/ - that will help the attack chain builder, 2 different views, tiered table and a flow chart. I think the flow chart is the best way to go, but the tiered table is a good backup.


## TODO
1. Create database calls for the crud operations. (In-progress)
2. Write tests
3. Think about struct creation for input structs and database structs
3. Create Seed Script for DB


## TODO Testing
1. pgTap for Postgresql Functions
2. Handlers
3. Database functions

## pgTAP Installation
1. make PERL=PATH to Perl


## UI Flows
1. Create Business
2. Create Threat
3. Create Action - check, Delete acts weird
4. Create Attack Chain
5. Create Liabilities
6. Create Impact
