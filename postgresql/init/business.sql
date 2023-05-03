CREATE OR REPLACE FUNCTION risky_public.businesses() 
RETURNS SETOF risky_public.business 
AS $$
    SELECT * FROM risky_public.business;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.get_business(fn_business_id uuid) 
RETURNS risky_public.business 
AS $$
    SELECT * FROM risky_public.business WHERE id = fn_business_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.delete_business(fn_business_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.business WHERE id = fn_business_id;
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.create_business(fn_name varchar, fn_revenue double precision) 
RETURNS void
AS $$
    INSERT INTO risky_public.business(name, revenue) values(fn_name, fn_revenue);
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_business(fn_business_id uuid, fn_name varchar, fn_revenue double precision) 
RETURNS void 
AS $$
    UPDATE risky_public.business SET name = fn_name, revenue = fn_revenue WHERE id = fn_business_id;
$$ LANGUAGE sql VOLATILE;
