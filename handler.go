package main

import (
	catzzzlogger_v1 "github.com/logston/k8s-catzzz-logger/pkg/apis/catzzzlogger/v1"
	log "github.com/sirupsen/logrus"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// CatzzzLoggerHandler is a sample implementation of Handler
type CatzzzLoggerHandler struct{}

// Init handles any handler initialization
func (t *CatzzzLoggerHandler) Init() error {
	log.Info("CatzzzLoggerHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *CatzzzLoggerHandler) ObjectCreated(obj interface{}) {
	log.Info("CatzzzLoggerHandler.ObjectCreated")

	cl, ok := obj.(*catzzzlogger_v1.CatzzzLogger)
	if !ok {
		log.Info("Could not cast object to CatzzzLogger")
		return
	}

	max := int(*cl.Spec.Count)

	for i := 0; i < max; i++ {
		log.Info("\n" + cl.Spec.AsciiCat)
	}
}

// ObjectDeleted is called when an object is deleted
func (t *CatzzzLoggerHandler) ObjectDeleted(obj interface{}) {
	log.Info("CatzzzLoggerHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *CatzzzLoggerHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("CatzzzLoggerHandler.ObjectUpdated")
}
