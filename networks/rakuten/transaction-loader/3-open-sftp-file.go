package rakuten

import (
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/sftp-client"
	"github.com/halorium/networks-example/state"
)

func openSFTPFile(state *state.State) *flaw.Error {
	logger.Debug("rakuten-read-sftp-file", state)

	sftpClient, flawError := sftpclient.New()

	if flawError != nil {
		return flawError.Wrap("cannot openSFTPFile")
	}

	defer sftpClient.Close()

	fileInfo, err := sftpClient.ReadDir(envvars.SFTPPathRakuten)

	if err != nil {
		return flaw.From(err).Wrap("cannot openSFTPFile")
	}

	for _, item := range fileInfo {
		if item.IsDir() {
			continue // skip directories
		}

		if state.Regexp.MatchString(item.Name()) {
			filePath := envvars.SFTPPathRakuten + item.Name()

			sftpFile, err := sftpClient.Open(filePath)

			if err != nil {
				return flaw.From(err).Wrap("cannot openSFTPFile")
			}

			state.VariantReadCloser = sftpFile

			break
		}
	}

	if state.VariantReadCloser == nil {
		logger.Warn("rakuten-read-sftp-file-not-found", state)

		return nil
	}

	return deserializeXML(state)
}
