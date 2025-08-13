--
-- PostgreSQL database dump
--
SET statement_timeout = 0;

SET lock_timeout = 0;

SET client_encoding = 'UTF8';

SET standard_conforming_strings = on;

SET check_function_bodies = false;

SET client_min_messages = warning;

--
-- Name: gin_public; Type: DATABASE; Schema: -; Owner: postgres
--
DO $$
DECLARE
    db_count INTEGER;
BEGIN
    SELECT COUNT(*) INTO db_count FROM pg_stat_activity WHERE datname = 'eoffice_aset';
    IF db_count = 0 THEN
        EXECUTE 'DROP DATABASE eoffice_aset';
    END IF;
EXCEPTION
    WHEN OTHERS THEN
        RAISE NOTICE 'Could not drop database, it may be in use.';
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'eoffice_aset') THEN
        CREATE DATABASE eoffice_aset WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';

        ALTER DATABASE eoffice_aset OWNER TO postgres;
    END IF;
END $$;

\connect eoffice_aset

SET statement_timeout = 0;

SET lock_timeout = 0;

SET client_encoding = 'UTF8';

SET standard_conforming_strings = on;

SET check_function_bodies = false;

SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;

--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';

CREATE FUNCTION created_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
	NEW.updated_at = EXTRACT(EPOCH FROM NOW());
	NEW.created_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$$;

ALTER FUNCTION public.created_at_column() OWNER TO postgres;

--
-- Name: update_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION update_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
    NEW.updated_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$$;

ALTER FUNCTION public.update_at_column() OWNER TO postgres;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;
--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE IF NOT EXISTS "user" (
    id integer NOT NULL,
    email character varying,
    password character varying,
    name character varying,
    updated_at integer,
    created_at integer
);

ALTER TABLE "user"
ALTER COLUMN email
SET DATA TYPE character varying,
ALTER COLUMN password
SET DATA TYPE character varying,
ALTER COLUMN name
SET DATA TYPE character varying,
ALTER COLUMN updated_at
SET DATA TYPE integer,
ALTER COLUMN created_at
SET DATA TYPE integer;

ALTER TABLE "user" OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE user_id_seq START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER TABLE user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user"
ALTER COLUMN id
SET DEFAULT nextval('user_id_seq'::regclass);
--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--


COPY "user" (id, email, password, name, updated_at, created_at) FROM stdin;
\.

--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval ('user_id_seq', 1, false);
--
-- Name: user_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY "user" ADD CONSTRAINT user_id PRIMARY KEY (id);
--
-- Name: user create_user_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_user_created_at BEFORE INSERT ON "user" FOR EACH ROW EXECUTE PROCEDURE created_at_column();
--
-- Name: user update_user_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_user_updated_at BEFORE UPDATE ON "user" FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- Name: article; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE IF NOT EXISTS article (
    id integer NOT NULL,
    user_id integer,
    title character varying,
    content text,
    updated_at integer,
    created_at integer
);

ALTER TABLE article
ALTER COLUMN user_id
SET NOT NULL,
ALTER COLUMN title
SET DATA TYPE character varying,
ALTER COLUMN content
SET DATA TYPE text,
ALTER COLUMN updated_at
SET DATA TYPE integer,
ALTER COLUMN created_at
SET DATA TYPE integer;

ALTER TABLE article OWNER TO postgres;

--
-- Name: article_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE article_id_seq START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER TABLE article_id_seq OWNER TO postgres;

--
-- Name: article_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE article_id_seq OWNED BY article.id;

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY article
ALTER COLUMN id
SET DEFAULT nextval('article_id_seq'::regclass);

--
-- Data for Name: article; Type: TABLE DATA; Schema: public; Owner: postgres
--


COPY article (id, user_id, title, content, updated_at, created_at) FROM stdin;
\.

--
-- Name: article_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval ('article_id_seq', 1, false);

--
-- Name: article_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY article ADD CONSTRAINT article_id PRIMARY KEY (id);

--
-- Name: article_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY article
ADD CONSTRAINT article_user_id FOREIGN KEY (user_id) REFERENCES "user" (id) ON UPDATE CASCADE ON DELETE CASCADE;
--
-- Name: article create_article_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_article_created_at BEFORE INSERT ON article FOR EACH ROW EXECUTE PROCEDURE created_at_column();
--
-- Name: article update_article_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_article_updated_at BEFORE UPDATE ON article FOR EACH ROW EXECUTE PROCEDURE update_at_column();
-- Previous schema definitions remain unchanged until we add our new tables

-- 1. Create reference tables for vehicle attributes

CREATE TABLE IF NOT EXISTS vehicle_wheels (
    id SERIAL PRIMARY KEY,
    count INTEGER NOT NULL UNIQUE,
    description TEXT,
    created_at INTEGER,
    updated_at INTEGER
);

CREATE TABLE IF NOT EXISTS vehicle_model (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at INTEGER,
    updated_at INTEGER
);

CREATE TABLE IF NOT EXISTS vehicle_color (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    hex_code VARCHAR(7),
    created_at INTEGER,
    updated_at INTEGER
);

