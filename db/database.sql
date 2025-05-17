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
    SELECT COUNT(*) INTO db_count FROM pg_stat_activity WHERE datname = 'eoffice_lke';
    IF db_count = 0 THEN
        EXECUTE 'DROP DATABASE eoffice_lke';
    END IF;
EXCEPTION
    WHEN OTHERS THEN
        RAISE NOTICE 'Could not drop database, it may be in use.';
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'eoffice_lke') THEN
        CREATE DATABASE eoffice_lke WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';

        ALTER DATABASE eoffice_lke OWNER TO postgres;
    END IF;
END $$;

\connect eoffice_lke

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
-- Name: lke_rekap; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE IF NOT EXISTS lke_rekap (
    id integer NOT NULL,
    user_id integer,
    id_opd integer,
    tahun integer,
    kelengkapan numeric,
    nilai_capaian numeric,
    predikat_akhir character varying,
    predikat character varying,
    kelengkapan_m numeric,
    nilai_capaian_m numeric,
    predikat_akhir_m character varying,
    predikat_m character varying,
    status_evaluasi character varying,
    id_verifikator integer,
    id_ketua integer,
    id_evaluator integer,
    id_pengendali integer,
    updated_at integer,
    created_at integer
);

ALTER TABLE lke_rekap
ALTER COLUMN user_id
SET NOT NULL,
ALTER COLUMN id_opd
SET DATA TYPE integer,
ALTER COLUMN tahun
SET DATA TYPE integer,
ALTER COLUMN kelengkapan
SET DATA TYPE numeric,
ALTER COLUMN nilai_capaian
SET DATA TYPE numeric,
ALTER COLUMN predikat_akhir
SET DATA TYPE character varying,
ALTER COLUMN predikat
SET DATA TYPE character varying,
ALTER COLUMN status_evaluasi
SET DATA TYPE character varying,
ALTER COLUMN id_verifikator
SET DATA TYPE integer,
ALTER COLUMN id_ketua
SET DATA TYPE integer,
ALTER COLUMN id_evaluator
SET DATA TYPE integer,
ALTER COLUMN id_pengendali
SET DATA TYPE integer,
ALTER COLUMN updated_at
SET DATA TYPE integer,
ALTER COLUMN created_at
SET DATA TYPE integer;

ALTER TABLE lke_rekap OWNER TO postgres;

ALTER TABLE lke_rekap ADD COLUMN IF NOT EXISTS kelengkapan_m numeric;

ALTER TABLE lke_rekap
ADD COLUMN IF NOT EXISTS nilai_capaian_m numeric;

ALTER TABLE lke_rekap
ADD COLUMN IF NOT EXISTS predikat_akhir_m character varying;

ALTER TABLE lke_rekap
ADD COLUMN IF NOT EXISTS predikat_m character varying;

--
-- Name: lke_rekap_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE lke_rekap_id_seq START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER TABLE lke_rekap_id_seq OWNER TO postgres;

--
-- Name: lke_rekap_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE lke_rekap_id_seq OWNED BY lke_rekap.id;

--
-- Name: lke_rekomendasi; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE IF NOT EXISTS lke_rekomendasi (
    id integer NOT NULL,
    parent_id integer UNIQUE NOT NULL REFERENCES lke_rekap (id) ON UPDATE CASCADE ON DELETE CASCADE,
    ta1a text,
    ta1b text,
    ta1c text,
    ta2a text,
    ta2b text,
    ta2c text,
    tb1 text,
    tb2 text,
    tb3 text,
    tc1 text,
    tc2 text,
    tc3 text,
    td1 text,
    td2 text,
    td3 text,
    updated_at integer,
    created_at integer
);

ALTER TABLE lke_rekomendasi
ALTER COLUMN parent_id
SET NOT NULL,
ALTER COLUMN updated_at
SET DATA TYPE integer,
ALTER COLUMN created_at
SET DATA TYPE integer;

ALTER TABLE lke_rekomendasi OWNER TO postgres;

--
-- Name: lke_rekomendasi_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE lke_rekomendasi_id_seq START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER TABLE lke_rekomendasi_id_seq OWNER TO postgres;

--
-- Name: lke_rekomendasi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE lke_rekomendasi_id_seq OWNED BY lke_rekomendasi.id;

--
-- Name: lke_evaluasi; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE IF NOT EXISTS lke_evaluasi (
    id integer NOT NULL,
    lke_rekap_id integer NOT NULL,
    user_id integer,
    kode_evaluasi character varying,
    jawaban text,
    berkas text,
    catatan text,
    updated_at integer,
    created_at integer
);

ALTER TABLE lke_evaluasi
ALTER COLUMN lke_rekap_id
SET NOT NULL,
ALTER COLUMN user_id
SET DATA TYPE integer,
ALTER COLUMN kode_evaluasi
SET DATA TYPE character varying,
ALTER COLUMN jawaban
SET DATA TYPE text,
ALTER COLUMN berkas
SET DATA TYPE text,
ALTER COLUMN catatan
SET DATA TYPE text,
ALTER COLUMN updated_at
SET DATA TYPE integer,
ALTER COLUMN created_at
SET DATA TYPE integer;

ALTER TABLE lke_evaluasi OWNER TO postgres;

-- Add evaluasi column to lke_evaluasi table
ALTER TABLE lke_evaluasi ADD COLUMN IF NOT EXISTS evaluasi text;

--
-- Name: lke_evaluasi_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE lke_evaluasi_id_seq START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER TABLE lke_evaluasi_id_seq OWNER TO postgres;

