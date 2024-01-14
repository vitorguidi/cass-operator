/*
Copyright 2023.

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

package autoscaling

import (
	"context"
	"fmt"
	cassandradatastaxcomv1beta1 "github.com/k8ssandra/cass-operator/apis/autoscaling/v1beta1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// CassandraDataCenterAutoscalerReconciler reconciles a CassandraDataCenterAutoscaler object
type CassandraDataCenterAutoscalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cassandra.datastax.com.cassandra.datastax.com,resources=cassandradatacenterautoscalers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cassandra.datastax.com.cassandra.datastax.com,resources=cassandradatacenterautoscalers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cassandra.datastax.com.cassandra.datastax.com,resources=cassandradatacenterautoscalers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CassandraDataCenterAutoscaler object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *CassandraDataCenterAutoscalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Bem vindo ao maravilhoso mundo de autoscaling. Recebi o evento do recurso %s", "resource name", req.NamespacedName)
	autoscaler := &cassandradatastaxcomv1beta1.CassandraDataCenterAutoscaler{}
	err := r.Get(ctx, req.NamespacedName, autoscaler)
	if err != nil {
		logger.Error(err, "Could not find autoscaler crd in question", "namespaced name", req.NamespacedName)
		return ctrl.Result{}, err
	}
	// TODO(user): your logic here
	logger.Info(fmt.Sprintf("Datacenter Ref = %s", autoscaler.Spec.CassandraDatacenterRef))
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CassandraDataCenterAutoscalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cassandradatastaxcomv1beta1.CassandraDataCenterAutoscaler{}).
		Complete(r)
}
