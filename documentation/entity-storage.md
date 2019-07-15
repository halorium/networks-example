## Entity Storage

All (items, objects, resources, entities) are stored in the entites PostgreSql table as JSONB data with their HTTP resource path and ID as the unique identifier. The code for this can be found in /entities. This is the interface for putting, getting, and deleting from the DB. This allows us to keep the code DRY by having a single place for SQL and a single interface for the DB.
