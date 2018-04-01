# MollyDB

MollyDB is a **file system data store** whose information. is stored into the 
file system. MollyDB works exclusively under **Graphql**

## Purpose
 

    So far, the purpose of mollydb is mainly being used for experimental 
    purposes. 


## Understanding mollyDB

### Diagram

####


## Running 

### Installating

Not available yet.

### From the code

- Clone the repository into your local machine: 

        git clone https://github.com/wesovilabs/mollydb.git
        
- Checkout the chosen tag.

        git checkout <mollydb.tag>
        
- Run any of the following command: 
    
        make run
or 
        make docker-run 
 


To launch yout queries, just open  [GrapIQL](http://localhost:9090) on your 
browser


### Docker

Docker images are published as far as a new tag on master branch is created. 
To get the latest version you can do it by running the below commands

    docker pull wesovilabs/mollydb
    docker run  wesovilabs/mollydb
    
In case of you want a specific version of mongodb you just need to set it on 
the name 
    
    export MOLLYDB_VERSION=<version>
    docker pull wesovilabs/mollydb:$MOLLYDB_VERSION
    docker run wesovilabs/mollydb:$MOLLYDB_VERSION


## Links

MollyDB documentation can be found on **Wiki**:

[Storage Hierarchy](https://github.com/wesovilabs/mollydb/wiki/Storage-hierarchy)

[GraphQL](https://github.com/wesovilabs/mollydb/wiki/GraphQL)
