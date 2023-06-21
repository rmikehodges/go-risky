DROP FUNCTION  risky_public.liabilities;
CREATE FUNCTION risky_public.liabilities(fn_business_id uuid) 
RETURNS SETOF risky_public.liability 
AS $$
    SELECT * FROM risky_public.liability WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.liabilities_by_threat_id;
CREATE FUNCTION risky_public.liabilities_by_threat_id;(fn_business_id uuid, fn_threat_id uuid) 
RETURNS SETOF risky_public.liability 
AS $$
       SELECT * FROM risky_public.liability WHERE business_id = fn_business_id AND threat_id = fn_threat_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.get_liabilities_by_impact_id;
CREATE FUNCTION risky_public.liabilities_by_impact_id;(fn_business_id uuid, fn_threat_id uuid) 
RETURNS SETOF risky_public.liability 
AS $$
    SELECT * FROM risky_public.liability WHERE business_id = fn_business_id AND impact_id = fn_impact_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.get_liabilities_by_mitigation_id;
CREATE FUNCTION risky_public.liabilities_by_impact_id;(fn_business_id uuid, fn_mitigation_id uuid) 
RETURNS SETOF risky_public.liability 
AS $$
    SELECT * FROM risky_public.liability WHERE business_id = fn_business_id AND mitigation_id = fn_mitigation_id;
$$ LANGUAGE sql;

-- DROP FUNCTION risky_public.get_liability_by_resource_id;
-- CREATE FUNCTION risky_public.get_liability_by_impact_id(fn_impact_id uuid) 
-- RETURNS risky_public.liability 
-- AS $$
--     SELECT * FROM risky_public.liability WHERE impact_id = fn_impact_id;
-- $$ LANGUAGE sql;


DROP FUNCTION risky_public.get_liability;
CREATE FUNCTION risky_public.get_liability(fn_liability_id uuid) 
RETURNS risky_public.liability 
AS $$
    SELECT * FROM risky_public.liability WHERE id = fn_liability_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.delete_liability;
CREATE FUNCTION risky_public.delete_liability(fn_liability_id uuid) 
RETURNS void
AS $$
    DELETE FROM risky_public.liability WHERE id = fn_liability_id ;
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.create_liability;
CREATE FUNCTION risky_public.create_liability(fn_name varchar, fn_description varchar, fn_quantity DOUBLE PRECISION, fn_business_id uuid, fn_mitigation_id uuid, fn_resource_id uuid, fn_detection_id uuid, fn_impact_id uuid, fn_threat_id uuid) 
RETURNS uuid
AS $$
    declare
        v_resource_cost DOUBLE PRECISION;
        v_total_cost DOUBLE PRECISION;
        v_liability_id uuid;
    begin
        IF fn_resource_id IS NOT NULL AND fn_quantity IS NOT NULL 
        THEN
            SELECT cost INTO v_resource_cost FROM resource WHERE id = fn_resource_id;
            v_total_cost := v_resource_cost * fn_quantity;
        ELSE
            v_total_cost := 0.0;
        END IF;

        INSERT INTO risky_public.liability(name, description, quantity, cost, business_id, mitigation_id, detection_id, resource_id, impact_id, threat_id) values(fn_name, fn_description, fn_quantity, v_total_cost, fn_business_id, fn_mitigation_id, fn_resource_id, fn_detection_id, fn_impact_id, fn_threat_id) RETURNING id INTO v_liability_id ;
    
        RETURN v_liability_id;
    end;
$$ LANGUAGE plpgsql VOLATILE;

DROP FUNCTION risky_public.update_liability;
CREATE FUNCTION risky_public.update_liability(fn_liability_id uuid,fn_name varchar, fn_description varchar, fn_quantity DOUBLE PRECISION, fn_business_id uuid, fn_mitigation_id uuid, fn_resource_id uuid, fn_detection_id uuid, fn_impact_id uuid, fn_threat_id uuid)
RETURNS void 
AS $$
    declare
        v_resource_cost DOUBLE PRECISION;
        v_total_cost DOUBLE PRECISION;
    begin
        IF fn_resource_id IS NOT NULL AND fn_quantity IS NOT NULL 
        THEN
            SELECT cost INTO v_resource_cost FROM resource WHERE id = fn_resource_id;
            v_total_cost := v_resource_cost * fn_quantity;
        ELSE
            v_total_cost := 0.0;
        END IF;
        UPDATE risky_public.liability SET name = fn_name, description = fn_description, quantity = fn_quantity, cost = v_total_cost, business_id = fn_business_id, mitigation_id = fn_mitigation_id, resource_id = fn_resource_id, detection_id = fn_detection_id, impact_id = fn_impact_id, threat_id = fn_threat_id WHERE id = fn_liability_id;
        RETURN;
    end;
$$ LANGUAGE plpgsql VOLATILE;
