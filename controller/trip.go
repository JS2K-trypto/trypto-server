package controller

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"trypto-server/model"

	_ "trypto-server/docs"

	"github.com/gin-gonic/gin"
)

var (
	tripPlan    model.TripPlan
	searchQuery string
)

// CreateTripPlan godoc

// @BasePath				/v01
// @Summary					Enter the wallet account, title, country, departure date, arrival date, etc. days is an array containing the items.
// @Tags					CreateTripPlan(Create my itinerary)
// @Description				DAYS has items in DAY1 and DAY2, and each DAY1 has time and imtes, and you can enter start time, end time, image, title, description, and notes. After inputting, a travel plan is created.
// @Accept					json
// @Produce					json
// @Param					walletAccount			path	string 		true	"walletAccount",
// @Param					tripTitle				path	string	 	true	"tripTitle",
// @Param					tripCountry				path	string		true	"tripCountry",
// @Param					tripDeparture			path	string		true	"tripDeparture",
// @Param					tripArrival				path	string		true	"tripArrival"
// @Param					dayItems					path	string 		true	"dayItems"
// @Router					/v01/trip/myplan	[post]
// @Success					200	{array} model.TripPlan
func (p *Controller) CreateTripPlan(c *gin.Context) {
	if err := c.ShouldBindJSON(&tripPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	empty := []string{" "}
	fmt.Println("tripPlan", tripPlan)
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")

	tripPlan.TripDeparture = custom
	tripPlan.TripArrival = custom
	fmt.Println("before res")
	res := p.md.InsertTripPlan(&tripPlan)
	fmt.Println("after res")
	if res != nil {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Trip plan creation failed.": empty})
	}
}

// GetMyTrip godoc
//
//	@BasePath		/v01
//	@Summary		Enter your wallet account to import your trip plans
//	@Tags			GetMyTrip(Importing my trip plans)
//	@Description	Import the itinerary you created from MongoDB.
//	@name			GetMyTrip
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string 	 	true	"walletAccount"
//	@Router			/v01/trip/myplan [get]
//	@Success		200	{array} model.TripPlan
func (p *Controller) GetMyTrip(c *gin.Context) {
	tripPlan.WalletAccount = c.Query("walletAccount")
	empty := []string{" "}
	fmt.Println("tripPlan", tripPlan)
	res := p.md.SelectMyTrip(tripPlan.WalletAccount)
	// fmt.Println("len, len(res.Arr)", len(res.Arr))
	// c.JSON(http.StatusOK, res)

	if len(res) > 0 {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"The result is empty. ": empty})
	}
}

// GetAllTrip godoc
//
//	@BasePath		/v01
//	@Summary		Import all trip plans.
//	@Tags			GetAllTrip(Import all  trip plans.)
//	@Description    Function to fetch all trip plans from MongoDB. No parameters. Retrieves all.
//	@name			GetAllTrip
//	@Accept			json
//	@Produce		json
//	@Router			/v01/trip/allplan [get]
//	@Success		200	{array} model.TripPlan
func (p *Controller) GetAllTrip(c *gin.Context) {
	empty := []string{" "}
	res := p.md.SelectAllTrip()
	fmt.Println(len(res))
	if len(res) > 0 {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"The result is empty. ": empty})
	}
}

// SearchTrip godoc
//
//	@BasePath		/v01
//	@Summary		In Q, type the keyword you want to search for.
//	@Tags			SearchTrip(Search your travel plans by keyword)
//	@Description	A search API that returns content for matching strings in the title of a travel plan, implemented on a word-by-word basis, e.g. q="South Korea".
//	@name			SearchTrip
//	@Accept			json
//	@Produce		json
//	@Param			q	path	string 	 	true	"q"
//	@Router			/v01/trip/search [get]
//	@Success		200	{array} model.TripPlan
func (p *Controller) SearchTrip(c *gin.Context) {
	searchQuery = c.Query("q")
	empty := []string{" "}
	res := p.md.SearchTrip(searchQuery)
	fmt.Println("res", res)
	if len(res) > 0 {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"The result is empty.": empty})
	}
}

// CreateSimpleTripPlan godoc

