package route

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
	"github.com/IceWhaleTech/CasaOS-MessageBus/service"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"go.uber.org/goleak"
	"gotest.tools/assert"
)

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

func TestEventRoute(t *testing.T) {
	defer goleak.VerifyNone(t)

	sourceID := "Foo"
	name := "Bar"

	expectedEventType := model.EventType{
		SourceID:         sourceID,
		Name:             name,
		PropertyTypeList: []model.PropertyType{{Name: "Property1"}, {Name: "Property2"}},
	}

	eventTypeJSON, err := json2.Marshal(expectedEventType)
	assert.NilError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repository, err := repository.NewInMemoryRepository(&ctx)
	assert.NilError(t, err)
	defer repository.Close()

	services := service.NewServices(&ctx, repository)
	services.Start()

	apiRoute := NewAPIRoute(services)

	e := echo.New()

	// register event type
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(eventTypeJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	err = apiRoute.RegisterEventType(e.NewContext(req, rec))
	assert.NilError(t, err)
	assert.Equal(t, rec.Code, http.StatusOK)

	var actualEventType model.EventType
	err = json2.UnmarshalFromString(rec.Body.String(), &actualEventType)
	assert.NilError(t, err)
	assert.DeepEqual(t, actualEventType, expectedEventType)

	// get event types
	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/", nil)

	err = apiRoute.GetEventTypes(e.NewContext(req, rec))
	assert.NilError(t, err)
	assert.Equal(t, rec.Code, http.StatusOK)

	var actualEventTypes []model.EventType
	err = json2.UnmarshalFromString(rec.Body.String(), &actualEventTypes)
	assert.NilError(t, err)
	assert.DeepEqual(t, actualEventTypes, []model.EventType{expectedEventType})

	// get event type
	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/", nil)

	err = apiRoute.GetEventType(e.NewContext(req, rec), sourceID, name)
	assert.NilError(t, err)
	assert.Equal(t, rec.Code, http.StatusOK)

	err = json2.UnmarshalFromString(rec.Body.String(), &actualEventType)
	assert.NilError(t, err)
	assert.DeepEqual(t, actualEventType, expectedEventType)

	// subscribe event type - TODO
}
