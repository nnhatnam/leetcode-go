SELECT T0.FirstName, T0.LastName, T1.City, T1.State
FROM Person T0
LEFT JOIN Address T1 ON T0.PersonId = T1.PersonId
