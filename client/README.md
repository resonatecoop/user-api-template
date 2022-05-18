# User API Swagger client

## Install goswagger

[https://goswagger.io/install.html](https://goswagger.io/install.html)

## Generate swagger client

```sh
swagger generate client -f ./third_party/OpenAPI/user/user.swagger.json -t ./client -a "user-api-client"
```

## Code example

```go
package main

import (
	"fmt"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/resonatecoop/user-api/client/client"

	"github.com/resonatecoop/user-api-client/client/usergroups"
	"github.com/resonatecoop/user-api-client/models"
)

var (
	insecureSkipVerify = true
	basepath           = ""
	schemes            = []string{}
)

func main() {
	httpClient, err := httptransport.TLSClient(httptransport.TLSClientOptions{
		InsecureSkipVerify: insecureSkipVerify,
	})

	if err != nil {
		panic(err)
	}

	hostname := fmt.Sprintf("%s%s", "0.0.0.0", ":11000") // replace with your user api hostname
	transport := httptransport.NewWithClient(hostname, basepath, schemes, httpClient)

	client := apiclient.New(transport, strfmt.Default)

	bearer := httptransport.BearerToken("4f29a260-141e-4238-a6c0-921e1e842fcd") // replace with valid token

	params := usergroups.NewResonateUserAddUserGroupParams()

	params.WithID("4e4a2187-2d7c-49ac-978e-0656c2d4b050") // replace with valid user id

	params.Body = &models.UserUserGroupCreateRequest{
		DisplayName: "Burial",
		GroupType:   "persona",
	}

	result, err := client.Usergroups.ResonateUserAddUserGroup(params, bearer)

	if err != nil {
		if casted, ok := err.(*usergroups.ResonateUserAddUserGroupDefault); ok {
			if ok {
				fmt.Println(casted)
			}
		}
	}

	if result == nil {
		panic("User API not started?")
	}

	fmt.Println(result.Payload)
	return
}
```
