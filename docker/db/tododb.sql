--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.8
-- Dumped by pg_dump version 9.5.8

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

-- Creation des roles

CREATE ROLE todo_admin;
ALTER ROLE todo_admin WITH SUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'md5a83577d2c545e15f643ef5bb3b901729';

CREATE ROLE todo_user;
ALTER ROLE todo_user WITH NOSUPERUSER INHERIT CREATEROLE CREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'md5c54b9606b650ec765c12bacb71db5b36';

GRANT ALL PRIVILEGES ON DATABASE tododb TO todo_admin;
GRANT ALL PRIVILEGES ON DATABASE tododb TO todo_user;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: task; Type: TABLE; Schema: public; Owner: todo_admin
--

CREATE TABLE task (
    id integer NOT NULL,
    creation timestamp without time zone DEFAULT now(),
    description text
);


ALTER TABLE task OWNER TO todo_admin;
GRANT ALL PRIVILEGES ON TABLE task TO todo_user;

--
-- Name: task_id_seq; Type: SEQUENCE; Schema: public; Owner: todo_admin
--

CREATE SEQUENCE task_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE task_id_seq OWNER TO todo_admin;
GRANT ALL PRIVILEGES ON SEQUENCE task_id_seq TO todo_user;

--
-- Name: task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: todo_admin
--

ALTER SEQUENCE task_id_seq OWNED BY task.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: todo_admin
--

ALTER TABLE ONLY task ALTER COLUMN id SET DEFAULT nextval('task_id_seq'::regclass);


--
-- Data for Name: task; Type: TABLE DATA; Schema: public; Owner: todo_admin
--

COPY task (id, creation, description) FROM stdin;
3	2017-09-10 11:00:11.690247	Apprendre a utiliser PostgreSQL
4	2017-09-10 11:00:49.071926	Creer la base tododb
6	2017-09-10 11:04:07.419318	Creer un app type en ligne de commande pour acceder a tododb
8	2017-09-10 11:05:06.737861	Creer les images docker et les scripts compose pour deployer API
\.


--
-- Name: task_id_seq; Type: SEQUENCE SET; Schema: public; Owner: todo_admin
--

SELECT pg_catalog.setval('task_id_seq', 20, true);


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

