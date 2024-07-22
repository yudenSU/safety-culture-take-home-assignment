package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

/*
These test assume sample data is not representative of real data source, and can be used as mock data.
*/
func Test_GetAllFolders(t *testing.T) {
	t.Run("test_success_case_666_folders", func(t *testing.T) {

		// Define UUIDs and folders for testing
		uuidMock, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		if err != nil {
			t.Fatalf("Error parsing uuidMock: %v", err)
		}

		// Create a request with the test UUID
		req := &folders.FetchFolderRequest{OrgID: uuidMock}

		// Act
		response, err := folders.GetAllFolders(req)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response.Folders, 666)

	})
	t.Run("test_success_case_0_folders", func(t *testing.T) {

		// Define UUIDs and folders for testing
		uuidMock, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-c155aaa6ae17")
		if err != nil {
			t.Fatalf("Error parsing uuidMock: %v", err)
		}

		// Create a request with the test UUID
		req := &folders.FetchFolderRequest{OrgID: uuidMock}

		// Act
		response, err := folders.GetAllFolders(req)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response.Folders, 0)

	})
}

func Test_GetAllFoldersPaginzed(t *testing.T) {
	t.Run("test_success_case_all_folders", func(t *testing.T) {

		// Define UUIDs and folders for testing
		uuidMock, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		if err != nil {
			t.Fatalf("Error parsing uuidMock: %v", err)
		}

		// Create a request with the test UUID
		req := &folders.FetchFolderRequest{OrgID: uuidMock}

		// Act
		response, next, err := folders.GetFoldersPaginized(req, -1, -1)

		// Assert
		assert.Equal(t, 666, next)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response.Folders, 666)

	})
	t.Run("test_success_case_0_folders", func(t *testing.T) {

		// Define UUIDs and folders for testing
		uuidMock, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-c155aaa6ae17")
		if err != nil {
			t.Fatalf("Error parsing uuidMock: %v", err)
		}

		// Create a request with the test UUID
		req := &folders.FetchFolderRequest{OrgID: uuidMock}

		// Act
		response, next, err := folders.GetFoldersPaginized(req, -1, -1)

		// Assert
		assert.Equal(t, 0, next)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response.Folders, 0)

	})
	t.Run("test_success_case_retrieves_subset_if_no_items_left", func(t *testing.T) {

		// Define UUIDs and folders for testing
		uuidMock, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		if err != nil {
			t.Fatalf("Error parsing uuidMock: %v", err)
		}

		// Create a request with the test UUID
		req := &folders.FetchFolderRequest{OrgID: uuidMock}

		// Act
		response, next, err := folders.GetFoldersPaginized(req, 660, 10)

		// Assert
		assert.Equal(t, 666, next)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response.Folders, 6)

	})
	t.Run("test_succes_no_overlaps", func(t *testing.T) {

		// Define UUIDs and folders for testing
		uuidMock, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		if err != nil {
			t.Fatalf("Error parsing uuidMock: %v", err)
		}

		// Create a request with the test UUID
		req := &folders.FetchFolderRequest{OrgID: uuidMock}

		// Act
		response, next, err := folders.GetFoldersPaginized(req, 0, 10)

		// Assert
		assert.Equal(t, 10, next)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response.Folders, 10)

		// Act 2
		response2, next, err := folders.GetFoldersPaginized(req, next, 10)

		// Assert 2
		assert.Equal(t, 20, next)
		assert.NoError(t, err)
		assert.NotNil(t, response2)
		assert.Len(t, response2.Folders, 10)

		// Check for no common elements between response.Folders and response2.Folders
		responseFolderIDs := make(map[uuid.UUID]struct{})
		for _, folder := range response.Folders {
			responseFolderIDs[folder.Id] = struct{}{}
		}

		for _, folder := range response2.Folders {
			_, exists := responseFolderIDs[folder.Id]
			assert.False(t, exists, "Found a common element in response and response2")
		}

	})
	t.Run("test_error_bad_input", func(t *testing.T) {

		// Define UUIDs and folders for testing
		uuidMock, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		if err != nil {
			t.Fatalf("Error parsing uuidMock: %v", err)
		}

		// Create a request with the test UUID
		req := &folders.FetchFolderRequest{OrgID: uuidMock}

		// Act
		_, _, err = folders.GetFoldersPaginized(req, 0, -1)

		// Assert
		assert.Error(t, err)

		// Act2
		_, _, err = folders.GetFoldersPaginized(req, -1, 10)

		// Assert2
		assert.Error(t, err)

		// Act3
		_, _, err = folders.GetFoldersPaginized(req, -2, -2)

		// Assert3
		assert.Error(t, err)

	})

}
