CREATE FUNCTION risky_public.attack_chains(fn_business_id uuid) 
RETURNS SETOF risky_public.attack_chain 
AS $$
    SELECT * FROM risky_public.attack_chain WHERE business_id = fn_business_id;
$$ LANGUAGE sql VOLATILE;

CREATE FUNCTION risky_public.get_attack_chain(fn_attack_chain_id uuid) 
RETURNS risky_public.attack_chain 
AS $$
    SELECT * FROM risky_public.attack_chain WHERE id = fn_attack_chain_id;
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.delete_attack_chain(fn_attack_chain_id uuid)
RETURNS uuid 
AS $$
    DELETE FROM risky_public.attack_chain WHERE id = fn_attack_chain_id RETURNING fn_attack_chain_id;
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.create_attack_chain(fn_name varchar, fn_description varchar, fn_business_id uuid, fn_threat_id uuid) 
RETURNS uuid
AS $$
    INSERT INTO risky_public.attack_chain(name, description, business_id, threat_id) VALUES (fn_name, fn_description, fn_business_id, fn_threat_id) RETURNING id;
 $$ LANGUAGE sql VOLATILE;


DROP FUNCTION risky_public.update_attack_chain;
CREATE FUNCTION risky_public.update_attack_chain(fn_attack_chain_id uuid, fn_name varchar, fn_description varchar, fn_business_id uuid, fn_threat_id uuid) 
RETURNS void
AS $$
    UPDATE risky_public.attack_chain SET name = fn_name, description = fn_description, threat_id = fn_threat_id, business_id = fn_business_id WHERE id = fn_attack_chain_id;
$$ LANGUAGE sql VOLATILE;

