package fizzbuzz

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"bodotech.net/fizzbuzz/httpcommons"
	apiModels "bodotech.net/fizzbuzz/models/api"
)

// GetFizzBuzz returns a list of strings, and follows this instructions :
//
//	all multiples of int1 are replaced by string1.
//	all multiples of int2 are replaced by string2.
//	all multiples of int1 and int2 are replaced by string1string2.
//
// Returns an badRequest error if int1 or int2 or limit is not a valid integer.
func (c controller) GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"route": fmt.Sprintf("/%s", c.name), "method": "GET"})
	params := fizzbuzzparams{
		string1: r.URL.Query().Get("string1"),
		string2: r.URL.Query().Get("string2"),
		int1:    r.URL.Query().Get("int1"),
		int2:    r.URL.Query().Get("int2"),
		limit:   r.URL.Query().Get("limit"),
	}
	res, err := getFizzBuzz(params)
	resAPI := apiModels.FizzBuzzRes{
		Result: res,
	}
	if err != nil {
		logger.WithError(err).Error()
		if errors.Is(err, strconv.ErrSyntax) {
			resAPI.Error = err.Error()
			httpcommons.Respond(w, resAPI, http.StatusBadRequest)
			return
		}
		httpcommons.InternalServerError(w, err)
		return
	}

	httpcommons.Respond(w, resAPI, http.StatusOK)
}

type fizzbuzzparams struct {
	string1 string
	string2 string
	int1    string
	int2    string
	limit   string
}

func getFizzBuzz(params fizzbuzzparams) (res []string, err error) {
	var int1, int2, limit int
	int1, err = strconv.Atoi(params.int1)
	if err != nil {
		err = fmt.Errorf("error during int1 conversion: %w", err)
		return
	}
	int2, err = strconv.Atoi(params.int2)
	if err != nil {
		err = fmt.Errorf("error during int2 conversion: %w", err)
		return
	}
	limit, err = strconv.Atoi(params.limit)
	if err != nil {
		err = fmt.Errorf("error during limit conversion: %w", err)
		return
	}
	for i := 1; i <= limit; i++ {
		switch {
		case i%int1 == 0 && i%int2 == 0:
			res = append(res, params.string1+params.string2)
		case i%int1 == 0:
			res = append(res, params.string1)
		case i%int2 == 0:
			res = append(res, params.string2)
		default:
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	return
}
