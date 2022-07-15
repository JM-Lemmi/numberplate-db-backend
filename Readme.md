# Numberplate Database Backend

This project is supposed to be a private Numberplate Database. You can set it up and record numberplates you see in it.

This project is centered on Germany, so many specialities are specifically for German numberplates.

## Idea

### Database Structure

![dbschema](/.readme/dbschema.png)

See also `createdb.sql`

#### calculated values

some values should be precalculated like [this](https://stackoverflow.com/q/58772806/9397749)

```
product BIGINT GENERATED ALWAYS AS (int1 * int2) STORED
```

but you can't calculate cross tables. so i'll have to see.

#### data collection

Das Bundesamt für Kartographie und Geodäsie bietet [hier](https://gdz.bkg.bund.de/index.php/default/kfz-kennzeichen-1-250-000-kfz250.html) einen Datensatz mit KFZ-Kennzeichen zum Download an.

https://daten.gdz.bkg.bund.de/produkte/sonstige/kfz250/aktuell/kfz250.gk3.csv.zip

### API features

The API is supposed to be written in Go.

- [x] GET numberplates
- PUT numberplates
- GET numberplates/recent only the most recent 100 (by me)
- GET numberplates/:id
- DELETE numberplates/:id
- GET numberplates/:id/image (gets the image from the most recent meet)
- GET numberplates/:id/meets
- GET numberplates/bycity/:city-id
- GET numberplates/bycountry/:country-id
- GET numberplates/search/:query

- GET meet/:uuid
- GET meet/:uuid/image will get the image directly. alias to GET image/:uuid
- POST meet/:uuid
- POST meet/:uuid/image will upload with the correct uuid. Alias to POST image/:uuid
- GET meet/:uuid/distancetohome this will return the distance from its home city/country in km

- GET image/:uuid
- POST image/:uuid

I also want to be able to receive some stats. Don't know how to do this yet.

- GET stats/numberplates
- GET stats/meets
- GET stats/cities/:id
- GET stats/countries/:id

- POST compress I also want an endpoint that can compress an image. There is probably already a utility that implements this. Just to save some storage from the images.

- GET country/:id/font gets the ttf file for the font of the country. Some will be shared, but not all
- GET font/:id

#### garbage collection

orphaned images are deleted after a certain amount of time.

## Unsorted Thoughts

Alles vor den Siegeln wird als "city" gehandhabt. Dh BWL4 ist eine Stadt, BWL1 ist eine andere Stadt.

## Dev

```
sudo apt-get install -y postgresql-client
psql --host=localhost --username=postgres --password
```

```
curl -d '{"plate":"HGJL1999","country":"DEU","owner":"Julian Lemmerich"}' http://localhost/numberplates/new
```
