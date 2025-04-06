SELECT id, name 
FROM menu
UNION ALL 
SELECT id, name 
FROM person
ORDER BY object_id, object_name;