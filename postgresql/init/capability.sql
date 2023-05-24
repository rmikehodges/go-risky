CREATE OR REPLACE FUNCTION risky_public.capabilities(fn_business_id uuid) 
RETURNS SETOF risky_public.capability 
AS $$
    SELECT * FROM risky_public.capability WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.get_capability(fn_capability_id uuid) 
RETURNS risky_public.capability 
AS $$
    SELECT * FROM risky_public.capability WHERE id = fn_capability_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.delete_capability(fn_capability_id uuid) 
RETURNS uuid 
AS $$
    DELETE FROM risky_public.capability WHERE id = fn_capability_id RETURNING fn_capability_id;
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.create_capability(fn_name varchar, fn_description varchar) 
RETURNS uuid
AS $$
    INSERT INTO risky_public.capability(name, description) values(fn_name, fn_description) RETURNING id;
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_capability(fn_capability_id uuid, fn_name varchar,  fn_description varchar) 
RETURNS uuid 
AS $$
    UPDATE risky_public.capability SET name = fn_name, description = fn_description WHERE id = fn_capability_id RETURNING fn_capability_id;
$$ LANGUAGE sql VOLATILE;
