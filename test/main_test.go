package main

import (
	"github/lucasjct/response"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

// specified the base url that will be used during the test.
func BaseURL() string {
	URL := "https://reqres.in/api"
	return URL
}

// create an instance of httpExpect and set the URL.
func HTTPexpectinstance(t *testing.T) *httpexpect.Expect {

	return httpexpect.Default(t, BaseURL())

}

// test to list all resources.
func TestListResources(t *testing.T) {
	e := HTTPexpectinstance(t)
	test := e.GET("/{resource}?page=1&per_page=5").
		Expect().
		Status(http.StatusOK)

	test.Body().IsEqual(response.Resource)

}

// test to list a specific resource using userID as a path parameter.
func TestListResourceByID(t *testing.T) {
	e := HTTPexpectinstance(t)

	userId := "2"
	test := e.GET("/{resource}" + "/" + userId).
		Expect().
		Status(http.StatusOK)
	test.Body().IsEqual(response.ResourceByID)
}

// test to update just one resource using userID as a path parameter.
func TestUpdateSpecificResource(t *testing.T) {
	nameUpdated := "fuchsia teste"
	userData := map[string]interface{}{
		"name": nameUpdated,
	}
	e := HTTPexpectinstance(t)
	userId := "2"
	test := e.PUT("/{resource}" + "/" + userId).
		WithJSON(userData).
		Expect().Status(http.StatusOK)
	test.Body().Contains(nameUpdated)
}

// test to update a specific resource.
func TestUpdateResourceURL(t *testing.T) {
	urlUpdated := "http://reqres.in/#support-heading"
	userData := map[string]interface{}{
		"url": urlUpdated,
	}
	e := HTTPexpectinstance(t)
	userId := "2"

	test := e.PATCH("/{resource}" + "/" + userId).
		WithJSON(userData).
		Expect().Status(http.StatusOK)
	test.Body().Contains(urlUpdated)
}

// test to delete specific data using userID as a path parameter.
func TestDeleteResourceCreated(t *testing.T) {
	e := HTTPexpectinstance(t)
	userId := "2"
	e.DELETE("/{resource}" + "/" + userId).
		Expect().Status(http.StatusNoContent)
}
