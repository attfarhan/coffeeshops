package main

import "github.com/attfarhan/go-foursquare"

// Use this URL to find coffeeshops within 500m of Sourcegraph
exploreSourcegraphCoffeeshops := "https://api.foursquare.com/v2/venues/explore?client_id=ZWDWKYBU4HGW45Z5HJTIJJV3BOMATGVBCBBTKFME3RO3NEN4&client_secret=CM1W21GT2TJXKKWOGFQIH4NQ2FEOWJOZT3K1SHB5CZVRHLYV&ll=37.787830,-122.3992&v=20130815&query=coffee&radius=500"

// Put your response into this struct
var resp foursquare.ExploreResponse
// Get a list of venues using this method
resp.GetVenues()
