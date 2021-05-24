## XLSX to DB Schema

### 1. Getting Started:

### 2. Supported format files:

- .xlsx, .csv

### 3. XLSX to DB Schema APIs:

#### 3.1. Execute:

`
./console -table_name=corporation_salary_systems -apdapter={:adapter} -source={:sourcePATH} -dest={:destPATH}
`

Example:

`
./console -adapter=mysql -table_name=corporation_salary_systems -source=/Users/admin/Desktop/workspace/ponos/payroll_apis_design/corporation_salary_systems.csv -dest=/Users/admin/Desktop/test.sql
`

- adapter: mysql or postgresql
- sourcePATH: Storage location of csv or xslx files.
- destPATH: Storage location of output db schema.

#### 3.2. Help:

`
./console help
`

```
Type of consoles:
    - exec: Execute to generate db schema.
    - help: Exetute to show options of console.
For execte:
    -adapter: Support types of DBMS (mysql, postgresql).
    -sourcePATH: Storage location of csv or xslx files.
    -destPATH: Storage location of output db schema.
```
