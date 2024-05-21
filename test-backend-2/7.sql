select rank() over (order by total_sales) as ranking, name, total_sales
from employees
         join (select employee_id, sum(sales) as total_sales from sales_data group by employee_id) temp
              on temp.employee_id = employees.employee_id
order by total_sales;