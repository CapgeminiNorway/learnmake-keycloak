 Golang with KEYCLOAK  
=====================  

Golang app using KEYCLOAK!    

### Prerequisites

In order to build and run this app you need to have a couple of things installed:  

- The [Go SDK](https://golang.org)              
- An IDE for the development, like [Atom](https://atom.io) or IntelliJ/Goland      
- The Docker Toolbox or native Docker, whatever you prefer. See [Docker](https://docs.docker.com) and [Docker-Compose](https://docs.docker.com/compose/)       

### Building & Running this App    
```bash

# Clone this repo and verify compile  
$ git clone GIT_REPO_URL   
$ cd path/to/keycloak-golang    

# Run as App 
$ go run demoserver.go

# build it then run it
$ ./demoserver

```
Now it should be available via [http://localhost:8026](http://localhost:8026)    
  

### Containerization with Docker  

Building, publishing and running via _Docker_ and _Docker-Compose_:       
```bash
# set env vars for ease-of-use
# NOTE! please just replace 'zeusbaba' with your user  
$ export dockerhubUser=zeusbaba \
  export appName=keycloak-golang \
  export appVersion=1.0.0
$ export dockerImage=${dockerhubUser}/${appName}:${appVersion}

  # required for compatibility
$ GOOS=linux GOARCH=amd64 go build demoserver.go

## using Docker!!!       
# build a docker image  
$ docker build \
  -t ${dockerImage} \
  --rm --no-cache .    
$ docker images  	
# (optional) publish the image to docker hub  
$ docker push ${dockerImage}  

# (optional) run the docker image locally    
$ docker run \
	-p 8026:8026 \
	${dockerImage}  
```
Now it should be available via [http://localhost:8026](http://localhost:8026)  
