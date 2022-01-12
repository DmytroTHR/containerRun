## Make program in Docker container run some copies of another program in several Docker containers

To run simply do this in the root directory of a project:
```
docker-compose up --build
```

After that you should be able to see results by entering the following in your browser:
```
http://localhost:808x/scooter
```
where x = [0:4]

####Eg.
You should see this text on http://localhost:8083/scooter
```
Hello scooter, your id = 113
```

###Important
By now you should find & stop generated containers manually by:
```
docker ps
//find your containers (has "service_b" in command)
docker stop container_name
```
or with the help of your IDE.