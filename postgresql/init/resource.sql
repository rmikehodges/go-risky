CREATE OR REPLACE FUNCTION risky_public.resources(fn_business_id uuid) 
RETURNS SETOF risky_public.resource
AS $$
    SELECT * FROM risky_public.resource WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.get_resource(fn_resource_id uuid) 
RETURNS risky_public.resource 
AS $$
    SELECT *  FROM risky_public.resource WHERE id = fn_resource_id;
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.delete_resource(fn_resource_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.resource WHERE id = fn_resource_id;
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.create_resource(fn_name varchar, fn_cost DOUBLE PRECISION, fn_unit TEXT, fn_total DOUBLE, fn_resouce_type risky_public.resources_type, fn_business_id uuid) 
RETURNS uuid
AS $$
    INSERT INTO risky_public.resource(name, cost, unit, total, resouce_type, business_id)  values(fn_name, fn_cost, fn_unit, fn_total, fn_resouce_type, fn_business_id) RETURNING id;
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_resource(fn_resource_id uuid, fn_name varchar, fn_cost DOUBLE PRECISION, fn_unit TEXT, fn_total DOUBLE, fn_resouce_type risky_public.resources_type, fn_business_id uuid)
RETURNS void
AS $$
    UPDATE risky_public.resource SET name = fn_name, cost = fn_cost, unit = fn_unit, total = fn_total, resouce_type = fn_resouce_type, fn_business_id = fn_business_id WHERE id = fn_resource_id;
$$ LANGUAGE sql VOLATILE;
