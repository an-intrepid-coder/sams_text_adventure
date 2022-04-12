package main

type Combat struct {
    // Owner is the Actor which this component belongs to:
    Owner *Actor
    // Engaged is the Actor being attacked by this one:
    Engaged *Actor
    // Health is HP/MaxHP:
    Health int
    MaxHealth int
    // NOTE: Will add MP/MaxMP as well later.
    // Attack is a bonus to D20 roll to hit:
    Attack int
    MinDamage int
    MaxDamage int
    // Armor is dmg mitigation:
    Armor int
    // Defense is vs. a D20 roll to hit:
    Defense int
}

