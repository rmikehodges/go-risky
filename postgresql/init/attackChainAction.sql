CREATE OR REPLACE FUNCTION risky_public.attack_chain_actions(fn_business_id uuid) 
RETURNS SETOF risky_public.attack_chain_action 
AS $$
    SELECT * FROM risky_public.attack_chain_action WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.get_attack_chain_action(fn_attack_chain_id uuid, fn_action_id uuid) 
RETURNS risky_public.attack_chain_action
AS $$
    SELECT * FROM risky_public.attack_chain_action WHERE attack_chain_id = fn_attack_chain_id OR action_id = fn_action_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.delete_attack_chain_action(fn_attack_chain_id uuid, fn_action_id uuid)
RETURNS void 
AS $$
    DELETE FROM risky_public.attack_chain_action WHERE attack_chain_id = fn_attack_chain_id OR action_id = fn_action_id;
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.create_attack_chain_action(fn_attack_chain_id uuid, fn_action_id uuid, fn_position integer, fn_business_id uuid) 
RETURNS void
AS $$
    INSERT INTO risky_public.attack_chain_action(attack_chain_id, action_id, position, business_id) values(fn_attack_chain_id, fn_action_id, fn_position, fn_business_id);
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_attack_chain_action(fn_attack_chain_id uuid, fn_action_id uuid, fn_position integer, fn_business_id uuid) 
RETURNS void 
AS $$
    UPDATE risky_public.attack_chain_action SET attack_chain_id = fn_attack_chain_id, action_id = fn_action_id, position = fn_position, business_id = fn_business_id WHERE attack_chain_id = fn_attack_chain_id AND action_id = fn_action_id;
$$ LANGUAGE sql VOLATILE;
