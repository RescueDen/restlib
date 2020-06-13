package images

import "net/http"

//go:generate mockgen -destination=../mocks/mock_image_host.go -package=mocks github.com/reaction-eng/restlib/images Host

type Host interface {
	Location() http.Dir
	PrepareImage(imageId string) error
}
