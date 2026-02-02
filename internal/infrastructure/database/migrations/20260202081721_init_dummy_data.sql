-- +goose Up
-- +goose StatementBegin

-- Insert dummy categories (root categories first)
INSERT INTO categories (id, name, description, parent_id, status, created_at, updated_at) VALUES
-- Root categories
('550e8400-e29b-41d4-a716-446655440001', 'Electronics', 'Electronic devices and gadgets', NULL, 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440002', 'Clothing', 'Apparel and fashion items', NULL, 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440003', 'Home & Garden', 'Home improvement and gardening supplies', NULL, 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440004', 'Sports & Outdoors', 'Sports equipment and outdoor activities', NULL, 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440005', 'Books & Media', 'Books, magazines, and media content', NULL, 'active', NOW(), NOW());

-- Sub categories
INSERT INTO categories (id, name, description, parent_id, status, created_at, updated_at) VALUES
-- Electronics subcategories
('550e8400-e29b-41d4-a716-446655440011', 'Smartphones', 'Mobile phones and accessories', '550e8400-e29b-41d4-a716-446655440001', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440012', 'Laptops', 'Portable computers and accessories', '550e8400-e29b-41d4-a716-446655440001', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440013', 'Audio', 'Headphones, speakers, and audio equipment', '550e8400-e29b-41d4-a716-446655440001', 'active', NOW(), NOW()),
-- Clothing subcategories
('550e8400-e29b-41d4-a716-446655440021', 'Men''s Clothing', 'Clothing for men', '550e8400-e29b-41d4-a716-446655440002', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440022', 'Women''s Clothing', 'Clothing for women', '550e8400-e29b-41d4-a716-446655440002', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440023', 'Shoes', 'Footwear for all ages', '550e8400-e29b-41d4-a716-446655440002', 'active', NOW(), NOW());

