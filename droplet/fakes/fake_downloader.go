// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/krujos/download_droplet_plugin/droplet"
)

type FakeDownloader struct {
	GetDropletStub        func(guid string) ([]byte, error)
	getDropletMutex       sync.RWMutex
	getDropletArgsForCall []struct {
		guid string
	}
	getDropletReturns struct {
		result1 []byte
		result2 error
	}
	SaveDropletToFileStub        func(filePath string, data []byte) error
	saveDropletToFileMutex       sync.RWMutex
	saveDropletToFileArgsForCall []struct {
		filePath string
		data     []byte
	}
	saveDropletToFileReturns struct {
		result1 error
	}
}

func (fake *FakeDownloader) GetDroplet(guid string) ([]byte, error) {
	fake.getDropletMutex.Lock()
	fake.getDropletArgsForCall = append(fake.getDropletArgsForCall, struct {
		guid string
	}{guid})
	fake.getDropletMutex.Unlock()
	if fake.GetDropletStub != nil {
		return fake.GetDropletStub(guid)
	} else {
		return fake.getDropletReturns.result1, fake.getDropletReturns.result2
	}
}

func (fake *FakeDownloader) GetDropletCallCount() int {
	fake.getDropletMutex.RLock()
	defer fake.getDropletMutex.RUnlock()
	return len(fake.getDropletArgsForCall)
}

func (fake *FakeDownloader) GetDropletArgsForCall(i int) string {
	fake.getDropletMutex.RLock()
	defer fake.getDropletMutex.RUnlock()
	return fake.getDropletArgsForCall[i].guid
}

func (fake *FakeDownloader) GetDropletReturns(result1 []byte, result2 error) {
	fake.GetDropletStub = nil
	fake.getDropletReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloader) SaveDropletToFile(filePath string, data []byte) error {
	fake.saveDropletToFileMutex.Lock()
	fake.saveDropletToFileArgsForCall = append(fake.saveDropletToFileArgsForCall, struct {
		filePath string
		data     []byte
	}{filePath, data})
	fake.saveDropletToFileMutex.Unlock()
	if fake.SaveDropletToFileStub != nil {
		return fake.SaveDropletToFileStub(filePath, data)
	} else {
		return fake.saveDropletToFileReturns.result1
	}
}

func (fake *FakeDownloader) SaveDropletToFileCallCount() int {
	fake.saveDropletToFileMutex.RLock()
	defer fake.saveDropletToFileMutex.RUnlock()
	return len(fake.saveDropletToFileArgsForCall)
}

func (fake *FakeDownloader) SaveDropletToFileArgsForCall(i int) (string, []byte) {
	fake.saveDropletToFileMutex.RLock()
	defer fake.saveDropletToFileMutex.RUnlock()
	return fake.saveDropletToFileArgsForCall[i].filePath, fake.saveDropletToFileArgsForCall[i].data
}

func (fake *FakeDownloader) SaveDropletToFileReturns(result1 error) {
	fake.SaveDropletToFileStub = nil
	fake.saveDropletToFileReturns = struct {
		result1 error
	}{result1}
}

var _ droplet.Downloader = new(FakeDownloader)
