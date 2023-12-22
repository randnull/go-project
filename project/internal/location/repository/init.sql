CREATE TABLE IF NOT EXISTS drivers (
    id SERIAL PRIMARY KEY,
    lat DOUBLE PRECISION,
    lng DOUBLE PRECISION,
    name VARCHAR(255),
    auto VARCHAR(255)
);

INSERT INTO drivers (lat, lng, name, auto) VALUES (123.23, 123.42, 'ivan', 'toyota');
INSERT INTO drivers (lat, lng, name, auto) VALUES (125.23, 121.42, 'misha', 'bmw');
INSERT INTO drivers (lat, lng, name, auto) VALUES (130.0, 135.0, 'anna', 'honda');
INSERT INTO drivers (lat, lng, name, auto) VALUES (140.0, 145.0, 'john', 'audi');
INSERT INTO drivers (lat, lng, name, auto) VALUES (150.0, 155.0, 'emily', 'mercedes');
INSERT INTO drivers (lat, lng, name, auto) VALUES (160.0, 165.0, 'alex', 'volkswagen');
INSERT INTO drivers (lat, lng, name, auto) VALUES (170.0, 175.0, 'olivia', 'subaru');
INSERT INTO drivers (lat, lng, name, auto) VALUES (180.0, 185.0, 'liam', 'porsche');
INSERT INTO drivers (lat, lng, name, auto) VALUES (40.7128, -74.0060, 'misha', 'toyota');
INSERT INTO drivers (lat, lng, name, auto) VALUES (34.0522, -118.2437, 'sasha', 'bmw');
INSERT INTO drivers (lat, lng, name, auto) VALUES (51.509865, -0.118092, 'Ya', 'honda');
INSERT INTO drivers (lat, lng, name, auto) VALUES (48.8566, 2.3522, 'Kirill', 'audi');