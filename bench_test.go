package gendoc_test

import (
	gendoc "github.com/n404an/protoc-gen-doc"
	"testing"

	"github.com/pseudomuto/protokit/utils"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	set, _ := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	plugin := new(gendoc.Plugin)

	for i := 0; i < b.N; i++ {
		plugin.Generate(req)
	}
}
