package character

import (
	"github.com/mulatinho/golabs/littlegame/utils"
)

const (
	PACKAGE_NAME = "character"
	LOOP_MSG     = "loop is running"
)

type (
	HealthStatusType int
	DamageType       int
)

const (
	HEALTH_STATUS_DEAD HealthStatusType = iota
	HEALTH_STATUS_UNCONSCIOUS_SEVERE
	HEALTH_STATUS_UNCONSCIOUS_HIGH
	HEALTH_STATUS_UNCONSCIOUS_MEDIUM
	HEALTH_STATUS_UNCONSCIOUS_LOW
	HEALTH_STATUS_INCAPACITATED_HIGH
	HEALTH_STATUS_INCAPACITATED_MEDIUM
	HEALTH_STATUS_INCAPACITATED_LOW
	HEALTH_STATUS_DAMAGED_HIGH
	HEALTH_STATUS_DAMAGED_MEDIUM
	HEALTH_STATUS_DAMAGED_LOW
	HEALTH_STATUS_INJURED_HIGH
	HEALTH_STATUS_INJURED_MEDIUM
	HEALTH_STATUS_INJURED_LOW
	HEALTH_STATUS_NORMAL
	HEALTH_STATUS_TOTAL
)

const (
	DAMAGE_TYPE_BASHING DamageType = 1
	DAMAGE_TYPE_LETHAL
	DAMAGE_TYPE_AGGRAVATE
	DAMAGE_TYPE_HEADSHOT  = DamageType(HEALTH_STATUS_TOTAL)
	DAMAGE_TYPE_EXPLOSION = DamageType(HEALTH_STATUS_TOTAL) * 100
)

type Player struct {
	Id     int
	Name   string
	Health HealthStatusType
	isDead bool
}

func NewPlayer() *Player {
	return &Player{
		Id:     utils.GetPlayerID(),
		Name:   utils.GenerateName(utils.NAME_TYPE_MAN),
		Health: HEALTH_STATUS_NORMAL,
		isDead: false,
	}
}

func (p *Player) HealthStatus() HealthStatusType {
	return p.Health
}

func (p *Player) ApplyDamage(damageType DamageType) HealthStatusType {
	p.Health -= HealthStatusType(damageType)
	if p.Health <= 0 {
		p.Health = HEALTH_STATUS_DEAD
		p.isDead = true
	}

	utils.LogMessage(utils.LOG_TYPE_DEBUG, p)
	return p.Health
}
