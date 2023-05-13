DROP FUNCTION risky_public.actions;
CREATE FUNCTION risky_public.actions(fn_business_id uuid) 
RETURNS SETOF risky_public.action  
AS $$
    SELECT * FROM risky_public.action WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.get_action;
CREATE FUNCTION risky_public.get_action(fn_action_id uuid) 
RETURNS risky_public.action 
AS $$
    SELECT * FROM risky_public.action WHERE id = fn_action_id;
$$ LANGUAGE sql;


DROP FUNCTION risky_public.delete_action;
CREATE FUNCTION risky_public.delete_action(fn_action_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.action WHERE id = fn_action_id;
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.create_action;
CREATE FUNCTION risky_public.create_action(fn_name varchar, fn_description varchar, fn_capability_id uuid, fn_vulnerability_id uuid, fn_business_id uuid, fn_complexity risky_public.action_complexity, fn_asset_id uuid) 
RETURNS void
AS $$
    INSERT INTO risky_public.action(name, description, capability_id, vulnerability_id, business_id, complexity, asset_id) values(fn_name, fn_description, fn_capability_id,fn_vulnerability_id, fn_business_id, fn_complexity, fn_asset_id);
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.update_action;
CREATE FUNCTION risky_public.update_action(fn_action_id uuid,fn_name varchar, fn_description varchar, fn_capability_id uuid, fn_vulnerability_id uuid, fn_business_id uuid, fn_complexity risky_public.action_complexity, fn_asset_id uuid) 
RETURNS void 
AS $$
    UPDATE risky_public.action SET name = fn_name, description = fn_description, capability_id = fn_capability_id, vulnerability_id = fn_vulnerability_id, business_id = fn_business_id, complexity = fn_complexity, asset_id = fn_asset_id WHERE id = fn_action_id;
$$ LANGUAGE sql VOLATILE;
