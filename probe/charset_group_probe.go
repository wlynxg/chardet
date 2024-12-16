package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type CharSetGroupProbe struct {
	CharSetProbe

	filter         consts.LangFilter
	activeNum      int
	bestGuessProbe Probe
	probes         []Probe
}

func NewCharSetGroupProbe(filter consts.LangFilter, probes []Probe) CharSetGroupProbe {
	p := CharSetGroupProbe{
		CharSetProbe:   NewCharSetProbe(filter),
		filter:         filter,
		activeNum:      0,
		bestGuessProbe: nil,
		probes:         probes,
	}
	p.Reset()
	return p
}

func (c *CharSetGroupProbe) Reset() {
	c.CharSetProbe.Reset()
	c.activeNum = 0
	for _, probe := range c.probes {
		if probe != nil {
			probe.Reset()
			probe.SetActive(true)
			c.activeNum++
		}
	}
	c.bestGuessProbe = nil
}

func (c *CharSetGroupProbe) CharSetName() string {
	if c.bestGuessProbe == nil {
		c.GetConfidence()
		if c.bestGuessProbe == nil {
			return ""
		}
	}
	return c.bestGuessProbe.CharSetName()
}

func (c *CharSetGroupProbe) Language() string {
	if c.bestGuessProbe == nil {
		c.GetConfidence()
		if c.bestGuessProbe == nil {
			return ""
		}
	}
	return c.bestGuessProbe.Language()
}

func (c *CharSetGroupProbe) Feed(buf []byte) consts.ProbingState {
	for _, probe := range c.probes {
		if probe == nil {
			continue
		}

		if !probe.IsActive() {
			continue
		}

		state := probe.Feed(buf)
		switch state {
		case consts.FoundItProbingState:
			c.bestGuessProbe = probe
			c.state = consts.FoundItProbingState
			return c.state
		case consts.NotMeProbingState:
			probe.SetActive(false)
			c.activeNum--
			if c.activeNum <= 0 {
				c.state = consts.NotMeProbingState
				return c.state
			}
		default:
		}
	}
	return c.state
}

func (c *CharSetGroupProbe) GetConfidence() float64 {
	state := c.CharSetProbe.State()
	switch state {
	case consts.FoundItProbingState:
		return 0.99
	case consts.NotMeProbingState:
		return 0.01
	default:
	}

	bestConf := 0.0
	c.bestGuessProbe = nil
	for _, probe := range c.probes {
		if probe == nil {
			continue
		}

		if !probe.IsActive() {
			continue
		}

		conf := probe.GetConfidence()
		if bestConf < conf {
			bestConf = conf
			c.bestGuessProbe = probe
		}
	}

	if c.bestGuessProbe == nil {
		return 0.0
	}
	return bestConf
}

func (c *CharSetGroupProbe) Probes() []Probe {
	return c.probes
}
