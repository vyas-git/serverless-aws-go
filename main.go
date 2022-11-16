package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/valyala/fastjson"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	switch req.HTTPMethod {
	case "GET":
		// get payload from query string params
		name := req.QueryStringParameters["name"]

		if name != "" {
			ApiResponse = events.APIGatewayProxyResponse{
				Body:       "Hey " + name + "How are you!",
				StatusCode: 200,
			}
		} else {
			ApiResponse = events.APIGatewayProxyResponse{
				Body:       "Error: Missing name query string param",
				StatusCode: 500,
			}
		}
	case "POST":
		err := fastjson.Validate(req.Body)

		if err != nil {
			body := "Error : Invalid JSON payload err @@@@ " + fmt.Sprint(err) + " Payload is " + req.Body
			ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 500}

		} else {
			ApiResponse = events.APIGatewayProxyResponse{Body: req.Body, StatusCode: 200}

		}

	}
	return ApiResponse, nil
}
