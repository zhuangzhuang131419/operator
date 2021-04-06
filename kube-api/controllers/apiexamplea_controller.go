/*
Copyright 2021 zhuang.

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

package controllers

import (
	"context"
	"log"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	groupav1beta1 "github.com/operator-demo/kube-api/api/v1beta1"
)

// ApiExampleAReconciler reconciles a ApiExampleA object
type ApiExampleAReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=groupa.k8s.zhuang.com,resources=apiexampleas,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=groupa.k8s.zhuang.com,resources=apiexampleas/status,verbs=get;update;patch

func (r *ApiExampleAReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("apiexamplea", req.NamespacedName)

	obja := &groupav1beta1.ApiExampleA{}
	if err := r.Get(ctx, req.NamespacedName, obja); err != nil {
		log.Println(err, "unable to fetch New Object")
	} else {
		log.Println("fetch New Object:", obja.Spec.FirstName, obja.Spec.SecondName)
	}

	return ctrl.Result{}, nil
}

func (r *ApiExampleAReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&groupav1beta1.ApiExampleA{}).
		Complete(r)
}
