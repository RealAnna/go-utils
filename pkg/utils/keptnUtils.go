package utils

import (
	"log"
	"regexp"

	"github.com/keptn/go-utils/pkg/configuration-service/utils"
	"github.com/keptn/go-utils/pkg/models"
	"gopkg.in/yaml.v2"
)

// KeptnHandler provides an interface to keptn resources
type KeptnHandler struct {
	ResourceHandler *utils.ResourceHandler
}

// NewKeptnHandler returns a new KeptnHandler instance
func NewKeptnHandler(rh *utils.ResourceHandler) *KeptnHandler {
	return &KeptnHandler{
		ResourceHandler: rh,
	}
}

// GetShipyard returns the shipyard definition of a project
func (k *KeptnHandler) GetShipyard(project string) (*models.Shipyard, error) {
	shipyardResource, err := k.ResourceHandler.GetProjectResource(project, "shipyard.yaml")
	if err != nil {
		return nil, err
	}

	shipyard := models.Shipyard{}
	err = yaml.Unmarshal([]byte(shipyardResource.ResourceContent), &shipyard)
	if err != nil {
		return nil, err
	}
	return &shipyard, nil
}

// ValidateKeptnEntityName checks whether the provided name represents a valid
// project, service, or stage name
func ValidateKeptnEntityName(name string) bool {
	reg, err := regexp.Compile(`[a-z][a-z0-9-]+[a-z0-9]`)
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.FindString(name)
	return len(processedString) == len(name)
}
