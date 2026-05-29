package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Next_URL *string
var URL string
func (c *Client) ListLocationAreas()(LocationAreasResponse,error){
	endpoint:="/location-area"
	if Next_URL ==nil{URL = baseURL + endpoint}else {URL= *Next_URL}
	req,err := http.NewRequest("GET",URL,nil)
	if err!=nil {
		return  LocationAreasResponse{},err
	}
	res,err :=c.httpClient.Do(req)
	if err!=nil {
		return  LocationAreasResponse{},err
	}
	
	defer res.Body.Close()
	
	if res.StatusCode > 399{
		return LocationAreasResponse{},fmt.Errorf("Bad status code: %v",res.StatusCode)

	}
	data ,err := io.ReadAll(res.Body)
	if err!=nil {
		return  LocationAreasResponse{},err
	}
	LocationAreasResp :=LocationAreasResponse{}
	
	err = json.Unmarshal(data,&LocationAreasResp)
	
	if err!=nil {
		return  LocationAreasResponse{},err
	}
	Next_URL = LocationAreasResp.Next
	return LocationAreasResp,nil

}