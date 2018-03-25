# MollyDB

MollyDB is a **file system data store** that will provide us with data stored
in configuration files. So far, **json** and **yaml** files are handled by 
mollyDB but others like **toml** or **conf** are in the road map.

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
    

## Endpoints

| Uri | Description |
| --- | --- |
| POST /v0/folder:create | Create a new folder |
| POST /v0/folder:create | Create a new folder |
| POST /v0/folder:create | Create a new folder |
| git diff | Show file differences that haven't been staged |

## Hooks


