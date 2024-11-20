# HTTPexpect


The tests to API Rest, has used HTTPExpect, is a Golang Framework to send http requests. It is simple to use and fast to run. We can do the same assertion like Bruno.       



See details:  [httpexpect](https://github.com/gavv/httpexpect)  .


Run the workflow on Git Hub Action: [action to httpexpect](https://github.com/lucasjct/bruno-api-testing/actions/workflows/run-test-http-expect.yaml)  .



See below the example to send and test a simple request with GET method.    



```go
func TestListResourceByID(t *testing.T) {
	e := HTTPexpectinstance(t)

	userId := "2"
	test := e.GET("/{resource}" + "/" + userId).
		Expect().
		Status(http.StatusOK)
	test.Body().IsEqual(response.ResourceByID)
}
```    

Check more examples in [main_test](main_test.go) .

## How can I execute it?  
* Golang installed is mandatory  .
* run this command on root directory: `make test-api` .