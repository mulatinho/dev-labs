package character

import (
	"testing"

	"github.com/mulatinho/golabs/mlt"
)

func TestPlayerCreation(t *testing.T) {
	newPlayer := NewPlayer()
	mlt.Equals(t, newPlayer.HealthStatus(), HEALTH_STATUS_NORMAL)
}

func TestPlayerDamage(t *testing.T) {
	playerOne := NewPlayer()
	mlt.Equals(t, playerOne.Health, HEALTH_STATUS_NORMAL)
	playerOne.ApplyDamage(DAMAGE_TYPE_BASHING)
	mlt.Assert(t, playerOne.HealthStatus() == HEALTH_STATUS_INJURED_LOW)
	playerOne.ApplyDamage(DAMAGE_TYPE_LETHAL)
	mlt.Assert(t, playerOne.Health > HEALTH_STATUS_DEAD)
	playerOne.ApplyDamage(DAMAGE_TYPE_EXPLOSION)
	mlt.Equals(t, playerOne.isDead, true)
	mlt.Equals(t, playerOne.Health, HEALTH_STATUS_DEAD)
}
