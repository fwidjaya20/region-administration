CREATE TABLE IF NOT EXISTS public."provinces"
(
    "id" INT8 NOT NULL GENERATED ALWAYS AS IDENTITY,
    "code" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" TEXT NOT NULL,
    "updated_at" TIMESTAMPTZ,
    "updated_by" TEXT
);

CREATE TABLE IF NOT EXISTS public."regencies"
(
    "id" INT8 NOT NULL GENERATED ALWAYS AS IDENTITY,
    "code" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" TEXT NOT NULL,
    "updated_at" TIMESTAMPTZ,
    "updated_by" TEXT
);

CREATE TABLE IF NOT EXISTS public."districts"
(
    "id" INT8 NOT NULL GENERATED ALWAYS AS IDENTITY,
    "code" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" TEXT NOT NULL,
    "updated_at" TIMESTAMPTZ,
    "updated_by" TEXT
);

CREATE TABLE IF NOT EXISTS public."villages"
(
    "id" INT8 NOT NULL GENERATED ALWAYS AS IDENTITY,
    "code" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "postal" CHAR(5) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" TEXT NOT NULL,
    "updated_at" TIMESTAMPTZ,
    "updated_by" TEXT
);
