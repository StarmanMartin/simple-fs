# simple-fs

A simple file system tool

## What is simple-fs

Simple-fs is a collection of filesystem methods. It has grown by the challenges of the developer to the file system.

## Install simple-fs

```bash
go get github.com/starmanmaritn/simple-fs
```

```go
import "github.com/starmanmaritn/simple-fs"

func main() {
    fs.CopyFolder("src", "dest")    
}
```

## Methods

* [`SyncFolder`](#SyncFolder)
* [`SyncFolderAndIngonre`](#SyncFolderAndIngonre)
* [`SyncFile`](#SyncFile)
* [`ReadLines`](#ReadLines)
* [`CheckIfFolderUpdated`](#CheckIfFolderUpdated)

### SyncFolder (source, dest string) (err error)

With the SyncFolder method it is possible to copy elements from a *destination* folder to a *source* folder.
SyncFolder does not just copy files and folders. It first checks if the file/folder exists at the *destination*
folder. If the file exists it compares the last changed timestamp on the *destination* and the *source*. Only
if the *source* timestamp is newer the *source* gets copied

#### Parameter

* `source` *string* Absolute path to source folder
* `dest` *string* Absolute path to destination folder

#### return

* `error` nil if success. Else some error

### SyncFolderAndIngonre (source, dest string, ignoreList ...string) (err error)

With the SyncFolderAndIngonre method is the same as the SyncFolder but it is possible to ignore a list of folders.
it is only allowed to add a global folder name as ignored folder. For example *.git* so all *.git* folder get ignored

#### Parameter

* `source` *string* Absolute path to source folder
* `dest` *string* Absolute path to destination folder
* `ignoreList` *...string* A list of ignored folder names (Just global folder names)

#### return

* `error` nil if success. Else some error

### SyncFile (source string, dest string) (err error)

With the SyncFile method it is possible to copy a *source* element to a *destination* folder.
SyncFile first checks if the file/folder exists at the *destination*
folder. If the file exists it compares the last changed timestamp on the *destination* and the *source*. Only
if the *source* timestamp is newer the *source* gets copied

#### Parameter

* `source` *string* Absolute path to source file
* `dest` *string* Absolute path to destination folder

#### return

* `error` nil if success. Else some error

### ReadLines(path string, lineCounts int) ([]string, error)

ReadLines reads a number of lines in a *source* file.

#### Parameter

* `path` *string* Absolute path to source file
* `lineCounts` *int* number of lines to return. *-1* to return whole file

#### return

* `[]string` lines of a flie
* `error` nil if success. Else some error

### CheckIfFolderUpdated(source string, lastCheck time.Time) (bool, error)

CheckIfFolderUpdated checks if a folders last change is newer then a given time.
It compares all sub files and folder to make sure if the folder ist up to date.

#### Parameter

* `source` *string* Absolute path to source folder
* `lastCheck` *time.Time* time to compare last change timestamp of folder

#### return

* `bool` `true` if the folder and all the sub files and sub folders are up to date
* `error` nil if success. Else some error