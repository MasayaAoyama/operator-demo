/*
Copyright 2020 amsy810.

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

package v1beta1

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// log is for logging in this package.
var webserverlog = logf.Log.WithName("webserver-resource")

func (r *WebServer) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-servers-amsy810-dev-v1beta1-webserver,mutating=true,failurePolicy=fail,groups=servers.amsy810.dev,resources=webservers,verbs=create;update,versions=v1beta1,name=mwebserver.kb.io

var _ webhook.Defaulter = &WebServer{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *WebServer) Default() {
	webserverlog.Info("default", "name", r.Name)

	if r.Spec.Content == "" {
		r.Spec.Content = fmt.Sprintf("This CR was defaulting at %s", time.Now())
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-servers-amsy810-dev-v1beta1-webserver,mutating=false,failurePolicy=fail,groups=servers.amsy810.dev,resources=webservers,versions=v1beta1,name=vwebserver.kb.io

var _ webhook.Validator = &WebServer{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *WebServer) ValidateCreate() error {
	webserverlog.Info("validate create", "name", r.Name)

	var allErrs field.ErrorList

	if r.Spec.Replicas == 0 {
		err := field.Invalid(
			field.NewPath("spec").Child("schedule"),
			r.Spec.Replicas,
			"replicas must be >0")
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return errors.NewInvalid(
		schema.GroupKind{Group: "servers.amsy.dev", Kind: "WebServer"},
		r.Name, allErrs)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *WebServer) ValidateUpdate(old runtime.Object) error {
	webserverlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *WebServer) ValidateDelete() error {
	webserverlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
