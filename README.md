# Risky

## Where I Left Off
Finished all the base CRUD operations but I realized the usefulness in having different structs for input, output,  and the data model. Namely, selectively returning information to the user and doing input validation in the conversion functions. Only Action, Asset, Resource, Threat, and Vulnerability have been done. Need to do attackChain, attackChainAction, Capability, Detection, Impact, Liability, and Mitigation.

## TODO
1. Create database calls for the crud operations. (In-progress)
2. Write tests
3. Think about struct creation for input structs and database structs
3. Create Seed Script for DB


## TODO Testing
1. pgTap for Postgresql Functions
2. Handlers
3. Database functions