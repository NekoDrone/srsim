package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type AttackStartEventHandler = handler.EventHandler[AttackStartEvent]
type AttackStartEvent struct {
	Attacker     key.TargetID
	Targets      []key.TargetID
	AttackType   model.AttackType
	AttackEffect model.AttackEffect
	DamageType   model.DamageType
}

type AttackEndEventHandler = handler.EventHandler[AttackEndEvent]
type AttackEndEvent struct {
	Attacker     key.TargetID
	Targets      []key.TargetID
	AttackType   model.AttackType
	AttackEffect model.AttackEffect
	DamageType   model.DamageType
}

type BeforeHitEventHandler = handler.EventHandler[BeforeHitEvent]
type BeforeHitEvent struct {
	Attacker key.TargetID
	Defender key.TargetID
	Hit      *info.Hit
}

type DamageResultEventHandler = handler.EventHandler[DamageResultEvent]
type DamageResultEvent struct {
	Attacker        key.TargetID
	Defender        key.TargetID
	AttackType      model.AttackType
	DamageType      model.DamageType
	AttackEffect    model.AttackEffect
	BaseDamage      float64
	BonusDamage     float64
	TotalDamage     float64
	HPDamage        float64
	ShieldDamage    float64
	HealthRemaining float64
	IsCrit          bool
}

type AfterHitEventHandler = handler.EventHandler[AfterHitEvent]
type AfterHitEvent struct {
	Attacker     key.TargetID
	Defender     key.TargetID
	AttackType   model.AttackType
	DamageType   model.DamageType
	AttackEffect model.AttackEffect
	IsCrit       bool
}
