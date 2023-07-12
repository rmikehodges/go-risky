CREATE extension if not exists "uuid-ossp";

CREATE schema risky_public;
CREATE schema risky_private;
CREATE schema risky_secret;

CREATE TYPE risky_public.action_complexity AS ENUM ('LOW', 'MEDIUM', 'HIGH', 'EXTREME');

CREATE TYPE risky_public.resource_type AS ENUM ('REVENUE','OVERTIME','EMPLOYEE', 'LEGAL', 'COMPLIANCE', 'CONSULTING', 'CASH', 'OTHER');

CREATE TYPE risky_public.liability_type AS ENUM ('EXPLICIT', 'BUSINESS INTERRUPTION LOSS');

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
    resource_type risky_public.resource_type NOT NULL,
    business_id uuid references risky_public.business(id) NOT NULL,
    created_at       timestamp default now()
);


CREATE TABLE risky_public.mitigation (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
    implemented boolean default FALSE,
    created_at       timestamp default now()
);

CREATE TABLE risky_public.detection (
    id uuid primary key default gen_random_uuid(),
    name varchar not null,
    description varchar default '',
    business_id uuid references risky_public.business(id) NOT NULL,
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
    quantity DOUBLE PRECISION,
    cost DOUBLE precision,
    type risky_public.liability_type NOT NULL,
    resource_type risky_public.resource_type,
    business_id uuid references risky_public.business(id) NOT NULL,
    detection_id uuid references risky_public.detection(id),
    mitigation_id uuid references risky_public.mitigation(id),
    resource_id uuid references risky_public.resource(id),
    threat_id uuid references risky_public.threat(id),
    impact_id uuid references risky_public.impact(id),
    created_at       timestamp default now()
);

CREATE TABLE risky_public.attack_chain_step (
    id uuid primary key default gen_random_uuid(),
    attack_chain_id uuid constraint attack_chain_mapping_attack_chain_id_fkey references risky_public.attack_chain (id) NOT NULL,
    action_id uuid  constraint attack_chain_mapping_action_id_fkey references risky_public.action (id) NOT NULL,
    asset_id uuid  constraint attack_chain_mapping_asset_id_fkey references risky_public.asset (id),
    business_id uuid references risky_public.business(id) NOT NULL,
    detection_id uuid references risky_public.detection(id),
    mitigation_id uuid references risky_public.mitigation(id),
    previous_step uuid constraint attack_chain_step_mapping_previous_step_id_fkey references risky_public.attack_chain_step (id),
    next_step uuid constraint attack_chain_step_mapping_next_step_id_fkey references risky_public.attack_chain_step (id),
    created_at       timestamp default now()
);