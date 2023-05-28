CREATE OR REPLACE FUNCTION risky_public.impacts(fn_business_id uuid) 
RETURNS SETOF risky_public.impact 
AS $$
    SELECT * FROM risky_public.impact WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.get_impact(fn_impact_id uuid) 
RETURNS risky_public.impact 
AS $$
    SELECT * FROM risky_public.impact WHERE id = fn_impact_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.delete_impact(fn_impact_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.impact WHERE id = fn_impact_id;
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.create_impact(fn_name varchar, fn_business_id uuid, fn_threat_id uuid) 
RETURNS uuid
AS $$
    INSERT INTO risky_public.impact(name, business_id, threat_id) values(fn_name, fn_business_id, fn_threat_id) RETURNING id;
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_impact(fn_impact_id uuid, fn_name varchar, fn_business_id uuid, fn_threat_id uuid) 
RETURNS void 
AS $$
    declare
        v_mitigation_cost DOUBLE PRECISION;
        v_exploitation_cost DOUBLE PRECISION;
    begin
    v_mitigation_cost := impact_mitigation_cost(fn_impact_id);
    v_exploitation_cost := impact_exploitation_cost(fn_impact_id);
    UPDATE risky_public.impact SET name = fn_name, business_id = fn_business_id, threat_id = fn_threat_id, exploitation_cost = v_exploitation_cost, mitigation_cost = v_mitigation_cost  WHERE id = fn_impact_id;
    RETURN;
    end;
$$ LANGUAGE plpgsql VOLATILE;


CREATE OR REPLACE FUNCTION impact_exploitation_cost(fn_impact_id uuid) 
RETURNS DOUBLE PRECISION 
AS $$
  declare
    v_threat_cost uuid;
    v_threat_id uuid;
  begin
    SELECT threat_id INTO v_threat_id FROM risky_public.impact WHERE id = fn_impact_id;
    SELECT SUM(total) INTO v_threat_cost FROM risky_public.liability WHERE threat_id = v_threat_id;
    RETURN v_threat_cost;
  end;
$$ LANGUAGE plpgsql STABLE;

CREATE OR REPLACE FUNCTION impact_mitigation_cost(fn_impact_id uuid)
RETURNS DOUBLE PRECISION
AS $$
    declare
        v_attack_chain_id uuid;
        v_mitigation_ids uuid[];
        v_detection_ids uuid[];
        v_action_id uuid;
        v_cost DOUBLE PRECISION;
        v_threat_id uuid;
    begin
        SELECT threat_id INTO v_threat_id FROM risky_public.impact WHERE id = fn_impact_id;

        SELECT id INTO v_attack_chain_id FROM risky_public.attack_chain WHERE threat_id = v_threat_id;

        SELECT action_id INTO v_action_id from risky_public.attack_chain_action WHERE attack_chain_id = v_attack_chain_id;

        v_mitigation_ids := ARRAY(SELECT id FROM risky_public.mitigation WHERE action_id = v_action_id AND implemented = FALSE);

        v_detection_ids := ARRAY(SELECT id FROM risky_public.detection WHERE action_id = v_action_id AND implemented = FALSE);

        SELECT SUM(cost) INTO v_cost FROM risky_public.liability WHERE mitigation_id IN v_mitigation_ids OR detection_id IN v_detection_ids;

        RETURN v_cost;
    end;
$$ LANGUAGE plpgsql STABLE;
