-- Modern ELT approach: We load raw data (Staging), then transform in-place.
-- This ensures idempotency (No duplicates).

MERGE INTO `analytics.fact_orders` AS target
USING `staging.orders_raw` AS source
ON target.order_id = source.order_id
WHEN MATCHED THEN
  UPDATE SET 
    target.status = source.status,
    target.updated_at = source.updated_at
WHEN NOT MATCHED THEN
  INSERT (order_id, customer_id, amount, status)
  VALUES (source.order_id, source.customer_id, source.amount, source.status);
