package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func main() {

}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf(event.HTTPMethod)

	return events.APIGatewayProxyResponse{}, nil
}
