// root package for all core logic that powers srsim
package engine

//go:generate mockgen -destination=../mock/mock_engine.go -package=mock github.com/simimpact/srsim/pkg/engine Engine

// only event & info are allowed to be imported from engine here
import (
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Interface that defines how implementations can interact with all engine components. This should
// be thought of as the API definition for all things within the engine. Any implementations will
// be given access to the engine implementation (simulation, outside of the engine package) and call
// these methods to interact.
type Engine interface {
	// Accessor to all event handlers for event subscription and emission
	Events() *event.System

	// Random number generator
	Rand() *rand.Rand

	Modifier
	Attribute
	Combat
	Shield
	Turn
	Validator
	Info
	Target

	// TODO: Execute Queue
	//	For callback + skill methods, need an "AttackState" passed in which allows you to do operations
	//	such as decide when `AttackEnd` happens (if left uncalled, will happen after all logic executes)
	// TODO: Skill Point (Boost Point). Other sim metadata calls?
}

type Modifier interface {
	// Adds a new modifier to the given target. At minimum, instance must specify the name of the
	// modifier and the source.
	AddModifier(target key.TargetID, instance info.Modifier) error

	// Removes all instances of a modifier from the target
	RemoveModifier(target key.TargetID, modifier key.Modifier)

	// Removes all instances of a modifier from the target, only where source matches the given source
	RemoveModifierFromSource(target, source key.TargetID, modifier key.Modifier)

	// Extends the duration of all instances of the given modifier by the amount
	ExtendModifierDuration(target key.TargetID, modifier key.Modifier, amt int)

	// Extends the count of all instances of the given modifier by the amount. Will not extend past
	// the modifiers MaxCount
	ExtendModifierCount(target key.TargetID, modifier key.Modifier, amt float64)

	// Returns true if the target has at least one instance of the modifier
	HasModifier(target key.TargetID, modifier key.Modifier) bool

	// Returns the total count of modifiers that are of the given StatusType (Buff or Debuff)
	ModifierCount(target key.TargetID, statusType model.StatusType) int

	// Returns true if the target has the given behavior flag from an attached modifier. If multiple
	// flags are passed, will return true if at least one is attached
	HasBehaviorFlag(target key.TargetID, flags ...model.BehaviorFlag) bool
}

type Attribute interface {
	// Gets a snapshot of the current target's stats. Any modifications to these stats will
	// only be applied to the snapshot.
	Stats(target key.TargetID) *info.Stats

	// TODO: Energy, HP, Stance

	// Adds energy to the given target using the specified EnergyAdd logic
	AddEnergy(target key.TargetID, addType model.EnergyAdd, amt float64)
}

type Combat interface {
	// Performs the given attack where Source is the attacker and Targets are all targets that
	// are being hit
	Attack(atk info.Attack)
	Heal()
}

type Shield interface {
	AddShield()
	RemoveShield()
}

type Turn interface {
	ModifyGauge(target key.TargetID, modifyType model.ModifyGauge, amt float64)
	SetGauge(target key.TargetID, amt float64)
	// TODO: need ModifyCurrentSkillDelayCost? (in dm used to modify gauge for next turn, during current turn)
}

type Validator interface {
	// Check if the given TargetID is valid
	IsValid(target key.TargetID) bool
}

type Info interface {
	// Metadata for the given character, such as their current level, ascension, traces, etc.
	CharacterInfo(target key.TargetID) (info.Character, error)

	// Metadata for the given enemy, such as their current level and weaknesses.
	EnemyInfo(target key.TargetID) (model.Enemy, error)
}

type Target interface {
	// returns true if the given TargetID is for a character
	IsCharacter(target key.TargetID) bool

	// returns true if the given TargetID is for an enemy
	IsEnemy(target key.TargetID) bool

	// returns the ids of targets that are adjacent to the given targent (empty if there are none)
	AdjacentTo(target key.TargetID) []key.TargetID

	// returns a list of all character target ids
	Characters() []key.TargetID

	// returns a list of all enemy target ids
	Enemies() []key.TargetID

	// returns a list of all neutral target ids (these are special cases, such as the Lightning-Lord)
	Neutrals() []key.TargetID

	// TODO: target type, (Light, Dark, Neutral)
	AddTarget() key.TargetID
}
