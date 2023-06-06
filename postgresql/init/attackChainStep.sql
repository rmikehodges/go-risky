DROP FUNCTION risky_public.attack_chain_steps;
CREATE FUNCTION risky_public.attack_chain_steps(fn_business_id uuid, fn_attack_chain_id uuid) 
RETURNS SETOF risky_public.attack_chain_step 
AS $$
    SELECT * FROM risky_public.attack_chain_step WHERE business_id = fn_business_id AND attack_chain_id = fn_attack_chain_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.get_attack_chain_step;
CREATE FUNCTION risky_public.get_attack_chain_step(fn_attack_chain_step_id uuid) 
RETURNS risky_public.attack_chain_step
AS $$
    SELECT * FROM risky_public.attack_chain_step WHERE id = fn_attack_chain_step_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.delete_attack_chain_step;
CREATE FUNCTION risky_public.delete_attack_chain_step(fn_attack_chain_step_id uuid)
RETURNS void 
AS $$
    DELETE FROM risky_public.attack_chain_step WHERE id = fn_attack_chain_step_id;
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.create_attack_chain_step;
CREATE FUNCTION risky_public.create_attack_chain_step(fn_attack_chain_id uuid, fn_action_id uuid, fn_asset_id uuid, fn_position integer, fn_business_id uuid) 
RETURNS risky_public.attack_chain_step
AS $$
    INSERT INTO risky_public.attack_chain_step(attack_chain_id, action_id, asset_id, position, business_id) values(fn_attack_chain_id, fn_action_id, fn_asset_id,fn_position, fn_business_id) RETURNING *;
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.update_attack_chain_step;
CREATE FUNCTION risky_public.update_attack_chain_step(fn_attack_chain_step_id uuid, fn_attack_chain_id uuid, fn_action_id uuid,  fn_asset_id uuid, fn_position integer, fn_business_id uuid) 
RETURNS void 
AS $$
    UPDATE risky_public.attack_chain_step SET attack_chain_id = fn_attack_chain_id, action_id = fn_action_id, asset_id = fn_asset_id, position = fn_position, business_id = fn_business_id WHERE id = fn_attack_chain_step_id;
$$ LANGUAGE sql VOLATILE;
