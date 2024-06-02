// Package suman - SUSE Manager support functions
package suman

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_sumanUseCase "ecp-golang-cm/pkg/usecases/susemanager"
	returnCodes "ecp-golang-cm/pkg/util/returnCodes"
	"github.com/pkg/errors"
)

// CheckIfSumaServer - check if server is a SUSE Manager Server
//
// return:
func CheckIfSumaServer() bool {
	_, err := os.Stat("/etc/rhn/rhn.conf")
	return err == nil
}

// GetCredentials - get credentiales needed for SUSE Manager Server
//
// param: fileName
// return:
func GetCredentials(fileName string) (_sumanUseCase.SumanConfig, error) {
	var sumancfg _sumanUseCase.SumanConfig
	sumancfg.Insecure = true
	fileName = filepath.Clean(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return sumancfg, errors.New(returnCodes.ErrOpeningFile)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), "=")
		if l[0] == "username" {
			sumancfg.Login = l[1]
		}
		if l[0] == "password" {
			sumancfg.Password = l[1]
		}
	}
	cmd := exec.Command("/bin/hostname", "-f")
	var out, stdErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stdErr
	err = cmd.Run()
	if err != nil {
		return sumancfg, err
	}
	if len(stdErr.String()) != 0 {
		return sumancfg, errors.New(stdErr.String())
	}
	if len(out.String()) == 0 {
		return sumancfg, errors.New("error retrieving the FQDN of the server")
	}
	sumancfg.Host = out.String()
	sumancfg.Host = sumancfg.Host[:len(sumancfg.Host)-1]
	return sumancfg, nil
}

func GetCredentialsUyuni(fileName string) (_sumanUseCase.SumanConfig, error) {
	var sumancfg _sumanUseCase.SumanConfig
	sumancfg.Insecure = true
	fileName = filepath.Clean(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return sumancfg, errors.New(returnCodes.ErrOpeningFile)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), ":")
		if strings.TrimSpace(l[0]) == "user" {
			sumancfg.Login = strings.TrimSpace(l[1])
		}
		if strings.TrimSpace(l[0]) == "password" {
			sumancfg.Password = strings.TrimSpace(l[1])
		}
		if strings.TrimSpace(l[0]) == "hubmaster" {
			sumancfg.Host = strings.TrimSpace(l[1])
		}
	}
	return sumancfg, nil
}
