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
    

## Hierarchy

### Storage
A storage can be defined as a place where documents are hosted. 
So far, only local directories are supported but in the near future other 
places such as **ssh directory**, **database collections** or even git 
repository could be understand as **a storage**

The storage data model below:

- **documents**: The list of documents that belong to this storage
- **len**: The number of documents in the storage
- **name**: The name of a storage 

### Documents
As is said in the previous paragraph, a storage contains documents. So far, 
only yaml documents are supported but json and properties files are in road 
map too.

The document data model below:

- **name**: The name of the document. 
- **len**: The number of properties in the document.
- **properties**: List of properties of the document

### Properties
A document is composed by properties. Properties are defined in documents.

The property data model below:
 
- **key**: The key of the property
- **path**: The path of the property
- **value**: The value of the property