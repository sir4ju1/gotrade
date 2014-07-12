package indicators

import (
	"github.com/thetruetrade/gotrade"
)

type TSF struct {
	*LinearRegWithoutStorage

	// public variables
	Data []float64
}

func NewTSF(timePeriod int, selectData gotrade.DataSelectionFunc) (indicator *TSF, err error) {
	newInd := TSF{}
	newInd.LinearRegWithoutStorage, err = NewLinearRegWithoutStorage(timePeriod, selectData,
		func(dataItem float64, slope float64, intercept float64, streamBarIndex int) {
			result := intercept + slope*float64(timePeriod)

			if result > newInd.LinearRegWithoutStorage.maxValue {
				newInd.LinearRegWithoutStorage.maxValue = result
			}

			if result < newInd.LinearRegWithoutStorage.minValue {
				newInd.LinearRegWithoutStorage.minValue = result
			}

			newInd.Data = append(newInd.Data, result)
		})

	return &newInd, err
}

func NewTSFForStream(priceStream *gotrade.DOHLCVStream, timePeriod int, selectData gotrade.DataSelectionFunc) (indicator *TSF, err error) {
	newInd, err := NewTSF(timePeriod, selectData)
	priceStream.AddTickSubscription(newInd)
	return newInd, err
}