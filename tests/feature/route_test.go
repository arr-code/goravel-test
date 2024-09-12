package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"goravel/tests"
)

type RouteTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestRouteTestSuite(t *testing.T) {
	suite.Run(t, new(RouteTestSuite))
}

func (s *RouteTestSuite) SetupTest() {
	s.RefreshDatabase()
}

func (s *RouteTestSuite) TearDownTest() {
	// Add any necessary cleanup here if required
}

func (s *RouteTestSuite) TestUsers() {
	// s.True(true)
}
