SELECT p.id, pz.id 
FROM person p 
CROSS JOIN
pizzeria pz  
ORDER BY p.id, pz.id;