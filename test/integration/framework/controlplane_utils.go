/*
Copyright 2014 The Kubernetes Authors.

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

package framework

import (
	"path"

	"github.com/google/uuid"

	"k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	// "k8s.io/kubernetes/pkg/api/legacyscheme"
	// "k8s.io/kubernetes/pkg/generated/openapi"
)

const (
	UnprivilegedUserToken = "unprivileged-user"
)

// MinVerbosity determines the minimum klog verbosity when running tests that
// involve the apiserver.  This overrides the -v value from the command line,
// i.e. -v=0 has no effect when MinVerbosity is 4 (the default).  Tests can opt
// out of this by setting MinVerbosity to zero before starting the control
// plane or choose some different minimum verbosity.
var MinVerbosity = 4

// // DefaultOpenAPIConfig returns an openapicommon.Config initialized to default values.
// func DefaultOpenAPIConfig() *openapicommon.Config {
// 	openAPIConfig := genericapiserver.DefaultOpenAPIConfig(openapi.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(legacyscheme.Scheme))
// 	openAPIConfig.Info = &spec.Info{
// 		InfoProps: spec.InfoProps{
// 			Title:   "Kubernetes",
// 			Version: "unversioned",
// 		},
// 	}
// 	openAPIConfig.DefaultResponse = &spec.Response{
// 		ResponseProps: spec.ResponseProps{
// 			Description: "Default Response.",
// 		},
// 	}
// 	openAPIConfig.GetDefinitions = utilopenapi.GetOpenAPIDefinitionsWithoutDisabledFeatures(openapi.GetOpenAPIDefinitions)

// 	return openAPIConfig
// }

// // DefaultOpenAPIV3Config returns an openapicommon.Config initialized to default values.
// func DefaultOpenAPIV3Config() *openapicommon.Config {
// 	openAPIConfig := genericapiserver.DefaultOpenAPIV3Config(openapi.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(legacyscheme.Scheme))
// 	openAPIConfig.Info = &spec.Info{
// 		InfoProps: spec.InfoProps{
// 			Title:   "Kubernetes",
// 			Version: "unversioned",
// 		},
// 	}
// 	openAPIConfig.DefaultResponse = &spec.Response{
// 		ResponseProps: spec.ResponseProps{
// 			Description: "Default Response.",
// 		},
// 	}
// 	openAPIConfig.GetDefinitions = utilopenapi.GetOpenAPIDefinitionsWithoutDisabledFeatures(openapi.GetOpenAPIDefinitions)

// 	return openAPIConfig
// }

// DefaultEtcdOptions are the default EtcdOptions for use with integration tests.
func DefaultEtcdOptions() *options.EtcdOptions {
	// This causes the integration tests to exercise the etcd
	// prefix code, so please don't change without ensuring
	// sufficient coverage in other ways.
	etcdOptions := options.NewEtcdOptions(storagebackend.NewDefaultConfig(uuid.New().String(), nil))
	etcdOptions.StorageConfig.Transport.ServerList = []string{GetEtcdURL()}
	return etcdOptions
}

// SharedEtcd creates a storage config for a shared etcd instance, with a unique prefix.
func SharedEtcd() *storagebackend.Config {
	cfg := storagebackend.NewDefaultConfig(path.Join(uuid.New().String(), "registry"), nil)
	cfg.Transport.ServerList = []string{GetEtcdURL()}
	return cfg
}
