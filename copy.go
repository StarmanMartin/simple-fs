package fs

import (
	"io"
	"os"
	"time"
)

//CopyFolder is deprecated
func CopyFolder(source, dest string) (err error) {
	return SyncFolder(source, dest)
}

//CopyFolderAndIngonre is deprecated
func CopyFolderAndIngonre(source, dest string, ingnoreList ...string) (err error) {
    return CopyFolderAndIngonre(source, dest, ingnoreList...)
}

//SyncFolder method is to copy elements from a *destination* folder to a *source* folder.
//SyncFolder does not just copy files and folders. It first checks if the file/folder exists at the *destination*
//folder. If the file exists it compares the last changed timestamp on the *destination* and the *source*. Only
//if the *source* timestamp is newer the *source* gets copied
//
//Parameter
//
// `source` *string* Absolute path to source folder
// `dest` *string* Absolute path to destination folder
//
// return
//
// `error` nil if success. Else some error
func SyncFolder(source, dest string) (err error) {
	return CopyFolderAndIngonre(source, dest)
}

//SyncFolderAndIngonre method is the same as the SyncFolder but it is possible to ignore a list of folders.
//it is only allowed to add a global folder name as ignored folder. For example *.git* so all *.git* folder get ignored
//SyncFolder method is to copy elements from a *destination* folder to a *source* folder.
//SyncFolder does not just copy files and folders. It first checks if the file/folder exists at the *destination*
//folder. If the file exists it compares the last changed timestamp on the *destination* and the *source*. Only
//if the *source* timestamp is newer the *source* gets copied
//
//Parameter
//
//`source *string* Absolute path to source folder
//`dest` *string* Absolute path to destination folder
//`ignoreList` *...string* A list of ignored folder names (Just global folder names)
//
//return
//`error` nil if success. Else some error
func SyncFolderAndIngonre(source, dest string, ingnoreList ...string) (err error) {
	if sourceinfo, err := os.Stat(source); err != nil {
		return err
	} else if !sourceinfo.IsDir() {
		return SyncFile(source, dest)
	}

	return _copyFolderAndIngonre(source, dest, ingnoreList)
}

func _copyFolderAndIngonre(source, dest string, ingnoreList []string) (err error) {
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
			if !contains(ingnoreList, obj.Name()) {
				err = _copyFolderAndIngonre(sourcefilepointer, destinationfilepointer, ingnoreList)
				if err != nil {
					return
				}
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

//SyncFile method is to copy a *source* element to a *destination* folder.
//SyncFile first checks if the file/folder exists at the *destination*
//folder. If the file exists it compares the last changed timestamp on the *destination* and the *source*. Only
//if the *source* timestamp is newer the *source* gets copied
//
//Parameter
//
//`source` *string* Absolute path to source file
//`dest` *string* Absolute path to destination folder
//
//return
//
//`error` nil if success. Else some error
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

//CheckIfFolderUpdated checks if a folders last change is newer then a given time.
//It compares all sub files and folder to make sure if the folder ist up to date.
//
//Parameter
//
//`source` *string* Absolute path to source folder
//`lastCheck` *time.Time* time to compare last change timestamp of folder
//
//return
//
//`bool` `true` if the folder and all the sub files and sub folders are up to date
//`error` nil if success. Else some error
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
