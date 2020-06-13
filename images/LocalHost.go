package images

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/reaction-eng/restlib/file"
)

type LocalHost struct {
	tmpDir string

	//Store a map of last time it was updated
	lastUpdate map[string]time.Time

	//We need to download the files
	storage file.Storage
}

func NewLocalHost(storage file.Storage) (*LocalHost, error) {
	//Create a temporary directory
	dir, err := ioutil.TempDir("", "prefix")
	if err != nil {
		return nil, err
	}
	imageHost := &LocalHost{
		tmpDir:     dir,
		lastUpdate: make(map[string]time.Time, 0),
		storage:    storage,
	}

	return imageHost, nil
}

func (host *LocalHost) Location() http.Dir {
	return http.Dir(host.tmpDir)
}

func (host *LocalHost) PrepareImage(imageId string) error {
	//See when the the last time updated
	needsUpdate := false

	//Check if there is a time
	lastUpdate, found := host.lastUpdate[imageId]

	//If not found we need to update
	if !found {
		needsUpdate = true
	} else {
		needsUpdate = lastUpdate.Sub(time.Now()).Hours() > 1
	}

	//Now get the file from google
	if needsUpdate {
		rep, err := host.storage.GetArbitraryFile(imageId)
		if err != nil {
			return err
		}
		defer rep.Close()

		//Open a new file
		out, err := os.Create(host.tmpDir + "/" + imageId)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, rep)
		if err != nil {
			fmt.Println(err)
			return err
		}

		//Store the update time
		host.lastUpdate[imageId] = time.Now()
	}
	return nil
}
