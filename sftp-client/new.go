package sftpclient

import (
	"os"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/halt"
	"github.com/halorium/sftp"
	"golang.org/x/crypto/ssh"
)

var inSFTPHost string
var inSFTPKey string
var inSFTPPort string
var inSFTPUser string

func init() {
	inSFTPHost = os.Getenv("IN_SFTP_HOST")

	if inSFTPHost == "" {
		halt.PanicWith("IN_SFTP_HOST environment variable must be set")
	}

	inSFTPKey = os.Getenv("IN_SFTP_KEY")

	if inSFTPKey == "" {
		halt.PanicWith("IN_SFTP_KEY environment variable must be set")
	}

	inSFTPPort = os.Getenv("IN_SFTP_PORT")

	if inSFTPPort == "" {
		halt.PanicWith("IN_SFTP_PORT environment variable must be set")
	}

	inSFTPUser = os.Getenv("IN_SFTP_USER")

	if inSFTPUser == "" {
		halt.PanicWith("IN_SFTP_USER environment variable must be set")
	}
}

type SFTPClient struct {
	*sftp.Client
}

func New() (*SFTPClient, *flaw.Error) {
	buf := []byte(inSFTPKey)
	signer, err := ssh.ParsePrivateKey(buf)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot New")
	}

	sshClientConfig := &ssh.ClientConfig{
		User: inSFTPUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	hostAndPort := inSFTPHost + ":" + inSFTPPort

	sshClient, err := ssh.Dial("tcp", hostAndPort, sshClientConfig)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot New")
	}

	sftpClient, err := sftp.NewClient(sshClient)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot New")
	}

	return &SFTPClient{sftpClient}, nil
}
