## XLSX to DB Schema

### 1. Getting Started:

### 2. Supported format files:

- .xlsx, .csv

### 3. XLSX to DB Schema APIs:

#### 3.1. Execute:

`
./console -apdapter={:adapter} -source={:sourcePATH} -dest={:destPATH}
`

Example:

`
./console exec -apdapter=mysql -source=workspace/example.xlsx -dest=workspace/db_schemas
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