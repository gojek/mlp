// Code generated by mockery v2.0.0-alpha.14. DO NOT EDIT.

package mocks

import (
	models "github.com/gojek/mlp/api/models"
	mock "github.com/stretchr/testify/mock"
)

// SecretService is an autogenerated mock type for the SecretService type
type SecretService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: secretId, projectId
func (_m *SecretService) Delete(secretId models.Id, projectId models.Id) error {
	ret := _m.Called(secretId, projectId)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Id, models.Id) error); ok {
		r0 = rf(secretId, projectId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByIdAndProjectId provides a mock function with given fields: secretId, projectId
func (_m *SecretService) FindByIdAndProjectId(secretId models.Id, projectId models.Id) (*models.Secret, error) {
	ret := _m.Called(secretId, projectId)

	var r0 *models.Secret
	if rf, ok := ret.Get(0).(func(models.Id, models.Id) *models.Secret); ok {
		r0 = rf(secretId, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Id, models.Id) error); ok {
		r1 = rf(secretId, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecret provides a mock function with given fields: projectId
func (_m *SecretService) ListSecret(projectId models.Id) ([]*models.Secret, error) {
	ret := _m.Called(projectId)

	var r0 []*models.Secret
	if rf, ok := ret.Get(0).(func(models.Id) []*models.Secret); ok {
		r0 = rf(projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Id) error); ok {
		r1 = rf(projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: secret
func (_m *SecretService) Save(secret *models.Secret) (*models.Secret, error) {
	ret := _m.Called(secret)

	var r0 *models.Secret
	if rf, ok := ret.Get(0).(func(*models.Secret) *models.Secret); ok {
		r0 = rf(secret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Secret) error); ok {
		r1 = rf(secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
