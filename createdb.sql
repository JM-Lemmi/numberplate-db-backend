CREATE TABLE "numberplates" (
  "plate" string PRIMARY KEY,
  "country" string,
  "city" string,
  "owner" string,
  "notes" string,
  "meets" array[uuid]
);

CREATE TABLE "meets" (
  "id" uuid PRIMARY KEY,
  "location" geo,
  "time" timestamp,
  "image" bool
);

CREATE TABLE "countries" (
  "id" string PRIMARY KEY,
  "location" geo
);

CREATE TABLE "city" (
  "id" string PRIMARY KEY,
  "country" string,
  "location" geo
);

ALTER TABLE "meets" ADD FOREIGN KEY ("id") REFERENCES "numberplates" ("meets");

ALTER TABLE "numberplates" ADD FOREIGN KEY ("country") REFERENCES "countries" ("id");

ALTER TABLE "countries" ADD FOREIGN KEY ("id") REFERENCES "city" ("country");

ALTER TABLE "numberplates" ADD FOREIGN KEY ("city") REFERENCES "city" ("id");
