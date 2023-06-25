DROP FUNCTION risky_public.resources;
CREATE FUNCTION risky_public.resources(fn_business_id uuid) 
RETURNS SETOF risky_public.resource
AS $$
    SELECT * FROM risky_public.resource WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.get_resource;
CREATE FUNCTION risky_public.get_resource(fn_resource_id uuid) 
RETURNS risky_public.resource 
AS $$
    SELECT *  FROM risky_public.resource WHERE id = fn_resource_id;
$$ LANGUAGE sql VOLATILE;


DROP FUNCTION risky_public.delete_resource;
CREATE FUNCTION risky_public.delete_resource(fn_resource_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.resource WHERE id = fn_resource_id;
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.create_resource;
CREATE FUNCTION risky_public.create_resource(fn_name varchar, fn_description varchar, fn_cost DOUBLE PRECISION, fn_unit TEXT, fn_total DOUBLE PRECISION, fn_resource_type risky_public.resource_type, fn_business_id uuid) 
RETURNS uuid
AS $$
    INSERT INTO risky_public.resource(name, description, cost, unit, total, resource_type, business_id)  values(fn_name, fn_description, fn_cost, fn_unit, fn_total, fn_resource_type, fn_business_id) RETURNING id;
$$ LANGUAGE sql VOLATILE;


DROP FUNCTION risky_public.update_resource;
CREATE FUNCTION risky_public.update_resource(fn_resource_id uuid, fn_name varchar, fn_description varchar, fn_cost DOUBLE PRECISION, fn_unit TEXT, fn_total DOUBLE PRECISION, fn_resource_type risky_public.resource_type, fn_business_id uuid)
RETURNS void
AS $$
    UPDATE risky_public.resource SET name = fn_name, description = fn_description, cost = fn_cost, unit = fn_unit, total = fn_total, resource_type = fn_resource_type, business_id = fn_business_id WHERE id = fn_resource_id;
$$ LANGUAGE sql VOLATILE;
