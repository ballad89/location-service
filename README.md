# location-service

[![](https://images.microbadger.com/badges/version/ballad/location-service.svg)](https://microbadger.com/images/ballad/location-service "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/image/ballad/location-service.svg)](https://microbadger.com/images/ballad/location-service "Get your own image badge on microbadger.com")

This is a microservice which takes a public ip address and returns information about the country matching the ip

```bash
docker run -p 1989:3000 ballad/location-service
```

```bash
curl http://localhost:1989/location?client_ip=194.182.7.70
```

```json

{"name":{"common":"Denmark","official":"Kingdom of Denmark","Native":{"dan":{"common":"Danmark","official":"Kongeriget Danmark"}}},"EuMember":true,"LandLocked":false,"Nationality":"","tld":[".dk"],"Languages":{"dan":"Danish"},"Translations":{"FIN":{"common":"Tanska","official":"Tanskan kuningaskunta"},"FRA":{"common":"Danemark","official":"Royaume du Danemark"},"HRV":{"common":"Danska","official":"Kraljevina Danska"},"ITA":{"common":"Danimarca","official":"Regno di Danimarca"},"JPN":{"common":"デンマーク","official":"デンマーク王国"},"NLD":{"common":"Denemarken","official":"Koninkrijk Denemarken"},"POR":{"common":"Dinamarca","official":"Reino da Dinamarca"},"RUS":{"common":"Дания","official":"Королевство Дания"},"SPA":{"common":"Dinamarca","official":"Reino de Dinamarca"}},"currency":["DKK"],"Borders":["DEU"],"cca2":"DK","cca3":"DNK","CIOC":"DEN","CCN3":"208","callingCode":["45"],"InternationalPrefix":"00","region":"Europe","subregion":"Northern Europe","Continent":"Europe","capital":"Copenhagen","Area":43094,"longitude":"10 00 E","latitude":"56 00 N","MinLongitude":4.516667,"MinLatitude":53.583332,"MaxLongitude":18,"MaxLatitude":64,"Latitude":56.10176,"Longitude":9.555907}

```
