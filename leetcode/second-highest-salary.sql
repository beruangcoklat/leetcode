# Write your MySQL query statement below


select ifnull((select distinct e.Salary from Employee e order by e.Salary desc limit 1 offset 1), null) as 'SecondHighestSalary'
