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

INSERT INTO public.categories (id, name) VALUES (1, 'Fighter aircrafts');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, username, password, role, token, is_active) VALUES (3, 'expert2', '$2a$10$siQbgPqRp73XO5XrzCJXA.qRhJZ8o81X0eBFsxQ.dZKumr.zXsGG6', 'Expert', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTY4OTA3NjI3MiwiaWQiOjN9.VtXtZHfMbLKeBeedAxOhayQrEByCWqiiDmMpdYKZcXg', true);
INSERT INTO public.users (id, username, password, role, token, is_active) VALUES (5, 'user2', '$2a$10$De72U25sFMXE8D./8SFDO.TEiPhAvW1S5ciW2nRS5RkudfnQs3pam', 'Airline', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTY4OTA3NzkwOCwiaWQiOjV9.Ku68UGio-usWfKgIX08Sv2Iz5sejA48bwdvvnrTvad0', true);
INSERT INTO public.users (id, username, password, role, token, is_active) VALUES (4, 'matin', '$2a$10$c5NLMAH6NEyN.R/YJ5V7MuD5YXeR05ClP42vh/YkuGH4k40Zvhx7G', 'Matin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTY4OTA4MDAwMywiaWQiOjR9.eLsWi1IdUJPugSnRcNCy81MM18FQ_U1TikPMXch63XQ', true);
INSERT INTO public.users (id, username, password, role, token, is_active) VALUES (6, 'admin', '$2a$10$Ns0EFT5hrBEwHo9pisXTMu.Ge4ME0U9HbfVo/cPXqVfEXrNJN43iy', 'Admin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTY4OTA4MDE5OCwiaWQiOjZ9.bLeCAE9ywy8FwgQmREA7TTVUV7sd0IdJofwQlyQGXOQ', true);
INSERT INTO public.users (id, username, password, role, token, is_active) VALUES (1, 'expert1', '$2a$10$zM2Q5PNMg6si1oeSeZijD.bLyaeLEfORnQ8btxGewGN/xA5aszSam', 'Expert', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTY4OTA4MDQxMiwiaWQiOjF9.bSl_tAGPSyAntx9wNl1yVZJLue0q1BDD8MctIu0m-uM', true);
INSERT INTO public.users (id, username, password, role, token, is_active) VALUES (2, 'user1', '$2a$10$bt7LM.ARJb9S4iFnLdPJE.0AT2vle6Nrv/zUgbVGIb3nRy7JylOca', 'Airline', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTY4OTA4MzE4NSwiaWQiOjJ9.9Wc83df5zazq_OxLC-lfTtT3VXjE6d9M618vvLx3IJY', true);


--
-- Data for Name: ads; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.ads (id, user_id, image, description, subject, price, category_id, status, fly_time, airplane_model, repair_check, expert_check, plane_age) VALUES (1, 2, NULL, 'Test description', 'Ad 1', 50000, 1, 'Active', NULL, NULL, false, false, NULL);
INSERT INTO public.ads (id, user_id, image, description, subject, price, category_id, status, fly_time, airplane_model, repair_check, expert_check, plane_age) VALUES (2, 2, NULL, NULL, 'Ad 2', 200000, 1, 'Active', NULL, NULL, false, false, NULL);


--
-- Data for Name: bookmarks; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: expert_ads; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.expert_ads (id, user_id, expert_id, ads_id, report, status, created_at) VALUES (1, 2, 3, 1, 'Evrything is ok!', 'Done', '2023-07-09 10:04:59.386335');
INSERT INTO public.expert_ads (id, user_id, expert_id, ads_id, report, status, created_at) VALUES (3, 5, 1, 2, 'dfgdfg', 'Done', '2023-07-11 15:13:09.925577');
INSERT INTO public.expert_ads (id, user_id, expert_id, ads_id, report, status, created_at) VALUES (2, 2, 1, 2, 'dfgdfg', 'Done', '2023-07-09 10:34:21.826941');


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
-- Data for Name: user_bookmark; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Name: ads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ads_id_seq', 2, true);


--
-- Name: bookmarks_ads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookmarks_ads_id_seq', 1, false);


--
-- Name: bookmarks_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookmarks_user_id_seq', 1, false);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 1, true);


--
-- Name: expert_ads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.expert_ads_id_seq', 3, true);


--
-- Name: repair_request_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.repair_request_id_seq', 1, false);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 1, false);


--
-- Name: user_bookmark_bookmark_ads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_bookmark_bookmark_ads_id_seq', 1, false);


--
-- Name: user_bookmark_bookmark_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_bookmark_bookmark_user_id_seq', 1, false);


--
-- Name: user_bookmark_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_bookmark_user_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 6, true);


--
-- PostgreSQL database dump complete
--

