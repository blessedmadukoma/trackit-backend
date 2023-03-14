BEGIN;


CREATE TABLE IF NOT EXISTS public."user"
(
    firstname "char" NOT NULL,
    lastname "char" NOT NULL,
    email "char" NOT NULL GENERATED ALWAYS AS IDENTITY,
    password "char" NOT NULL,
    subdomain "char",
    userid bigserial NOT NULL,
    date_created date NOT NULL,
    user_type boolean NOT NULL,
    PRIMARY KEY (userid)
);

CREATE TABLE IF NOT EXISTS public.userprofile
(
    profileid bigserial NOT NULL,
    userid "char" NOT NULL,
    picture "char",
    phone "char",
    title "char",
    bio "char"[],
    city "char",
    state "char",
    coutry "char",
    isactive bigint NOT NULL,
    date_updated date NOT NULL,
    PRIMARY KEY (profileid)
);

CREATE TABLE IF NOT EXISTS public.classroom
(
    classroomid bigserial NOT NULL,
    userid "char" NOT NULL,
    name "char" NOT NULL,
    about "char"[] NOT NULL,
    date_created date NOT NULL,
    PRIMARY KEY (classroomid)
);

CREATE TABLE IF NOT EXISTS public.resource
(
    resourceid bigserial NOT NULL,
    userid "char" NOT NULL,
    filename "char" NOT NULL,
    fileref "char" NOT NULL,
    date_created date NOT NULL,
    PRIMARY KEY (resourceid)
);

CREATE TABLE IF NOT EXISTS public.scheduler
(
    schedulerid bigserial NOT NULL,
    subject "char" NOT NULL,
    starttime date NOT NULL,
    endtime date NOT NULL,
    location "char" NOT NULL,
    description "char"[] NOT NULL,
    userid "char" NOT NULL,
    status "char" NOT NULL,
    PRIMARY KEY (schedulerid)
);

ALTER TABLE IF EXISTS public."user"
    ADD FOREIGN KEY (userid)
    REFERENCES public.userprofile (userid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public."user"
    ADD FOREIGN KEY (userid)
    REFERENCES public.classroom (userid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public."user"
    ADD FOREIGN KEY (userid)
    REFERENCES public.resource (userid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public."user"
    ADD FOREIGN KEY (userid)
    REFERENCES public.scheduler (userid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

END;