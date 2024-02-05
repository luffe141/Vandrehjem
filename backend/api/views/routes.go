package views

import (
	"backend/api/models"
	"backend/bekrouter"
	"net/http"
)

func AddRoutes() *http.ServeMux {
	router := bekrouter.New()

	// Restful routes
	router.AddRestfulJSONRoute("/api/activities/", models.ActivityModel)
	router.AddRestfulJSONRoute("/api/about-us/", models.AboutUsModel)
	router.AddRestfulJSONRoute("/api/contact/", models.ContactModel)
	router.AddRestfulJSONRoute("/api/events/", models.EventsModel)
	router.AddRestfulJSONRoute("/api/gallery/", models.GalleryModel)
	router.AddRestfulJSONRoute("/api/review/", models.ReviewModel)
	router.AddRestfulJSONRoute("/api/room/", models.RoomModel)
	router.AddRestfulJSONRoute("/api/slider/", models.SliderModel)
	router.AddRestfulJSONRoute("/api/train-bar/", models.TrainBarModel)

	return router.Mux
}
