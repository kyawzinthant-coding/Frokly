-- +goose Up
ALTER TABLE products ALTER COLUMN price TYPE BIGINT;

ALTER TABLE menu_categories ADD COLUMN archived_at TIMESTAMPTZ;
ALTER TABLE products ADD COLUMN archived_at TIMESTAMPTZ;

ALTER TABLE menu_categories DROP CONSTRAINT menu_categories_vendor_id_name_key;
CREATE UNIQUE INDEX menu_categories_vendor_id_name_active_key
    ON menu_categories (vendor_id, lower(name)) WHERE archived_at IS NULL;

ALTER TABLE products DROP CONSTRAINT products_vendor_id_name_key;
CREATE UNIQUE INDEX products_vendor_id_name_active_key
    ON products (vendor_id, lower(name)) WHERE archived_at IS NULL;

ALTER TABLE menu_categories ADD CONSTRAINT menu_categories_id_vendor_id_key UNIQUE (id, vendor_id);
ALTER TABLE products DROP CONSTRAINT products_menu_categories_id_fkey;
ALTER TABLE products ADD CONSTRAINT products_category_same_vendor_fkey
    FOREIGN KEY (menu_categories_id, vendor_id) REFERENCES menu_categories (id, vendor_id) ON DELETE RESTRICT;

-- +goose Down
ALTER TABLE products DROP CONSTRAINT products_category_same_vendor_fkey;
ALTER TABLE products ADD CONSTRAINT products_menu_categories_id_fkey
    FOREIGN KEY (menu_categories_id) REFERENCES menu_categories (id) ON DELETE CASCADE;
ALTER TABLE menu_categories DROP CONSTRAINT menu_categories_id_vendor_id_key;

DROP INDEX products_vendor_id_name_active_key;
ALTER TABLE products ADD CONSTRAINT products_vendor_id_name_key UNIQUE (vendor_id, name);

DROP INDEX menu_categories_vendor_id_name_active_key;
ALTER TABLE menu_categories ADD CONSTRAINT menu_categories_vendor_id_name_key UNIQUE (vendor_id, name);

ALTER TABLE products DROP COLUMN archived_at;
ALTER TABLE menu_categories DROP COLUMN archived_at;

ALTER TABLE products ALTER COLUMN price TYPE INT;
