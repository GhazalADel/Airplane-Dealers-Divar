--
-- PostgreSQL database dump
--

-- Dumped from database version 14.8 (Ubuntu 14.8-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.8 (Ubuntu 14.8-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories (id, name) VALUES (1, 'Big');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, username, password, role) VALUES (1, 'expert_1', 'gsdgfsdfg', 2);
INSERT INTO public.users (id, username, password, role) VALUES (2, 'adel_airplane', 'sdfsdfdsf', 4);
INSERT INTO public.users (id, username, password, role) VALUES (3, 'expert_2', 'asdfasdf', 2);


--
-- Data for Name: ads; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.ads (id, user_id, image, description, subject, price, category_id, status, fly_time, airplane_model, repair_check, expert_check, plane_age) VALUES (1, 2, NULL, 'Test description', 'Ad 1', 50000, 1, 'published', NULL, NULL, false, false, NULL);
INSERT INTO public.ads (id, user_id, image, description, subject, price, category_id, status, fly_time, airplane_model, repair_check, expert_check, plane_age) VALUES (2, 2, NULL, NULL, 'Ad 2', 200000, 1, 'published', NULL, NULL, false, false, NULL);


--
-- Data for Name: bookmarks; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: expert_ads; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.expert_ads (id, user_id, expert_id, ads_id, report, status, created_at) VALUES (1, 2, 3, 1, 'Evrything is ok!', 'Done', '2023-07-09 10:04:59.386335');
INSERT INTO public.expert_ads (id, user_id, expert_id, ads_id, report, status, created_at) VALUES (2, 2, 1, 2, NULL, 'In progress', '2023-07-09 10:34:21.826941');


--
-- Data for Name: repair_request; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.schema_migrations (version, dirty) VALUES (7, false);


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Name: ads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ads_id_seq', 2, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 1, true);


--
-- Name: expert_ads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.expert_ads_id_seq', 2, true);


--
-- Name: repair_request_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.repair_request_id_seq', 1, false);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 3, true);


--
-- PostgreSQL database dump complete
--

