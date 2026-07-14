package main

// OrderWorkflow coordinates services without a central "Bus" or XML.
// It manages the "Saga" pattern: If payment fails, it triggers the refund automatically.
func OrderWorkflow(ctx Context, orderDetails Order) error {
	// Step 1: Inventory
	if err := InventoryService.Reserve(ctx, orderDetails.Items); err != nil {
		return err // Workflow engine handles retry/state persistence
	}

	// Step 2: Payment (The stateful operation)
	if err := PaymentService.Charge(ctx, orderDetails.Amount); err != nil {
		// Compensation logic (Saga Pattern)
		InventoryService.Release(ctx, orderDetails.Items)
		return err
	}

	// Step 3: Shipping
	return ShippingService.Ship(ctx, orderDetails.Address)
}
