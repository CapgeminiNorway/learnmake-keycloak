 SpringBoot with KEYCLOAK  
==========================  

SpringBoot app using KEYCLOAK!    

*NB! This codebase is an adaptation from this [baeldung-tutorial](https://www.baeldung.com/spring-boot-keycloak)*   
	 
### Building & Running this App    
```

# Clone this repo and verify compile  
  git clone GIT_REPO_URL   
  cd path/to/keycloak-springboot  
  mvn clean compile  

# Run as spring boot app 
  mvn clean spring-boot:run

```
Now it should be available via [http://localhost:8022](http://localhost:8022)  


### Containerization with Docker  

Building, publishing and running via _Docker_ and _Docker-Compose_:       
```bash
# set env vars for ease-of-use
# NOTE! please just replace 'zeusbaba' with your user  
$ export dockerhubUser=zeusbaba \
  export appName=keycloak-springboot \
  export appVersion=1.0.0
$ export dockerImage=${dockerhubUser}/${appName}:${appVersion}

# build the app package, aka fat jar 
$ mvn clean package spring-boot:repackage

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
	-p 8022:8022 \
	${dockerImage}  
```
Now it should be available via [http://localhost:8022](http://localhost:8022)  
