-- Create a tables
CREATE TABLE IF NOT EXISTS public.users
(
    user_id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    renting boolean NOT NULL,
    bike_id text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

-- Insert initial data
-- INSERT INTO public.users (user_id, name, password, renting, bike_id) VALUES ('1', 'TestUser1', 'password1', false, '0'), ('2', 'TestUser2', 'password1', false, '0'), ('3', 'TestUser3', 'password1', false, '0'); 

CREATE TABLE IF NOT EXISTS public.bike_renting_system
(
    bike_id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default",
    latitude numeric,
    longtitude numeric,
    rented boolean,
    user_name text COLLATE pg_catalog."default",
    CONSTRAINT bike_renting_system_pkey PRIMARY KEY (bike_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.bike_renting_system
    OWNER to postgres;

-- Insert initial data
INSERT INTO public.bike_renting_system (bike_id, name, latitude, longtitude, rented, user_name) VALUES ('1', 'Skippy', 50, 20, false, '0'), ('2', 'Blinky', 51, 17, false, '0'), ('3', 'Delorean', 60, -20, false, '0'); 
