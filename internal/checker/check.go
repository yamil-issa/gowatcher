package checker

import (
	"net/http"
	"time"
)

type CheckResult struct {
	Target string
	Status int
	Err    error
}

func CheckURL(url string) CheckResult {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return CheckResult{
			Target: url,
			Err: &UnreachableUrlError{
				Url: url,
				Err: err,
			},
		}
	}
	defer resp.Body.Close()

	return CheckResult{
		Target: url,
		Status: resp.StatusCode,
	}
}
