package githubprovider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	restclient "github.com/dieg0code/go-microservices/src/api/clients/restClient"
	"github.com/dieg0code/go-microservices/src/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)

}

func CreateRepo(accesToken string, request github.CreateRepoRequest) (*github.CreateRepoRequest, *github.GitHubErrorResponse) {
	header := getAuthorizationHeader(accesToken)
	headers := http.Header{}
	headers.Set(headerAuthorization, header)

	response, err := restclient.Post(urlCreateRepo, request, headers)

	if err != nil {
		log.Printf("error when trying to create new repo in github: %s", err.Error())
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GitHubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body",
			}
		}

		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshal github create repo successful response: %s", err.Error())
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when trying to unmarshal github create repo successful response",
		}
	}

	return &github.CreateRepoRequest{}, nil
}
