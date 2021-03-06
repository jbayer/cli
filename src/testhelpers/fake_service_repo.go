package testhelpers

import (
	"cf"
	"cf/net"
	"net/http"
)

type FakeServiceRepo struct {
	ServiceOfferings []cf.ServiceOffering

	CreateServiceInstanceName string
	CreateServiceInstancePlan cf.ServicePlan

	CreateUserProvidedServiceInstanceName string
	CreateUserProvidedServiceInstanceParameters map[string]string

	FindInstanceByNameName string
	FindInstanceByNameServiceInstance cf.ServiceInstance

	BindServiceServiceInstance cf.ServiceInstance
	BindServiceApplication cf.Application
	BindServiceErrorCode string

	UnbindServiceServiceInstance cf.ServiceInstance
	UnbindServiceApplication cf.Application

	DeleteServiceServiceInstance cf.ServiceInstance
}

func (repo *FakeServiceRepo) GetServiceOfferings() (offerings []cf.ServiceOffering, apiErr *net.ApiError) {
	offerings = repo.ServiceOfferings
	return
}

func (repo *FakeServiceRepo) CreateServiceInstance(name string, plan cf.ServicePlan) (apiErr *net.ApiError) {
	repo.CreateServiceInstanceName = name
	repo.CreateServiceInstancePlan = plan
	return
}

func (repo *FakeServiceRepo) CreateUserProvidedServiceInstance(name string, params map[string]string) (apiErr *net.ApiError) {
	repo.CreateUserProvidedServiceInstanceName = name
	repo.CreateUserProvidedServiceInstanceParameters = params
	return
}

func (repo *FakeServiceRepo) FindInstanceByName(name string) (instance cf.ServiceInstance, apiErr *net.ApiError) {
	repo.FindInstanceByNameName = name
	instance = repo.FindInstanceByNameServiceInstance
	return
}

func (repo *FakeServiceRepo) BindService(instance cf.ServiceInstance, app cf.Application) (apiErr *net.ApiError) {
	repo.BindServiceServiceInstance = instance
	repo.BindServiceApplication = app

	if repo.BindServiceErrorCode != "" {
		apiErr = net.NewApiError("Error binding service", repo.BindServiceErrorCode, http.StatusBadRequest)
	}

	return
}

func (repo *FakeServiceRepo) UnbindService(instance cf.ServiceInstance, app cf.Application) (apiErr *net.ApiError) {
	repo.UnbindServiceServiceInstance = instance
	repo.UnbindServiceApplication = app
	return
}

func (repo *FakeServiceRepo) DeleteService(instance cf.ServiceInstance) (apiErr *net.ApiError) {
	repo.DeleteServiceServiceInstance = instance
	return
}
