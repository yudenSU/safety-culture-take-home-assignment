package folders

import (
	"fmt"

	"github.com/gofrs/uuid"
)

/*
Functionality:
This function takes in a fetchFolderRequest struct, uses the UUID
and calls FetchAllFoldersByOrgID, and returns the Folders and error.

Assumptions:
FetchFolderRequest is currently just a UUID in a struct. Although it
may seem inefficient to not take the value directly, storing the request
parameters as a type provides an interface easily shared with other services.
It also adheres to the open-closed SOLID principle, as it is easier to add new
parameters to the request object without breaking GetAllFOlders. Thus, we will keep it
as is.

Improvements:
- Unuseed variable declarations err, f1, fs
  - err should be used to communicate the err but is currently not used (more on this later)
  - f1 is unused, it seems that it may have been intended to work with Folder objects individually,
    but this is currently not present in the code.
  - fs, a slice of pointers to Folder structs. this may have been to store folder results from
    FetchAllFoldersByOrgID directly, but this is not done in the code. Instead f is used to do this.

- Unnecessary slices, f and fp. and unnecesary loop
  - f and fp is used unnecessary to build the FetchFolderResponse object
  - it seems that 2 loops are used to append folders into f and than append them to fp, which is unnecsary

- Error handling ignored:
  - The error response is hard-coded to nil
  - The error from fetchAllFolders is ignored.
  - Ignoring errors is bad practise as this makes errors hard to track and does not inform
    services using this function of errors that occur.
*/
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// unnecesary variable declarations
	// var (
	// 	err error
	// 	f1  Folder
	// 	fs  []*Folder
	// )

	// f := []Folder{}

	// Error handling ignored
	// r, _ := FetchAllFoldersByOrgID(req.OrgID)
	folders, err := FetchAllFoldersByOrgID(req.OrgID)

	//Premature return if an error is encountered.
	if err != nil {
		return nil, err
	}

	// Unnecessary loops
	// for k, v := range r {
	// 	f = append(f, *v)
	// }
	// var fp []*Folder
	// for k1, v1 := range f {
	// 	fp = append(fp, &v1)
	// }

	// var ffr *FetchFolderResponse
	// ffr = &FetchFolderResponse{Folders: fp}

	//Create FetchFolderResponse to for succesfull return case
	response := &FetchFolderResponse{Folders: folders}
	return response, nil
}

/*
Functionality:
This function takes in a parameter orgID which is a UUID, calls GetSampleData(), and than filters the retrieved folders by the provided orgID.
The function responds with a slice of pointers to folders.

Assumptions:
We assume that GetSampleData() is mimicking a API call or a database query.

Sorting by UUID on application code is typically not ideal, as this does not scale well in the case where we have a large number
of organisations, GetSampleData in this case, will return a huge set of data that will be filtered very slowly.
Additionally OrgID as a property stored in the Folder struct, means that filtering is done iteratively as it is in the provided code, this
will take O(N) time (N = number of returned foders from GetSampleData).

This can improved by having the filtering logic performed by a system more optimised for this kind of task, such as a database management
system in the persistance layer.

We assume that this is out of the scope of this assignment based on the following:
- disclaimer in static.go stating that there is no need to update GetSampleData
- The sample data being store in sample.json, which may represent the schema of the data is not in a format that can be filtered faster tan o(n) time. (id as a property in the object and not an index)
- GetSampleData may represent a response from an API that we do not have authority or a need to modify
- This may be representative of a scenario where we never expect a large amount of organisation and the above optimizations will bring in more unnececsary overhead.

Improvements:
  - Error catching for GetSampleData
  - We do not check if the response from GetSampleData is working, and the error responses is hard-coded to Nil.
*/
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	//Initate all required variables
	var (
		folders   []*Folder
		resFolder []*Folder
		err       error
	)

	//Defer function to record any errors made by sample data
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to fetch sample data: %v", r)
		}
	}()

	folders = GetSampleData()

	// Premature return in the case that GetSampleData has am error
	if err != nil {
		return nil, err
	}

	//Filter and return folders
	resFolder = []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
