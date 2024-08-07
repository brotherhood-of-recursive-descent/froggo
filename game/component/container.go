package component

import "reflect"

type Container interface {
	GetComponent(Component) (bool, Component)
	AddComponent(...Component)
	RemoveComponent(Component)
	AllComponents() map[string]Component
}

// DefaultContainer satisfies the Container interface and can be used as a starting point for other entities, by embedding the Base
type DefaultContainer struct {
	components map[string]Component
}

func (container *DefaultContainer) AddComponent(c ...Component) {
	container.components = lazyLoadMap(container.components)

	for _, v := range c {
		key := getType(v)
		container.components[key] = v
	}
}

func (container *DefaultContainer) GetComponent(c Component) (bool, Component) {
	container.components = lazyLoadMap(container.components)

	key := getType(c)
	comp := container.components[key]
	if comp != nil {
		return true, comp
	}
	return false, nil
}

func (container *DefaultContainer) RemoveComponent(c Component) {
	container.components = lazyLoadMap(container.components)

	key := getType(c)
	delete(container.components, key)
}

func (container *DefaultContainer) AllComponents() map[string]Component {
	return container.components
}

func lazyLoadMap(m map[string]Component) map[string]Component {
	if m == nil {
		return make(map[string]Component)
	} else {
		return m
	}
}

func getType(v interface{}) string {
	return reflect.TypeOf(v).String()
}
