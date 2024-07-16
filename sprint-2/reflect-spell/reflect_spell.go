package reflect_spell

import "reflect"

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func ChangeFieldByName(v interface{}, fname string, newval int) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}

	field := val.FieldByName(fname)
	if field.IsValid() {
		if field.CanSet() {
			field.SetInt(int64(newval))
		}
	}
}

func CastTo(spell Spell, object interface{}) {
	fieldSelector := 0
	if reflect.ValueOf(object).String() == "main.Player" {
		fieldSelector = 1
	}
	fieldName := reflect.ValueOf(object).Elem().Type().Field(fieldSelector).Name
	fieldValue := reflect.ValueOf(object).Elem().Field(0)
	if spell.Char() == fieldName {
		ChangeFieldByName(object, fieldName, int(fieldValue.Int())+spell.Value())
	}
}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	// nolint: unused
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}
