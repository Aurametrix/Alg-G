g := globe.New()
g.DrawGraticule(10.0)
g.SavePNG("graticule.png", 400)

g := globe.New()
g.DrawGraticule(10.0)
g.DrawLandBoundaries()
g.CenterOn(51.453349, -2.588323)
g.SavePNG("land.png", 400)

shops, err := LoadCoffeeShops("./starbucks.json")
if err != nil {
	log.Fatal(err)
}

green := color.NRGBA{0x00, 0x64, 0x3c, 192}
g := globe.New()
g.DrawGraticule(10.0)
for _, s := range shops {
	g.DrawDot(s.Lat, s.Lng, 0.05, globe.Color(green))
}
g.CenterOn(40.645423, -73.903879)
err = g.SavePNG("starbucks.png", 400)
if err != nil {
	log.Fatal(err)
}
