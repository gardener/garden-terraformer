// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package terraformer_test

import (
	"context"
	"testing"

	"github.com/gardener/gardener/test/framework"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	runtimelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func TestTerraformer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Terraformer Suite")
}

var (
	ctx        context.Context
	testEnv    *envtest.Environment
	restConfig *rest.Config
	testClient client.Client
)

var _ = BeforeSuite(func() {
	ctx = context.Background()
	runtimelog.SetLogger(zap.New(zap.UseDevMode(true), zap.WriteTo(GinkgoWriter)))

	By("starting test environment")
	testEnv = &envtest.Environment{}

	var err error
	restConfig, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(restConfig).ToNot(BeNil())

	testClient, err = client.New(restConfig, client.Options{})
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	By("running cleanup actions")
	framework.RunCleanupActions()
	gexec.CleanupBuildArtifacts()

	By("stopping test environment")
	Expect(testEnv.Stop()).To(Succeed())
})
