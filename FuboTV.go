package mediaunlocktest

import "net/http"

func FuboTV(c http.Client) Result {
	return Result{Status: StatusNo}
}
