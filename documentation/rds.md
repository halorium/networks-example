## Relational Database Service

AWS RDS is used for our PostgreSQL database.

### Migrations
Database Migrations are done using the migrator service found in /migrator/service.
Individual migrations are located in the migrations sub-directory of the service.
This service runs upon deployment and checks the migrations DB table to validate if the migration has already run or not and only runs new migrations not in the table.
Migration file names should begin with an incrementing number following the last migration number as well as a meaningful name describing the action of the migration separated by hyphens. The migration files are executable bash scripts using the psql command which uses the PostgreSQL credentials stored in the environment variables.
