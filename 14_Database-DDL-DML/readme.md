# Table modelling
pemodelan tabel database dapat menggunakan ERD Model menggunakan aplikasi seperti draw.io, mysql workbench, atau aplikasi serupa

## Relation
### 1 to 1
1 data hanya dapat dimiliki oleh 1 data dari tabel lainnya  
contoh: 
- 1 foto profil dapat dimiliki oleh 1 user

### 1 to Many
1 data dapat digunakan oleh banyak data dari tabel lainnya  
contoh:
- 1 user dapat membuat banyak post, post hanya dapat dibuat oleh 1 user

### Many to Many
1 data dapat digunakan oleh banyak data dari tabel lainnya, berlaku juga kebalikannya  
contoh:
- 1 user dapat memiliki banyak follower, follower tersebut juga dapat memiliki banyak follower lainnya

# RDBMS
salah satu tipe database yang menggunakan relation model, seperti mySQL, postgreSQL, dsb

# DDL Command
## CREATE TABLE:
membuat table baru
```sql
CREATE table table_name
(
    Column1 datatype (size),
    column2 datatype (size),
    .
    .
    columnN datatype(size)
);
```
## DROP Table
menghapus table
```sql
DROP TABLE table_name;
```

## ALTER TABLE
mengubah struktur table
```sql
ALTER TABLE table_name
    ADD (Columnname_1  datatype,
        Columnname_2  datatype,
        …
        Columnname_n  datatype
    )
    ;
```
# DML Command
## INSERT
menambahkan data baru
```sql
INSERT INTO table_name (field...) VALUES (field_value...);
```

## SELECT
mengambil data dari suatu table dengan kriteria yang ditentukan
```sql
SELECT field_name FROM table_name WHERE option... 
```

## UPDATE
mengubah data dari suatu table dengan kriteria yang ditentukan
```sql
UPDATE table_name SET ...
```

## DELETE
menghapus data dari suatu table dengan kriteria yang ditentukan
```sql
DELETE FROM table_name WHERE option...
```

## DML option
control flow untuk perintah dml
- LIKE/BETWEEN
- AND/OR
- LIMIT
- ORDER BY