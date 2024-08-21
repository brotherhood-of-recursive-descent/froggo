package lib

import (
	"errors"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	GetComponent(Component) (bool, Component)
	AddComponent(...Component)
	RemoveComponent(Component)
	AllComponents() map[string]Component

	Draw(*ebiten.Image)
	Update() error
}

// BaseEntity implements the Entity interface
type BaseEntity struct {
	components map[string]Component
}

func (e *BaseEntity) AddComponent(c ...Component) {
	e.components = lazyLoadMap(e.components)

	for _, v := range c {
		key := getType(v)
		e.components[key] = v
	}
}

func (e *BaseEntity) GetComponent(c Component) (bool, Component) {
	e.components = lazyLoadMap(e.components)

	key := getType(c)
	comp := e.components[key]
	if comp != nil {
		return true, comp
	}
	return false, nil
}

func (e *BaseEntity) RemoveComponent(c Component) {
	e.components = lazyLoadMap(e.components)

	key := getType(c)
	delete(e.components, key)
}

func (e *BaseEntity) AllComponents() map[string]Component {
	return e.components
}

func lazyLoadMap(m map[string]Component) map[string]Component {
	if m == nil {
		return make(map[string]Component)
	} else {
		return m
	}
}

func (e *BaseEntity) Update() error {

	for k, c := range e.AllComponents() {
		if c.IsActive() {
			if err := c.Update(); err != nil {
				return errors.Join(errors.New("failed calling Update() on: '"+k), err)
			}
		}
	}

	return nil
}

func (e *BaseEntity) Draw(screen *ebiten.Image) {

	for _, c := range e.AllComponents() {
		if c.IsActive() {
			c.Draw(screen)
		}
	}
}

func getType(v interface{}) string {
	return reflect.TypeOf(v).String()
}
