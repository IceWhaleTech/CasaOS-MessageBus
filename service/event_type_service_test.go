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

	// new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// new repository
	repository, err := repository.NewInMemoryRepository(&ctx)
	assert.NilError(t, err)
	defer repository.Close()

	// new service
	service := NewEventTypeService(&ctx, repository)

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

	go service.Start()

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

	_, err = service.Publish(expectedEvent)
	assert.NilError(t, err)

	actualEvent, ok := <-outputChannel
	assert.Equal(t, ok, true)
	assert.Equal(t, actualEvent.Name, expectedEvent.Name)
	assert.Equal(t, actualEvent.SourceID, expectedEvent.SourceID)

	for i, property := range actualEvent.Properties {
		assert.Equal(t, property.Name, expectedEvent.Properties[i].Name)
		assert.Equal(t, property.Value, expectedEvent.Properties[i].Value)
	}
}
