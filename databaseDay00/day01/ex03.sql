SELECT person_order.order_date, person_order.person_id
FROM person_order, person_visits
WHERE person_order.order_date = person_visits.visit_date AND person_order.person_id = person_visits.person_id
ORDER BY person_order.order_date ASC, person_order.person_id DESC;