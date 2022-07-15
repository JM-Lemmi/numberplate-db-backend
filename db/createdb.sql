CREATE TABLE "meets" (
  "id" UUID,
  "location" POINT,
  "time" TimestampTZ,
  "image" BOOL,
  PRIMARY KEY ("id"),
  UNIQUE ("id")
);

CREATE TABLE "countries" (
  "id" VARCHAR(3),
  "location" POLYGON,
  "font" INT,
  "flag" TEXT,
  PRIMARY KEY ("id"),
  UNIQUE ("id")
);

CREATE TABLE "cities" (
  "name" TEXT,
  "country" TEXT REFERENCES "countries"("id"),
  "location" POLYGON,
  "coatofarms" TEXT,
  PRIMARY KEY ("name")
);

CREATE TABLE "numberplates" (
  "plate" VARCHAR(8),
  "country" TEXT REFERENCES "countries"("id"),
  "city" TEXT REFERENCES "cities"("name"),
  "owner" TEXT,
  "notes" TEXT,
  "meets" UUID REFERENCES "meets"("id"),
  PRIMARY KEY ("plate"),
  UNIQUE ("plate")
);

ALTER TABLE "numberplates" ADD FOREIGN KEY ("meets") REFERENCES "meets" ("id");

ALTER TABLE "numberplates" ADD FOREIGN KEY ("country") REFERENCES "countries" ("id");

ALTER TABLE "cities" ADD FOREIGN KEY ("country") REFERENCES "countries" ("id");

ALTER TABLE "numberplates" ADD FOREIGN KEY ("city") REFERENCES "cities" ("name");

-- deutsche unterscheidungszeichen

--CREATE TABLE de_distinctions (
--  "kfz" string PRIMARY KEY,
--  "name" string,
--  "plz" int
--);
--COPY de_distinctions FROM '/tmp/kfz250.gk3.csv/kfz250/KFZ250.csv' WITH (FORMAT csv);
