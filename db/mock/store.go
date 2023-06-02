// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockStore) ChangePassword(arg0 context.Context, arg1 db.ChangePasswordParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockStoreMockRecorder) ChangePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockStore)(nil).ChangePassword), arg0, arg1)
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 string) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateExercise mocks base method.
func (m *MockStore) CreateExercise(arg0 context.Context, arg1 db.CreateExerciseParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExercise", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExercise indicates an expected call of CreateExercise.
func (mr *MockStoreMockRecorder) CreateExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExercise", reflect.TypeOf((*MockStore)(nil).CreateExercise), arg0, arg1)
}

// CreateLift mocks base method.
func (m *MockStore) CreateLift(arg0 context.Context, arg1 db.CreateLiftParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLift", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLift indicates an expected call of CreateLift.
func (mr *MockStoreMockRecorder) CreateLift(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLift", reflect.TypeOf((*MockStore)(nil).CreateLift), arg0, arg1)
}

// CreateMuscleGroup mocks base method.
func (m *MockStore) CreateMuscleGroup(arg0 context.Context, arg1 string) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMuscleGroup", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMuscleGroup indicates an expected call of CreateMuscleGroup.
func (mr *MockStoreMockRecorder) CreateMuscleGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMuscleGroup", reflect.TypeOf((*MockStore)(nil).CreateMuscleGroup), arg0, arg1)
}

// CreateProfile mocks base method.
func (m *MockStore) CreateProfile(arg0 context.Context, arg1 db.CreateProfileParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockStoreMockRecorder) CreateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockStore)(nil).CreateProfile), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateStockExercise mocks base method.
func (m *MockStore) CreateStockExercise(arg0 context.Context, arg1 db.CreateStockExerciseParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStockExercise", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStockExercise indicates an expected call of CreateStockExercise.
func (mr *MockStoreMockRecorder) CreateStockExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStockExercise", reflect.TypeOf((*MockStore)(nil).CreateStockExercise), arg0, arg1)
}

// CreateTemplate mocks base method.
func (m *MockStore) CreateTemplate(arg0 context.Context, arg1 db.CreateTemplateParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemplate", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTemplate indicates an expected call of CreateTemplate.
func (mr *MockStoreMockRecorder) CreateTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemplate", reflect.TypeOf((*MockStore)(nil).CreateTemplate), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateWorkout mocks base method.
func (m *MockStore) CreateWorkout(arg0 context.Context, arg1 db.CreateWorkoutParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkout", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkout indicates an expected call of CreateWorkout.
func (mr *MockStoreMockRecorder) CreateWorkout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkout", reflect.TypeOf((*MockStore)(nil).CreateWorkout), arg0, arg1)
}

// DeleteCategory mocks base method.
func (m *MockStore) DeleteCategory(arg0 context.Context, arg1 int32) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1)
}

// DeleteExercise mocks base method.
func (m *MockStore) DeleteExercise(arg0 context.Context, arg1 db.DeleteExerciseParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExercise", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExercise indicates an expected call of DeleteExercise.
func (mr *MockStoreMockRecorder) DeleteExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExercise", reflect.TypeOf((*MockStore)(nil).DeleteExercise), arg0, arg1)
}

// DeleteLift mocks base method.
func (m *MockStore) DeleteLift(arg0 context.Context, arg1 db.DeleteLiftParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLift", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLift indicates an expected call of DeleteLift.
func (mr *MockStoreMockRecorder) DeleteLift(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLift", reflect.TypeOf((*MockStore)(nil).DeleteLift), arg0, arg1)
}

// DeleteMuscleGroup mocks base method.
func (m *MockStore) DeleteMuscleGroup(arg0 context.Context, arg1 int32) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMuscleGroup", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMuscleGroup indicates an expected call of DeleteMuscleGroup.
func (mr *MockStoreMockRecorder) DeleteMuscleGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMuscleGroup", reflect.TypeOf((*MockStore)(nil).DeleteMuscleGroup), arg0, arg1)
}

