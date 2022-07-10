CREATE TABLE "numberplates" (
  "plate" string PRIMARY KEY,
  "country" string,
  "city" string,
  "owner" string,
  "notes" string,
  "meets" array[uuid],
  "special" id
);

CREATE TABLE "meets" (
  "id" uuid PRIMARY KEY,
  "location" geo,
  "time" timestamp,
  "image" bool
);

CREATE TABLE "countries" (
  "id" string PRIMARY KEY,
  "location" geo,
  "font" int,
  "flag" url
);

CREATE TABLE "city" (
  "id" string PRIMARY KEY,
  "country" string,
  "location" geo,
  "coatofarms" url
);

ALTER TABLE "meets" ADD FOREIGN KEY ("id") REFERENCES "numberplates" ("meets");

ALTER TABLE "numberplates" ADD FOREIGN KEY ("country") REFERENCES "countries" ("id");

ALTER TABLE "countries" ADD FOREIGN KEY ("id") REFERENCES "city" ("country");

ALTER TABLE "numberplates" ADD FOREIGN KEY ("city") REFERENCES "city" ("id");
