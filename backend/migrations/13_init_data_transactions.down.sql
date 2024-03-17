DO $$ BEGIN TRUNCATE TABLE refunds;
TRUNCATE TABLE transactions;
COMMIT;
EXCEPTION
WHEN OTHERS THEN ROLLBACK;
RAISE NOTICE 'Error: %',
SQLERRM;
END $$