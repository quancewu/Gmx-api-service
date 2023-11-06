<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://github.com/quancewu/Gmx-api-service/blob/main/picture/favicon.svg" alt="Project logo"></a>
</p>

<h3 align="center">gmx-api-service</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/quancewu/Gmx-api-service.svg)](https://github.com/quancewu/Gmx-api-service/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/quancewu/Gmx-api-service.svg)](https://github.com/quancewu/Gmx-api-service/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> GMX5xx Series SQL database using golang for RESTfull API
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

GMX5xx Series SQL database

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

Get Linux golang devlop envirments

```
example@server$ go version
go version go1.20 linux/amd64
```

### Installing

Install docker on linux system

```
docker compose up -build -d
```


## 🔧 Running the tests <a name = "tests"></a>

check docker run

```
xecuting task: docker compose -f "docker-compose.yaml" up -d --build 

[+] Building 2.8s (8/8) FINISHED
 => [internal] load build definition from dockerfile                                                                                                                    
 => => transferring dockerfile: 32B                                                                                                                                     
 => [internal] load .dockerignore                                                                                                                                       
 => => transferring context: 2B                                                                                                                                         
 => [internal] load metadata for docker.io/library/golang:1.19                                                                                                          
 => [internal] load build context                                                                                                                                       
 => => transferring context: 125B                                                                                                                                       
 => [base 1/1] FROM docker.io/library/golang:1.19@sha256:3025bf670b8363ec9f1b4c4f27348e6d9b7fec607c47e401e40df816853e743a                                               
 => CACHED [dev 1/2] WORKDIR /opt/app/api                                                                                                                               
 => CACHED [dev 2/2] COPY ./gmx5xx-api-sv/bin/gmx5xx-api-sv /opt/app/api                                                                                                
 => exporting to image                                                                                                                                                  
 => => exporting layers                                                                                                                                                 
 => => writing image sha256:ff5d506e4402f7199c9038ca87e6f179de8a90a164f43cd04b6a618c9148e40e                                                                            
 => => naming to docker.io/library/gmx_api_service-app                                                                                                                  
[+] Running 3/3
 - Network gmx_api_net         Created                                                                                                                                  
 - Container gmx_met_database  Started                                                                                                                                  
 - Container gmx_api_app       Started
```

### Break down into end to end tests


### And coding style tests



## 🎈 Usage <a name="usage"></a>

Add notes about how to use the system.

## 🚀 Deployment <a name = "deployment"></a>

Add additional notes about how to deploy this on a live system.

## ⛏️ Built Using <a name = "built_using"></a>

## ✍️ Authors <a name = "authors"></a>

- [@quancewu](https://github.com/quancewu) - Idea & Initial work

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

