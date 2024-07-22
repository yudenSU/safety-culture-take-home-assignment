package folders

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func GetFoldersPaginized(req *FetchFolderRequest, startIndex int, pageSize int) (*FetchFolderResponse, int, error) {

	folders, nextIndex, err := FetchAllFoldersByOrgIDPaginzed(req.OrgID, startIndex, pageSize)

	if err != nil {
		return nil, -1, err
	}

	response := &FetchFolderResponse{Folders: folders}
	return response, nextIndex, nil
}

func FetchAllFoldersByOrgIDPaginzed(orgID uuid.UUID, startIndex int, pageSize int) ([]*Folder, int, error) {
	var (
		folders   []*Folder
		resFolder []*Folder
		err       error
	)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to fetch sample data: %v", r)
		}
	}()

	folders = GetSampleData()

	if err != nil {
		return nil, -1, err
	}

	resFolder = []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	if startIndex < 0 || pageSize < 0 {
		if startIndex == -1 && pageSize == -1 {
			return resFolder, len(resFolder), nil
		}
		return nil, -1, fmt.Errorf("invalid pagination parameters: startIndex=%d, pageSize=%d", startIndex, pageSize)
	}

	// Compute end index
	endIndex := startIndex + pageSize
	if startIndex > len(resFolder) {
		return nil, -1, fmt.Errorf("startIndex %d is beyond the length of the folder list", startIndex)
	}
	if endIndex > len(resFolder) {
		endIndex = len(resFolder)
	}

	return resFolder[startIndex:endIndex], endIndex, nil
}
