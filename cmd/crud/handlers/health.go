package handlers

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/internal/platform/db"
	"github.com/ardanlabs/service/internal/platform/web"
	"github.com/pkg/errors"
)

// Health represents the User API method handler set.
type Health struct {
	MasterDB *db.DB
}

// Check validates the service is ready and healthy to accept requests.
func (h *Health) Check(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer dbConn.Close()

	if err := dbConn.StatusCheck(); err != nil {
		return err
	}

	web.Respond(ctx, w, []byte("ok"), http.StatusOK)
	return nil
}