--
-- Name: lke_evaluasi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE lke_evaluasi_id_seq OWNED BY lke_evaluasi.id;

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY article
ALTER COLUMN id
SET DEFAULT nextval('article_id_seq'::regclass);

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user"
ALTER COLUMN id
SET DEFAULT nextval('user_id_seq'::regclass);

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lke_rekap
ALTER COLUMN id
SET DEFAULT nextval('lke_rekap_id_seq'::regclass);

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lke_rekomendasi
ALTER COLUMN id
SET DEFAULT nextval(
    'lke_rekomendasi_id_seq'::regclass
);

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lke_evaluasi
ALTER COLUMN id
SET DEFAULT nextval(
    'lke_evaluasi_id_seq'::regclass
);

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
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--


COPY "user" (id, email, password, name, updated_at, created_at) FROM stdin;
\.

--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval ('user_id_seq', 1, false);

--
-- Name: article_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY article ADD CONSTRAINT article_id PRIMARY KEY (id);

--
-- Name: user_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY "user" ADD CONSTRAINT user_id PRIMARY KEY (id);

--
-- Name: lke_rekap_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY lke_rekap
ADD CONSTRAINT lke_rekap_id PRIMARY KEY (id);

--
-- Name: lke_rekomendasi_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY lke_rekomendasi
ADD CONSTRAINT lke_rekomendasi_id PRIMARY KEY (id);
--
-- Name: lke_evaluasi_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY lke_evaluasi
ADD CONSTRAINT lke_evaluasi_id PRIMARY KEY (id);

--
-- Name: lke_komponen; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--

CREATE TABLE IF NOT EXISTS lke_komponen (
    id integer NOT NULL,
    kode_evaluasi character varying NOT NULL,
    bobot numeric NOT NULL,
    komponen text,
    eviden text,
    level character varying,
    updated_at integer,
    created_at integer
);

ALTER TABLE lke_komponen
ALTER COLUMN kode_evaluasi
SET NOT NULL,
ALTER COLUMN bobot
SET NOT NULL,
ALTER COLUMN updated_at
SET DATA TYPE integer,
ALTER COLUMN created_at
SET DATA TYPE integer;

ALTER TABLE lke_komponen OWNER TO postgres;

--
-- Name: lke_komponen_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE lke_komponen_id_seq START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER TABLE lke_komponen_id_seq OWNER TO postgres;

--
-- Name: lke_komponen_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE lke_komponen_id_seq OWNED BY lke_komponen.id;

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lke_komponen
ALTER COLUMN id
SET DEFAULT nextval(
    'lke_komponen_id_seq'::regclass
);

--
-- Name: lke_komponen_id; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace:
--

ALTER TABLE ONLY lke_komponen
ADD CONSTRAINT lke_komponen_id PRIMARY KEY (id);

--
-- Name: article_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY article
ADD CONSTRAINT article_user_id FOREIGN KEY (user_id) REFERENCES "user" (id) ON UPDATE CASCADE ON DELETE CASCADE;

--
-- Name: lke_rekap_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lke_rekap
ADD CONSTRAINT lke_rekap_user_id FOREIGN KEY (user_id) REFERENCES "user" (id) ON UPDATE CASCADE ON DELETE CASCADE;

--
-- Name: lke_evaluasi_lke_rekap_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lke_evaluasi
ADD CONSTRAINT lke_evaluasi_lke_rekap_id FOREIGN KEY (lke_rekap_id) REFERENCES lke_rekap (id) ON UPDATE CASCADE ON DELETE CASCADE;

--
-- Name: lke_evaluasi_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lke_evaluasi
ADD CONSTRAINT lke_evaluasi_user_id FOREIGN KEY (user_id) REFERENCES "user" (id) ON UPDATE CASCADE ON DELETE CASCADE;

--
-- Name: article create_article_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_article_created_at BEFORE INSERT ON article FOR EACH ROW EXECUTE PROCEDURE created_at_column();

--
-- Name: create_lke_rekap_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_lke_rekap_created_at BEFORE INSERT ON lke_rekap FOR EACH ROW EXECUTE PROCEDURE created_at_column();

--
-- Name: create_lke_evaluasi_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_lke_evaluasi_created_at BEFORE INSERT ON lke_evaluasi FOR EACH ROW EXECUTE PROCEDURE created_at_column();

--
-- Name: create_lke_komponen_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_lke_komponen_created_at BEFORE INSERT ON lke_komponen FOR EACH ROW EXECUTE PROCEDURE created_at_column();

--
-- Name: create_lke_rekomendasi_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_lke_rekomendasi_created_at BEFORE INSERT ON lke_rekomendasi FOR EACH ROW EXECUTE PROCEDURE created_at_column();

--
-- Name: user create_user_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_user_created_at BEFORE INSERT ON "user" FOR EACH ROW EXECUTE PROCEDURE created_at_column();

--
-- Name: article update_article_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_article_updated_at BEFORE UPDATE ON article FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- Name: update_lke_rekap_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_lke_rekap_updated_at BEFORE UPDATE ON lke_rekap FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- Name: update_lke_evaluasi_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_lke_evaluasi_updated_at BEFORE UPDATE ON lke_evaluasi FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- Name: update_lke_komponen_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_lke_komponen_updated_at BEFORE UPDATE ON lke_komponen FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- Name: update_lke_rekomendasi_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_lke_rekomendasi_updated_at BEFORE UPDATE ON lke_rekomendasi FOR EACH ROW EXECUTE PROCEDURE update_at_column();

--
-- Name: user update_user_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_user_updated_at BEFORE UPDATE ON "user" FOR EACH ROW EXECUTE PROCEDURE update_at_column();

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