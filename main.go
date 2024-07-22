package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	// res, err := folders.GetAllFolders(req)
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	return
	// }

	res, nextIndex, err := folders.GetFoldersPaginized(req, 0, 10)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	folders.PrettyPrint(res)
	fmt.Printf("Next index is: %d", nextIndex)
}
