SELECT 1;

INSERT INTO master_account_types (id, "name") VALUES  
    (1, 'Assets'),
    (2, 'Liabilities'),
    (3, 'Equity'),
    (4, 'Revenue'),
    (5, 'Expenses');

INSERT INTO master_account_categories (id, account_type_id, "name", "description") VALUES
    (1, 1, 'Cash', 'Kas'),
    (2, 1, 'Accounts Receivable', 'Piutang Usaha'),
    (3, 1, 'Inventory', 'Persediaan'),
    (4, 1, 'Equipment', 'Peralatan'),
    (5, 2, 'Accounts Payable', 'Hutang Usaha'),
    (6, 2, 'Notes Payable', 'Hutang Bank'),
    (7, 2, 'Wages Payable', 'Hutang Gaji'),
    (8, 2, 'Taxes Payable', 'Pajak Yang Masih Harus Dibayar'),
    (9, 3, 'Capital', 'Modal'),
    (10, 3, 'Retained Earnings', 'Laba Ditahan'),
    (11, 4, 'Sales', 'Penjualan'),
    (12, 4, 'Other Revenues', 'Pendapatan Lainnya'),
    (13, 5, 'Electricity Expense', 'Biaya Listrik'),
    (14, 5, 'Wages Expense', 'Biaya Gaji'),
    (15, 5, 'Rent Expense', 'Biaya Sewa'),
    (16, 5, 'Taxes Expense', 'Biaya Pajak'),
    (17, 5, 'Raw Material Expense', 'Biaya Bahan Baku'),
    (18, 5, 'Cost of Good Sold', 'Biaya Pokok Penjualan'),
    (19, 5, 'Operating Expenses', 'Biaya Operasional');

INSERT INTO master_accounts (id, "name", category_id) VALUES
    (1, 'Petty Cash', 1),
    (2, 'Cash in Bank', 1),
    (3, 'Raw Materials', 3),
    (4, 'Work in Progress', 3),
    (5, 'Finished Goods', 3),
    (6, 'Production Machines', 4),
    (7, 'Vehicles', 4),
    (8, 'Computer and Software', 4),
    (9, 'Owner''s Equity', 9),
    (10, 'Stockholder''s Equity', 9),
    (11, 'Salary Expense', 14),
    (12, 'Wage Expense', 14),
    (13, 'Overtime Expense', 14),
    (14, 'Payroll Tax Expense', 14),
    (15, 'Men''s Shirt', 3),
    (16, 'Accounts Payable to Shirt Supplier', 5),
    (17, 'Sales - Retail', 11),
    (18, 'Direct Materials Purchases', 18),
    (19, 'Accounts Receivable - Customers', 2),
    (20, 'Admin Fee Revenue', 12),
    (21, 'Admin Fee - Marketplace T', 19);

INSERT INTO master_event_types (id, "name", "description", "rules") VALUES
    (1, 'Purchase of merchandise (Shirt)', 'Pembelian barang dagangan dari pemasok atau produsen kemeja', '[{"account_id":15,"transaction_type":"DEBIT","rule_type":"PERCENTAGE","amount":100},{"account_id":16,"transaction_type":"CREDIT","rule_type":"PERCENTAGE","amount":100}]'),
    (2, 'Sales of shirt to customers (COGS 60%)', 'Penjualan barang dagangan (kemeja) kepada pelanggan dengan COGS sebesar 60%', '[{"account_id":15,"transaction_type":"DEBIT","rule_type":"PERCENTAGE","amount":60},{"account_id":17,"transaction_type":"CREDIT","rule_type":"PERCENTAGE","amount":100},{"account_id":18,"transaction_type":"CREDIT","rule_type":"PERCENTAGE","amount":60},{"account_id":2,"transaction_type":"DEBIT","rule_type":"PERCENTAGE","amount":100}]'),
    (3, 'Online Sales of shirt With Admin Fee 150000', 'Penjualan online barang dagangan (kemeja) dengan admin fee sebesar Rp. 150,000,00', '[{"account_id":17,"transaction_type":"CREDIT","rule_type":"PERCENTAGE","amount":100},{"account_id":20,"transaction_type":"CREDIT","rule_type":"FLAT","amount":150000},{"account_id":19,"transaction_type":"DEBIT","rule_type":"PERCENTAGE","amount":100}]'),
    (4, 'Test Case 4', 'Test case percentage and flat with more than one account', '[{"account_id":15,"transaction_type":"DEBIT","rule_type":"PERCENTAGE","amount":60},{"account_id":3,"transaction_type":"DEBIT","rule_type":"PERCENTAGE","amount":40},{"account_id":21,"transaction_type":"DEBIT","rule_type":"FLAT","amount":200000},{"account_id":16,"transaction_type":"CREDIT","rule_type":"PERCENTAGE","amount":100}]');