# Numberplate Database Backend

This project is supposed to be a private Numberplate Database. You can set it up and record numberplates you see in it.

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

### API features

The API is supposed to be written in Rust.

- GET numberplates
- GET numberplates/recent only the most recent 100 (by me)
- GET numberplates/:id
- POST numberplates/new
- PUT numberplates/:id
- DELETE numberplates/:id
- GET numberplates/:id/image (gets the image from the most recent meet)
- GET numberplates/:id/meets
- GET numberplates/bycity/:city-id
- GET numberplates/bycountry/:country-id
- GET numberplates/search/:query

- GET meet/:id
- GET meet/:id/image
- POST meet/:id
- POST meet/:id/image
- GET meet/:id/distancetohome this will return the distance from its home city/country in km

I also want to be able to receive some stats. Don't know how to do this yet.

- GET stats/numberplates
- GET stats/meets
- GET stats/cities/:id
- GET stats/countries/:id

I also want an endpoint that can compress an image. There is probably already a utility that implements this. Just to save some storage from the images.

#### garbage collection

orphaned images are deleted after a certain amount of time.
