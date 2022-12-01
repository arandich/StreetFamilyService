BEGIN;

-- EXTENSIONS --

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --
-- CREATE TABLE public.shop
-- (
--     "id" SERIAL PRIMARY KEY,
--     "shop_name" varchar not null,
--     "city_coordinate" varchar not null,
--     "shop_phone" varchar not null
-- );

CREATE TABLE public.shop_name_category
(
    "id" SERIAL PRIMARY KEY,
    "name" varchar not null
);

CREATE TABLE public.shop_name_catalog
(
    "id" SERIAL PRIMARY KEY,
    "category_id" INT,
    "name" varchar not null,
    "value" int not null,
    "src" varchar not null,
    "weight" int not null,
    "description" varchar not null,
    FOREIGN KEY (category_id) REFERENCES shop_name_category (id)
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
INSERT INTO public.shop_name_category (name) values ('Бергеры');
INSERT INTO public.shop_name_category (name) values ('Пицца');
INSERT INTO public.shop_name_catalog (category_id, name, value, src, weight, description) values ('1','Бычара','125','','200','Жесткий бурегр');
INSERT INTO public.shop_name_catalog (category_id, name, value, src, weight, description) values ('2','Пипирони','299','','400','Пипирони с сиром');
COMMIT;
