package chucknorris

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type RoundTripLogger struct {
	logger io.Writer
	next   http.RoundTripper
}

func (rtl *RoundTripLogger) RoundTrip(r *http.Request) (*http.Response, error) {
	startTime := time.Now()
	fmt.Fprintf(rtl.logger, "%s: Request started \n", startTime.Format(time.ANSIC))
	resp, err := rtl.next.RoundTrip(r)
	finishTime := time.Now()
	if err == nil {
		fmt.Fprintf(rtl.logger, "%s: Request finished in %s with code %d\n", finishTime.Format(time.ANSIC), finishTime.Sub(startTime), resp.StatusCode)
	} else {
		fmt.Fprintf(rtl.logger, "%s: Request finished with error in %s. Error: %s", finishTime.Format(time.ANSIC), finishTime.Sub(startTime), err)
	}

	return resp, err
}
