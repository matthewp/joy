package audionode

import "github.com/matthewmueller/golly/js"

// js:"IIRFilterNode,omit"
type IIRFilterNode struct {
	AudioNode
}

// GetFrequencyResponse
func (*IIRFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	js.Rewrite("$<.GetFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}
