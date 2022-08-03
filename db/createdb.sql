CREATE TABLE "countries" (
  "id" TEXT,
  "name" TEXT,
  "location" POLYGON,
  "font" INT,
  "flag" TEXT,
  PRIMARY KEY ("id"),
  UNIQUE ("id")
);

CREATE TABLE "kfz_de" (
  "kfz" VARCHAR(4),
  "ableitung" TEXT,
  "landkreis" TEXT,
  "country" TEXT REFERENCES "countries"("id"),
  PRIMARY KEY ("kfz"),
  UNIQUE ("kfz")
);

CREATE TABLE "numberplates" (
  "plate" VARCHAR(9),
  "country" TEXT REFERENCES "countries"("id"),
  "owner" TEXT,
  "notes" TEXT,
  PRIMARY KEY ("plate"),
  UNIQUE ("plate")
);

CREATE TABLE "meets" (
  "id" UUID,
  "plate" VARCHAR(9) REFERENCES "numberplates"("plate"),
  "lat" FLOAT,
  "lon" FLOAT,
  "time" TimestampTZ,
  "image" BOOL,
  PRIMARY KEY ("id"),
  UNIQUE ("id")
);

COPY countries ("id", "name") FROM '/tmp/europa.csv' DELIMITER ';' CSV;
