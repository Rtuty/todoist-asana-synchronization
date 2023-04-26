package asana

import (
	"fmt"
	"os"

	"bitbucket.org/mikehouston/asana-go"
)

func GetAllUsers() {
	token, exists := os.LookupEnv("ASANA_TOKEN")
	if !exists {
		panic("Asana API token not found")
	}

	client := asana.NewClientWithAccessToken(token)

	pretty := true

	mass, err := client.AllWorkspaces(&asana.Options{Pretty: &pretty})
	if err != nil {
		panic(err)
	}

	for _, v := range mass {
		u, err := v.AllUsers(client)
		if err != nil {
			panic(err)
		}

		for _, v2 := range u {
			fmt.Printf("%v", v2)
		}
	}

}
