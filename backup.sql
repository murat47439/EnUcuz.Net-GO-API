-- ========================
-- MARKALAR
-- ========================
CREATE TABLE brands (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- ========================
-- ORTAK ÜRÜN TABLOSU
-- ========================
CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(20) NOT NULL, -- phone, car, ac
    brand_id BIGINT NOT NULL,
    launch_announced DATE,
    launch_released DATE,
    launch_status VARCHAR(50),
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);

CREATE INDEX idx_products_category ON products(category);
CREATE INDEX idx_products_brand ON products(brand_id);

-- ========================
-- TELEFON DETAYLARI
-- ========================
CREATE TABLE phone_details (
    product_id BIGINT PRIMARY KEY,
    current_os VARCHAR(100),
    upgradable_to VARCHAR(100),
    chipset VARCHAR(100),
    cpu VARCHAR(255),
    gpu VARCHAR(100),
    body_dimensions VARCHAR(100),
    body_weight VARCHAR(50),
    body_build TEXT,
    sim_info VARCHAR(100),
    network_technology VARCHAR(100),
    network_speed VARCHAR(100),
    network_2g BOOLEAN,
    network_3g BOOLEAN,
    network_4g BOOLEAN,
    network_5g BOOLEAN,
    positioning_gps BOOLEAN,
    nfc BOOLEAN,
    radio BOOLEAN,
    wlan TEXT,
    bluetooth TEXT,
    usb TEXT,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_phone_network_5g ON phone_details(network_5g);
CREATE INDEX idx_phone_os ON phone_details(current_os);

-- ========================
-- ARABA DETAYLARI
-- ========================
CREATE TABLE car_details (
    product_id BIGINT PRIMARY KEY,
    engine_type VARCHAR(100),
    horsepower INT,
    torque INT,
    transmission VARCHAR(50),
    fuel_type VARCHAR(50),
    seating_capacity INT,
    weight INT,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_car_engine ON car_details(engine_type);
CREATE INDEX idx_car_fuel ON car_details(fuel_type);

-- ========================
-- KLİMA DETAYLARI
-- ========================
CREATE TABLE ac_details (
    product_id BIGINT PRIMARY KEY,
    cooling_capacity INT,
    heating_capacity INT,
    noise_level INT,
    energy_efficiency VARCHAR(20),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_ac_noise_level ON ac_details(noise_level);

-- ========================
-- ESNEK ÖZELLİKLER
-- ========================
CREATE TABLE product_attributes (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    attribute_name VARCHAR(100),
    attribute_value VARCHAR(255),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_attr_product_name ON product_attributes(product_id, attribute_name);

-- ========================
-- TELEFON ÖZEL TABLOLARI
-- ========================

-- BATARYA
CREATE TABLE battery (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    technology VARCHAR(50),
    capacity VARCHAR(50),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE battery_charging (
    id BIGSERIAL PRIMARY KEY,
    battery_id BIGINT NOT NULL,
    type VARCHAR(50),
    power VARCHAR(50),
    description VARCHAR(255),
    FOREIGN KEY (battery_id) REFERENCES battery(id) ON DELETE CASCADE
);

-- MEMORY
CREATE TABLE memory_options (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    storage VARCHAR(50),
    ram VARCHAR(50),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- KAMERA
CREATE TABLE cameras (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    type VARCHAR(20), -- main/selfie
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE camera_lenses (
    id BIGSERIAL PRIMARY KEY,
    camera_id BIGINT NOT NULL,
    megapixels VARCHAR(20),
    aperture VARCHAR(20),
    focal_length VARCHAR(20),
    sensor_size VARCHAR(20),
    type VARCHAR(50),
    pixel_size VARCHAR(20),
    other_features TEXT,
    lens_order INT,
    FOREIGN KEY (camera_id) REFERENCES cameras(id) ON DELETE CASCADE
);

CREATE TABLE camera_features (
    id BIGSERIAL PRIMARY KEY,
    camera_id BIGINT NOT NULL,
    feature VARCHAR(100),
    FOREIGN KEY (camera_id) REFERENCES cameras(id) ON DELETE CASCADE
);

CREATE TABLE camera_video (
    id BIGSERIAL PRIMARY KEY,
    camera_id BIGINT NOT NULL,
    video_spec VARCHAR(50),
    FOREIGN KEY (camera_id) REFERENCES cameras(id) ON DELETE CASCADE
);

-- DISPLAY
CREATE TABLE display (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    panel_type VARCHAR(50),
    size_inches VARCHAR(20),
    resolution VARCHAR(50),
    aspect_ratio VARCHAR(20),
    other_features TEXT,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- SENSÖRLER
CREATE TABLE sensors (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    sensor_name VARCHAR(100),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- RENKLER
CREATE TABLE colors (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    color_name_en VARCHAR(50),
    color_name_tr VARCHAR(50),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
