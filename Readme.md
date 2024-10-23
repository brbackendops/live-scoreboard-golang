# Real-Time ScoreBoard Api (fasthttp) ddd pattern / Redis / Websocket / React /

* [X]  dev
  * [X]  database (postgresql-sqlx)

    * [X]  connection pooling
  * [X]  fasthttp
  * [X]  redis
  * [X]  websocket
  * [X]  repo
  * [X]  controller
  * [X]  middlewares

    * [X]  cors
    * [X]  rate limit
* [X]  react

## Production

* can use `kafka` for real time updates  instead of websockets
* cache using redis
* websocket scalling with redis or using `centrifugo`
* background tasks for time consuming tasks
* monitoring
