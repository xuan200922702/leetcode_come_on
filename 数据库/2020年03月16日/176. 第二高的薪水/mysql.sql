select max(Salary)  as SecondHighestSalary from Employee
where Salary < (select max(Salary) from Employee)  ;


-- 使用临时表
select (select distinct(Salary)  from Employee order by Salary  desc limit 1,1)
as SecondHighestSalary;

-- 使用ifnull函数
select ifnull((select distinct(Salary)  from Employee order by Salary
 desc limit 1,1),null )as SecondHighestSalary;