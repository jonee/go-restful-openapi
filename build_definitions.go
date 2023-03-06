package restfulspec

import (
	"log"
	"reflect"

	restful "github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"
)

func buildDefinitions(ws *restful.WebService, cfg Config) (definitions spec.Definitions) {
	definitions = spec.Definitions{}
	for _, each := range ws.Routes() {
		val, hasMetadataApiIgnore := each.Metadata[MetadataApiIgnore]
		if !hasMetadataApiIgnore || !val.(bool) {
			addDefinitionsFromRouteTo(each, cfg, definitions)
		} else {
			log.Printf("Api ignore: %v %v", each.Method, each.Path)
		}
	}
	return
}

func addDefinitionsFromRouteTo(r restful.Route, cfg Config, d spec.Definitions) {
	builder := definitionBuilder{Definitions: d, Config: cfg}
	if r.ReadSample != nil {
		builder.addModel(reflect.TypeOf(r.ReadSample), "")
	}
	if r.WriteSample != nil {
		builder.addModel(reflect.TypeOf(r.WriteSample), "")
	}
	for _, v := range r.ResponseErrors {
		if v.Model == nil {
			continue
		}
		builder.addModel(reflect.TypeOf(v.Model), "")
	}
}