// DeleteProfile mocks base method.
func (m *MockStore) DeleteProfile(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockStoreMockRecorder) DeleteProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockStore)(nil).DeleteProfile), arg0, arg1)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), arg0, arg1)
}

// DeleteStockExercise mocks base method.
func (m *MockStore) DeleteStockExercise(arg0 context.Context, arg1 int32) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStockExercise", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStockExercise indicates an expected call of DeleteStockExercise.
func (mr *MockStoreMockRecorder) DeleteStockExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStockExercise", reflect.TypeOf((*MockStore)(nil).DeleteStockExercise), arg0, arg1)
}

// DeleteTemplate mocks base method.
func (m *MockStore) DeleteTemplate(arg0 context.Context, arg1 db.DeleteTemplateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTemplate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTemplate indicates an expected call of DeleteTemplate.
func (mr *MockStoreMockRecorder) DeleteTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTemplate", reflect.TypeOf((*MockStore)(nil).DeleteTemplate), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteWorkout mocks base method.
func (m *MockStore) DeleteWorkout(arg0 context.Context, arg1 db.DeleteWorkoutParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkout", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkout indicates an expected call of DeleteWorkout.
func (mr *MockStoreMockRecorder) DeleteWorkout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkout", reflect.TypeOf((*MockStore)(nil).DeleteWorkout), arg0, arg1)
}

// GetCategory mocks base method.
func (m *MockStore) GetCategory(arg0 context.Context, arg1 int32) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockStoreMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockStore)(nil).GetCategory), arg0, arg1)
}

