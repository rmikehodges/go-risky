
-- Create Seed Data for Business
INSERT INTO risky_public.business (name, revenue) VALUES ('ACME', 1000000);
INSERT INTO risky_public.business (name, revenue) VALUES ('Globex', 1000000);
INSERT INTO risky_public.business (name, revenue) VALUES ('Soylent', 1000000);

-- Create Seed Data for Vulnerability using the output ofr the business table seed for the business_id
INSERT INTO risky_public.vulnerability (name, description, business_id) VALUES ('Cross Site Scripting', 'XSS is a type of attack that can be carried out to compromise users of a website. The attacker injects malicious code into a web page that then is executed by the victim.', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.vulnerability (name, description, business_id) VALUES ('SQL Injection', 'SQL Injection is a type of attack that can be carried out to compromise users of a website. The attacker injects malicious SQL code into a web page that then is executed by the victim.', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.vulnerability (name, description, business_id) VALUES ('Cross Site Request Forgery', 'CSRF is a type of attack that can be carried out to compromise users of a website. The attacker forces the victim to make a request to a website that they are already logged into, which then executes the request on behalf of the victim.', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.vulnerability (name, description, business_id) VALUES ('Broken Authentication', 'Broken Authentication is a type of attack that can be carried out to compromise users of a website. The attacker exploits a vulnerability in the authentication protocol to gain access to a user account.', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');


-- Create Seed Data for Capability using common actions someone would take in technology such as logging on
INSERT INTO risky_public.capability (name, description, business_id) VALUES ('Logon', 'The ability to logon to a system', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.capability (name, description, business_id) VALUES ('Read', 'The ability to read data from a system', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.capability (name, description, business_id) VALUES ('Write', 'The ability to write data to a system', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.capability (name, description, business_id) VALUES ('Delete', 'The ability to delete data from a system', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');


-- Create Seed Data for Resource using common resources someone would have in technology such as an employee's time or the businesses' money
INSERT INTO risky_public.resource (name, description, business_id) VALUES ('Employee Time', 'The time of an employee', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.resource (name, description, business_id) VALUES ('Business Money', 'The money of the business', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');

-- Create Seed Data for action using common actions a hacker would take attacking a system
INSERT INTO risky_public.action (name, description, business_id) VALUES ('Exploit', 'The ability to exploit a vulnerability', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.action (name, description, business_id) VALUES ('Use', 'The ability to use a capability', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');

-- Create seed data for threat using the common threats to a system such as deletion of data or theft of data
INSERT INTO risky_public.threat (name, description, business_id) VALUES ('Theft', 'The ability to steal data from a system', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.threat (name, description, business_id) VALUES ('Deletion', 'The ability to delete data from a system', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');


--Create seed data for asset
INSERT INTO risky_public.asset (name, description, business_id) VALUES ('Customer Data', 'The data of the customer', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.asset (name, description, business_id) VALUES ('Employee Data', 'The data of the employee', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');

--Create seed data for attackChain
INSERT INTO risky_public.attack_chain (name, description, business_id) VALUES ('XSS Attack Chain', 'The attack chain for XSS', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.attack_chain (name, description, business_id) VALUES ('SQL Injection Attack Chain', 'The attack chain for SQL Injection', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.attack_chain (name, description, business_id) VALUES ('CSRF Attack Chain', 'The attack chain for CSRF', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');

--Create seed data for attackChainStep
INSERT INTO risky_public.attack_chain_step (name, description, attack_chain_id, business_id) VALUES ('Step 1', 'Step 1 of the attack chain', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.attack_chain_step (name, description, attack_chain_id, business_id) VALUES ('Step 2', 'Step 2 of the attack chain', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');

-- Create seed data for detection using data describing how to detect an attack
INSERT INTO risky_public.detection (name, description, business_id) VALUES ('Detect 1', 'Detect 1 of the attack chain', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.detection (name, description, business_id) VALUES ('Detect 2', 'Detect 2 of the attack chain', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');

-- Create seed data for risky_public.liability using data describing the liability of an attack
INSERT INTO risky_public.liability (name, description, business_id) VALUES ('Liability 1', 'Liability 1 of the attack chain', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.liability (name, description, business_id) VALUES ('Liability 2', 'Liability 2 of the attack chain', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');

-- create seed data for risky_public.mitigation using data describing the mitigation of an attack
INSERT INTO risky_public.mitigation (name, description, business_id) VALUES ('Mitigation 1', 'Mitigation 1 of the attack chain', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');
INSERT INTO risky_public.mitigation (name, description, business_id) VALUES ('Mitigation 2', 'Mitigation 2 of the attack chain', 'a6b3c0a1-3d1a-4a8b-bc0d-9e9b3f14c8c7');


