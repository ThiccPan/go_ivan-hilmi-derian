# Code Competence Golang - Ivan Hilmi  Derian
## Installation
```bash
# 1. install app container
docker compose up -d --build
# 2. open db container
docker exec -it alta_cc_db bash
# 3. open mysql client
mysql -u admin -p123 alta_cc_db
# 4. insert genre
INSERT INTO categories (name) VALUES ("*category_name*");
# done
```
## Api spec
https://documenter.getpostman.com/view/23637484/2s93m34jHL

## app screenshot
- In docs folder