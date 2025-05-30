package operations

import (
	"errors"
	"net/http"

	"github.com/treeverse/lakefs/pkg/block"
	gatewayerrors "github.com/treeverse/lakefs/pkg/gateway/errors"
	"github.com/treeverse/lakefs/pkg/graveler"
	"github.com/treeverse/lakefs/pkg/logging"
	"github.com/treeverse/lakefs/pkg/permissions"
)

type DeleteObject struct{}

func (controller *DeleteObject) RequiredPermissions(_ *http.Request, repoID, _, path string) (permissions.Node, error) {
	return permissions.Node{
		Permission: permissions.Permission{
			Action:   permissions.DeleteObjectAction,
			Resource: permissions.ObjectArn(repoID, path),
		},
	}, nil
}

func (controller *DeleteObject) HandleAbortMultipartUpload(w http.ResponseWriter, req *http.Request, o *PathOperation) {
	o.Incr("abort_mpu", o.Principal, o.Repository.Name, o.Reference)
	query := req.URL.Query()
	uploadID := query.Get(QueryParamUploadID)

	ctx := req.Context()
	mpu, err := o.MultipartTracker.Get(ctx, uploadID)
	if err != nil {
		o.Log(req).WithError(err).Error("upload id not found in tracker")
		_ = o.EncodeError(w, req, err, gatewayerrors.Codes.ToAPIErr(gatewayerrors.ErrNoSuchKey))
		return
	}
	if mpu.Path != o.Path {
		o.Log(req).Error("could not match multipart upload with multipart tracker record")
		_ = o.EncodeError(w, req, err, gatewayerrors.Codes.ToAPIErr(gatewayerrors.ErrNoSuchKey))
		return
	}

	req = req.WithContext(logging.AddFields(ctx, logging.Fields{logging.UploadIDFieldKey: uploadID}))
	err = o.BlockStore.AbortMultiPartUpload(ctx, block.ObjectPointer{
		StorageID:        o.Repository.StorageID,
		StorageNamespace: o.Repository.StorageNamespace,
		IdentifierType:   block.IdentifierTypeRelative,
		Identifier:       mpu.PhysicalAddress,
	}, uploadID)
	if err != nil {
		o.Log(req).WithError(err).Error("could not abort multipart upload")
		_ = o.EncodeError(w, req, err, gatewayerrors.Codes.ToAPIErr(gatewayerrors.ErrInternalError))
		return
	}

	if err := o.MultipartTracker.Delete(ctx, uploadID); err != nil {
		o.Log(req).WithError(err).Warn("could not delete multipart record")
	}

	w.WriteHeader(http.StatusNoContent)
}

func (controller *DeleteObject) Handle(w http.ResponseWriter, req *http.Request, o *PathOperation) {
	if o.HandleUnsupported(w, req, "tagging", "acl", "torrent") {
		return
	}
	if o.Repository.ReadOnly {
		_ = o.EncodeError(w, req, nil, gatewayerrors.Codes.ToAPIErr(gatewayerrors.ErrReadOnlyRepository))
		return
	}
	query := req.URL.Query()
	if query.Has(QueryParamUploadID) {
		controller.HandleAbortMultipartUpload(w, req, o)
		return
	}

	o.Incr("delete_object", o.Principal, o.Repository.Name, o.Reference)
	lg := o.Log(req).WithField("key", o.Path)
	err := o.Catalog.DeleteEntry(req.Context(), o.Repository.Name, o.Reference, o.Path)
	switch {
	case errors.Is(err, graveler.ErrNotFound):
		lg.WithError(err).Debug("could not delete object, it doesn't exist")
	case errors.Is(err, graveler.ErrWriteToProtectedBranch):
		_ = o.EncodeError(w, req, err, gatewayerrors.Codes.ToAPIErr(gatewayerrors.ErrWriteToProtectedBranch))
		return
	case errors.Is(err, graveler.ErrReadOnlyRepository):
		_ = o.EncodeError(w, req, err, gatewayerrors.Codes.ToAPIErr(gatewayerrors.ErrReadOnlyRepository))
		return
	case err != nil:
		lg.WithError(err).Error("could not delete object")
		_ = o.EncodeError(w, req, err, gatewayerrors.Codes.ToAPIErr(gatewayerrors.ErrInternalError))
		return
	default:
		lg.Debug("object set for deletion")
	}
	w.WriteHeader(http.StatusNoContent)
}
