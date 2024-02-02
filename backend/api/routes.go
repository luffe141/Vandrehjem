package api

import (
	"backend/api/controllers/mongodb"
	"backend/api/handlers"
	"net/http"
)

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	activity := &mongodb.Activity{}
	aboutUs := &mongodb.AboutUs{}
	contact := &mongodb.Contact{}
	events := &mongodb.Events{}
	gallery := &mongodb.Gallery{}
	review := &mongodb.Review{}
	room := &mongodb.Room{}
	slider := &mongodb.Slider{}
	trainBar := &mongodb.TrainBar{}

	mux.HandleFunc("GET /api/activities/", JsonWrapper(handlers.HandleGet(activity)))
	mux.HandleFunc("GET /api/activities/{id}/", JsonWrapper(handlers.HandleGetById(activity)))
	mux.HandleFunc("POST /api/activities/", JsonWrapper(handlers.HandlePost(activity)))
	mux.HandleFunc("PUT /api/activities/{id}/", JsonWrapper(handlers.HandlePut(activity)))
	mux.HandleFunc("DELETE /api/activities/{id}/", JsonWrapper(handlers.HandleDelete(activity)))

	mux.HandleFunc("GET /api/aboutus/", JsonWrapper(handlers.HandleGet(aboutUs)))
	mux.HandleFunc("GET /api/aboutus/{id}/", JsonWrapper(handlers.HandleGetById(aboutUs)))
	mux.HandleFunc("POST /api/aboutus/", JsonWrapper(handlers.HandlePost(aboutUs)))
	mux.HandleFunc("PUT /api/aboutus/{id}/", JsonWrapper(handlers.HandlePut(aboutUs)))
	mux.HandleFunc("DELETE /api/aboutus/{id}/", JsonWrapper(handlers.HandleDelete(aboutUs)))

	mux.HandleFunc("GET /api/contact/", JsonWrapper(handlers.HandleGet(contact)))
	mux.HandleFunc("GET /api/contact/{id}/", JsonWrapper(handlers.HandleGetById(contact)))
	mux.HandleFunc("POST /api/contact/", JsonWrapper(handlers.HandlePost(contact)))
	mux.HandleFunc("PUT /api/contact/{id}/", JsonWrapper(handlers.HandlePut(contact)))
	mux.HandleFunc("DELETE /api/contact/{id}/", JsonWrapper(handlers.HandleDelete(contact)))

	mux.HandleFunc("GET /api/events/", JsonWrapper(handlers.HandleGet(events)))
	mux.HandleFunc("GET /api/events/{id}/", JsonWrapper(handlers.HandleGetById(events)))
	mux.HandleFunc("POST /api/events/", JsonWrapper(handlers.HandlePost(events)))
	mux.HandleFunc("PUT /api/events/{id}/", JsonWrapper(handlers.HandlePut(events)))
	mux.HandleFunc("DELETE /api/events/{id}/", JsonWrapper(handlers.HandleDelete(events)))

	mux.HandleFunc("GET /api/review/", JsonWrapper(handlers.HandleGet(review)))
	mux.HandleFunc("GET /api/review/{id}/", JsonWrapper(handlers.HandleGetById(review)))
	mux.HandleFunc("POST /api/review/", JsonWrapper(handlers.HandlePost(review)))
	mux.HandleFunc("PUT /api/review/{id}/", JsonWrapper(handlers.HandlePut(review)))
	mux.HandleFunc("DELETE /api/review/{id}/", JsonWrapper(handlers.HandleDelete(review)))

	mux.HandleFunc("GET /api/room/", JsonWrapper(handlers.HandleGet(room)))
	mux.HandleFunc("GET /api/room/{id}/", JsonWrapper(handlers.HandleGetById(room)))
	mux.HandleFunc("POST /api/room/", JsonWrapper(handlers.HandlePost(room)))
	mux.HandleFunc("PUT /api/room/{id}/", JsonWrapper(handlers.HandlePut(room)))
	mux.HandleFunc("DELETE /api/room/{id}/", JsonWrapper(handlers.HandleDelete(room)))

	mux.HandleFunc("GET /api/gallery/", JsonWrapper(handlers.HandleGet(gallery)))
	mux.HandleFunc("GET /api/gallery/{id}/", JsonWrapper(handlers.HandleGetById(gallery)))
	mux.HandleFunc("POST /api/gallery/", JsonWrapper(handlers.HandlePost(gallery)))
	mux.HandleFunc("PUT /api/gallery/{id}/", JsonWrapper(handlers.HandlePut(gallery)))
	mux.HandleFunc("DELETE /api/gallery/{id}/", JsonWrapper(handlers.HandleDelete(gallery)))

	//
	mux.HandleFunc("GET /api/slider/", JsonWrapper(handlers.HandleGet(slider)))
	mux.HandleFunc("GET /api/slider/{id}/", JsonWrapper(handlers.HandleGetById(slider)))
	mux.HandleFunc("POST /api/slider/", JsonWrapper(handlers.HandlePost(slider)))
	mux.HandleFunc("PUT /api/slider/{id}/", JsonWrapper(handlers.HandlePut(slider)))
	mux.HandleFunc("DELETE /api/slider/{id}/", JsonWrapper(handlers.HandleDelete(slider)))

	mux.HandleFunc("GET /api/trainbar/", JsonWrapper(handlers.HandleGet(trainBar)))
	mux.HandleFunc("GET /api/trainbar/{id}/", JsonWrapper(handlers.HandleGetById(trainBar)))
	mux.HandleFunc("POST /api/trainbar/", JsonWrapper(handlers.HandlePost(trainBar)))
	mux.HandleFunc("PUT /api/trainbar/{id}/", JsonWrapper(handlers.HandlePut(trainBar)))
	mux.HandleFunc("DELETE /api/trainbar/{id}/", JsonWrapper(handlers.HandleDelete(trainBar)))

	return mux
}
