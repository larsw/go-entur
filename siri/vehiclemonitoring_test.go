package siri

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestUnmarsalServiceDeliveryMessage(t *testing.T) {
	dat, err := ioutil.ReadFile("examples/vm1.json")
	if err != nil {
		t.Error(err)
	}

	siri := Siri{}
	err = json.Unmarshal([]byte(dat), &siri)

	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, dat)

	t.Log("Iterating over vehicle activities:")

	for _, vmd := range siri.ServiceDelivery.VehicleMonitoringDelivery {
		for _, va := range vmd.VehicleActivity {
			t.Log(va.ItemIdentifier)
		}
	}
}
