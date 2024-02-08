package views

import (
	"backend/api/mvc/models"
	"github.com/lmbek/bekrouter"
	"net/http"
)

func AddRoutes() *http.ServeMux {
	router := bekrouter.New()

	// Restful routes
	router.AddRestfulJSONRoute("/api/activities/", models.ActivityModel)
	router.AddRestfulJSONRoute("/api/about-us/", models.AboutUsModel)
	router.AddRestfulJSONRoute("/api/contacts/", models.ContactModel)
	router.AddRestfulJSONRoute("/api/events/", models.EventModel)
	router.AddRestfulJSONRoute("/api/galleries/", models.GalleryModel)
	router.AddRestfulJSONRoute("/api/reviews/", models.ReviewModel)
	router.AddRestfulJSONRoute("/api/rooms/", models.RoomModel)
	router.AddRestfulJSONRoute("/api/sliders/", models.SliderModel)
	router.AddRestfulJSONRoute("/api/train-bars/", models.TrainBarModel)

	return router.Mux
}
