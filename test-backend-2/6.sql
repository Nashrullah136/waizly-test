select name, salary, avgSal
from employees join (select top 1 department, avg(salary) as avgSal from employees group by department order by avg(salary)) temp
on employees.department = temp.department;

