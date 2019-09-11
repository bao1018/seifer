# Seifer 

Related repos:
> https://github.com/bao1018/gpt-2-master

## Overall Tech Arch Diagram

![Image of Arch Design](https://i.imgur.com/I1Y3GiG.png)

### Install
> Golang 1.12 and make go_mod be ON

1. Create a DB called selfile in Postgres
2. Run the DDL SQL to create the DB table
```
/migration/000001_create_story_table.up.sql
```
3. Add `config.json` in root directory with below reference
```json
{
    "addr": "127.0.0.1",
    "port": "2333",
    "database": {
        "connection": "postgresql://pguser:pgpassword@localhost/seifer"
    }
}
```
4. Run below commands to install/run the go server
```shell
go run main.go
```



