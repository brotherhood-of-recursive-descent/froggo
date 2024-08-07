package component

import "testing"

func TestLazyLoadMap_initializedNilMap(t *testing.T) {

	var notInitializedMap map[string]Component
	initializedMap := make(map[string]Component)

	notInitializedMap = lazyLoadMap(notInitializedMap)
	if notInitializedMap == nil {
		t.Errorf("expected map to be initialized, got %v\n", notInitializedMap)
	}

	initializedMap = lazyLoadMap(initializedMap)
	if initializedMap == nil {
		t.Errorf("expected map to be initialized, got %v\n", initializedMap)
	}

}

func TestLazyLoadMap_retainsMap(t *testing.T) {

	initializedMap := make(map[string]Component)
	initializedMap["test"] = nil

	initializedMap = lazyLoadMap(initializedMap)
	if _, ok := initializedMap["test"]; !ok {
		t.Errorf("expected map to be not changed, got %v\n", initializedMap)
	}

}
