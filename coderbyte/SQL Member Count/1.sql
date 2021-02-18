/* write your SQL query below */

SELECT
  x.ReportsTo,
  COUNT(1) as Members,
  ROUND(AVG(x.age)) as 'Average AGE'
FROM 
  maintable_M1N9W x
WHERE
  x.ReportsTo IS NOT NULL
GROUP BY
  x.ReportsTo
ORDER BY
  x.ReportsTo