// GetExercise mocks base method.
func (m *MockStore) GetExercise(arg0 context.Context, arg1 db.GetExerciseParams) (db.Exercise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExercise", arg0, arg1)
	ret0, _ := ret[0].(db.Exercise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExercise indicates an expected call of GetExercise.
func (mr *MockStoreMockRecorder) GetExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExercise", reflect.TypeOf((*MockStore)(nil).GetExercise), arg0, arg1)
}

// GetExerciseByName mocks base method.
func (m *MockStore) GetExerciseByName(arg0 context.Context, arg1 db.GetExerciseByNameParams) (db.Exercise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExerciseByName", arg0, arg1)
	ret0, _ := ret[0].(db.Exercise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExerciseByName indicates an expected call of GetExerciseByName.
func (mr *MockStoreMockRecorder) GetExerciseByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExerciseByName", reflect.TypeOf((*MockStore)(nil).GetExerciseByName), arg0, arg1)
}

// GetInactiveUser mocks base method.
func (m *MockStore) GetInactiveUser(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInactiveUser", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInactiveUser indicates an expected call of GetInactiveUser.
func (mr *MockStoreMockRecorder) GetInactiveUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInactiveUser", reflect.TypeOf((*MockStore)(nil).GetInactiveUser), arg0, arg1)
}

// GetLift mocks base method.
func (m *MockStore) GetLift(arg0 context.Context, arg1 db.GetLiftParams) (db.Lift, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLift", arg0, arg1)
	ret0, _ := ret[0].(db.Lift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLift indicates an expected call of GetLift.
func (mr *MockStoreMockRecorder) GetLift(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLift", reflect.TypeOf((*MockStore)(nil).GetLift), arg0, arg1)
}

// GetMaxLifts mocks base method.
func (m *MockStore) GetMaxLifts(arg0 context.Context, arg1 db.GetMaxLiftsParams) ([]db.Lift, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxLifts", arg0, arg1)
	ret0, _ := ret[0].([]db.Lift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxLifts indicates an expected call of GetMaxLifts.
func (mr *MockStoreMockRecorder) GetMaxLifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxLifts", reflect.TypeOf((*MockStore)(nil).GetMaxLifts), arg0, arg1)
}

// GetMaxLiftsByExercise mocks base method.
func (m *MockStore) GetMaxLiftsByExercise(arg0 context.Context, arg1 db.GetMaxLiftsByExerciseParams) ([]db.Lift, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxLiftsByExercise", arg0, arg1)
	ret0, _ := ret[0].([]db.Lift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxLiftsByExercise indicates an expected call of GetMaxLiftsByExercise.
func (mr *MockStoreMockRecorder) GetMaxLiftsByExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxLiftsByExercise", reflect.TypeOf((*MockStore)(nil).GetMaxLiftsByExercise), arg0, arg1)
}

// GetMaxLiftsByMuscleGroup mocks base method.
func (m *MockStore) GetMaxLiftsByMuscleGroup(arg0 context.Context, arg1 db.GetMaxLiftsByMuscleGroupParams) ([]db.GetMaxLiftsByMuscleGroupRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxLiftsByMuscleGroup", arg0, arg1)
	ret0, _ := ret[0].([]db.GetMaxLiftsByMuscleGroupRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxLiftsByMuscleGroup indicates an expected call of GetMaxLiftsByMuscleGroup.
func (mr *MockStoreMockRecorder) GetMaxLiftsByMuscleGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxLiftsByMuscleGroup", reflect.TypeOf((*MockStore)(nil).GetMaxLiftsByMuscleGroup), arg0, arg1)
}

// GetMaxRepLifts mocks base method.
func (m *MockStore) GetMaxRepLifts(arg0 context.Context, arg1 db.GetMaxRepLiftsParams) ([]db.Lift, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxRepLifts", arg0, arg1)
	ret0, _ := ret[0].([]db.Lift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxRepLifts indicates an expected call of GetMaxRepLifts.
func (mr *MockStoreMockRecorder) GetMaxRepLifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxRepLifts", reflect.TypeOf((*MockStore)(nil).GetMaxRepLifts), arg0, arg1)
}

// GetMuscleGroup mocks base method.
func (m *MockStore) GetMuscleGroup(arg0 context.Context, arg1 int32) (db.MuscleGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMuscleGroup", arg0, arg1)
	ret0, _ := ret[0].(db.MuscleGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMuscleGroup indicates an expected call of GetMuscleGroup.
func (mr *MockStoreMockRecorder) GetMuscleGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMuscleGroup", reflect.TypeOf((*MockStore)(nil).GetMuscleGroup), arg0, arg1)
}

// GetMuscleGroupByName mocks base method.
func (m *MockStore) GetMuscleGroupByName(arg0 context.Context, arg1 string) (db.MuscleGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMuscleGroupByName", arg0, arg1)
	ret0, _ := ret[0].(db.MuscleGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMuscleGroupByName indicates an expected call of GetMuscleGroupByName.
func (mr *MockStoreMockRecorder) GetMuscleGroupByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMuscleGroupByName", reflect.TypeOf((*MockStore)(nil).GetMuscleGroupByName), arg0, arg1)
}

// GetProfile mocks base method.
func (m *MockStore) GetProfile(arg0 context.Context, arg1 string) (db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0, arg1)
	ret0, _ := ret[0].(db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockStoreMockRecorder) GetProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockStore)(nil).GetProfile), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 string) (db.GetSessionRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.GetSessionRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetStockExercise mocks base method.
func (m *MockStore) GetStockExercise(arg0 context.Context, arg1 int32) (db.StockExercise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStockExercise", arg0, arg1)
	ret0, _ := ret[0].(db.StockExercise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStockExercise indicates an expected call of GetStockExercise.
func (mr *MockStoreMockRecorder) GetStockExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockExercise", reflect.TypeOf((*MockStore)(nil).GetStockExercise), arg0, arg1)
}

// GetTemplate mocks base method.
func (m *MockStore) GetTemplate(arg0 context.Context, arg1 db.GetTemplateParams) (db.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplate", arg0, arg1)
	ret0, _ := ret[0].(db.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplate indicates an expected call of GetTemplate.
func (mr *MockStoreMockRecorder) GetTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplate", reflect.TypeOf((*MockStore)(nil).GetTemplate), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (db.GetUserByEmailRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.GetUserByEmailRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 string) (db.GetUserByIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.GetUserByIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// GetWorkout mocks base method.
func (m *MockStore) GetWorkout(arg0 context.Context, arg1 db.GetWorkoutParams) (db.Workout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkout", arg0, arg1)
	ret0, _ := ret[0].(db.Workout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkout indicates an expected call of GetWorkout.
func (mr *MockStoreMockRecorder) GetWorkout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkout", reflect.TypeOf((*MockStore)(nil).GetWorkout), arg0, arg1)
}

// InsertInactiveUser mocks base method.
func (m *MockStore) InsertInactiveUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertInactiveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertInactiveUser indicates an expected call of InsertInactiveUser.
func (mr *MockStoreMockRecorder) InsertInactiveUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertInactiveUser", reflect.TypeOf((*MockStore)(nil).InsertInactiveUser), arg0, arg1)
}

// ListCategories mocks base method.
func (m *MockStore) ListCategories(arg0 context.Context) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", arg0)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCategories indicates an expected call of ListCategories.
func (mr *MockStoreMockRecorder) ListCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockStore)(nil).ListCategories), arg0)
}

// ListExercises mocks base method.
func (m *MockStore) ListExercises(arg0 context.Context, arg1 string) ([]db.Exercise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExercises", arg0, arg1)
	ret0, _ := ret[0].([]db.Exercise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExercises indicates an expected call of ListExercises.
func (mr *MockStoreMockRecorder) ListExercises(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExercises", reflect.TypeOf((*MockStore)(nil).ListExercises), arg0, arg1)
}

// ListLiftsFromWorkout mocks base method.
func (m *MockStore) ListLiftsFromWorkout(arg0 context.Context, arg1 db.ListLiftsFromWorkoutParams) ([]db.Lift, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLiftsFromWorkout", arg0, arg1)
	ret0, _ := ret[0].([]db.Lift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLiftsFromWorkout indicates an expected call of ListLiftsFromWorkout.
func (mr *MockStoreMockRecorder) ListLiftsFromWorkout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLiftsFromWorkout", reflect.TypeOf((*MockStore)(nil).ListLiftsFromWorkout), arg0, arg1)
}

// ListMuscleGroups mocks base method.
func (m *MockStore) ListMuscleGroups(arg0 context.Context) ([]db.MuscleGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMuscleGroups", arg0)
	ret0, _ := ret[0].([]db.MuscleGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMuscleGroups indicates an expected call of ListMuscleGroups.
func (mr *MockStoreMockRecorder) ListMuscleGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMuscleGroups", reflect.TypeOf((*MockStore)(nil).ListMuscleGroups), arg0)
}

// ListStockExercies mocks base method.
func (m *MockStore) ListStockExercies(arg0 context.Context) ([]db.StockExercise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStockExercies", arg0)
	ret0, _ := ret[0].([]db.StockExercise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStockExercies indicates an expected call of ListStockExercies.
func (mr *MockStoreMockRecorder) ListStockExercies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStockExercies", reflect.TypeOf((*MockStore)(nil).ListStockExercies), arg0)
}

// ListTemplates mocks base method.
func (m *MockStore) ListTemplates(arg0 context.Context, arg1 string) ([]db.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTemplates", arg0, arg1)
	ret0, _ := ret[0].([]db.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTemplates indicates an expected call of ListTemplates.
func (mr *MockStoreMockRecorder) ListTemplates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTemplates", reflect.TypeOf((*MockStore)(nil).ListTemplates), arg0, arg1)
}

// ListWorkouts mocks base method.
func (m *MockStore) ListWorkouts(arg0 context.Context, arg1 string) ([]db.Workout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkouts", arg0, arg1)
	ret0, _ := ret[0].([]db.Workout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkouts indicates an expected call of ListWorkouts.
func (mr *MockStoreMockRecorder) ListWorkouts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkouts", reflect.TypeOf((*MockStore)(nil).ListWorkouts), arg0, arg1)
}

// NewUserTx mocks base method.
func (m *MockStore) NewUserTx(arg0 context.Context, arg1 db.CreateUserParams) (db.NewUserTxResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUserTx", arg0, arg1)
	ret0, _ := ret[0].(db.NewUserTxResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUserTx indicates an expected call of NewUserTx.
func (mr *MockStoreMockRecorder) NewUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUserTx", reflect.TypeOf((*MockStore)(nil).NewUserTx), arg0, arg1)
}

// UpdateCategory mocks base method.
func (m *MockStore) UpdateCategory(arg0 context.Context, arg1 db.UpdateCategoryParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockStoreMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0, arg1)
}

// UpdateExercise mocks base method.
func (m *MockStore) UpdateExercise(arg0 context.Context, arg1 db.UpdateExerciseParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExercise", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExercise indicates an expected call of UpdateExercise.
func (mr *MockStoreMockRecorder) UpdateExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExercise", reflect.TypeOf((*MockStore)(nil).UpdateExercise), arg0, arg1)
}

// UpdateLift mocks base method.
func (m *MockStore) UpdateLift(arg0 context.Context, arg1 db.UpdateLiftParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLift", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLift indicates an expected call of UpdateLift.
func (mr *MockStoreMockRecorder) UpdateLift(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLift", reflect.TypeOf((*MockStore)(nil).UpdateLift), arg0, arg1)
}

// UpdateMuscleGroup mocks base method.
func (m *MockStore) UpdateMuscleGroup(arg0 context.Context, arg1 db.UpdateMuscleGroupParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMuscleGroup", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMuscleGroup indicates an expected call of UpdateMuscleGroup.
func (mr *MockStoreMockRecorder) UpdateMuscleGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMuscleGroup", reflect.TypeOf((*MockStore)(nil).UpdateMuscleGroup), arg0, arg1)
}

// UpdateProfile mocks base method.
func (m *MockStore) UpdateProfile(arg0 context.Context, arg1 db.UpdateProfileParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockStoreMockRecorder) UpdateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockStore)(nil).UpdateProfile), arg0, arg1)
}

// UpdateStockExercise mocks base method.
func (m *MockStore) UpdateStockExercise(arg0 context.Context, arg1 db.UpdateStockExerciseParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStockExercise", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStockExercise indicates an expected call of UpdateStockExercise.
func (mr *MockStoreMockRecorder) UpdateStockExercise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStockExercise", reflect.TypeOf((*MockStore)(nil).UpdateStockExercise), arg0, arg1)
}

// UpdateTemplate mocks base method.
func (m *MockStore) UpdateTemplate(arg0 context.Context, arg1 db.UpdateTemplateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTemplate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTemplate indicates an expected call of UpdateTemplate.
func (mr *MockStoreMockRecorder) UpdateTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTemplate", reflect.TypeOf((*MockStore)(nil).UpdateTemplate), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateWorkout mocks base method.
func (m *MockStore) UpdateWorkout(arg0 context.Context, arg1 db.UpdateWorkoutParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkout", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkout indicates an expected call of UpdateWorkout.
func (mr *MockStoreMockRecorder) UpdateWorkout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkout", reflect.TypeOf((*MockStore)(nil).UpdateWorkout), arg0, arg1)
}

// WorkoutTx mocks base method.
func (m *MockStore) WorkoutTx(arg0 context.Context, arg1 db.WorkoutTxParams) (db.Workout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkoutTx", arg0, arg1)
	ret0, _ := ret[0].(db.Workout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkoutTx indicates an expected call of WorkoutTx.
func (mr *MockStoreMockRecorder) WorkoutTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkoutTx", reflect.TypeOf((*MockStore)(nil).WorkoutTx), arg0, arg1)
}
