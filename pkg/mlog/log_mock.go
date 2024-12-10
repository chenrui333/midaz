// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/maxwelbm/Workspace/midaz/pkg/mlog/log.go
//
// Generated by this command:
//
//	mockgen -source=/Users/maxwelbm/Workspace/midaz/pkg/mlog/log.go -destination=/Users/maxwelbm/Workspace/midaz/pkg/mlog/log_mock.go -package mlog
//

// Package mlog is a generated GoMock package.
package mlog

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
	isgomock struct{}
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLogger) Debug(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), args...)
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Debugln mocks base method.
func (m *MockLogger) Debugln(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugln", varargs...)
}

// Debugln indicates an expected call of Debugln.
func (mr *MockLoggerMockRecorder) Debugln(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugln", reflect.TypeOf((*MockLogger)(nil).Debugln), args...)
}

// Error mocks base method.
func (m *MockLogger) Error(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), args...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Errorln mocks base method.
func (m *MockLogger) Errorln(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorln", varargs...)
}

// Errorln indicates an expected call of Errorln.
func (mr *MockLoggerMockRecorder) Errorln(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorln", reflect.TypeOf((*MockLogger)(nil).Errorln), args...)
}

// Fatal mocks base method.
func (m *MockLogger) Fatal(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockLoggerMockRecorder) Fatal(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), args...)
}

// Fatalf mocks base method.
func (m *MockLogger) Fatalf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockLoggerMockRecorder) Fatalf(format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockLogger)(nil).Fatalf), varargs...)
}

// Fatalln mocks base method.
func (m *MockLogger) Fatalln(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalln", varargs...)
}

// Fatalln indicates an expected call of Fatalln.
func (mr *MockLoggerMockRecorder) Fatalln(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalln", reflect.TypeOf((*MockLogger)(nil).Fatalln), args...)
}

// Info mocks base method.
func (m *MockLogger) Info(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), args...)
}

// Infof mocks base method.
func (m *MockLogger) Infof(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Infoln mocks base method.
func (m *MockLogger) Infoln(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infoln", varargs...)
}

// Infoln indicates an expected call of Infoln.
func (mr *MockLoggerMockRecorder) Infoln(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infoln", reflect.TypeOf((*MockLogger)(nil).Infoln), args...)
}

// Sync mocks base method.
func (m *MockLogger) Sync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockLoggerMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockLogger)(nil).Sync))
}

// Warn mocks base method.
func (m *MockLogger) Warn(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerMockRecorder) Warn(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger)(nil).Warn), args...)
}

// Warnf mocks base method.
func (m *MockLogger) Warnf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockLoggerMockRecorder) Warnf(format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockLogger)(nil).Warnf), varargs...)
}

// Warnln mocks base method.
func (m *MockLogger) Warnln(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnln", varargs...)
}

// Warnln indicates an expected call of Warnln.
func (mr *MockLoggerMockRecorder) Warnln(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnln", reflect.TypeOf((*MockLogger)(nil).Warnln), args...)
}

// WithDefaultMessageTemplate mocks base method.
func (m *MockLogger) WithDefaultMessageTemplate(message string) Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDefaultMessageTemplate", message)
	ret0, _ := ret[0].(Logger)
	return ret0
}

// WithDefaultMessageTemplate indicates an expected call of WithDefaultMessageTemplate.
func (mr *MockLoggerMockRecorder) WithDefaultMessageTemplate(message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDefaultMessageTemplate", reflect.TypeOf((*MockLogger)(nil).WithDefaultMessageTemplate), message)
}

// WithFields mocks base method.
func (m *MockLogger) WithFields(fields ...any) Logger {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithFields", varargs...)
	ret0, _ := ret[0].(Logger)
	return ret0
}

// WithFields indicates an expected call of WithFields.
func (mr *MockLoggerMockRecorder) WithFields(fields ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithFields", reflect.TypeOf((*MockLogger)(nil).WithFields), fields...)
}
