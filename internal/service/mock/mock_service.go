// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/murat96k/kitaptar.kz/api"
	entity "github.com/murat96k/kitaptar.kz/internal/entity"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateAuthor mocks base method.
func (m *MockService) CreateAuthor(ctx context.Context, req *api.AuthorRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthor", ctx, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuthor indicates an expected call of CreateAuthor.
func (mr *MockServiceMockRecorder) CreateAuthor(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthor", reflect.TypeOf((*MockService)(nil).CreateAuthor), ctx, req)
}

// CreateBook mocks base method.
func (m *MockService) CreateBook(ctx context.Context, req *api.BookRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", ctx, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockServiceMockRecorder) CreateBook(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockService)(nil).CreateBook), ctx, req)
}

// CreateFilePath mocks base method.
func (m *MockService) CreateFilePath(ctx context.Context, req *api.FilePathRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFilePath", ctx, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFilePath indicates an expected call of CreateFilePath.
func (mr *MockServiceMockRecorder) CreateFilePath(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilePath", reflect.TypeOf((*MockService)(nil).CreateFilePath), ctx, req)
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(ctx context.Context, u *entity.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, u)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), ctx, u)
}

// DeleteAuthor mocks base method.
func (m *MockService) DeleteAuthor(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthor", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthor indicates an expected call of DeleteAuthor.
func (mr *MockServiceMockRecorder) DeleteAuthor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthor", reflect.TypeOf((*MockService)(nil).DeleteAuthor), ctx, id)
}

// DeleteBook mocks base method.
func (m *MockService) DeleteBook(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockServiceMockRecorder) DeleteBook(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockService)(nil).DeleteBook), ctx, id)
}

// DeleteFilePath mocks base method.
func (m *MockService) DeleteFilePath(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFilePath", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFilePath indicates an expected call of DeleteFilePath.
func (mr *MockServiceMockRecorder) DeleteFilePath(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFilePath", reflect.TypeOf((*MockService)(nil).DeleteFilePath), ctx, id)
}

// DeleteUser mocks base method.
func (m *MockService) DeleteUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockServiceMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockService)(nil).DeleteUser), ctx, id)
}

// GetAllAuthors mocks base method.
func (m *MockService) GetAllAuthors(ctx context.Context) ([]entity.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAuthors", ctx)
	ret0, _ := ret[0].([]entity.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAuthors indicates an expected call of GetAllAuthors.
func (mr *MockServiceMockRecorder) GetAllAuthors(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAuthors", reflect.TypeOf((*MockService)(nil).GetAllAuthors), ctx)
}

// GetAllBooks mocks base method.
func (m *MockService) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks", ctx)
	ret0, _ := ret[0].([]entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockServiceMockRecorder) GetAllBooks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockService)(nil).GetAllBooks), ctx)
}

// GetAllFilePaths mocks base method.
func (m *MockService) GetAllFilePaths(ctx context.Context) ([]entity.FilePath, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilePaths", ctx)
	ret0, _ := ret[0].([]entity.FilePath)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilePaths indicates an expected call of GetAllFilePaths.
func (mr *MockServiceMockRecorder) GetAllFilePaths(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilePaths", reflect.TypeOf((*MockService)(nil).GetAllFilePaths), ctx)
}

// GetAuthorById mocks base method.
func (m *MockService) GetAuthorById(ctx context.Context, id string) (*entity.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorById", ctx, id)
	ret0, _ := ret[0].(*entity.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorById indicates an expected call of GetAuthorById.
func (mr *MockServiceMockRecorder) GetAuthorById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorById", reflect.TypeOf((*MockService)(nil).GetAuthorById), ctx, id)
}

// GetBookById mocks base method.
func (m *MockService) GetBookById(ctx context.Context, id string) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", ctx, id)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockServiceMockRecorder) GetBookById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockService)(nil).GetBookById), ctx, id)
}

// GetFilePathById mocks base method.
func (m *MockService) GetFilePathById(ctx context.Context, id string) (*entity.FilePath, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilePathById", ctx, id)
	ret0, _ := ret[0].(*entity.FilePath)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilePathById indicates an expected call of GetFilePathById.
func (mr *MockServiceMockRecorder) GetFilePathById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilePathById", reflect.TypeOf((*MockService)(nil).GetFilePathById), ctx, id)
}

// GetUserBooks mocks base method.
func (m *MockService) GetUserBooks(email string) ([]entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBooks", email)
	ret0, _ := ret[0].([]entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBooks indicates an expected call of GetUserBooks.
func (mr *MockServiceMockRecorder) GetUserBooks(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBooks", reflect.TypeOf((*MockService)(nil).GetUserBooks), email)
}

// GetUserByEmail mocks base method.
func (m *MockService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockServiceMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockService)(nil).GetUserByEmail), ctx, email)
}

// GetUserById mocks base method.
func (m *MockService) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockServiceMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockService)(nil).GetUserById), ctx, id)
}

// Login mocks base method.
func (m *MockService) Login(ctx context.Context, email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockServiceMockRecorder) Login(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockService)(nil).Login), ctx, email, password)
}

// UpdateAuthor mocks base method.
func (m *MockService) UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthor", ctx, id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthor indicates an expected call of UpdateAuthor.
func (mr *MockServiceMockRecorder) UpdateAuthor(ctx, id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthor", reflect.TypeOf((*MockService)(nil).UpdateAuthor), ctx, id, req)
}

// UpdateBook mocks base method.
func (m *MockService) UpdateBook(ctx context.Context, id string, req *api.BookRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", ctx, id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockServiceMockRecorder) UpdateBook(ctx, id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockService)(nil).UpdateBook), ctx, id, req)
}

// UpdateFilePath mocks base method.
func (m *MockService) UpdateFilePath(ctx context.Context, id string, req *api.FilePathRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilePath", ctx, id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFilePath indicates an expected call of UpdateFilePath.
func (mr *MockServiceMockRecorder) UpdateFilePath(ctx, id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilePath", reflect.TypeOf((*MockService)(nil).UpdateFilePath), ctx, id, req)
}

// UpdateUser mocks base method.
func (m *MockService) UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockServiceMockRecorder) UpdateUser(ctx, id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockService)(nil).UpdateUser), ctx, id, req)
}

// VerifyToken mocks base method.
func (m *MockService) VerifyToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockServiceMockRecorder) VerifyToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockService)(nil).VerifyToken), token)
}
