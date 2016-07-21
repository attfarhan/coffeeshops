package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"

	"github.com/JustinBeckwith/go-yelp/yelp"
	"github.com/attfarhan/go-foursquare"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "p",
			Usage: "flag to sort by popularity",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "flag to sort by distance",
		},
		cli.BoolFlag{
			Name:  "y",
			Usage: "flag to use yelp",
		},
	}
	app.Action = func(c *cli.Context) error {
		foursquareExplore := "https://api.foursquare.com/v2/venues/explore?client_id=ZWDWKYBU4HGW45Z5HJTIJJV3BOMATGVBCBBTKFME3RO3NEN4&client_secret=CM1W21GT2TJXKKWOGFQIH4NQ2FEOWJOZT3K1SHB5CZVRHLYV&ll=37.787830,-122.3992&v=20130815&query=coffee&radius=500"

		var o yelp.AuthOptions
		data, err := ioutil.ReadFile("./config/yelp-config.json")
		json.Unmarshal(data, &o)
		if c.Bool("y") {
			client := yelp.New(&o, nil)
			options := yelp.SearchOptions{
				GeneralOptions: &{
					Term:"Coffee"
				},
				yelp.LocationOptions{
					Location: "121 2nd St, San Francisco, CA",
					CoordinateOptions: &CoordinateOptions{
						Latitude: 37.787830,
						Longitude: -122.3992,
					}
				}
			}
			client.DoSearch(options)
		}
		exploreSourcegraphCoffeeshops := "https://api.foursquare.com/v2/venues/explore?client_id=ZWDWKYBU4HGW45Z5HJTIJJV3BOMATGVBCBBTKFME3RO3NEN4&client_secret=CM1W21GT2TJXKKWOGFQIH4NQ2FEOWJOZT3K1SHB5CZVRHLYV&ll=37.787830,-122.3992&v=20130815&query=coffee&radius=500"
		r, _ := http.Get(exploreSourcegraphCoffeeshops)
		var resp foursquare.ExploreResponse
		defer r.Body.Close()
		dec := json.NewDecoder(r.Body)
		dec.Decode(&resp)
		v := resp.GetVenues()
		if c.Bool("p") {
			sort.Sort(VenuesPop(v))
			for _, ven := range v {
				fmt.Println(ven.Name)
			}
			return nil
		} else if c.Bool("d") {
			sort.Sort(VenuesDist(v))
			for _, ven := range v {
				fmt.Println(ven.Name)
			}
			return nil
		}
		for _, ven := range v {
			fmt.Println(ven.Name)
		}

		return nil
	}
	app.Run(os.Args)
}

type Configuration struct {
	client_id     string
	client_secret string
}
type VenuesPop []foursquare.Venue

func (v VenuesPop) Len() int           { return len(v) }
func (v VenuesPop) Less(i, j int) bool { return v[i].Stats.CheckInsCount < v[j].Stats.CheckInsCount }
func (v VenuesPop) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

type VenuesDist []foursquare.Venue

func (v VenuesDist) Len() int           { return len(v) }
func (v VenuesDist) Less(i, j int) bool { return v[i].Location.Distance > v[j].Location.Distance }
func (v VenuesDist) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
