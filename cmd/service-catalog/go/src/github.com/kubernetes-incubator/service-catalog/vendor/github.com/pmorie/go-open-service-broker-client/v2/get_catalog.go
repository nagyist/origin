package v2

import (
	"fmt"
	"net/http"
)

func (c *client) GetCatalog() (*CatalogResponse, error) {
	fullURL := fmt.Sprintf(catalogURL, c.URL)

	response, err := c.prepareAndDo(http.MethodGet, fullURL, nil /* params */, nil /* request body */)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		catalogResponse := &CatalogResponse{}
		if err := c.unmarshalResponse(response, catalogResponse); err != nil {
			return nil, HTTPStatusCodeError{StatusCode: response.StatusCode, ResponseError: err}
		}

		if !c.EnableAlphaFeatures {
			for ii := range catalogResponse.Services {
				for jj := range catalogResponse.Services[ii].Plans {
					catalogResponse.Services[ii].Plans[jj].AlphaParameterSchemas = nil
				}
			}
		}

		return catalogResponse, nil
	default:
		return nil, c.handleFailureResponse(response)
	}
}
