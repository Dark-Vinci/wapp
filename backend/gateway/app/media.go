package app

import (
	"context"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"

	"github.com/dark-vinci/wapp/backend/gateway/model"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

// UploadSingleMedia this will call media endpoint on media server
func (a *App) UploadSingleMedia(ctx context.Context, userID uuid.UUID, file *multipart.FileHeader) (*model.FileContent, error) {
	logger := a.logger.With().
		Str(constants.MethodStrHelper, packageName).
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	f := model.FileContent{
		Key:      uuid.New().String(),
		Name:     file.Filename,
		Size:     file.Size,
		FileType: path.Ext(file.Filename),
		Headers:  make(map[string]string),
		UserID:   userID,
	}

	for j, val := range file.Header {
		f.Headers[j] = strings.Join(val, "|>")
	}

	opFile, err := file.Open()
	if err != nil {
		logger.Err(err).Msg("AppFailure: unable to open file")
		return nil, err
	}

	osFile, ok := opFile.(*os.File)
	if !ok {
		logger.Warn().Msg("AppFailure: unable to cast to os.File")
		return nil, err
	}

	defer func() {
		_ = osFile.Close()
	}()

	up, err := a.ss3.Upload(osFile, f.Key, "uploads")
	if err != nil {
		logger.Err(err).Msg("AppFailure: unable to upload file to file storage")
		return nil, err
	}

	f.UploadID = *up

	return &f, nil
}

func (a *App) UploadMultipleMedia(ctx context.Context, files []*multipart.FileHeader) ([]model.FileContent, error) {
	logger := a.logger.With().
		Str(constants.MethodStrHelper, packageName).
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	fileResult := make([]model.FileContent, 0)

	for _, file := range files {
		f := model.FileContent{
			Name:     file.Filename,
			Size:     file.Size,
			FileType: path.Ext(file.Filename),
			Headers:  make(map[string]string),
		}

		for K, val := range file.Header {
			f.Headers[K] = strings.Join(val, "|>")
		}

		opFile, err := file.Open()
		if err != nil {
			logger.Err(err).Msg("AppFailure: unable to open file")
			return nil, err
		}

		osFile, ok := opFile.(*os.File)
		if !ok {
			logger.Warn().Msg("AppFailure: unable to cast to os.File")
			return nil, err
		}

		up, err := a.ss3.Upload(osFile, f.Key, "uploads")
		if err != nil {
			logger.Err(err).Msg("AppFailure: unable to upload file to file storage")
			return nil, err
		}

		f.UploadID = *up

		_ = osFile.Close()

		fileResult = append(fileResult, f)
	}

	return fileResult, nil
}
