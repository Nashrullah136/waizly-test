CREATE PROCEDURE GetEmployeesDept
@DepartmentName varchar(64)
AS
BEGIN
    SELECT *
    FROM Employees
    WHERE department = @DepartmentName;
END;