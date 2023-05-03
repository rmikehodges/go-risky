CREATE OR REPLACE FUNCTION risky_public.detections(fn_business_id uuid) 
RETURNS SETOF risky_public.detection 
AS $$
    SELECT * FROM risky_public.detection WHERE business_id = fn_business_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.get_detection(fn_detection_id uuid) 
RETURNS risky_public.detection 
AS $$
    SELECT * FROM risky_public.detection WHERE id = fn_detection_id;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION risky_public.delete_detection(fn_detection_id uuid) 
RETURNS void 
AS $$
    DELETE FROM risky_public.detection WHERE id = fn_detection_id;
$$ LANGUAGE sql VOLATILE;

CREATE OR REPLACE FUNCTION risky_public.create_detection(fn_name varchar, fn_description varchar, fn_business_id uuid,  fn_action_id uuid, fn_implemented boolean)
RETURNS void
AS $$
    INSERT INTO risky_public.detection(name, description, business_id, action_id, implemented) values(fn_name, fn_description, fn_business_id, fn_action_id, fn_implemented);
$$ LANGUAGE sql VOLATILE;


CREATE OR REPLACE FUNCTION risky_public.update_detection(fn_detection_id uuid, fn_name varchar, fn_description varchar, fn_business_id uuid,  fn_action_id uuid, fn_implemented boolean) 
RETURNS void 
AS $$
    UPDATE risky_public.detection SET name = fn_name, description = fn_description, action_id = fn_action_id, implemented = fn_implemented WHERE id = fn_detection_id;
$$ LANGUAGE sql VOLATILE;
