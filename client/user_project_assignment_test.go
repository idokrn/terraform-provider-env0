package client_test

import (
	"errors"

	. "github.com/env0/terraform-provider-env0/client"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Agent Project Assignment", func() {
	projectId := "projectId"

	assignPayload := &AssignUserToProjectPayload{
		UserId: "userId",
		Role:   Admin,
	}

	expectedResponse := &UserProjectAssignment{
		Id: "id",
	}

	errorMock := errors.New("error")

	Describe("AssignUserToProject", func() {

		Describe("Successful", func() {
			var actualResult *UserProjectAssignment
			var err error

			BeforeEach(func() {
				httpCall = mockHttpClient.EXPECT().
					Post("/permissions/projects/"+projectId, assignPayload, gomock.Any()).
					Do(func(path string, request interface{}, response *UserProjectAssignment) {
						*response = *expectedResponse
					}).Times(1)
				actualResult, err = apiClient.AssignUserToProject(projectId, assignPayload)

			})

			It("Should send POST request with params", func() {
				httpCall.Times(1)
			})

			It("should return the POST result", func() {
				Expect(*actualResult).To(Equal(*expectedResponse))
			})

			It("Should not return error", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("Failure", func() {
			var actualResult *UserProjectAssignment
			var err error

			BeforeEach(func() {
				httpCall = mockHttpClient.EXPECT().
					Post("/permissions/projects/"+projectId, gomock.Any(), gomock.Any()).Return(errorMock)
				actualResult, err = apiClient.AssignUserToProject(projectId, assignPayload)
			})

			It("Should fail if API call fails", func() {
				Expect(err).To(Equal(errorMock))
			})

			It("Should not return results", func() {
				Expect(actualResult).To(BeNil())
			})
		})
	})

	Describe("RemoveUserFromProject", func() {
		BeforeEach(func() {
			httpCall = mockHttpClient.EXPECT().Delete("/permissions/projects/" + projectId + "/users/" + expectedResponse.Id)
			apiClient.RemoveUserFromProject(projectId, expectedResponse.Id)
		})

		It("Should send DELETE request with assignment id", func() {
			httpCall.Times(1)
		})
	})

})