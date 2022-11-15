BEGIN;

-- EXTENSIONS --

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --
CREATE TABLE public.shop
(
    "id" SERIAL PRIMARY KEY,
    "shop_name" varchar not null,
    "city_coordinate" varchar not null,
    "shop_phone" varchar not null
);

CREATE TABLE public.catalog_shop_name
(
    "id" SERIAL PRIMARY KEY,
    "name" varchar not null,
    "value" float not null,
    "src" varchar not null,
    "weight" float not null,
    "description" varchar not null
);


-- CREATE TABLE public.users
-- (
--         "id" SERIAL PRIMARY KEY,
--         "first_name" varchar not null,
--         "last_name" varchar not null,
--         "phone_number" varchar not null,
--         "email" varchar not null,
--         "username" varchar not null,
--         "password" varchar not null,
--         "user_passive_id" INTEGER REFERENCES public.user_passive (id),
--         "city_id" INTEGER REFERENCES public.city (id),
--         "exp_value" INTEGER not null
-- );


-- INSERT INTO public.city (city_name, city_coordinate) values ('СПБ','23');
-- INSERT INTO public.city (city_name, city_coordinate) values ('МСК','53');
-- INSERT INTO public.user (email,username,password,city_id) values ('emaiasd@email.ru','useasdser','dsadasd',1);
INSERT INTO public.shop (shop_name, city_coordinate, shop_phone) values ('Bylka','24241','81239213');
INSERT INTO public.shop (shop_name, city_coordinate, shop_phone) values ('bylyka','24232141','812319213');
COMMIT;