CREATE TABLE IF NOT EXISTS vehicle_fuel (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at INTEGER,
    updated_at INTEGER
);

CREATE TABLE IF NOT EXISTS utilization_company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    address TEXT,
    contact VARCHAR(20),
    created_at INTEGER,
    updated_at INTEGER
);

CREATE TABLE IF NOT EXISTS vehicle_owning (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at INTEGER,
    updated_at INTEGER
);

CREATE TABLE IF NOT EXISTS vehicle_brand (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    country VARCHAR(100),
    founded INTEGER,
    created_at INTEGER,
    updated_at INTEGER
);

-- 2. Create main asset table
CREATE TABLE IF NOT EXISTS vehicle_asset (
    id SERIAL PRIMARY KEY,
    license_plate VARCHAR(20) NOT NULL UNIQUE,
    stnk_status VARCHAR(20) CHECK (
        stnk_status IN ('AVAILABLE', 'NOT_AVAILABLE')
    ),
    bpkb_number VARCHAR(50),
    bpkb_status VARCHAR(20) CHECK (
        bpkb_status IN ('AVAILABLE', 'NOT_AVAILABLE')
    ),
    chassis_number VARCHAR(50),
    machine_number VARCHAR(50),
    wheels_id INTEGER REFERENCES vehicle_wheels (id),
    model_id INTEGER REFERENCES vehicle_model (id),
    type_name VARCHAR(200) NOT NULL,
    color_id INTEGER REFERENCES vehicle_color (id),
    fuel_id INTEGER REFERENCES vehicle_fuel (id),
    owning_id INTEGER REFERENCES vehicle_owning (id),
    brand_id INTEGER REFERENCES vehicle_brand (id),
    cc_capacity INTEGER,
    manufacture_year INTEGER,
    tax_due_date INTEGER,
    last_tax_payment_date INTEGER,
    current_owner VARCHAR(100),
    company_id INTEGER,
    stnk_photo_path TEXT,
    stnk_photo_verified BOOLEAN DEFAULT FALSE,
    bpkb_photo_path TEXT,
    bpkb_photo_verified BOOLEAN DEFAULT FALSE,
    owner_id_photo_path TEXT,
    owner_id_photo_verified BOOLEAN DEFAULT FALSE,
    vehicle_photo_path TEXT,
    vehicle_photo_verified BOOLEAN DEFAULT FALSE,
    payment_billing_photo_path TEXT,
    payment_billing_date INTEGER,
    recommendation_document_path TEXT,
    e_sign_status BOOLEAN DEFAULT FALSE,
    notes TEXT,
    status VARCHAR(100),
    created_at INTEGER,
    updated_at INTEGER,
    created_by INTEGER REFERENCES "user" (id)
);

-- Add verified columns if they don't exist (for existing databases)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='vehicle_asset' AND column_name='stnk_photo_verified') THEN
        ALTER TABLE vehicle_asset ADD COLUMN stnk_photo_verified BOOLEAN DEFAULT FALSE;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='vehicle_asset' AND column_name='bpkb_photo_verified') THEN
        ALTER TABLE vehicle_asset ADD COLUMN bpkb_photo_verified BOOLEAN DEFAULT FALSE;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='vehicle_asset' AND column_name='owner_id_photo_verified') THEN
        ALTER TABLE vehicle_asset ADD COLUMN owner_id_photo_verified BOOLEAN DEFAULT FALSE;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='vehicle_asset' AND column_name='vehicle_photo_verified') THEN
        ALTER TABLE vehicle_asset ADD COLUMN vehicle_photo_verified BOOLEAN DEFAULT FALSE;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='vehicle_asset' AND column_name='payment_billing_date') THEN
        ALTER TABLE vehicle_asset ADD COLUMN payment_billing_date INTEGER;
    END IF;
END $$;

-- Create indexes for better query performance
CREATE INDEX idx_vehicle_asset_company ON vehicle_asset (company_id);

CREATE INDEX idx_vehicle_asset_plate ON vehicle_asset (license_plate);

CREATE INDEX idx_vehicle_asset_owner ON vehicle_asset (current_owner);

-- Add check constraint for company_id to ensure it's a positive integer
ALTER TABLE vehicle_asset
ADD CONSTRAINT chk_vehicle_asset_company_id_positive CHECK (company_id > 0);

-- Add triggers for timestamp columns
CREATE TRIGGER create_vehicle_asset_created_at
BEFORE INSERT ON vehicle_asset
FOR EACH ROW EXECUTE PROCEDURE created_at_column();

CREATE TRIGGER update_vehicle_asset_updated_at
BEFORE UPDATE ON vehicle_asset
FOR EACH ROW EXECUTE PROCEDURE update_at_column();

-- Rest of existing schema remains unchanged
--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;

REVOKE ALL ON SCHEMA public FROM postgres;

GRANT ALL ON SCHEMA public TO postgres;

GRANT ALL ON SCHEMA public TO PUBLIC;

--
-- PostgreSQL database dump complete
--