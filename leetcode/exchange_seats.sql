#换座位
select (
           case
               when mod(id,2) != 0 and count != id then id+1
               when mod(id,2) != 0 and count = id then id
               else id-1
               end) AS id,
       student
from seat AS seat_info,
     (select count(*) AS count
      from seat) AS seat_count order by  id asc