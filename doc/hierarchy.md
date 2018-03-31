## Storage hierarchy

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