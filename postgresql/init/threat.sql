CREATE OR REPLACE FUNCTION risky_public.threats(fn_business_id uuid) 
RETURNS SETOF risky_public.threat 
AS $$
    SELECT * FROM risky_public.threat WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.get_threat(fn_threat_id uuid) 
RETURNS risky_public.threat 
AS $$
    SELECT * FROM risky_public.threat WHERE id = fn_threat_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.delete_threat(fn_threat_id uuid) 
RETURNS uuid 
AS $$
    DELETE FROM risky_public.threat WHERE id = fn_threat_id; 
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.create_threat(fn_name varchar, fn_description varchar, fn_business_id uuid) 
RETURNS uuid
AS $$
    INSERT INTO risky_public.threat(name, description, business_id) values(fn_name, fn_description,fn_business_id) RETURNING id;
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_threat(fn_threat_id uuid, fn_name varchar, fn_description varchar, fn_business_id uuid) 
RETURNS uuid 
AS $$
    UPDATE risky_public.threat SET name = fn_name, description = fn_description,business_id = fn_business_id  WHERE id = fn_threat_id RETURNING fn_threat_id;
$$ LANGUAGE sql VOLATILE;
