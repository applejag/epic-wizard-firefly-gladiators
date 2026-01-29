package racebattle

import (
	"github.com/applejag/firefly-go-math/ffmath"
)

var path = Path{
	ffmath.V(40, 414),
	ffmath.V(42, 434),
	ffmath.V(56, 451),
	ffmath.V(125, 504),
	ffmath.V(194, 573),
	ffmath.V(257, 588),
	ffmath.V(337, 583),
	ffmath.V(428, 575),
	ffmath.V(438, 541),
	ffmath.V(434, 513),
	ffmath.V(405, 456),
	ffmath.V(420, 415),
	ffmath.V(465, 410),
	ffmath.V(532, 365),
	ffmath.V(566, 294),
	ffmath.V(547, 247),
	ffmath.V(512, 212),
	ffmath.V(470, 195),
	ffmath.V(438, 110),
	ffmath.V(395, 90),
	ffmath.V(349, 114),
	ffmath.V(321, 192),
	ffmath.V(288, 195),
	ffmath.V(284, 241),
	ffmath.V(245, 268),
	ffmath.V(207, 261),
	ffmath.V(198, 230),
	ffmath.V(215, 183),
	ffmath.V(243, 150),
	ffmath.V(245, 91),
	ffmath.V(234, 69),
	ffmath.V(188, 62),
	ffmath.V(143, 68),
	ffmath.V(115, 84),
	ffmath.V(101, 122),
	ffmath.V(69, 138),
	ffmath.V(57, 165),
	ffmath.V(56, 199),
	ffmath.V(85, 232),
	ffmath.V(82, 271),
	ffmath.V(67, 292),
	ffmath.V(74, 312),
	ffmath.V(82, 331),
	ffmath.V(75, 350),
	ffmath.V(41, 361),
	ffmath.V(36, 380),
	ffmath.V(40, 405),
}

type Path []ffmath.Vec

type PathTracker struct {
	path     Path
	index    int
	previous ffmath.Vec
	current  ffmath.Vec
	next     ffmath.Vec
}

func NewPathTracker(path Path) PathTracker {
	return PathTracker{
		path:     path,
		previous: path[len(path)-1],
		current:  path[0],
		next:     path[1],
	}
}

func (p PathTracker) PeekPrevious() ffmath.Vec {
	return p.previous
}

func (p PathTracker) PeekCurrent() ffmath.Vec {
	return p.current
}

func (p PathTracker) PeekNext() ffmath.Vec {
	return p.next
}

func (p *PathTracker) goNext() {
	p.previous = p.current
	p.index = (p.index + 1) % len(p.path)
	p.current = p.path[p.index]
	p.next = p.path[(p.index+1)%len(p.path)]
}

func (p *PathTracker) PeekSoftNext(currentPos ffmath.Vec) ffmath.Vec {
	currentTarget := p.PeekCurrent()
	nextTarget := p.PeekNext()
	prevTarget := p.PeekPrevious()

	// TODO: this implementation is a little buggy
	// it's not smooth at all when it switches checkpoints.
	// Maybe if we checked the distance to a point that's projected on a line
	// that's perpendicular with the prev->current line and that crosses the current target.
	distSqToCurrent := currentTarget.Sub(currentPos).RadiusSquared()
	distSqFromPrev := currentTarget.Sub(prevTarget).RadiusSquared()
	distWeight := 1 - min(distSqToCurrent/distSqFromPrev, 1)

	return ffmath.V(
		ffmath.Lerp(currentTarget.X, nextTarget.X, distWeight),
		ffmath.Lerp(currentTarget.Y, nextTarget.Y, distWeight),
	)
}

// Progress returns the percentage (0.0-1.0) of progress made throughout the
// path. The "pos" is used to calculate fractional progress between checkpoints.
func (p *PathTracker) Progress(pos ffmath.Vec) float32 {
	prev := p.PeekPrevious()
	current := p.PeekCurrent()

	distSqToCurrent := current.Sub(pos).RadiusSquared()
	distSqFromPrev := current.Sub(prev).RadiusSquared()
	distWeight := 1 - min(distSqToCurrent/distSqFromPrev, 1)

	return float32(p.index+1)/float32(len(p.path)) + distWeight/float32(len(p.path))
}

func (p *PathTracker) Update(pos ffmath.Vec) PathTrackerResult {
	curr := p.PeekCurrent()
	prev := p.PeekPrevious()
	distSqToCurr := curr.Sub(pos).RadiusSquared()
	distSqToPrev := pos.Sub(prev).RadiusSquared()
	distSqBetweenPoints := curr.Sub(prev).RadiusSquared()
	switch {
	case distSqToPrev < distSqBetweenPoints:
		// haven't gotten far enough away from previous point
		return PathTrackerKeepCurrent
	case distSqToPrev < distSqToCurr:
		// moving backwards
		return PathTrackerMovingBackwards
	}
	p.goNext()
	if p.index == 0 {
		return PathTrackerLooped
	}
	return PathTrackerNextCheckpoint
}

type PathTrackerResult byte

const (
	PathTrackerKeepCurrent PathTrackerResult = iota
	PathTrackerMovingBackwards
	PathTrackerNextCheckpoint
	PathTrackerLooped
)
