# Dummy Data Documentation

## Overview
Dummy data telah berhasil dibuat untuk testing dan development. Migration file `20260202081721_init_dummy_data.sql` berisi data sample yang realistis untuk menguji semua fitur aplikasi inventory.

## Data Structure

### 1. Categories (11 total)

#### Root Categories (5):
- **Electronics** (`550e8400-e29b-41d4-a716-446655440001`)
  - Electronic devices and gadgets
- **Clothing** (`550e8400-e29b-41d4-a716-446655440002`)
  - Apparel and fashion items
- **Home & Garden** (`550e8400-e29b-41d4-a716-446655440003`)
  - Home improvement and gardening supplies
- **Sports & Outdoors** (`550e8400-e29b-41d4-a716-446655440004`)
  - Sports equipment and outdoor activities
- **Books & Media** (`550e8400-e29b-41d4-a716-446655440005`)
  - Books, magazines, and media content

#### Sub Categories (6):
**Electronics Subcategories:**
- **Smartphones** (`550e8400-e29b-41d4-a716-446655440011`)
- **Laptops** (`550e8400-e29b-41d4-a716-446655440012`)
- **Audio** (`550e8400-e29b-41d4-a716-446655440013`)

**Clothing Subcategories:**
- **Men's Clothing** (`550e8400-e29b-41d4-a716-446655440021`)
- **Women's Clothing** (`550e8400-e29b-41d4-a716-446655440022`)
- **Shoes** (`550e8400-e29b-41d4-a716-446655440023`)

### 2. Products (15 total)

#### Electronics (6 products):
| SKU | Name | Category | Price | Stock | Min/Max Stock |
|-----|------|----------|-------|-------|---------------|
| `IPHONE-15-PRO` | iPhone 15 Pro | Smartphones | $999.99 | 50 | 10/100 |
| `SAMSUNG-S24` | Samsung Galaxy S24 | Smartphones | $899.99 | 75 | 15/150 |
| `MACBOOK-PRO-16` | MacBook Pro 16" | Laptops | $2,499.99 | 25 | 5/50 |
| `DELL-XPS-13` | Dell XPS 13 | Laptops | $1,299.99 | 30 | 8/60 |
| `AIRPODS-PRO` | AirPods Pro 2nd Gen | Audio | $249.99 | 100 | 20/200 |
| `SONY-WH1000XM5` | Sony WH-1000XM5 | Audio | $399.99 | 60 | 15/120 |

#### Clothing (4 products):
| SKU | Name | Category | Price | Stock | Min/Max Stock |
|-----|------|----------|-------|-------|---------------|
| `NIKE-TSHIRT-M` | Nike Men's T-Shirt | Men's Clothing | $29.99 | 200 | 50/500 |
| `LEVI-JEANS-32` | Levi's 501 Jeans | Men's Clothing | $79.99 | 150 | 30/300 |
| `ZARA-DRESS-S` | Zara Summer Dress | Women's Clothing | $59.99 | 80 | 20/160 |
| `ADIDAS-SNEAKER-42` | Adidas Ultra Boost | Shoes | $179.99 | 120 | 25/250 |

#### Home & Garden (2 products):
| SKU | Name | Price | Stock | Min/Max Stock |
|-----|------|-------|-------|---------------|
| `IKEA-CHAIR-001` | IKEA Office Chair | $199.99 | 40 | 10/80 |
| `DYSON-VACUUM` | Dyson V15 Detect | $749.99 | 35 | 8/70 |

#### Sports & Outdoors (2 products):
| SKU | Name | Price | Stock | Min/Max Stock |
|-----|------|-------|-------|---------------|
| `WILSON-TENNIS` | Wilson Tennis Racket | $149.99 | 45 | 12/90 |
| `YOGA-MAT-PRO` | Premium Yoga Mat | $79.99 | 100 | 25/200 |

#### Books & Media (1 product):
| SKU | Name | Price | Stock | Min/Max Stock |
|-----|------|-------|-------|---------------|
| `BOOK-GOLANG` | Go Programming Guide | $49.99 | 80 | 20/160 |

### 3. Transactions (15 total)

#### Stock In Transactions (7):
- **Initial Stock Transactions** (5 items) - 30-18 days ago
  - iPhone 15 Pro: +50 units (PO-2025-001)
  - Samsung Galaxy S24: +75 units (PO-2025-002)
  - MacBook Pro: +25 units (PO-2025-003)
  - Nike T-Shirts: +200 units (PO-2025-004)
  - Adidas Sneakers: +120 units (PO-2025-005)

- **Recent Restocks** (2 items)
  - Go Programming books: +80 units (PO-2025-006) - 2 days ago
  - Premium yoga mats: +100 units (PO-2025-007) - Today

#### Stock Out Transactions (5):
- iPhone 15 Pro: -5 units (SALE-2025-001) - 15 days ago
- Samsung Galaxy S24: -8 units (SALE-2025-002) - 12 days ago
- Nike T-Shirts: -25 units (SALE-2025-003) - 10 days ago
- AirPods Pro: -15 units (SALE-2025-004) - 8 days ago
- Adidas sneakers: -12 units (SALE-2025-005) - 5 days ago
- Zara summer dress: -3 units (SALE-2025-006) - 1 day ago

#### Stock Adjustments (2):
- MacBook Pro: -2 units (ADJ-2025-001) - damaged units - 7 days ago
- Sony headphones: +5 units (ADJ-2025-002) - found missing units - 3 days ago

## Usage for Testing

### API Testing Scenarios:

1. **List Products**: Test pagination with 15 products
2. **Search Products**: Test search with various keywords (Nike, Apple, etc.)
3. **Low Stock Alerts**: Some products have stock near min_stock threshold
4. **Category Hierarchy**: Test parent-child category relationships
5. **Transaction History**: Each product has realistic transaction history
6. **Stock Management**: Test stock in/out/adjustment operations

### Test Cases:

```bash
# Get all products
GET /api/v1/products

# Search products
GET /api/v1/products/search?q=iPhone

# Get low stock products  
GET /api/v1/products/low-stock

# Get product by ID
GET /api/v1/products/660e8400-e29b-41d4-a716-446655440001
```

## Sample User ID

For transactions, a sample user ID is used:
- `880e8400-e29b-41d4-a716-446655440001`

In a real application, this would reference actual users from a users table.

## Migration Commands

```bash
# Run migrations (including dummy data)
make migrate-up

# Rollback dummy data
make migrate-down

# Check migration status
make migrate-status
```

## Cleanup

The migration includes proper rollback (`-- +goose Down`) section that removes all dummy data in the correct order (transactions first, then products, then categories) to respect foreign key constraints.
