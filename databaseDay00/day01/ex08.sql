SELECT person_order.order_date, CONCAT(person.name, ' (age:', person.age, ')') AS information
FROM person_order
NATURAL JOIN person
ORDER BY person_order.order_date ASC, information ASC;