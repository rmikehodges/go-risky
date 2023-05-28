CREATE OR REPLACE FUNCTION risky_public.mitigations(fn_business_id uuid) 
RETURNS SETOF risky_public.mitigation 
AS $$
    SELECT * FROM risky_public.mitigation WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.get_mitigation(fn_mitigation_id uuid) 
RETURNS risky_public.mitigation 
AS $$
    SELECT * FROM risky_public.mitigation WHERE id = fn_mitigation_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.delete_mitigation(fn_mitigation_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.mitigation WHERE id = fn_mitigation_id;
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.create_mitigation(fn_name varchar, fn_description varchar, fn_business_id uuid,  fn_action_id uuid, fn_implemented boolean)
RETURNS uuid
AS $$
    INSERT INTO risky_public.mitigation(name, description, business_id, action_id, implemented) values(fn_name, fn_description, fn_business_id, fn_action_id, fn_implemented) RETURNING id;
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_mitigation(fn_mitigation_id uuid, fn_name varchar, fn_description varchar, fn_business_id uuid,  fn_action_id uuid, fn_implemented boolean) 
RETURNS void 
AS $$
    UPDATE risky_public.mitigation SET name = fn_name, description = fn_description, action_id = fn_action_id, implemented = fn_implemented WHERE id = fn_mitigation_id;
$$ LANGUAGE sql VOLATILE;
