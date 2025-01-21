package main

import (
	"context"
	"errors"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/vision"
	"github.com/progrium/darwinkit/objc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"unsafe"
)

type impl struct {
	UnimplementedExampleServer
}

func main() {
	srv := grpc.NewServer()
	var i impl
	RegisterExampleServer(srv, &i)
	l, err := net.Listen("tcp", ":10000")
	if err != nil {
		panic(err)
	}
	err = srv.Serve(l)
	if err != nil {
		panic(err)
	}
}

func (i *impl) Process(_ context.Context, req *ProcessRequest) (*ProcessResponse, error) {
	width := 1920
	height := 1080
	var err error
	objc.WithAutoreleasePool(func() {
		handler := vision.NewImageRequestHandler().InitWithDataOptions(req.Image, nil)

		req := vision.NewRecognizeTextRequest().InitWithCompletionHandler(func(request vision.Request, error foundation.Error) {
			if !error.IsNil() {
				err = errors.New(error.Description())
				return
			}

			// Get results and ensure they are handled correctly
			results := objc.Call[foundation.Array](request, objc.Sel("results"))

			// Iterate over results safely
			for _, result := range foundation.ArrayToSlice[objc.Object](results) {
				if objc.Call[bool](result, objc.Sel("isKindOfClass:"), objc.GetClass("VNRecognizedTextObservation")) {
					// Access the boundingBox as foundation.Rect
					boundingBox := objc.Call[foundation.Rect](result, objc.Sel("boundingBox"))

					// Convert Vision's bottom-left origin to top-left origin
					x0 := int(boundingBox.Origin.X * float64(width))
					y0 := int((1.0 - (boundingBox.Origin.Y + boundingBox.Size.Height)) * float64(height)) // Flip Y-axis
					x1 := int((boundingBox.Origin.X + boundingBox.Size.Width) * float64(width))
					y1 := int((1.0 - boundingBox.Origin.Y) * float64(height)) // Flip Y-axis

					// Get the recognized text
					recognizedText := objc.Call[foundation.Array](result, objc.Sel("topCandidates:"), 1)
					for _, textCandidate := range foundation.ArrayToSlice[objc.Object](recognizedText) {
						text := objc.Call[string](textCandidate, objc.Sel("string"))
						_, _, _, _, _ = x0, y0, x1, y1, text
					}
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
			errObj.Autorelease()
			err = errors.New(errObj.Description())
			return
		}
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ProcessResponse{}, nil
}
