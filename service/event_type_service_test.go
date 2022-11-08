package service

import (
	"context"
	"testing"

	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
	"go.uber.org/goleak"
	"gotest.tools/assert"
)

func TestEventTypeService(t *testing.T) {
	defer goleak.VerifyNone(t)

	// new repository
	repository, err := repository.NewDatabaseRepositoryInMemory()
	assert.NilError(t, err)
	defer repository.Close()

	// new service
	service := NewEventTypeService(&repository)

	// new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go service.Start(&ctx)

	sourceID := "Foo"
	name := "Bar"

	// register event type
	_, err = service.RegisterEventType(model.EventType{
		SourceID:         sourceID,
		Name:             name,
		PropertyTypeList: []model.PropertyType{{Name: "Property1"}, {Name: "Property2"}},
	})

	assert.NilError(t, err)

	// get event types
	eventTypes, err := service.GetEventTypes()
	assert.NilError(t, err)
	assert.Equal(t, len(eventTypes), 1)

	// get event types by source id
	eventTypes, err = service.GetEventTypesBySourceID(sourceID)
	assert.NilError(t, err)
	assert.Equal(t, len(eventTypes), 1)

	// get event type
	eventType, err := service.GetEventType(sourceID, name)
	assert.NilError(t, err)
	assert.Equal(t, eventType.SourceID, sourceID)
	assert.Equal(t, eventType.Name, name)

	// subscribe event type
	channel, err := service.Subscribe(sourceID, name)
	assert.NilError(t, err)

	outputChannel := make(chan model.Event)

	go func() {
		event, ok := <-channel
		if !ok {
			t.Error("channel closed")
		}
		outputChannel <- event
	}()

	expectedEvent := model.Event{
		SourceID: sourceID,
		Name:     name,
		Properties: []model.Property{
			{Name: "Property1", Value: "Value1"},
			{Name: "Property2", Value: "Value2"},
		},
	}

	actualEvent1, err := service.Publish(expectedEvent)
	assert.NilError(t, err)
	assert.DeepEqual(t, model.Event{SourceID: actualEvent1.SourceID, Name: actualEvent1.Name, Properties: actualEvent1.Properties}, expectedEvent)

	actualEvent2, ok := <-outputChannel
	assert.Equal(t, ok, true)
	assert.DeepEqual(t, actualEvent2, *actualEvent1)
}
