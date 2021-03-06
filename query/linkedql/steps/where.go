package steps

import (
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/query/linkedql"
	"github.com/cayleygraph/cayley/query/path"
	"github.com/cayleygraph/quad/voc"
)

func init() {
	linkedql.Register(&Where{})
}

var _ linkedql.PathStep = (*Where)(nil)

// Where corresponds to .where().
type Where struct {
	From  linkedql.PathStep   `json:"from"`
	Steps []linkedql.PathStep `json:"steps"`
}

// Description implements Step.
func (s *Where) Description() string {
	return "applies each provided step in steps in isolation on from"
}

// BuildPath implements linkedql.PathStep.
func (s *Where) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	p := fromPath
	for _, step := range s.Steps {
		stepPath, err := step.BuildPath(qs, ns)
		if err != nil {
			return nil, err
		}
		p = p.And(stepPath.Reverse())
	}
	return p, nil
}
