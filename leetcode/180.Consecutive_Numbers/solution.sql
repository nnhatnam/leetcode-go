select distinct x.num as ConsecutiveNums
from
    (
        select num,
               LEAD(num, 1) over(order by id) as next_num,
               LEAD(num, 2) over(order by id) as next_next_num
        from logs
    ) x
where x.num = x.next_num
and x.num = x.next_next_num