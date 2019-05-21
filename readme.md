## QR CODE GENERATOR FOR IMAGES RETURNING BLOBS

#### Installation
1. `Git Clone`
2. in project folders use `go mod tidy`
3. `go run main.go`

#### Usage
1. Method `GET`
2. Parameters `text`
3. Return `image_blob`
4. > HTML IMAGES <img src="data:image/jpg;base64,/9j/{{image_blob}}">