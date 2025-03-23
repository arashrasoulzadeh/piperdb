// Package persist /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package persist

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/arashrasoulzadeh/piperdb/src/config"
	"go.uber.org/zap"
	"io"
	"os"
	"path"
)

func saveToFile(filename string, data []byte) error {
	filename = config.Current().Instance.Path + "/" + filename
	err := createDirsIfNotExists(filename)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zap.L().Error("failed to close file", zap.Error(err))
		}
	}(file)

	if config.Current().Instance.Encode {
		encoder := base64.NewEncoder(base64.StdEncoding, file)
		defer func(encoder io.WriteCloser) {
			err := encoder.Close()
			if err != nil {
				zap.L().Error("failed to close file", zap.Error(err))
			}
		}(encoder)

		_, err = encoder.Write(data)
	} else {
		writer := bufio.NewWriter(file)
		_, err = writer.Write(data)
		if err != nil {
			return err
		}
		return writer.Flush()
	}

	return err
}

func loadFromFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zap.L().Error("failed to close file", zap.Error(err))
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func createDirsIfNotExists(filePath string) error {
	dir := path.Dir(filePath)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directories: %v", err)
	}

	_, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	return nil
}
