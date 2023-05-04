# Risky

## Where I Left Off
Rough unit tests are started, need actual data to test. Seeding sql statements are there, but need to be tied back to business id. Probably need to put seeding into a function to be able to use the variable of the business_id since its required.

Optional fields are handled using pointers to the original values so that they can possibly be nil. This is a little weird, but it works. I'm not sure if this is the best way to do it, but it works for now. Its the cleanest way I can think of doing this without having to create a new struct for each optional field.


## TODO
1. Create database calls for the crud operations. (In-progress)
2. Write tests
3. Think about struct creation for input structs and database structs
3. Create Seed Script for DB


## TODO Testing
1. pgTap for Postgresql Functions
2. Handlers
3. Database functions