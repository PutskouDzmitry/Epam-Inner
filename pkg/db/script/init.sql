CREATE TABLE IF NOT EXISTS public.visits_scud
(
    id integer NOT NULL,
    _id character varying COLLATE pg_catalog."default" NOT NULL,
    "employeeId" character varying COLLATE pg_catalog."default",
    "timePACS" timestamp without time zone,
    "timeUTC" timestamp with time zone,
    access character varying COLLATE pg_catalog."default",
    "codeDevice" character varying COLLATE pg_catalog."default",
    device character varying COLLATE pg_catalog."default",
    card character varying COLLATE pg_catalog."default",
    personal character varying COLLATE pg_catalog."default",
    "visitorType" integer,
    "position" character varying COLLATE pg_catalog."default",
    "spaceId" character varying COLLATE pg_catalog."default",
    "eventType" character varying COLLATE pg_catalog."default",
    "osmUnitId" character varying COLLATE pg_catalog."default",
    "timeDeviceLocal" timestamp without time zone,
    CONSTRAINT "PK_f36d36d3db624d1569617a67178" PRIMARY KEY (id),
    CONSTRAINT "UQ_4f62e10c4769366f4b6a1136af9" UNIQUE (_id, device)
);





CREATE INDEX "scud_visits_employeeId_index"
    ON public.visits_scud USING btree
        ("employeeId" COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: scud_visits_genID_index

-- DROP INDEX public."scud_visits_genID_index";

CREATE INDEX "scud_visits_genID_index"
    ON public.visits_scud USING btree
        (_id COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: scud_visits_personal_index

-- DROP INDEX public.scud_visits_personal_index;

CREATE INDEX scud_visits_personal_index
    ON public.visits_scud USING btree
        (personal COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: scud_visits_spaceId_index

-- DROP INDEX public."scud_visits_spaceId_index";

CREATE INDEX "scud_visits_spaceId_index"
    ON public.visits_scud USING btree
        ("spaceId" COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: scud_visits_timeDeviceLocal_index

-- DROP INDEX public."scud_visits_timeDeviceLocal_index";

CREATE INDEX "scud_visits_timeDeviceLocal_index"
    ON public.visits_scud USING btree
        ("timeDeviceLocal" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: scud_visits_timeLocale_index

-- DROP INDEX public."scud_visits_timeLocale_index";

CREATE INDEX "scud_visits_timeLocale_index"
    ON public.visits_scud USING btree
        ("timeUTC" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: scud_visits_timePACS_index

-- DROP INDEX public."scud_visits_timePACS_index";

CREATE INDEX "scud_visits_timePACS_index"
    ON public.visits_scud USING btree
        ("timePACS" ASC NULLS LAST)
    TABLESPACE pg_default;