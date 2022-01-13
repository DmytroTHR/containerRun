## Make program in Docker container run some copies of another program in several Docker containers

To run simply do this in the root directory of a project:
```
make run-app
```

After that you should be able to see results by entering the following in your browser:
```
http://localhost:808x/scooter
```
where x = [0:4]

To stop simply run
```
make stop-app 
```

### Eg.
You should see this text on http://localhost:8083/scooter
```
Hello scooter, your id = 113
```
