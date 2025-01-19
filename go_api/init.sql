-- public.tracking definition

-- Drop table
-- DROP TABLE public.tracking;

CREATE TABLE public.tracking (
    tracking_id varchar NOT NULL,
    "timestamp" time NOT NULL,
    input_path varchar NOT NULL,
    output_path varchar NOT NULL,
    status varchar NOT NULL,
    CONSTRAINT tracking_pk PRIMARY KEY (tracking_id)
);

-- Permissions
ALTER TABLE public.tracking OWNER TO "user";
GRANT ALL ON TABLE public.tracking TO "user";
