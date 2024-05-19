package http

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/egnptr/dating-app/model"
	"github.com/egnptr/dating-app/usecase"
)

type controller struct {
	Usecase usecase.Usecases
}

func NewPostController(service usecase.Usecases) *controller {
	return &controller{
		Usecase: service,
	}
}

func (c *controller) SignUp(w http.ResponseWriter, r *http.Request) {
	var (
		startTime      = time.Now()
		ctx            = r.Context()
		user           model.User
		response       responseDefault
		httpStatusCode = http.StatusOK
	)

	defer func() {
		response.Header.ProcessTime = float64(time.Since(startTime))
		w.WriteHeader(httpStatusCode)
		json.NewEncoder(w).Encode(response)
	}()

	w.Header().Set("Content-type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httpStatusCode = http.StatusBadRequest
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error unmarshaling the request"}
		return
	}

	err := c.Usecase.CreateUser(ctx, user)
	if err != nil {
		httpStatusCode = http.StatusInternalServerError
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error creating user"}
		return
	}

	response.Header.Messages = []string{"User is created successfully"}
}

func (c *controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var (
		startTime      = time.Now()
		ctx            = r.Context()
		req            model.LoginRequest
		response       responseDefault
		httpStatusCode = http.StatusOK
	)

	defer func() {
		response.Header.ProcessTime = float64(time.Since(startTime))
		w.WriteHeader(httpStatusCode)
		json.NewEncoder(w).Encode(response)
	}()

	w.Header().Set("Content-type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpStatusCode = http.StatusBadRequest
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error unmarshaling the request"}
		return
	}

	err := c.Usecase.Login(ctx, req)
	if err == model.UnauthorizedErr {
		httpStatusCode = http.StatusUnauthorized
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error unauthorized log in"}
		return
	} else if err != nil && err != model.UnauthorizedErr {
		httpStatusCode = http.StatusInternalServerError
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error logging in"}
		return
	}

	response.Header.Messages = []string{"Logged in successfully"}
}

func (c *controller) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	var (
		startTime      = time.Now()
		ctx            = r.Context()
		req            model.SubscribeRequest
		response       responseDefault
		httpStatusCode = http.StatusOK
	)

	defer func() {
		response.Header.ProcessTime = float64(time.Since(startTime))
		w.WriteHeader(httpStatusCode)
		json.NewEncoder(w).Encode(response)
	}()

	w.Header().Set("Content-type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpStatusCode = http.StatusBadRequest
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error unmarshaling the request"}
		return
	}

	req.Subscribe = true
	if strings.Contains(r.URL.Path, "unsubscribe") {
		req.Subscribe = false
	}

	err := c.Usecase.UpdateSubscription(ctx, req)
	if err != nil {
		httpStatusCode = http.StatusInternalServerError
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error update subscription"}
		return
	}

	response.Header.Messages = []string{"Subscription successfully"}
}

func (c *controller) GetProfiles(w http.ResponseWriter, r *http.Request) {
	var (
		startTime      = time.Now()
		ctx            = r.Context()
		req            model.GetRelatedUserRequest
		response       responseDefault
		httpStatusCode = http.StatusOK
	)

	defer func() {
		response.Header.ProcessTime = float64(time.Since(startTime))
		w.WriteHeader(httpStatusCode)
		json.NewEncoder(w).Encode(response)
	}()

	w.Header().Set("Content-type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpStatusCode = http.StatusBadRequest
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error unmarshaling the request"}
		return
	}

	data, err := c.Usecase.GetProfiles(ctx, req)
	if err != nil {
		httpStatusCode = http.StatusInternalServerError
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error fetching related profiles"}
		return
	}

	response.Data = data
}

func (c *controller) Swipe(w http.ResponseWriter, r *http.Request) {
	var (
		startTime      = time.Now()
		ctx            = r.Context()
		req            model.SwipeRequest
		response       responseDefault
		httpStatusCode = http.StatusOK
	)

	defer func() {
		response.Header.ProcessTime = float64(time.Since(startTime))
		w.WriteHeader(httpStatusCode)
		json.NewEncoder(w).Encode(response)
	}()

	w.Header().Set("Content-type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpStatusCode = http.StatusBadRequest
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error unmarshaling the request"}
		return
	}

	err := c.Usecase.Swipe(ctx, req)
	if err != nil {
		httpStatusCode = http.StatusInternalServerError
		response.Header.Reason = http.StatusText(httpStatusCode)
		response.Header.Messages = []string{"Error swiping profile"}
		return
	}

	response.Header.Messages = []string{"Swipe successful"}
}
