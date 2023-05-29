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

CREATE OR REPLACE FUNCTION risky_public.create_impact(fn_name varchar, fn_description varchar,  fn_business_id uuid, fn_threat_id uuid) 
RETURNS uuid
AS $$
    INSERT INTO risky_public.impact(name, description, business_id, threat_id) values(fn_name, fn_description, fn_business_id, fn_threat_id) RETURNING id;
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.update_impact;
CREATE FUNCTION risky_public.update_impact(fn_impact_id uuid, fn_name varchar, fn_description varchar, fn_business_id uuid, fn_threat_id uuid) 
RETURNS void 
AS $$
    declare
        v_mitigation_cost DOUBLE PRECISION;
        v_exploitation_cost DOUBLE PRECISION;
    begin
    IF fn_threat_id IS NOT NULL 
    THEN
    v_mitigation_cost := impact_mitigation_cost(fn_threat_id);
    v_exploitation_cost := impact_exploitation_cost(fn_threat_id);
    ELSE
        v_mitigation_cost := 0.0;
        v_exploitation_cost := 0.0;
    END IF;

    UPDATE risky_public.impact SET name = fn_name, description = fn_description, business_id = fn_business_id, threat_id = fn_threat_id, exploitation_cost = v_exploitation_cost, mitigation_cost = v_mitigation_cost  WHERE id = fn_impact_id;
    RETURN;
    end;
$$ LANGUAGE plpgsql VOLATILE;

DROP FUNCTION risky_public.impact_exploitation_cost;
CREATE FUNCTION risky_public.impact_exploitation_cost(fn_threat_id uuid) 
RETURNS DOUBLE PRECISION 
AS $$
  declare
    v_threat_cost DOUBLE PRECISION;
  begin
    SELECT SUM(cost) INTO v_threat_cost FROM risky_public.liability WHERE threat_id = fn_threat_id;
    RETURN v_threat_cost;
  end;
$$ LANGUAGE plpgsql STABLE;

DROP FUNCTION risky_public.impact_mitigation_cost;
CREATE FUNCTION risky_public.impact_mitigation_cost(fn_threat_id uuid)
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
        SELECT id INTO v_attack_chain_id FROM risky_public.attack_chain WHERE threat_id = fn_threat_id;

        SELECT action_id INTO v_action_id from risky_public.attack_chain_step WHERE attack_chain_id = v_attack_chain_id;

        v_mitigation_ids := ARRAY(SELECT id FROM risky_public.mitigation WHERE action_id = v_action_id AND implemented = FALSE);

        v_detection_ids := ARRAY(SELECT id FROM risky_public.detection WHERE action_id = v_action_id AND implemented = FALSE);

        SELECT SUM(cost) INTO v_cost FROM risky_public.liability WHERE mitigation_id = ANY(v_mitigation_ids) OR detection_id = ANY(v_detection_ids);

        RETURN v_cost;
    end;
$$ LANGUAGE plpgsql STABLE;
