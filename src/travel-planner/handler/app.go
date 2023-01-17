package handler

import (
	"fmt"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Sample API handler")

	// user := r.Context().Value("user")
	// claims := user.(*jwt.Token).Claims
	// username := claims.(jwt.MapClaims)["username"]

	// app := model.App{
	// 	Id:          uuid.New(),
	// 	User:        username.(string),
	// 	Title:       r.FormValue("title"),
	// 	Description: r.FormValue("description"),
	// }

	// price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	// fmt.Printf("%v,%T", price, price)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// app.Price = int(price * 100.0)

	// file, _, err := r.FormFile("media_file")
	// if err != nil {
	// 	http.Error(w, "Media file is not available", http.StatusBadRequest)
	// 	fmt.Printf("Media file is not available %v\n", err)
	// 	return
	// }

	// err = service.SaveApp(&app, file)
	// if err != nil {
	// 	http.Error(w, "Failed to save app to backend", http.StatusInternalServerError)
	// 	fmt.Printf("Failed to save app to backend %v\n", err)
	// 	return
	// }

	// fmt.Println("App is saved successfully.")
}
