CREATE TABLE free_market (
    id SERIAL PRIMARY KEY NOT NULL,
    latitude decimal NOT NULL,
    longitude decimal NOT NULL,
    set_cens VARCHAR(15) NOT NULL,
    areap VARCHAR(13) NOT NULL,
    region_code integer NOT NULL,
    region VARCHAR(18) NOT NULL,
    subprefecture_code integer NOT NULL,
    subprefecture VARCHAR(25) NOT NULL,
    region_five VARCHAR(6) NOT NULL,
    region_eight VARCHAR(7) NOT NULL,
    market_name VARCHAR(30) NOT NULL,
    register VARCHAR(6) NOT NULL,
    adress VARCHAR(34) NOT NULL,
    adress_number VARCHAR(5),
    district VARCHAR(20) NOT NULL,
    reference VARCHAR(24)
);