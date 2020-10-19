
# 184. 部门工资最高的员工
select D.Name AS Department,E.Name AS Employee,E.Salary AS Salary
from Employee E
         join Department AS D on
        E.DepartmentId = D.ID
where (E.DepartmentId,E.Salary) in (
    select DepartmentId,Max(Salary) from Employee group by DepartmentId
)