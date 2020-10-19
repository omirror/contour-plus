package controllers

import (
	projectcontourv1 "github.com/projectcontour/contour/apis/projectcontour/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	// +kubebuilder:scaffold:imports
)

// ReconcilerOptions is a set of options for reconcilers
type ReconcilerOptions struct {
	ServiceKey        client.ObjectKey
	Prefix            string
	DefaultIssuerName string
	DefaultIssuerKind string
	CreateDNSEndpoint bool
	CreateCertificate bool
	IngressClassName  string
}

// SetupScheme initializes a schema
func SetupScheme(scm *runtime.Scheme) error {
	projectcontourv1.AddKnownTypes(scm)

	// for corev1.Service
	if err := corev1.AddToScheme(scm); err != nil {
		return err
	}

	// +kubebuilder:scaffold:scheme
	return nil
}

// SetupReconciler initializes reconcilers
func SetupReconciler(mgr manager.Manager, scheme *runtime.Scheme, opts ReconcilerOptions) error {
	httpProxyReconciler := &HTTPProxyReconciler{
		Client:            mgr.GetClient(),
		Log:               ctrl.Log.WithName("controllers").WithName("HTTPProxy"),
		Scheme:            scheme,
		ServiceKey:        opts.ServiceKey,
		Prefix:            opts.Prefix,
		DefaultIssuerName: opts.DefaultIssuerName,
		DefaultIssuerKind: opts.DefaultIssuerKind,
		CreateDNSEndpoint: opts.CreateDNSEndpoint,
		CreateCertificate: opts.CreateCertificate,
		IngressClassName:  opts.IngressClassName,
	}
	err := httpProxyReconciler.SetupWithManager(mgr)
	if err != nil {
		return err
	}

	// +kubebuilder:scaffold:builder
	return nil
}
