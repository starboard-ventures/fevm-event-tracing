package common

import (
	"context"
	"event-trace/pkg/models/fevm"
	"event-trace/pkg/utils"

	log "github.com/sirupsen/logrus"
)

func UpdateEventHeightCheckoutpoint(ctx context.Context, ec *fevm.EventHeightCheckpoint) error {
	b, err := utils.X.ID(ec.Id).MustCols("txn_type").Update(ec)
	if err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if b == 0 { // there aren't any records in the table
		_, err = utils.X.InsertOne(ec)
		if err != nil {
			log.Errorf("execute sql error: %v", err)
			return err
		}
	}

	return nil
}
