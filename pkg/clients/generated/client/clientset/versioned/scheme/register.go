// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package scheme

import (
	accesscontextmanagerv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/accesscontextmanager/v1alpha1"
	accesscontextmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/accesscontextmanager/v1beta1"
	alloydbv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/alloydb/v1beta1"
	apigatewayv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigateway/v1alpha1"
	apigeev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigee/v1alpha1"
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigee/v1beta1"
	appenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/appengine/v1alpha1"
	artifactregistryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/artifactregistry/v1beta1"
	beyondcorpv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/beyondcorp/v1alpha1"
	bigqueryv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigquery/v1alpha1"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigquery/v1beta1"
	bigqueryanalyticshubv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigqueryanalyticshub/v1alpha1"
	bigqueryconnectionv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigqueryconnection/v1alpha1"
	bigquerydatapolicyv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigquerydatapolicy/v1alpha1"
	bigquerydatatransferv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigquerydatatransfer/v1alpha1"
	bigqueryreservationv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigqueryreservation/v1alpha1"
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigtable/v1beta1"
	billingbudgetsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/billingbudgets/v1beta1"
	binaryauthorizationv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/binaryauthorization/v1beta1"
	certificatemanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/certificatemanager/v1beta1"
	cloudassetv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudasset/v1alpha1"
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudbuild/v1beta1"
	cloudfunctionsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudfunctions/v1beta1"
	cloudfunctions2v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudfunctions2/v1alpha1"
	cloudidentityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudidentity/v1beta1"
	cloudidsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudids/v1alpha1"
	cloudiotv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudiot/v1alpha1"
	cloudschedulerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudscheduler/v1beta1"
	cloudtasksv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudtasks/v1alpha1"
	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	configcontrollerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/configcontroller/v1beta1"
	containerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/container/v1beta1"
	containeranalysisv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/containeranalysis/v1alpha1"
	containeranalysisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/containeranalysis/v1beta1"
	containerattachedv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/containerattached/v1beta1"
	datacatalogv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datacatalog/v1alpha1"
	datacatalogv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datacatalog/v1beta1"
	dataflowv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dataflow/v1beta1"
	dataformv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dataform/v1alpha1"
	datafusionv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datafusion/v1beta1"
	dataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dataproc/v1beta1"
	datastorev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datastore/v1alpha1"
	datastreamv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datastream/v1alpha1"
	deploymentmanagerv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/deploymentmanager/v1alpha1"
	dialogflowv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dialogflow/v1alpha1"
	dialogflowcxv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dialogflowcx/v1alpha1"
	dlpv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dlp/v1beta1"
	dnsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dns/v1alpha1"
	dnsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dns/v1beta1"
	documentaiv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/documentai/v1alpha1"
	edgenetworkv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/edgenetwork/v1beta1"
	essentialcontactsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/essentialcontacts/v1alpha1"
	eventarcv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/eventarc/v1beta1"
	filestorev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/filestore/v1alpha1"
	filestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/filestore/v1beta1"
	firebasev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firebase/v1alpha1"
	firebasedatabasev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firebasedatabase/v1alpha1"
	firebasehostingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firebasehosting/v1alpha1"
	firebasestoragev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firebasestorage/v1alpha1"
	firestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firestore/v1beta1"
	gkebackupv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/gkebackup/v1alpha1"
	gkehubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/gkehub/v1beta1"
	healthcarev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/healthcare/v1alpha1"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	iapv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iap/v1beta1"
	identityplatformv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1alpha1"
	identityplatformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	kmsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/kms/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/kms/v1beta1"
	loggingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/logging/v1beta1"
	memcachev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/memcache/v1beta1"
	mlenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/mlengine/v1alpha1"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/monitoring/v1beta1"
	networkconnectivityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkconnectivity/v1beta1"
	networkmanagementv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkmanagement/v1alpha1"
	networksecurityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networksecurity/v1beta1"
	networkservicesv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1alpha1"
	networkservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1beta1"
	notebooksv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/notebooks/v1alpha1"
	orgpolicyv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/orgpolicy/v1alpha1"
	osconfigv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/osconfig/v1alpha1"
	osconfigv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/osconfig/v1beta1"
	osloginv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/oslogin/v1alpha1"
	privatecav1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/privateca/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsub/v1beta1"
	pubsublitev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsublite/v1alpha1"
	pubsublitev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsublite/v1beta1"
	recaptchaenterprisev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/recaptchaenterprise/v1beta1"
	redisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/redis/v1beta1"
	resourcemanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/resourcemanager/v1beta1"
	runv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/run/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/secretmanager/v1beta1"
	securitycenterv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/securitycenter/v1alpha1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/servicedirectory/v1beta1"
	servicenetworkingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/servicenetworking/v1beta1"
	serviceusagev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/serviceusage/v1alpha1"
	serviceusagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/serviceusage/v1beta1"
	sourcerepov1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/sourcerepo/v1beta1"
	spannerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/spanner/v1beta1"
	sqlv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/sql/v1beta1"
	storagev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storage/v1alpha1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storage/v1beta1"
	storagetransferv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storagetransfer/v1alpha1"
	storagetransferv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storagetransfer/v1beta1"
	tagsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/tags/v1beta1"
	tpuv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/tpu/v1alpha1"
	vertexaiv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/vertexai/v1alpha1"
	vpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/vpcaccess/v1beta1"
	workflowsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/workflows/v1alpha1"
	workstationsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/workstations/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	accesscontextmanagerv1beta1.AddToScheme,
	accesscontextmanagerv1alpha1.AddToScheme,
	alloydbv1beta1.AddToScheme,
	apigatewayv1alpha1.AddToScheme,
	apigeev1alpha1.AddToScheme,
	apigeev1beta1.AddToScheme,
	appenginev1alpha1.AddToScheme,
	artifactregistryv1beta1.AddToScheme,
	beyondcorpv1alpha1.AddToScheme,
	bigqueryv1alpha1.AddToScheme,
	bigqueryv1beta1.AddToScheme,
	bigqueryanalyticshubv1alpha1.AddToScheme,
	bigqueryconnectionv1alpha1.AddToScheme,
	bigquerydatapolicyv1alpha1.AddToScheme,
	bigquerydatatransferv1alpha1.AddToScheme,
	bigqueryreservationv1alpha1.AddToScheme,
	bigtablev1beta1.AddToScheme,
	billingbudgetsv1beta1.AddToScheme,
	binaryauthorizationv1beta1.AddToScheme,
	certificatemanagerv1beta1.AddToScheme,
	cloudassetv1alpha1.AddToScheme,
	cloudbuildv1beta1.AddToScheme,
	cloudfunctionsv1beta1.AddToScheme,
	cloudfunctions2v1alpha1.AddToScheme,
	cloudidentityv1beta1.AddToScheme,
	cloudidsv1alpha1.AddToScheme,
	cloudiotv1alpha1.AddToScheme,
	cloudschedulerv1beta1.AddToScheme,
	cloudtasksv1alpha1.AddToScheme,
	computev1alpha1.AddToScheme,
	computev1beta1.AddToScheme,
	configcontrollerv1beta1.AddToScheme,
	containerv1beta1.AddToScheme,
	containeranalysisv1alpha1.AddToScheme,
	containeranalysisv1beta1.AddToScheme,
	containerattachedv1beta1.AddToScheme,
	datacatalogv1alpha1.AddToScheme,
	datacatalogv1beta1.AddToScheme,
	dataflowv1beta1.AddToScheme,
	dataformv1alpha1.AddToScheme,
	datafusionv1beta1.AddToScheme,
	dataprocv1beta1.AddToScheme,
	datastorev1alpha1.AddToScheme,
	datastreamv1alpha1.AddToScheme,
	deploymentmanagerv1alpha1.AddToScheme,
	dialogflowv1alpha1.AddToScheme,
	dialogflowcxv1alpha1.AddToScheme,
	dlpv1beta1.AddToScheme,
	dnsv1alpha1.AddToScheme,
	dnsv1beta1.AddToScheme,
	documentaiv1alpha1.AddToScheme,
	edgenetworkv1beta1.AddToScheme,
	essentialcontactsv1alpha1.AddToScheme,
	eventarcv1beta1.AddToScheme,
	filestorev1alpha1.AddToScheme,
	filestorev1beta1.AddToScheme,
	firebasev1alpha1.AddToScheme,
	firebasedatabasev1alpha1.AddToScheme,
	firebasehostingv1alpha1.AddToScheme,
	firebasestoragev1alpha1.AddToScheme,
	firestorev1beta1.AddToScheme,
	gkebackupv1alpha1.AddToScheme,
	gkehubv1beta1.AddToScheme,
	healthcarev1alpha1.AddToScheme,
	iamv1beta1.AddToScheme,
	iapv1beta1.AddToScheme,
	identityplatformv1alpha1.AddToScheme,
	identityplatformv1beta1.AddToScheme,
	k8sv1alpha1.AddToScheme,
	kmsv1alpha1.AddToScheme,
	kmsv1beta1.AddToScheme,
	loggingv1beta1.AddToScheme,
	memcachev1beta1.AddToScheme,
	mlenginev1alpha1.AddToScheme,
	monitoringv1beta1.AddToScheme,
	networkconnectivityv1beta1.AddToScheme,
	networkmanagementv1alpha1.AddToScheme,
	networksecurityv1beta1.AddToScheme,
	networkservicesv1alpha1.AddToScheme,
	networkservicesv1beta1.AddToScheme,
	notebooksv1alpha1.AddToScheme,
	orgpolicyv1alpha1.AddToScheme,
	osconfigv1alpha1.AddToScheme,
	osconfigv1beta1.AddToScheme,
	osloginv1alpha1.AddToScheme,
	privatecav1beta1.AddToScheme,
	pubsubv1beta1.AddToScheme,
	pubsublitev1alpha1.AddToScheme,
	pubsublitev1beta1.AddToScheme,
	recaptchaenterprisev1beta1.AddToScheme,
	redisv1beta1.AddToScheme,
	resourcemanagerv1beta1.AddToScheme,
	runv1beta1.AddToScheme,
	secretmanagerv1beta1.AddToScheme,
	securitycenterv1alpha1.AddToScheme,
	servicedirectoryv1beta1.AddToScheme,
	servicenetworkingv1beta1.AddToScheme,
	serviceusagev1alpha1.AddToScheme,
	serviceusagev1beta1.AddToScheme,
	sourcerepov1beta1.AddToScheme,
	spannerv1beta1.AddToScheme,
	sqlv1beta1.AddToScheme,
	storagev1alpha1.AddToScheme,
	storagev1beta1.AddToScheme,
	storagetransferv1alpha1.AddToScheme,
	storagetransferv1beta1.AddToScheme,
	tagsv1beta1.AddToScheme,
	tpuv1alpha1.AddToScheme,
	vertexaiv1alpha1.AddToScheme,
	vpcaccessv1beta1.AddToScheme,
	workflowsv1alpha1.AddToScheme,
	workstationsv1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//	import (
//	  "k8s.io/client-go/kubernetes"
//	  clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//	  aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//	)
//
//	kclientset, _ := kubernetes.NewForConfig(c)
//	_ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
