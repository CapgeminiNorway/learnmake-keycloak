 LearnMake KEYCLOAK  
====================  

baseline setup and example apps using KEYCLOAK!    

## LEARN - MAKE - SHOW - SHARE   

### Prerequisites

In order to build and run this app you need to have a couple of things installed:  

- The Docker Toolbox or native Docker, whatever you prefer. See [Docker](https://docs.docker.com) and [Docker-Compose](https://docs.docker.com/compose/)  
- An IDE for the development, like [Atom](https://atom.io) or [IntelliJ](https://www.jetbrains.com/idea/)    
- [Node.js](https://nodejs.org/), [npm](https://www.npmjs.com/), and [Yarn](https://yarnpkg.com) installed, _see [package.json](package.json) for the required versions._    
- Get familiar with [VueJS](https://vuejs.org/)._                    
         

### KeyCloak instance with Docker    

Building and running via _Docker_ and _Docker-Compose_:         
```bash   
$ docker-compose up --build   

  # NOTE: in linux-env you might have permission problems with 'docker-data-*' folders      
  # to fix; stop docker-compose, set permissions as below, then run docker-compose again.    
$ sudo chown 1001:1001 -R docker-data-*  

  # to shut it down, ctrl+c and   
$ docker-compose down   
```
`docker-compose up` gets Keycloak instance up and running in your local environment.    
Now you can access KeyCloak admin panel via [http://localhost:8081](http://localhost:8081)      
default admin user+pass is inside [docker-compose.yaml](docker-compose.yaml)  


### Building example applications  
See list of reference implementations  
- [KEYCLOAK using Spring Boot](keycloak-springboot)  
- [KEYCLOAK using Node.js/Vue.js](keycloak-vuejs)  
- *more...MiP...*      


