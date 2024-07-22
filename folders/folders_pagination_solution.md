## Solution Justification

### Nomenclature
Client: refers to any consumer of the function, thus terms like "client's side" does not refer to the end-user's device.

### Key assumption
The order of items returned by GetSampleData is "stable", as in it's order is not shuffled. Thus, it is not necessary to order all values before returning a page to ensure we don't get overlapping folders.


### Description of expected behaviour
The client must input 2 additional values: 
- startIndex 
- pageSize

__Both of these values are of type int:__ this assumes that the expected input will never result in overflow.


This will retrieve a page starting from the item at the startIndex (first item is at index 0 to adhere to convention). PageSize will determine the size of the page given to the client.

The page is inclusive of the start index.

A providing a value of -1 to both startIndex and pageSize, returns an unpaginized view. Providing any other negative number to either input, returns an error. This is under the assumption that negative pages (such as going retrieving previous items, is not supported and such functionality must be done at the client's side).

### Notable cases:

__Case:__ Call function on an organisation with no files 

__Returns:__ returns no empty array of folders and next == 0


__Case:__ Call function and retrieve the last set of available files

__Returns:__ returns no last set of folders and next == last_index + 1 (total number of files)

__Case:__ pageSize > values left 

__Returns:__ returns all remaining items

### Error cases:

IF one of pageSize OR startIndex is < 0.

IF both ageSize AND startIndex is < 0 AND != -1

Panic from GetSampleData() is also handled in defer function

### Why two ints are used instead of a token system

The current solution aims to provide a simple and predictable interface for the client. By allowing the client to directly input the pageSize and startIndex, we allow more degrees of freedom for a client to directly query any subset of the data. e.g. in cases where the client may allow the user to jump to an arbitary page n. 

### return value:  startIndex
We assume the client does not need to retrieve the pageSize as the value does not change from the standard return case of the function, hence it is preserved by the client if they would like to make succesive calls.


