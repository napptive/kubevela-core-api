/*
Copyright 2021 The KubeVela Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/napptive/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/napptive/kubevela-core-api/pkg/generated/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CoreV1beta1Interface interface {
	RESTClient() rest.Interface
	ApplicationsGetter
	ApplicationRevisionsGetter
	ComponentDefinitionsGetter
	DefinitionRevisionsGetter
	PolicyDefinitionsGetter
	ResourceTrackersGetter
	ScopeDefinitionsGetter
	TraitDefinitionsGetter
	WorkflowStepDefinitionsGetter
	WorkloadDefinitionsGetter
}

// CoreV1beta1Client is used to interact with features provided by the core.oam.dev group.
type CoreV1beta1Client struct {
	restClient rest.Interface
}

func (c *CoreV1beta1Client) Applications(namespace string) ApplicationInterface {
	return newApplications(c, namespace)
}

func (c *CoreV1beta1Client) ApplicationRevisions(namespace string) ApplicationRevisionInterface {
	return newApplicationRevisions(c, namespace)
}

func (c *CoreV1beta1Client) ComponentDefinitions(namespace string) ComponentDefinitionInterface {
	return newComponentDefinitions(c, namespace)
}

func (c *CoreV1beta1Client) DefinitionRevisions(namespace string) DefinitionRevisionInterface {
	return newDefinitionRevisions(c, namespace)
}

func (c *CoreV1beta1Client) PolicyDefinitions(namespace string) PolicyDefinitionInterface {
	return newPolicyDefinitions(c, namespace)
}

func (c *CoreV1beta1Client) ResourceTrackers(namespace string) ResourceTrackerInterface {
	return newResourceTrackers(c, namespace)
}

func (c *CoreV1beta1Client) ScopeDefinitions(namespace string) ScopeDefinitionInterface {
	return newScopeDefinitions(c, namespace)
}

func (c *CoreV1beta1Client) TraitDefinitions(namespace string) TraitDefinitionInterface {
	return newTraitDefinitions(c, namespace)
}

func (c *CoreV1beta1Client) WorkflowStepDefinitions(namespace string) WorkflowStepDefinitionInterface {
	return newWorkflowStepDefinitions(c, namespace)
}

func (c *CoreV1beta1Client) WorkloadDefinitions(namespace string) WorkloadDefinitionInterface {
	return newWorkloadDefinitions(c, namespace)
}

// NewForConfig creates a new CoreV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*CoreV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CoreV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new CoreV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CoreV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CoreV1beta1Client for the given RESTClient.
func New(c rest.Interface) *CoreV1beta1Client {
	return &CoreV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CoreV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
