DROP FUNCTION risky_public.assets;
CREATE FUNCTION risky_public.assets(fn_business_id uuid) 
RETURNS SETOF risky_public.asset 
AS $$
    SELECT * FROM risky_public.asset WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.get_asset;
CREATE FUNCTION risky_public.get_asset(fn_asset_id uuid) 
RETURNS risky_public.asset 
AS $$
    SELECT * FROM risky_public.asset WHERE id = fn_asset_id;
$$ LANGUAGE sql;

DROP FUNCTION risky_public.delete_asset;
CREATE FUNCTION risky_public.delete_asset(fn_asset_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.asset WHERE id = fn_asset_id;
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.create_asset;
CREATE FUNCTION risky_public.create_asset(fn_name varchar, fn_description varchar, fn_business_id uuid) 
RETURNS uuid
AS $$
    INSERT INTO risky_public.asset(name, description, business_id) values(fn_name, fn_description, fn_business_id) RETURNING id;
$$ LANGUAGE sql VOLATILE;

DROP FUNCTION risky_public.update_asset;
CREATE FUNCTION risky_public.update_asset(fn_asset_id uuid, fn_name varchar, fn_description varchar, fn_business_id uuid) 
RETURNS void 
AS $$
    UPDATE risky_public.asset SET name = fn_name, description = fn_description, business_id = fn_business_id WHERE id = fn_asset_id;
$$ LANGUAGE sql VOLATILE;
