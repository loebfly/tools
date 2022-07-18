package sftpT

import (
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"os"
	"path"
	"time"
)

// Connect to the remote server and return the sftp client and ssh client
func (Enter) Connect(user, password, host string, port int) (*sftp.Client, *ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connect to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, nil, errors.New("sftp ssh connect error: " + err.Error())
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, nil, errors.New("sftp create failed: " + err.Error())
	}

	return sftpClient, sshClient, nil
}

// Close the sftp client and ssh client
func (Enter) Close(sftpClient *sftp.Client, sshClient *ssh.Client) error {
	err := sftpClient.Close()
	if err != nil {
		return errors.New("sftp close error: " + err.Error())
	}
	err = sshClient.Close()
	if err != nil {
		return errors.New("ssh close error: " + err.Error())
	}
	return nil
}

// UploadFile upload the local file to the remote server
func (Enter) UploadFile(sftpClient *sftp.Client, localFilePath string, remotePath string) error {
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		return errors.New("local file open error: " + err.Error())
	}
	defer func(srcFile *os.File) {
		e := srcFile.Close()
		if e != nil {
			fmt.Println("local file close error: " + e.Error())
		}
	}(srcFile)

	var remoteFileName = path.Base(localFilePath)

	dstFile, err := sftpClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		return errors.New("remote file create error: " + err.Error())

	}
	defer func(dstFile *sftp.File) {
		e := dstFile.Close()
		if e != nil {
			fmt.Println("remote file close error: " + e.Error())
		}
	}(dstFile)

	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		return errors.New("local file read error: " + err.Error())
	}
	_, err = dstFile.Write(ff)
	if err != nil {
		return err
	}
	return nil
}
