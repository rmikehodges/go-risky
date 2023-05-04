CREATE extension if not exists "uuid-ossp";

CREATE schema risky_public;
CREATE schema risky_private;
CREATE schema risky_secret;

CREATE TYPE risky_public.action_complexity AS ENUM ('LOW', 'MEDIUM', 'HIGH', 'EXTREME');

CREATE TYPE risky_public.resources_type AS ENUM ('DEVELOPER', 'LEGAL', 'COMPLIANCE', 'ADMINISTRATIVE', 'CASH', 'OTHER');

CREATE TABLE risky_public.business (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    revenue DOUBLE PRECISION NOT NULL,
    created_at       timestamp default now()
);


CREATE TABLE risky_public.vulnerability (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    created_at       timestamp default now()
);

CREATE TABLE risky_public.capability (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    created_at       timestamp default now()
);

CREATE TABLE risky_public.asset(
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    created_at       timestamp default now()
);


CREATE TABLE risky_public.action(
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    capability_id uuid references risky_public.capability(id),
    vulnerability_id uuid references risky_public.capability(id),
    business_id uuid references risky_public.business(id) NOT NULL,
    complexity risky_public.action_complexity,
    asset_id uuid references risky_public.asset(id),
    created_at       timestamp default now()
);


CREATE TABLE risky_public.resource (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    cost DOUBLE PRECISION NOT NULL,
    unit TEXT NOT NULL,
    total DOUBLE PRECISION NOT NULL,
    resource_type risky_public.resources_type NOT NULL,
    business_id uuid references risky_public.business(id) NOT NULL,
    created_at       timestamp default now()
);


CREATE TABLE risky_public.mitigation (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    action_id uuid references risky_public.action(id),
    implemented boolean default FALSE,
    created_at       timestamp default now()
);

CREATE TABLE risky_public.detection (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    action_id uuid references risky_public.action(id),
    implemented boolean default FALSE,
    created_at       timestamp default now()
);


CREATE TABLE risky_public.threat (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    created_at       timestamp default now()
);


CREATE TABLE risky_public.impact (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    threat_id uuid references risky_public.threat(id),
    exploitation_cost DOUBLE PRECISION,
    mitigation_cost DOUBLE PRECISION,
    created_at       timestamp default now()
);

CREATE TABLE risky_public.attack_chain (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    threat_id uuid references risky_public.threat(id),
    created_at       timestamp default now()
);


CREATE TABLE risky_public.liability (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    quantity DOUBLE PRECISION NOT NULL,
    cost DOUBLE precision,
    business_id uuid references risky_public.business(id) NOT NULL,
    mitigation_id uuid references risky_public.mitigation(id),
    resource_id uuid references risky_public.resource(id),
    threat_id uuid references risky_public.threat(id),
    impact_id uuid references risky_public.impact(id),
    created_at       timestamp default now()
);

CREATE TABLE risky_public.attack_chain_action (
    attack_chain_id uuid constraint attack_chain_mapping_attack_chain_id_fkey references risky_public.attack_chain (id) NOT NULL,
    action_id uuid  constraint attack_chain_mapping_action_id_fkey references risky_public.action (id) NOT NULL,
    business_id uuid references risky_public.business(id) NOT NULL,
    primary key (attack_chain_id, action_id),
    position INT NOT NULL,
    created_at       timestamp default now()
);

-- Create Seed Data for Business
INSERT INTO risky_public.business (name, revenue) VALUES ('ACME', 1000000);
INSERT INTO risky_public.business (name, revenue) VALUES ('Globex', 1000000);
INSERT INTO risky_public.business (name, revenue) VALUES ('Soylent', 1000000);

-- Create Seed Data for Vulnerability
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