-- Insert dummy products
INSERT INTO products (id, sku, name, description, category_id, price, cost, stock, min_stock, max_stock, status, created_at, updated_at) VALUES
-- Electronics products
('660e8400-e29b-41d4-a716-446655440001', 'IPHONE-15-PRO', 'iPhone 15 Pro', 'Latest Apple iPhone with advanced camera system', '550e8400-e29b-41d4-a716-446655440011', 999.99, 750.00, 50, 10, 100, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440002', 'SAMSUNG-S24', 'Samsung Galaxy S24', 'Premium Android smartphone with AI features', '550e8400-e29b-41d4-a716-446655440011', 899.99, 650.00, 75, 15, 150, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440003', 'MACBOOK-PRO-16', 'MacBook Pro 16"', 'Professional laptop with M3 Pro chip', '550e8400-e29b-41d4-a716-446655440012', 2499.99, 2000.00, 25, 5, 50, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440004', 'DELL-XPS-13', 'Dell XPS 13', 'Ultrabook with premium build quality', '550e8400-e29b-41d4-a716-446655440012', 1299.99, 900.00, 30, 8, 60, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440005', 'AIRPODS-PRO', 'AirPods Pro 2nd Gen', 'Wireless earbuds with active noise cancellation', '550e8400-e29b-41d4-a716-446655440013', 249.99, 150.00, 100, 20, 200, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440006', 'SONY-WH1000XM5', 'Sony WH-1000XM5', 'Premium noise-canceling headphones', '550e8400-e29b-41d4-a716-446655440013', 399.99, 250.00, 60, 15, 120, 'active', NOW(), NOW()),

-- Clothing products
('660e8400-e29b-41d4-a716-446655440007', 'NIKE-TSHIRT-M', 'Nike Men''s T-Shirt', 'Comfortable cotton t-shirt for daily wear', '550e8400-e29b-41d4-a716-446655440021', 29.99, 15.00, 200, 50, 500, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440008', 'LEVI-JEANS-32', 'Levi''s 501 Jeans', 'Classic straight-fit denim jeans', '550e8400-e29b-41d4-a716-446655440021', 79.99, 45.00, 150, 30, 300, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440009', 'ZARA-DRESS-S', 'Zara Summer Dress', 'Elegant floral dress for summer', '550e8400-e29b-41d4-a716-446655440022', 59.99, 30.00, 80, 20, 160, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440010', 'ADIDAS-SNEAKER-42', 'Adidas Ultra Boost', 'Running shoes with Boost technology', '550e8400-e29b-41d4-a716-446655440023', 179.99, 120.00, 120, 25, 250, 'active', NOW(), NOW()),

-- Home & Garden products
('660e8400-e29b-41d4-a716-446655440011', 'IKEA-CHAIR-001', 'IKEA Office Chair', 'Ergonomic office chair with lumbar support', '550e8400-e29b-41d4-a716-446655440003', 199.99, 120.00, 40, 10, 80, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440012', 'DYSON-VACUUM', 'Dyson V15 Detect', 'Cordless vacuum cleaner with laser technology', '550e8400-e29b-41d4-a716-446655440003', 749.99, 500.00, 35, 8, 70, 'active', NOW(), NOW()),

-- Sports products
('660e8400-e29b-41d4-a716-446655440013', 'WILSON-TENNIS', 'Wilson Tennis Racket', 'Professional tennis racket for tournaments', '550e8400-e29b-41d4-a716-446655440004', 149.99, 90.00, 45, 12, 90, 'active', NOW(), NOW()),
('660e8400-e29b-41d4-a716-446655440014', 'YOGA-MAT-PRO', 'Premium Yoga Mat', 'Non-slip yoga mat with alignment lines', '550e8400-e29b-41d4-a716-446655440004', 79.99, 35.00, 100, 25, 200, 'active', NOW(), NOW()),

-- Books
('660e8400-e29b-41d4-a716-446655440015', 'BOOK-GOLANG', 'Go Programming Guide', 'Complete guide to Go programming language', '550e8400-e29b-41d4-a716-446655440005', 49.99, 25.00, 80, 20, 160, 'active', NOW(), NOW());

-- Create a sample user ID for transactions (in real app, this would be from users table)
-- For demo purposes, we'll use a fixed UUID

-- Insert dummy transactions
INSERT INTO transactions (id, product_id, type, quantity, reference, notes, created_by, created_at) VALUES
-- Stock in transactions (initial inventory)
('770e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001', 'in', 50, 'PO-2025-001', 'Initial stock - iPhone 15 Pro shipment', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '30 days'),
('770e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440002', 'in', 75, 'PO-2025-002', 'Initial stock - Samsung Galaxy S24 shipment', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '28 days'),
('770e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440003', 'in', 25, 'PO-2025-003', 'Initial stock - MacBook Pro delivery', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '25 days'),
('770e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440007', 'in', 200, 'PO-2025-004', 'Initial stock - Nike T-Shirts bulk order', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '20 days'),
('770e8400-e29b-41d4-a716-446655440005', '660e8400-e29b-41d4-a716-446655440010', 'in', 120, 'PO-2025-005', 'Initial stock - Adidas sneakers shipment', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '18 days'),

-- Stock out transactions (sales)
('770e8400-e29b-41d4-a716-446655440006', '660e8400-e29b-41d4-a716-446655440001', 'out', 5, 'SALE-2025-001', 'Online sale - iPhone 15 Pro', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '15 days'),
('770e8400-e29b-41d4-a716-446655440007', '660e8400-e29b-41d4-a716-446655440002', 'out', 8, 'SALE-2025-002', 'Store sale - Samsung Galaxy S24', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '12 days'),
('770e8400-e29b-41d4-a716-446655440008', '660e8400-e29b-41d4-a716-446655440007', 'out', 25, 'SALE-2025-003', 'Bulk order - Nike T-Shirts', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '10 days'),
('770e8400-e29b-41d4-a716-446655440009', '660e8400-e29b-41d4-a716-446655440005', 'out', 15, 'SALE-2025-004', 'Online sale - AirPods Pro', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '8 days'),
('770e8400-e29b-41d4-a716-446655440010', '660e8400-e29b-41d4-a716-446655440010', 'out', 12, 'SALE-2025-005', 'Store sale - Adidas sneakers', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '5 days'),

-- Stock adjustments
('770e8400-e29b-41d4-a716-446655440011', '660e8400-e29b-41d4-a716-446655440003', 'adjustment', -2, 'ADJ-2025-001', 'Inventory adjustment - damaged units', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '7 days'),
('770e8400-e29b-41d4-a716-446655440012', '660e8400-e29b-41d4-a716-446655440006', 'adjustment', 5, 'ADJ-2025-002', 'Inventory adjustment - found missing units', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '3 days'),

-- Recent transactions
('770e8400-e29b-41d4-a716-446655440013', '660e8400-e29b-41d4-a716-446655440015', 'in', 80, 'PO-2025-006', 'New shipment - Go Programming books', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '2 days'),
('770e8400-e29b-41d4-a716-446655440014', '660e8400-e29b-41d4-a716-446655440009', 'out', 3, 'SALE-2025-006', 'Online sale - Zara summer dress', '880e8400-e29b-41d4-a716-446655440001', NOW() - INTERVAL '1 day'),
('770e8400-e29b-41d4-a716-446655440015', '660e8400-e29b-41d4-a716-446655440014', 'in', 100, 'PO-2025-007', 'Restock - Premium yoga mats', '880e8400-e29b-41d4-a716-446655440001', NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Remove dummy data in reverse order (transactions first due to foreign key constraints)
DELETE FROM transactions WHERE id IN (
    '770e8400-e29b-41d4-a716-446655440001',
    '770e8400-e29b-41d4-a716-446655440002',
    '770e8400-e29b-41d4-a716-446655440003',
    '770e8400-e29b-41d4-a716-446655440004',
    '770e8400-e29b-41d4-a716-446655440005',
    '770e8400-e29b-41d4-a716-446655440006',
    '770e8400-e29b-41d4-a716-446655440007',
    '770e8400-e29b-41d4-a716-446655440008',
    '770e8400-e29b-41d4-a716-446655440009',
    '770e8400-e29b-41d4-a716-446655440010',
    '770e8400-e29b-41d4-a716-446655440011',
    '770e8400-e29b-41d4-a716-446655440012',
    '770e8400-e29b-41d4-a716-446655440013',
    '770e8400-e29b-41d4-a716-446655440014',
    '770e8400-e29b-41d4-a716-446655440015'
);

-- Remove dummy products
DELETE FROM products WHERE id LIKE '660e8400-e29b-41d4-a716-44665544%';

-- Remove dummy categories (subcategories first, then root categories)
DELETE FROM categories WHERE parent_id IS NOT NULL AND id LIKE '550e8400-e29b-41d4-a716-44665544%';
DELETE FROM categories WHERE parent_id IS NULL AND id LIKE '550e8400-e29b-41d4-a716-44665544%';

-- +goose StatementEnd
