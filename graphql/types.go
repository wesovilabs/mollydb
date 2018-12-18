package graphql

import (
	"github.com/graphql-go/graphql"
	"mollydb/model"
)

const (
	storage  = "StorageList"
	document = "Document"
	property = "Property"
)

var propertyType, documentType, storageType *graphql.Object

func defineTypes() {
	defineDataType()
	defineDocumentType()
	defineStorageType()
}

func defineDataType() {
	fieldID := &graphql.Field{
		Name:        "key",
		Description: "The key of the property",
		Type:        graphql.NewNonNull(graphql.ID),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			field := p.Source.(*model.Property)
			return field.Key, nil
		},
	}

	fieldPath := &graphql.Field{
		Name:        "path",
		Description: "The path of the property",
		Type:        graphql.NewNonNull(graphql.String),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			field := p.Source.(*model.Property)
			return field.Path, nil
		},
	}

	fieldValue := &graphql.Field{
		Name:        "value",
		Type:        graphql.NewNonNull(graphql.String),
		Description: "The value of the property",
		Resolve: func(p graphql.ResolveParams) (
			interface{}, error) {
			field := p.Source.(*model.Property)
			return field.Value, nil
		},
	}

	propertyType = graphql.NewObject(graphql.ObjectConfig{
		Name:        property,
		Description: "Document content",
		Fields: graphql.Fields{
			"path":  fieldPath,
			"key":   fieldID,
			"value": fieldValue,
		},
	})
}

func defineDocumentType() {

	fieldID := &graphql.Field{
		Name:        "name",
		Description: "The name of the document",
		Type:        graphql.NewNonNull(graphql.ID),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			document := p.Source.(*model.Document)
			return document.Name, nil
		},
	}

	fieldLen := &graphql.Field{
		Type:        graphql.Int,
		Description: "The number of properties in the document",
		Resolve: func(p graphql.ResolveParams) (
			interface{}, error) {
			document := p.Source.(*model.Document)
			return len(document.Properties), nil
		},
	}

	fieldProperties := &graphql.Field{
		Type:        &graphql.List{OfType: propertyType},
		Description: "List of properties of the document",
		Args: map[string]*graphql.ArgumentConfig{
			"key": {
				Type:        graphql.String,
				Description: "The key of the property",
			},
		},
		Resolve: func(p graphql.ResolveParams) (
			interface{}, error) {
			filterKey := p.Args["key"]
			document := p.Source.(*model.Document)
			var result []*model.Property
			for _, d := range document.Properties {
				if filterKey == nil || d.Key == filterKey {
					result = append(result, d)
				}
			}
			return result, nil
		},
	}

	documentType = graphql.NewObject(graphql.ObjectConfig{
		Name:        document,
		Description: "A mollyDB document",
		Fields: graphql.Fields{
			"name":       fieldID,
			"len":        fieldLen,
			"properties": fieldProperties,
		},
	})
}

func defineStorageType() {
	fieldName := &graphql.Field{
		Name:        "name",
		Description: "The name of a storage",
		Type:        graphql.NewNonNull(graphql.ID),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			storage := p.Source.(*model.Storage)
			return storage.Name, nil
		},
	}

	fieldLen := &graphql.Field{
		Type:        graphql.Int,
		Description: "The number of documents in the storage",
		Resolve: func(p graphql.ResolveParams) (
			interface{}, error) {
			parent := p.Source.(*model.Storage)
			return len(parent.Documents), nil
		},
	}

	fieldDocuments := &graphql.Field{
		Type: &graphql.List{OfType: documentType},
		Args: map[string]*graphql.ArgumentConfig{
			"name": {
				Type:        graphql.String,
				Description: "The name of the document",
			},
		},
		Description: "The list of documents that belong to this storage",
		Resolve: func(p graphql.ResolveParams) (
			interface{}, error) {
			filterID := p.Args["name"]
			parent := p.Source.(*model.Storage)
			var result []*model.Document
			for _, r := range parent.Documents {
				if filterID == nil || r.Name == filterID {
					result = append(result, r)
				}
			}
			return result, nil
		},
	}

	storageType = graphql.NewObject(graphql.ObjectConfig{
		Name:        storage,
		Description: "A mollyDB storage",
		Fields: graphql.Fields{
			"name":      fieldName,
			"len":       fieldLen,
			"documents": fieldDocuments,
		},
	})
}
