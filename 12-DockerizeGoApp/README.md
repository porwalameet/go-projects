Observation regarding Docker image size:
Without multistage build, the same docker image is of size 304MB.
With multistage build by using scratch, the docker image is just 7MB.
Output:
```
REPOSITORY                                                                  TAG                       IMAGE ID       CREATED         SIZE
dockerizegoapp                                                              v1                        dc98959accc9   4 seconds ago   7MB
dockerizegoapp                                                              latest                    4f2d0caa1fcc   6 minutes ago   304MB
```