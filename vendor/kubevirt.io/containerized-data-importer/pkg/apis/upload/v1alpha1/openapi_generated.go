// +build !ignore_autogenerated

/*
Copyright 2018 The CDI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequest":       schema_pkg_apis_upload_v1alpha1_UploadTokenRequest(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequestList":   schema_pkg_apis_upload_v1alpha1_UploadTokenRequestList(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequestSpec":   schema_pkg_apis_upload_v1alpha1_UploadTokenRequestSpec(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequestStatus": schema_pkg_apis_upload_v1alpha1_UploadTokenRequestStatus(ref),
	}
}

func schema_pkg_apis_upload_v1alpha1_UploadTokenRequest(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UploadTokenRequest is the CR used to initiate a CDI upload",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Spec contains the parameters of the request",
							Ref:         ref("kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequestSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status contains the status of the request",
							Ref:         ref("kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequestStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequestSpec", "kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequestStatus"},
	}
}

func schema_pkg_apis_upload_v1alpha1_UploadTokenRequestList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UploadTokenRequestList contains a list of UploadTokenRequests",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "Items contains a list of UploadTokenRequests",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequest"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta", "kubevirt.io/containerized-data-importer/pkg/apis/upload/v1alpha1.UploadTokenRequest"},
	}
}

func schema_pkg_apis_upload_v1alpha1_UploadTokenRequestSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UploadTokenRequestSpec defines the parameters of the token request",
				Properties: map[string]spec.Schema{
					"pvcName": {
						SchemaProps: spec.SchemaProps{
							Description: "PvcName is the name of the PVC to upload to",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"pvcName"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_upload_v1alpha1_UploadTokenRequestStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UploadTokenRequestStatus stores the status of a token request",
				Properties: map[string]spec.Schema{
					"token": {
						SchemaProps: spec.SchemaProps{
							Description: "Token is a JWT token to be inserted in \"Authentication Bearer header\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
