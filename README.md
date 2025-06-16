Route Planning:
-Find shortest path between two stations (Dijkstra's algorithm) //check
-Multi-modal route suggestions (metro + bus combinations)  //check
    example: /api/v1/routes?from=lat,lng&to=lat,lng&modes=metro,bus,walk

Schedule Features:
-Next departures from a specific station //ui
-First/last train indicators //ui

Geospatial Features:
-Nearby stations by geolocation (radius search) // check
-Station-to-station distance calculator //check
-Walking directions to nearest station entrance //check

User Experience Enhancements:
-Estimated arrival times // kinda
-Save frequent routes // ui
-Recent searches history // ui
-Preferred transport modes // ui