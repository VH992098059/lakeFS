// Code generated by MockGen. DO NOT EDIT.
// Source: meta_range.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	graveler "github.com/treeverse/lakefs/pkg/graveler"
	committed "github.com/treeverse/lakefs/pkg/graveler/committed"
)

// MockIterator is a mock of Iterator interface.
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator.
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance.
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIterator)(nil).Close))
}

// Err mocks base method.
func (m *MockIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next))
}

// NextRange mocks base method.
func (m *MockIterator) NextRange() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextRange")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NextRange indicates an expected call of NextRange.
func (mr *MockIteratorMockRecorder) NextRange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextRange", reflect.TypeOf((*MockIterator)(nil).NextRange))
}

// SeekGE mocks base method.
func (m *MockIterator) SeekGE(id graveler.Key) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SeekGE", id)
}

// SeekGE indicates an expected call of SeekGE.
func (mr *MockIteratorMockRecorder) SeekGE(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeekGE", reflect.TypeOf((*MockIterator)(nil).SeekGE), id)
}

// Value mocks base method.
func (m *MockIterator) Value() (*graveler.ValueRecord, *committed.Range) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(*graveler.ValueRecord)
	ret1, _ := ret[1].(*committed.Range)
	return ret0, ret1
}

// Value indicates an expected call of Value.
func (mr *MockIteratorMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockIterator)(nil).Value))
}

// MockDiffIterator is a mock of DiffIterator interface.
type MockDiffIterator struct {
	ctrl     *gomock.Controller
	recorder *MockDiffIteratorMockRecorder
}

// MockDiffIteratorMockRecorder is the mock recorder for MockDiffIterator.
type MockDiffIteratorMockRecorder struct {
	mock *MockDiffIterator
}

// NewMockDiffIterator creates a new mock instance.
func NewMockDiffIterator(ctrl *gomock.Controller) *MockDiffIterator {
	mock := &MockDiffIterator{ctrl: ctrl}
	mock.recorder = &MockDiffIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiffIterator) EXPECT() *MockDiffIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDiffIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDiffIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDiffIterator)(nil).Close))
}

// Err mocks base method.
func (m *MockDiffIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockDiffIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockDiffIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockDiffIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockDiffIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockDiffIterator)(nil).Next))
}

// NextRange mocks base method.
func (m *MockDiffIterator) NextRange() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextRange")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NextRange indicates an expected call of NextRange.
func (mr *MockDiffIteratorMockRecorder) NextRange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextRange", reflect.TypeOf((*MockDiffIterator)(nil).NextRange))
}

// SeekGE mocks base method.
func (m *MockDiffIterator) SeekGE(id graveler.Key) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SeekGE", id)
}

// SeekGE indicates an expected call of SeekGE.
func (mr *MockDiffIteratorMockRecorder) SeekGE(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeekGE", reflect.TypeOf((*MockDiffIterator)(nil).SeekGE), id)
}

// Value mocks base method.
func (m *MockDiffIterator) Value() (*graveler.Diff, *committed.RangeDiff) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(*graveler.Diff)
	ret1, _ := ret[1].(*committed.RangeDiff)
	return ret0, ret1
}

// Value indicates an expected call of Value.
func (mr *MockDiffIteratorMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockDiffIterator)(nil).Value))
}

// MockMetaRangeManager is a mock of MetaRangeManager interface.
type MockMetaRangeManager struct {
	ctrl     *gomock.Controller
	recorder *MockMetaRangeManagerMockRecorder
}

// MockMetaRangeManagerMockRecorder is the mock recorder for MockMetaRangeManager.
type MockMetaRangeManagerMockRecorder struct {
	mock *MockMetaRangeManager
}

// NewMockMetaRangeManager creates a new mock instance.
func NewMockMetaRangeManager(ctrl *gomock.Controller) *MockMetaRangeManager {
	mock := &MockMetaRangeManager{ctrl: ctrl}
	mock.recorder = &MockMetaRangeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetaRangeManager) EXPECT() *MockMetaRangeManagerMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockMetaRangeManager) Exists(ctx context.Context, ns graveler.StorageNamespace, id graveler.MetaRangeID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, ns, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockMetaRangeManagerMockRecorder) Exists(ctx, ns, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockMetaRangeManager)(nil).Exists), ctx, ns, id)
}

