package fileAccess

import (
	"os"
	"time"
)

type FileAction struct {
	open   int
	read   int
	write  int
	delete int
}

type FileAccess struct {
	filename string
	action   FileAction
}

// read the contents of a file, leaving the file open for @duration
func (f FileAccess) readFile(duration int) (bool, error) {
	fHandle, err := os.Open(f.filename)
	if err != nil {
		return false, err
	}
	defer fHandle.Close()

	// sleep for the duration requested before closing the file
	time.Sleep(time.Duration(duration) * time.Second)
	return true, nil
}

func (f FileAccess) writeFile(content []byte, duration int) (bool, error) {
	fHandle, err := os.Create(f.filename)
	if err != nil {
		return false, err
	}
	defer fHandle.Close()

	// we want to write enough bytes so that we are not writing everything
	// then just leaving the file open for the @duration.
	sleepAmount := duration / len(content)
	for i := 0; i <= duration; i++ {
		// write the byte
		_, err := fHandle.Write([]byte{content[i]})
		if err != nil {
			return false, err
		}
		// sleep the specified amount
		time.Sleep(time.Duration(sleepAmount) * time.Second)
	}
	return true, nil
}

func (f FileAccess) deleteFile() (bool, error) {
	err := os.Remove(f.filename)
	if err != nil {
		return false, err
	}
	return true, nil
}
