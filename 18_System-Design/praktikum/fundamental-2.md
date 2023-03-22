## Redis: 
GET users *
## neo4j:
MATCH (users:Users)
RETURN users
## cassandra
SELECT * FROM db.users;