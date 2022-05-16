package feature

import "github.com/yut-kt/gomfcc"

func GetMFCC(samples []float64, sampleRate int) [][]float64 {
	mfcc := gomfcc.NewGoMFCC(samples, sampleRate)
	return mfcc.GetFeatureByMS(39, 25, 25, 10)
}
