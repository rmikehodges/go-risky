DROP FUNCTION risky_public.impacts;
CREATE FUNCTION risky_public.impacts(fn_business_id uuid) 
RETURNS SETOF risky_public.impact 
AS $$
    SELECT * FROM risky_public.impact WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.get_impact;
CREATE FUNCTION risky_public.get_impact(fn_impact_id uuid) 
RETURNS risky_public.impact 
AS $$
    SELECT * FROM risky_public.impact WHERE id = fn_impact_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.delete_impact;
CREATE FUNCTION risky_public.delete_impact(fn_impact_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.impact WHERE id = fn_impact_id;
$$ LANGUAGE sql VOLATILE;


CREATE TYPE risky_public.calculated_impact AS (exploitation_cost DOUBLE PRECISION, mitigation_cost DOUBLE PRECISION);

DROP FUNCTION risky_public.calculate_impact;
CREATE FUNCTION risky_public.calculate_impact(fn_business_id uuid, fn_threat_id uuid)
RETURNS risky_public.calculated_impact
AS $$
    declare
        v_mitigation_liability DOUBLE PRECISION;
        v_detection_liability DOUBLE PRECISION;
        v_calculated_impact risky_public.calculated_impact;
        v_detection_implemented boolean;
        v_mitigation_implemented boolean;
        v_attack_chains RECORD;
    begin
        v_calculated_impact.exploitation_cost := 0.0;
        v_calculated_impact.mitigation_cost := 0.0;

        SELECT SUM(cost) FROM risky_public.liability WHERE business_id = fn_business_id AND threat_id = fn_threat_id INTO v_calculated_impact.exploitation_cost;

        FOR v_attack_chains IN 
        SELECT risky_public.attack_chain_step.detection_id, risky_public.attack_chain_step.mitigation_id FROM risky_public.attack_chain LEFT JOIN risky_public.attack_chain_step ON risky_public.attack_chain.id = risky_public.attack_chain_step.attack_chain_id WHERE risky_public.attack_chain.threat_id = fn_threat_id
        LOOP
            SELECT implemented FROM risky_public.detection WHERE id = v_attack_chains.detection_id INTO v_detection_implemented ;
            IF  v_detection_implemented = FALSE
            THEN
                SELECT cost FROM risky_public.liability WHERE detection_id = v_attack_chains.detection_id INTO v_detection_liability;
                v_calculated_impact.mitigation_cost := v_calculated_impact.mitigation_cost + v_detection_liability;
            END IF;
            SELECT implemented FROM risky_public.mitigation WHERE id = v_attack_chains.mitigation_id INTO v_mitigation_implemented;
            IF  v_mitigation_implemented = FALSE
            THEN
                SELECT cost FROM risky_public.liability WHERE mitigation_id = v_attack_chains.mitigation_id INTO v_mitigation_liability;
                v_calculated_impact.mitigation_cost := v_calculated_impact.mitigation_cost + v_mitigation_liability;
            END IF;
        END LOOP;
        RETURN v_calculated_impact;
    end;
$$ LANGUAGE plpgsql VOLATILE;
    

DROP FUNCTION risky_public.create_impact;
CREATE FUNCTION risky_public.create_impact(fn_name varchar, fn_description varchar,  fn_business_id uuid, fn_threat_id uuid) 
RETURNS uuid
AS $$
    declare
        v_impact_id uuid;
        v_calculated_impact risky_public.calculated_impact;
    begin

    v_calculated_impact:= risky_public.calculate_impact(fn_business_id, fn_threat_id);

    INSERT INTO risky_public.impact(name, description, business_id, threat_id, exploitation_cost, mitigation_cost) values(fn_name, fn_description, fn_business_id, fn_threat_id, v_calculated_impact.exploitation_cost, v_calculated_impact.mitigation_cost) RETURNING id INTO v_impact_id;
    RETURN v_impact_id;
    end;
$$ LANGUAGE plpgsql VOLATILE;

DROP FUNCTION risky_public.update_impact;
CREATE FUNCTION risky_public.update_impact(fn_impact_id uuid, fn_name varchar, fn_description varchar, fn_business_id uuid, fn_threat_id uuid) 
RETURNS void 
AS $$
    declare
        v_calculated_impact risky_public.calculated_impact;
    begin
    v_calculated_impact:= risky_public.calculate_impact(fn_business_id, fn_threat_id);
    UPDATE risky_public.impact SET name = fn_name, description = fn_description, business_id = fn_business_id, threat_id = fn_threat_id, exploitation_cost = v_calculated_impact.exploitation_cost, mitigation_cost = v_calculated_impact.mitigation_cost  WHERE id = fn_impact_id;
    RETURN;
    end;
$$ LANGUAGE plpgsql VOLATILE;