// @BasePath				/v01
// @Summary					Enter the wallet account, title, country, departure date, arrival date, etc. days is an array containing the items. Here it is entered as an empty array.
// @Tags					CreateSimpleTripPlan(Create my simple trip plan)
// @Description				DAYS has items in DAY1 and DAY2, and each DAY1 has time and imtes, and you can enter start time, end time, image, title, description, and notes.
// @Accept					json
// @Produce					json
// @Param					walletAccount			path	string 		true	"walletAccount",
// @Param					triplTitle				path	string	 	true	"triplTitle",
// @Param					tripCountry				path	string		true	"tripCountry",
// @Param					tripDeparture			path	string		true	"tripDeparture",
// @Param					tripArrival				path	string		true	"tripArrival"
// @Param					dayItems				path	string 		true	"dayItems"
// @Router					/v01/trip/simpleplan	[post]
// @Success					200	{array} model.TripPlan
func (p *Controller) CreateSimpleTripPlan(c *gin.Context) {
	if err := c.ShouldBindJSON(&tripPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("tripPlan", tripPlan)
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")

	tripPlan.TripDeparture = custom
	tripPlan.TripArrival = custom
	fmt.Println("before res")
	res := p.md.InsertTripPlan(&tripPlan)
	fmt.Println("after res")
	if res != nil {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Trip plan creation failed."})
	}
}

// PatchSimpleTripPlan godoc

// @BasePath				/v01
// @Summary					Enter the wallet account, title, country, departure date, arrival date, etc. Add the data after the function that takes dayItems as input.
// @Tags					PatchSimpleTripPlan(Modify my simple trip plan with the patch)
// @Description				This API is used to make modifications such as patches based on existing trip IDs.
// @Accept					json
// @Produce					json
// @Param					tripId			path	string 		true	"tripId"
// @Param					dayItems					path	string 		true	"dayItems"
// @Router					/v01/trip/simpleplan	[patch]
// @Success					200	{array} model.TripPlan
func (p *Controller) PatchSimpleTripPlan(c *gin.Context) {
	if err := c.ShouldBindJSON(&tripPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("tripPlan", tripPlan)
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")
	empty := []string{" "}

	tripPlan.TripDeparture = custom
	tripPlan.TripArrival = custom
	fmt.Println("before res")
	res := p.md.PatchTripPlan(&tripPlan)
	fmt.Println("after res")
	if res != nil {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"The result is empty.": empty})
	}
}

// GetDetailTrip godoc

// @BasePath				/v01
// @Summary					Entering a tripid after detail will output the corresponding trip plan. ex) /detail/tripId
// @Tags					GetDetailTrip(Prints the details page of my trip plan)
// @Description				If you enter a trip id after the detail path, it will display the detail page of my trip plan.
// @Accept					json
// @Produce					json
// @Param					tripId			path	string 		true	"tripId"
// @Router					/v01/trip/detail/:num	[get]
// @Success					200	{array} model.TripPlan
func (p *Controller) GetDetailTrip(c *gin.Context) {
	num := c.Param("num")
	tripPlan.TripId, _ = strconv.ParseInt(num, 10, 64)

	empty := []string{" "}
	fmt.Println("num", num)
	res := p.md.SelectDetailTrip(tripPlan.TripId)
	// fmt.Println("len, len(res.Arr)", len(res.Arr))
	// c.JSON(http.StatusOK, res)

	if len(res) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"The result is empty.": empty})
	} else {
		c.JSON(http.StatusOK, res[0])
	}
}

// DeleteTrip godoc

// @BasePath				/v01
// @Summary					Entering a tripid after delete will delete posts for that tripid.
// @Tags					DeleteTrip(Delete my trip)
// @Description				Enter the /delete path followed by the trip id and the corresponding trip will be deleted.
// @Accept					json
// @Produce					json
// @Param					tripId			path	string 		true	"tripId"
// @Router					/v01/trip/delete/:num	[delete]
// @Success					200	{array} model.TripPlan
func (p *Controller) DeleteTrip(c *gin.Context) {
	num := c.Param("num")
	tripPlan.TripId, _ = strconv.ParseInt(num, 10, 64)

	empty := []string{" "}
	fmt.Println("num", num)
	res := p.md.DeleteTrip(tripPlan.TripId)
	fmt.Println(reflect.TypeOf(res))
	if res.DeletedCount != 0 {
		c.JSON(http.StatusOK, gin.H{"The delete was successful.": res})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"The delete failed. ": empty})
	}
}
