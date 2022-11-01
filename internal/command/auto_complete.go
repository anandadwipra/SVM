package svm_cli

import (
	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"
)

func (data *Options) auto_complete() {
	data.cmd = &complete.Command{
		Sub: map[string]*complete.Command{
			"create": {
				Sub: nil,
				Flags: map[string]complete.Predictor{
					"name":       predict.Something,
					"base_image": predict.Something,
					"network":    predict.Something,
					"cpu":        predict.Something,
					"mem":        predict.Something,
					"debug":      predict.Nothing,
				},
			},
			"delete": {
				Sub: nil,
				Flags: map[string]complete.Predictor{
					"name":                 predict.Something,
					"yes-i-really-mean-it": predict.Nothing,
				},
			},
			"init": {
				Sub: nil,
				Flags: map[string]complete.Predictor{
					"ssh_path": predict.Files("/"),
					"workdir":  predict.Dirs("~/"),
				},
			},
			"profile": {
				Sub:   nil,
				Flags: nil,
			},
			"list": {
				Sub:   nil,
				Flags: nil,
			},
		},
		Flags: map[string]complete.Predictor{
			"test": predict.Something,
		},
	}

}
