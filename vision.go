package main

import (
	"errors"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/vision"
	"github.com/progrium/darwinkit/objc"
	"unsafe"
)

func process(image []byte) {
	width := 1920
	height := 1080
	var err error
	objc.WithAutoreleasePool(func() {
		handler := vision.NewImageRequestHandler().InitWithDataOptions(image, nil)

		req := vision.NewRecognizeTextRequest().InitWithCompletionHandler(func(request vision.Request, error foundation.Error) {
			if !error.IsNil() {
				err = errors.New(error.Description())
				return
			}

			// Get results and ensure they are handled correctly
			results := objc.Call[[]vision.RecognizedTextObservation](request, objc.Sel("results"))

			// Iterate over results safely
			for _, result := range results {
				boundingBox := result.BoundingBox()

				// Convert Vision's bottom-left origin to top-left origin
				x0 := int(boundingBox.Origin.X * float64(width))
				y0 := int((1.0 - (boundingBox.Origin.Y + boundingBox.Size.Height)) * float64(height)) // Flip Y-axis
				x1 := int((boundingBox.Origin.X + boundingBox.Size.Width) * float64(width))
				y1 := int((1.0 - boundingBox.Origin.Y) * float64(height)) // Flip Y-axis

				// Get the recognized text
				for _, candidate := range result.TopCandidates(1) {
					_, _, _, _, _ = x0, y0, x1, y1, candidate.String()
				}
			}
		})

		req.SetMinimumTextHeight(1 / 100)
		req.SetRecognitionLevel(vision.RequestTextRecognitionLevelAccurate)
		req.SetRecognitionLanguages([]string{"en"})
		req.SetUsesLanguageCorrection(true)
		req.SetCustomWords([]string{"copy", "print", "waiting", "black", "white", "yellow", "magenta", "cyan", "light"})

		var errObj foundation.Error
		handler.PerformRequestsError([]vision.IRequest{req}, unsafe.Pointer(&errObj))
		if !errObj.IsNil() {
			err = errors.New(errObj.Description())
			return
		}
	})
	if err != nil {
		panic(err)
	}
}
