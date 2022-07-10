CREATE TABLE "numberplates" (
  "plate" string PRIMARY KEY,
  "country" string,
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

ALTER TABLE "meets" ADD FOREIGN KEY ("id") REFERENCES "numberplates" ("meets");
