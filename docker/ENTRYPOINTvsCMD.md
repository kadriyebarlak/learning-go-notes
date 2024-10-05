man 2 execve
execve(const char *path, char *const argv[], char *const envp[]);

```
date -u
Fri Sep 27 13:15:47 UTC 2024
/bin/date -u
Fri Sep 27 13:15:54 UTC 2024

/bin/echo hello 1 2 3
hello 1 2 3
/bin/echo 'hello 1 2 3'
hello 1 2 3

```

```
FROM alpine
ENTRYPOINT ["/bin/echo","0"]
```
docker run foo 1 2 3
0 1 2 3


```
kadriyebarlak@kadriyes-mbp docker-example-02 % cat Dockerfile 
FROM alpine
ENTRYPOINT ["/bin/echo"]
CMD ["0"]

kadriyebarlak@kadriyes-mbp docker-example-02 % docker run foo 1 2 
1 2
kadriyebarlak@kadriyes-mbp docker-example-02 % docker run foo     
0
```


```
kadriyebarlak@kadriyes-mbp docker-example-02 % cat Dockerfile 
FROM alpine
CMD ["/bin/echo", "0"]

kadriyebarlak@kadriyes-mbp docker-example-02 % docker run foo
0
kadriyebarlak@kadriyes-mbp docker-example-02 % docker run foo 1 2 3
docker: Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: exec: "1": executable file not found in $PATH: unknown.
kadriyebarlak@kadriyes-mbp docker-example-02 %
```
Entry Point is  /bin/sh -c  and CMD is added to entry point and result is /bin/sh -c '/bin/echo 0'
Now 1 is added to entry point and new result is "/bin/sh -c 1"


```
cat Dockerfile 
FROM alpine
RUN apk add --no-cache curl
ENTRYPOINT ["curl"]

kadriyebarlak@kadriyes-mbp docker-example-02 % docker run mycurl 
curl: try 'curl --help' or 'curl --manual' for more information
kadriyebarlak@kadriyes-mbp docker-example-02 % docker run mycurl http://example.com
```
Recommedn is ENTRYPOINT



```
cat Dockerfile 
FROM alpine
RUN apk add --no-cache curl
ENTRYPOINT curl

kadriyebarlak@kadriyes-mbp docker-example-02 % docker run mycurl example.com
curl: try 'curl --help' or 'curl --manual' for more information
kadriyebarlak@kadriyes-mbp docker-example-02 % /bin/sh -c "curl" "example.com" 
curl: try 'curl --help' or 'curl --manual' for more information
```


ENTRYPOINT ["python", "server.py"] // recommend

ENTRYPOINT ["python", "server.py", "--host", "0.0.0.0", "--port", "8080"]

ENTRYPOINT file1.sh && file2.sh \
                file3.sh && python server.py //not recommend

```
entrypoint.sh
#!/bin/sh
echo "Starting server..."
python server.py --host 0.0.0.0 --port 8080

Dockerfile
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
```








