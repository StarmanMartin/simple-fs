package fs

import (
	"io"
	"os"
	"time"
)

//CopyFolder Copys Folders
func CopyFolder(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = CopyFolder(sourcefilepointer, destinationfilepointer)
			if err != nil {
				return
			}
		} else {
			err = SyncFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				return
			}
		}

	}

	return
}

//SyncFile Copys Folders
func SyncFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return
	}

	si, err := os.Stat(source)
	if err != nil {
		return
	}

	di, err := os.Stat(dest)

	if err == nil {
		if si.ModTime().Before(di.ModTime()) {
			return
		}
	}

	defer sourcefile.Close()
	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}
	}

	return
}

//CheckIfFolderUpdated checks if a file in the Folders has updated
func CheckIfFolderUpdated(source string, lastCheck time.Time) (bool, error) {
	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
        tempUpdate := false
		if obj.IsDir() {
			tempUpdate, err = CheckIfFolderUpdated(sourcefilepointer, lastCheck)
			if tempUpdate || err != nil {
				return tempUpdate, err
			}
		} else {
			si, err := os.Stat(sourcefilepointer)
			if err != nil {
				return false, err
			}
            
			if si.ModTime().After(lastCheck) {
				return true, nil
			}
		}

    }
    
    return false, nil
}
