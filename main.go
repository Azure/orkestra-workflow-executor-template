package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/Azure/helmrelease-workflow-executor/pkg/actions"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// The set of executor actions which can be performed on a helmrelease object
	Install ExecutorAction = "install"
	Delete  ExecutorAction = "delete"
)

// ExecutorAction defines the set of executor actions which can be performed on a helmrelease object
type ExecutorAction string

func ParseExecutorAction(s string) (ExecutorAction, error) {
	a := ExecutorAction(s)
	switch a {
	case Install, Delete:
		return a, nil
	}
	return "", fmt.Errorf("invalid executor action: %v", s)
}

func main() {
	var actionStr string
	var timeoutStr string
	var intervalStr string

	flag.StringVar(&actionStr, "action", "", "Action to perform on the helmrelease object. Must be either install or delete")
	flag.StringVar(&timeoutStr, "timeout", "5m", "Timeout for the execution of the argo workflow task")
	flag.StringVar(&intervalStr, "interval", "10s", "Retry interval for the all actions by the executor")

	// Add your logic here
	// Provide input params to the executor as flags

	flag.Parse()

	action, err := ParseExecutorAction(actionStr)
	if err != nil {
		log.Fatalf("Failed to parse action as an executor action with %v", err)
	}
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Fatalf("Failed to parse timeout as a duration with %v", err)
	}
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		log.Fatalf("Failed to parse interval as a duration with %v", err)
	}
	log.Infof("Parsed the action: %v, the timeout: %v and the interval: %v", string(action), timeout.String(), interval.String())

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// (optional) Use kubeconfig to create a client.
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(), &clientcmd.ConfigOverrides{})
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		log.Fatalf("Failed to initialize the client config with %v", err)
	}

	// Add your logic here
	// Register the types with the scheme
	k8sScheme := scheme.Scheme

	// Example:
	// if err := fluxhelmv2beta1.AddToScheme(k8sScheme); err != nil {
	// 	log.Fatalf("Failed to add the flux helm scheme to the configuration scheme with %v", err)
	// }

	clientSet, err := client.New(config, client.Options{Scheme: k8sScheme})
	if err != nil {
		log.Fatalf("Failed to create the clientset with the given config with %v", err)
	}

	if action == Install {
		if err := actions.Install(ctx, cancel, clientSet, interval); err != nil {
			log.Fatalf("failed to install the helm release: %v", err)
		}
	} else if action == Delete {
		if err := actions.Delete(ctx, cancel, clientSet, interval); err != nil {
			log.Fatalf("failed to delete the helm release: %v", err)
		}
	}
}
