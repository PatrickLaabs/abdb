# Animal Blood Donation Bank - ABDB
(Ja, es muss noch ein besserer Name her..)

## Spin up the project:

run `docker-compose.yml` file inside root dir.
Latest working go code will be pulled and the mysql file with all the needed configuration, too.

./ngrock http 8080

---

# Documentation

## ABDB-db & ABDB-website folders

Inside these folders are docker-compose files to seperatly spin up the database (currently mysql) and the go code for 
our website.

This is good practice, if we want to develop our go code and have a db quickly up and running.
See `local development`

## local development
To spin up the database for a clean start, just go into the folder ./abdb-db and run the docker-compose file.
Now you can't develop on the go code itself and quickly test the interaction with the database.

If everythings working, push and release it to gh.

> // This connectionstring is needed, if we wanna push this code to a container.
> // db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+"172.19.0.2"+":"+dbPort+")/"+dbName)
>
> // This connectionstring is for local development purposes only.
> db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+"0.0.0.0"+":"+dbPort+")/"+dbName)

### Hot Reload of changes

use air - making lives easier :P 

### Todos:

- [ ] create pipeline from scm checkout to deploying on my raspberry or k8s in cloud
- [ ] improve documentation