// GetMetaRangeURI mocks base method.
func (m *MockMetaRangeManager) GetMetaRangeURI(ctx context.Context, ns graveler.StorageNamespace, metaRangeID graveler.MetaRangeID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetaRangeURI", ctx, ns, metaRangeID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaRangeURI indicates an expected call of GetMetaRangeURI.
func (mr *MockMetaRangeManagerMockRecorder) GetMetaRangeURI(ctx, ns, metaRangeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaRangeURI", reflect.TypeOf((*MockMetaRangeManager)(nil).GetMetaRangeURI), ctx, ns, metaRangeID)
}

// GetRangeByKey mocks base method.
func (m *MockMetaRangeManager) GetRangeByKey(ctx context.Context, ns graveler.StorageNamespace, id graveler.MetaRangeID, key graveler.Key) (*committed.Range, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRangeByKey", ctx, ns, id, key)
	ret0, _ := ret[0].(*committed.Range)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRangeByKey indicates an expected call of GetRangeByKey.
func (mr *MockMetaRangeManagerMockRecorder) GetRangeByKey(ctx, ns, id, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRangeByKey", reflect.TypeOf((*MockMetaRangeManager)(nil).GetRangeByKey), ctx, ns, id, key)
}

// GetRangeURI mocks base method.
func (m *MockMetaRangeManager) GetRangeURI(ctx context.Context, ns graveler.StorageNamespace, rangeID graveler.RangeID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRangeURI", ctx, ns, rangeID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRangeURI indicates an expected call of GetRangeURI.
func (mr *MockMetaRangeManagerMockRecorder) GetRangeURI(ctx, ns, rangeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRangeURI", reflect.TypeOf((*MockMetaRangeManager)(nil).GetRangeURI), ctx, ns, rangeID)
}

// GetValue mocks base method.
func (m *MockMetaRangeManager) GetValue(ctx context.Context, ns graveler.StorageNamespace, id graveler.MetaRangeID, key graveler.Key) (*graveler.ValueRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", ctx, ns, id, key)
	ret0, _ := ret[0].(*graveler.ValueRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValue indicates an expected call of GetValue.
func (mr *MockMetaRangeManagerMockRecorder) GetValue(ctx, ns, id, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockMetaRangeManager)(nil).GetValue), ctx, ns, id, key)
}

// NewMetaRangeIterator mocks base method.
func (m *MockMetaRangeManager) NewMetaRangeIterator(ctx context.Context, ns graveler.StorageNamespace, metaRangeID graveler.MetaRangeID) (committed.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMetaRangeIterator", ctx, ns, metaRangeID)
	ret0, _ := ret[0].(committed.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMetaRangeIterator indicates an expected call of NewMetaRangeIterator.
func (mr *MockMetaRangeManagerMockRecorder) NewMetaRangeIterator(ctx, ns, metaRangeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMetaRangeIterator", reflect.TypeOf((*MockMetaRangeManager)(nil).NewMetaRangeIterator), ctx, ns, metaRangeID)
}

// NewWriter mocks base method.
func (m *MockMetaRangeManager) NewWriter(ctx context.Context, ns graveler.StorageNamespace, metadata graveler.Metadata) committed.MetaRangeWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWriter", ctx, ns, metadata)
	ret0, _ := ret[0].(committed.MetaRangeWriter)
	return ret0
}

// NewWriter indicates an expected call of NewWriter.
func (mr *MockMetaRangeManagerMockRecorder) NewWriter(ctx, ns, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWriter", reflect.TypeOf((*MockMetaRangeManager)(nil).NewWriter), ctx, ns, metadata)
}

// MockMetaRangeWriter is a mock of MetaRangeWriter interface.
type MockMetaRangeWriter struct {
	ctrl     *gomock.Controller
	recorder *MockMetaRangeWriterMockRecorder
}

// MockMetaRangeWriterMockRecorder is the mock recorder for MockMetaRangeWriter.
type MockMetaRangeWriterMockRecorder struct {
	mock *MockMetaRangeWriter
}

// NewMockMetaRangeWriter creates a new mock instance.
func NewMockMetaRangeWriter(ctrl *gomock.Controller) *MockMetaRangeWriter {
	mock := &MockMetaRangeWriter{ctrl: ctrl}
	mock.recorder = &MockMetaRangeWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetaRangeWriter) EXPECT() *MockMetaRangeWriterMockRecorder {
	return m.recorder
}

// Abort mocks base method.
func (m *MockMetaRangeWriter) Abort() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Abort")
	ret0, _ := ret[0].(error)
	return ret0
}

// Abort indicates an expected call of Abort.
func (mr *MockMetaRangeWriterMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockMetaRangeWriter)(nil).Abort))
}

// Close mocks base method.
func (m *MockMetaRangeWriter) Close(arg0 context.Context) (*graveler.MetaRangeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(*graveler.MetaRangeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close.
func (mr *MockMetaRangeWriterMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMetaRangeWriter)(nil).Close), arg0)
}

// WriteRange mocks base method.
func (m *MockMetaRangeWriter) WriteRange(arg0 committed.Range) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRange indicates an expected call of WriteRange.
func (mr *MockMetaRangeWriterMockRecorder) WriteRange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRange", reflect.TypeOf((*MockMetaRangeWriter)(nil).WriteRange), arg0)
}

// WriteRecord mocks base method.
func (m *MockMetaRangeWriter) WriteRecord(arg0 graveler.ValueRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRecord", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRecord indicates an expected call of WriteRecord.
func (mr *MockMetaRangeWriterMockRecorder) WriteRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRecord", reflect.TypeOf((*MockMetaRangeWriter)(nil).WriteRecord), arg0)
}
