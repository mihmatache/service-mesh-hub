// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	settings_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/settings.smh.solo.io/v1alpha2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the Settings Resource
// DEPRECATED: Prefer reconciler pattern.
type SettingsEventHandler interface {
	CreateSettings(obj *settings_smh_solo_io_v1alpha2.Settings) error
	UpdateSettings(old, new *settings_smh_solo_io_v1alpha2.Settings) error
	DeleteSettings(obj *settings_smh_solo_io_v1alpha2.Settings) error
	GenericSettings(obj *settings_smh_solo_io_v1alpha2.Settings) error
}

type SettingsEventHandlerFuncs struct {
	OnCreate  func(obj *settings_smh_solo_io_v1alpha2.Settings) error
	OnUpdate  func(old, new *settings_smh_solo_io_v1alpha2.Settings) error
	OnDelete  func(obj *settings_smh_solo_io_v1alpha2.Settings) error
	OnGeneric func(obj *settings_smh_solo_io_v1alpha2.Settings) error
}

func (f *SettingsEventHandlerFuncs) CreateSettings(obj *settings_smh_solo_io_v1alpha2.Settings) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *SettingsEventHandlerFuncs) DeleteSettings(obj *settings_smh_solo_io_v1alpha2.Settings) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *SettingsEventHandlerFuncs) UpdateSettings(objOld, objNew *settings_smh_solo_io_v1alpha2.Settings) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *SettingsEventHandlerFuncs) GenericSettings(obj *settings_smh_solo_io_v1alpha2.Settings) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type SettingsEventWatcher interface {
	AddEventHandler(ctx context.Context, h SettingsEventHandler, predicates ...predicate.Predicate) error
}

type settingsEventWatcher struct {
	watcher events.EventWatcher
}

func NewSettingsEventWatcher(name string, mgr manager.Manager) SettingsEventWatcher {
	return &settingsEventWatcher{
		watcher: events.NewWatcher(name, mgr, &settings_smh_solo_io_v1alpha2.Settings{}),
	}
}

func (c *settingsEventWatcher) AddEventHandler(ctx context.Context, h SettingsEventHandler, predicates ...predicate.Predicate) error {
	handler := genericSettingsHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericSettingsHandler implements a generic events.EventHandler
type genericSettingsHandler struct {
	handler SettingsEventHandler
}

func (h genericSettingsHandler) Create(object runtime.Object) error {
	obj, ok := object.(*settings_smh_solo_io_v1alpha2.Settings)
	if !ok {
		return errors.Errorf("internal error: Settings handler received event for %T", object)
	}
	return h.handler.CreateSettings(obj)
}

func (h genericSettingsHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*settings_smh_solo_io_v1alpha2.Settings)
	if !ok {
		return errors.Errorf("internal error: Settings handler received event for %T", object)
	}
	return h.handler.DeleteSettings(obj)
}

func (h genericSettingsHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*settings_smh_solo_io_v1alpha2.Settings)
	if !ok {
		return errors.Errorf("internal error: Settings handler received event for %T", old)
	}
	objNew, ok := new.(*settings_smh_solo_io_v1alpha2.Settings)
	if !ok {
		return errors.Errorf("internal error: Settings handler received event for %T", new)
	}
	return h.handler.UpdateSettings(objOld, objNew)
}

func (h genericSettingsHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*settings_smh_solo_io_v1alpha2.Settings)
	if !ok {
		return errors.Errorf("internal error: Settings handler received event for %T", object)
	}
	return h.handler.GenericSettings(obj)
}
