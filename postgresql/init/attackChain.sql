CREATE FUNCTION risky_public.attack_chains(fn_business_id uuid) 
RETURNS SETOF risky_public.attack_chain 
AS $$
    SELECT * FROM risky_public.attack_chain WHERE business_id = fn_business_id;
$$ LANGUAGE sql VOLATILE;

CREATE FUNCTION risky_public.get_attack_chains(fn_attack_chain_id uuid) 
RETURNS risky_public.attack_chain 
AS $$
    SELECT * FROM risky_public.attack_chain WHERE id = fn_attack_chain_id;
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.delete_attack_chain(fn_attack_chain_id uuid)
RETURNS uuid 
AS $$
    DELETE FROM risky_public.attack_chain WHERE id = fn_attack_chain_id RETURNING fn_attack_chain_id;
$$ LANGUAGE sql VOLATILE;



create type action_map as (
  id uuid,
  position INT
);


create type attack_chain_patch as (
    name varchar,
    business_id uuid,
    actions action_map[]
);

CREATE OR REPLACE FUNCTION risky_public.create_attack_chain(fn_name varchar, fn_business_id uuid,fn_actions action_map[]) 
RETURNS risky_public.attack_chain
AS $$
    declare
        v_attack_chain risky_public.attack_chain;
        v_action action_map;
    begin
        insert into risky_public.attack_chain(business_id, name)
        values(fn_business_id, fn_name)
        RETURNING * INTO v_attack_chain;

        IF fn_actions IS NOT NULL THEN
            FOREACH v_action IN ARRAY fn_actions
            LOOP
                insert into risky_public.attack_chain_action(attack_chain_id, action_id, position)
                values(v_attack_chain.id,v_action.id, v_action.position);
            END LOOP;
        END IF;
        RETURN v_attack_chain;
    end;
 $$ LANGUAGE plpgsql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_attack_chain(fn_attack_chain_id uuid, fn_patch attack_chain_patch) 
RETURNS risky_public.attack_chain 
AS $$
    declare
        v_action action_map;
        v_attack_chain risky_public.attack_chain;
    begin
        update risky_public.attack_chain
        SET business_id = fn_patch.business_id,
            name = fn_patch.name
        WHERE id = fn_attack_chain_id
        RETURNING * INTO v_attack_chain;

        DELETE FROM risky_public.attack_chain_action
        WHERE attack_chain_id = fn_attack_chain_id;
        
        IF fn_patch.actions IS NOT NULL THEN
            FOREACH v_action IN ARRAY fn_patch.actions
            LOOP
                insert into risky_public.attack_chain_action(id, action_id, position)
                values(v_attack_chain.id,v_action.action_id, v_action.position);
            END LOOP;
        END IF;

        RETURN v_attack_chain;
    end;
 $$ LANGUAGE plpgsql VOLATILE STRICT;

