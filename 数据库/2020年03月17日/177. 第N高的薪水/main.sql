CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  if N<0 then
    RETURN (select min(Salary) from Employee);
  else
    set n =  N-1;
    RETURN (
   SELECT DISTINCT Salary FROM Employee ORDER BY Salary DESC LIMIT n,1
  );
    end if;
END



CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
  SELECT  IF(count<N,NULL,min)
  FROM
    (SELECT MIN(Salary) AS min, COUNT(1) AS count
    FROM
      (SELECT DISTINCT Salary
      FROM Employee ORDER BY Salary DESC LIMIT N) AS a
    ) as b
  );
END

