SELECT person_order.order_date, CONCAT(person.name, ' (age:', person.age, ')') AS information
FROM person_order
JOIN person ON person_order.person_id = person.id
ORDER BY person_order.order_date ASC, information ASC;
