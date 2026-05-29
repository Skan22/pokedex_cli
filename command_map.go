package main
import (
	"fmt"
	
	"github.com/Skan22/pokedex_cli/internal/pokeapi"
)
func commandMap()error{
	pokeapiClient := pokeapi.NewClient()
	resp,err := pokeapiClient.ListLocationAreas()
	if err!=nil {
		return err

	}
	fmt.Println("Location Areas : ")
	for _,area:=range(resp.Results){
		fmt.Println(area.Name)

	}

	
	return nil
}