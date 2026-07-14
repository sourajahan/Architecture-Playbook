-- Pattern: Maintaining History in DWH
-- We track changes to customer data with valid_from/valid_to dates.

CREATE TABLE dim_customer (
    customer_key SERIAL PRIMARY KEY,
    customer_id INT,
    address TEXT,
    valid_from TIMESTAMP,
    valid_to TIMESTAMP, -- NULL or '9999-12-31' means current version
    is_current BOOLEAN
);

-- When an address changes:
-- 1. Expire the old record
UPDATE dim_customer 
SET valid_to = CURRENT_TIMESTAMP, is_current = FALSE 
WHERE customer_id = 123 AND is_current = TRUE;

-- 2. Insert the new version
INSERT INTO dim_customer (customer_id, address, valid_from, valid_to, is_current)
VALUES (123, 'New Address', CURRENT_TIMESTAMP, NULL, TRUE);
