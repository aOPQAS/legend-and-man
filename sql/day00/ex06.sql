select 
  name, 
  case when name = 'Denis' then true else false end 
from 
  (
    select 
      (
        select 
          name 
        FROM 
          person 
        where 
          person.id = person_order.person_id
      ) 
    from 
      person_order 
    where 
      (
        menu_id = 13 
        or menu_id = 14 
        or menu_id = 18
      ) 
      AND order_date = '2022-01-07'
  ) as person
