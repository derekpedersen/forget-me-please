package twitter

import (
	"github.com/derekpedersen/forget-me-please/domain"
	"github.com/derekpedersen/forget-me-please/model"
	log "github.com/sirupsen/logrus"
)

func NewOptions(twt Tweets) domain.Options {
	opt := domain.Options{}
	// TODO: this should be driven by a db or a least a json file
	opt["L"] = model.Option{
		Key:     "L",
		Value:   "Unlike",
		Display: "Un(L)ike",
		Action:  twt.Unlike,
	}
	opt["R"] = model.Option{
		Key:     "R",
		Value:   "Unretweet",
		Display: "Un(R)etweet",
		Action:  twt.UnRetweet,
	}
	opt["D"] = model.Option{
		Key:     "D",
		Value:   "Delete",
		Display: "(D)elete Tweets",
		Action:  twt.Delete,
	}
	log.WithField("TwitterOptions", opt).Debug("NewTwitterOptions")
	return opt
}
