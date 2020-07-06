select max(Salary) AS SecondHighestSalary
from employee
where
salary<(select max(salary) from employee)