package actions

import (
	"context"
	"fmt"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Install handles the installation or setup/init action
// This could mean creating a kubernetes resource, or making a call to an external API, etc.

// Add your logic here
// Extend the function parameters to add more options and inputs for your action.
func Install(ctx context.Context, cancel context.CancelFunc, clientSet client.Client, interval time.Duration) error {
	defer cancel()

	// Add your logic here
	// Step 1: Implement the install logic

	// Add your logic here
	// Step 2: (optional) Poll for success condition
	if err := pollStatus(ctx, clientSet, interval, 5); err != nil {
		return fmt.Errorf("failed to poll status with: %w", err)
	}

	return nil
}

func pollStatus(ctx context.Context, clientSet client.Client, interval time.Duration, retrySeconds int) error {
	// Add your logic here
	// (optional) Implement a polling loop to check for success condition using the "status" package

	// Example:
	// if err := status.Poll(ctx, _ , interval); err != nil {
	// 	return fmt.Errorf("timed out waiting for condition")
	// }
	return nil
}
