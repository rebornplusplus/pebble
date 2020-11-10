// Copyright (c) 2014-2020 Canonical Ltd
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License version 3 as
// published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package standby

import (
	"time"

	"github.com/canonical/pebble/internal/overlord/state"
)

func (m *StandbyOpinions) SetStartTime(t time.Time) {
	m.startTime = t
}

func FakeStateRequestRestart(newStateRequestRestart func(*state.State, state.RestartType)) (restore func()) {
	oldStateRequestRestart := stateRequestRestart
	stateRequestRestart = newStateRequestRestart
	return func() {
		stateRequestRestart = oldStateRequestRestart
	}
}
