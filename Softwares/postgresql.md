## psql
connect: `psql -d [dbname] -U [username] -h [host]`
```
sudo -i -u postgres
sudo su - postgres

\l - Display database
\c - Connect to database
\dn - List schemas
\dt - List tables inside public schemas
\dt schema1.* - List tables inside a particular schema.
                For example: 'schema1'.
```

## errors
#### 1. Caused by: org.postgresql.util.PSQLException: ERROR: duplicate key value violates unique constraint 
>https://stackoverflow.com/a/21639138/8936864

get current sequence
```sql
SELECT currval(pg_get_serial_sequence('prod.user_subscription','id'));
```
> DONOT THIS: get next sequence
> ```sql
> SELECT nextval(pg_get_serial_sequence('schema_name.table_name','table_primary_key'));
> ```

update sequence
```sql
SELECT setval(pg_get_serial_sequence('the_schema.the_table', 'the_primary_key'), (SELECT MAX(the_primary_key) FROM the_table) + 1);
```

## dump
```
docker exec -t <container_id_or_name> pg_dump -U <username> -d <database_name> -t <table_name> > /path/on/host/machine/dump_file.sql
```
## dump data only
```
docker exec -t <container_id_or_name> pg_dump -U <username> -d <database_name> -t <table_name> --data-only --column-inserts > /path/on/host/machine/dump_file.sql
```