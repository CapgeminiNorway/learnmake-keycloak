 VueJS with KEYCLOAK  
=====================  

Vue.js app using KEYCLOAK!    


### Building & Running this App    
```

# Clone this repo and verify compile  
  git clone GIT_REPO_URL   
  cd path/to/keycloak-vuejs  
  yarn  

# Run as App 
  yarn serve

```
Now it should be available via [http://localhost:8024](http://localhost:8024)    
  

### Containerization with Docker  

Building, publishing and running via _Docker_ and _Docker-Compose_:       
```bash
# set env vars for ease-of-use
# NOTE! please just replace 'zeusbaba' with your user  
$ export dockerhubUser=zeusbaba \
  export appName=keycloak-vuejs \
  export appVersion=1.0.0
$ export dockerImage=${dockerhubUser}/${appName}:${appVersion}

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
	-p 8024:8024 \
	-e "NODE_ENV=development" \
	${dockerImage}  
```
Now it should be available via [http://localhost:8024](http://localhost:8024)  
