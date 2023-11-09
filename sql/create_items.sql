CREATE TABLE IF NOT EXISTS public.items
(
    id integer NOT NULL,
    customer_name character varying COLLATE pg_catalog."default" NOT NULL,
    order_date date NOT NULL,
    product character varying COLLATE pg_catalog."default" NOT NULL,
    quantity integer NOT NULL,
    price numeric NOT NULL,
    CONSTRAINT items_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.items
    OWNER to postgres;