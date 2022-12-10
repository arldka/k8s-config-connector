// Copyright 2022 Google LLC
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

package mockserviceusage

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb_http "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/api/serviceusage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "google.golang.org/genproto/googleapis/api/serviceusage/v1"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ExpectedHost = "serviceusage.googleapis.com"

// MockService represents a mocked cloud billing service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects   *projects.ProjectStore
	operations *operations.Operations

	serviceusagev1 *ServiceUsageV1
}

// New creates a MockResourceManager
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       env.GetKubeClient(),
		storage:    storage,
		projects:   env.GetProjects(),
		operations: operations.NewOperationsService(storage),
	}
	s.serviceusagev1 = &ServiceUsageV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "serviceusage.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterServiceUsageServer(grpcServer, s.serviceusagev1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()

	if err := pb_http.RegisterServiceUsageHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}