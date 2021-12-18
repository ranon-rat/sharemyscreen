package controllers

import (
	"image/png"
	"net/http"
)

const (
	boundaryWord = "MJPEGBOUNDARY"
	headerf      = "\r\n" +
		"--" + boundaryWord + "\r\n" +
		"Content-Type: image/png\r\n" +
		"X-Timestamp: 0.000000\r\n" +
		"\r\n"
)

func StreamingImage(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "multipart/x-mixed-replace; boundary="+boundaryWord)
	for {
		if _, err := res.Write([]byte(headerf)); err != nil {
			return
		}

		if err := png.Encode(res, gettingImage()); err != nil {
			return
		}

	}
}
