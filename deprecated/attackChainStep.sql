DROP FUNCTION risky_public.attack_chain_steps_by_attack_chain_id;
CREATE FUNCTION risky_public.attack_chain_steps_by_attack_chain_id(fn_business_id uuid, fn_attack_chain_id uuid) 
RETURNS SETOF risky_public.attack_chain_step 
AS $$
    SELECT * FROM risky_public.attack_chain_step WHERE business_id = fn_business_id AND attack_chain_id = fn_attack_chain_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.attack_chain_steps_by_action_id;
CREATE FUNCTION risky_public.attack_chain_steps_by_action_id(fn_business_id uuid, fn_action_id uuid) 
RETURNS SETOF risky_public.attack_chain_step 
AS $$
    SELECT * FROM risky_public.attack_chain_step WHERE business_id = fn_business_id AND action_id = fn_action_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.attack_chain_steps;
CREATE FUNCTION risky_public.attack_chain_steps(fn_business_id uuid) 
RETURNS SETOF risky_public.attack_chain_step 
AS $$
    SELECT * FROM risky_public.attack_chain_step WHERE business_id = fn_business_id;
$$ LANGUAGE sql;


DROP FUNCTION risky_public.get_attack_chain_step;
CREATE FUNCTION risky_public.get_attack_chain_step(fn_attack_chain_step_id uuid) 
RETURNS risky_public.attack_chain_step
AS $$
    SELECT * FROM risky_public.attack_chain_step WHERE id = fn_attack_chain_step_id;
$$ LANGUAGE sql;

-- Delete should remove the item and move the position of others in the chain downward
DROP FUNCTION risky_public.delete_attack_chain_step;
CREATE FUNCTION risky_public.delete_attack_chain_step(fn_attack_chain_step_id uuid)
RETURNS void 
AS $$
    declare
        v_next_step uuid;
        v_previous_step uuid;
    begin
    SELECT next_step, previous_step INTO v_next_step, v_previous_step FROM risky_public.attack_chain_step WHERE id = fn_attack_chain_step_id;
    UPDATE risky_public.attack_chain_step SET next_step = v_next_step WHERE next_step = fn_attack_chain_step_id;
    If v_previous_step = fn_attack_chain_step_id then
        UPDATE risky_public.attack_chain_step SET previous_step = NULL WHERE previous_step = fn_attack_chain_step_id;
    else
        UPDATE risky_public.attack_chain_step SET previous_step = v_previous_step WHERE previous_step = fn_attack_chain_step_id;
    end if;
    DELETE FROM risky_public.attack_chain_step WHERE id = fn_attack_chain_step_id;
    end;
$$ LANGUAGE plpgsql VOLATILE;

-- Create should move the position of others in the chain upward if its position is less than the new position
DROP FUNCTION risky_public.create_attack_chain_step;
-- CREATE FUNCTION risky_public.create_attack_chain_step(fn_attack_chain_id uuid, fn_action_id uuid, fn_asset_id uuid, fn_next_step uuid, fn_previous_step uuid, fn_business_id uuid) 
-- RETURNS uuid
-- AS $$
--     INSERT INTO risky_public.attack_chain_step(attack_chain_id, action_id, asset_id, next_step, previous_step, business_id) values(fn_attack_chain_id, fn_action_id, fn_asset_id,fn_next_step, fn_previous_step , fn_business_id) RETURNING id;
-- $$ LANGUAGE sql VOLATILE;

-- Update should move the position of others in the chain upward if its position is less than the new position
DROP FUNCTION risky_public.update_attack_chain_step;
-- CREATE FUNCTION risky_public.update_attack_chain_step(fn_attack_chain_step_id uuid, fn_attack_chain_id uuid, fn_action_id uuid,  fn_asset_id uuid, fn_next_step uuid, fn_previous_step uuid, fn_business_id uuid) 
-- RETURNS void 
-- AS $$
--     UPDATE risky_public.attack_chain_step SET attack_chain_id = fn_attack_chain_id, action_id = fn_action_id, asset_id = fn_asset_id, next_step = fn_next_step, previous_step = fn_previous_step, business_id = fn_business_id WHERE id = fn_attack_chain_step_id;
-- $$ LANGUAGE sql VOLATILE;
