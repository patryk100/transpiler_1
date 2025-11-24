package main

import (
	"fmt"
	"reflect"
)

type BaseResource struct {
	Name        string `tf:"resource_name"`
	Description string `tf:"resource_description"`
	Region      string `tf:"default_region"`
}
type S3Bucket struct {
	BaseResource
	BucketName string `tf:"bucket,required"`
	ACL        string `tf:"acl"`
	Versioning bool   `tf:"versioning_enabled"`
}

type EC2Instance struct {
	BaseResource
	AMI          string `tf:"ami,required"`
	InstanceType string `tf:"instance_type,required"`
	Versioning   bool   `tf:"versioning_enabled"`
}

func NewBaseResource(name string, desc string, region string) *BaseResource {
	return &BaseResource{Name: name, Description: desc, Region: region}
}

func main() {
	base := NewBaseResource("name", "desc", "us-east-1")
	v := reflect.ValueOf(base)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("Not a struct")
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(t.Field(i).Name,
			t.Field(i).Tag.Get("tf"),
			v.Field(i).Interface())
	}

}
