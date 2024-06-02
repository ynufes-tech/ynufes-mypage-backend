package question

import "errors"

type (
	DimensionSpec struct {
		Min, Max *int
		Eq       *int
	}

	// RatioSpec represents the ratio of width / height
	RatioSpec struct {
		Min, Max *float32
		Eq       *float32
	}
)

func NewDimensionSpec(min, max, eq *int) DimensionSpec {
	if eq != nil {
		return DimensionSpec{Eq: eq}
	}
	return DimensionSpec{Min: min, Max: max, Eq: nil}
}

func NewRatioSpec(min, max, eq *float32) RatioSpec {
	if eq != nil {
		return RatioSpec{Eq: eq}
	}
	return RatioSpec{Min: min, Max: max, Eq: nil}
}

func (s DimensionSpec) Export() map[string]interface{} {
	result := map[string]interface{}{}
	if s.Eq != nil {
		result[FileImageConstraintDimensionEq] = *s.Eq
	} else {
		if s.Min != nil {
			result[FileImageConstraintDimensionMin] = *s.Min
		}
		if s.Max != nil {
			result[FileImageConstraintDimensionMax] = *s.Max
		}
	}
	return result
}

func (s RatioSpec) Export() map[string]interface{} {
	result := map[string]interface{}{}
	if s.Eq != nil {
		result[FileImageConstraintRatioEq] = *s.Eq
	} else {
		if s.Min != nil {
			result[FileImageConstraintRatioMin] = *s.Min
		}
		if s.Max != nil {
			result[FileImageConstraintRatioMax] = *s.Max
		}
	}
	return result
}

func loadDimensionSpec(c map[string]interface{}, key string) (DimensionSpec, error) {
	target := DimensionSpec{}
	t, has := c[key]
	if !has {
		return DimensionSpec{}, nil
	}
	m, ok := t.(map[string]interface{})
	if !ok {
		return DimensionSpec{}, errors.New("invalid dimension spec")
	}
	if eqR, has := m[FileImageConstraintDimensionEq]; has && eqR != nil {
		eqV, ok := eqR.(int)
		if !ok {
			return DimensionSpec{}, errors.New("invalid eq dimension")
		}
		target.Eq = &eqV
	} else {
		if minR, has := m[FileImageConstraintDimensionMin]; has && minR != nil {
			minV, ok := minR.(int)
			if !ok {
				return DimensionSpec{}, errors.New("invalid min dimension")
			}
			target.Min = &minV
		}
		if maxR, has := m[FileImageConstraintDimensionMax]; has && maxR != nil {
			maxV, ok := maxR.(int)
			if !ok {
				return DimensionSpec{}, errors.New("invalid max dimension")
			}
			target.Max = &maxV
		}
	}
	return target, nil
}

func loadRatioSpec(c map[string]interface{}, key string) (RatioSpec, error) {
	target := RatioSpec{}
	t, has := c[key]
	if !has {
		return RatioSpec{}, nil
	}
	m, ok := t.(map[string]interface{})
	if !ok {
		return RatioSpec{}, errors.New("invalid ratio spec")
	}
	if eqR, has := m[FileImageConstraintRatioEq]; has && eqR != nil {
		eqV, ok := eqR.(float32)
		if !ok {
			return RatioSpec{}, errors.New("invalid eq ratio")
		}
		target.Eq = &eqV
	} else {
		if minR, has := m[FileImageConstraintRatioMin]; has && minR != nil {
			minV, ok := minR.(float32)
			if !ok {
				return RatioSpec{}, errors.New("invalid min ratio")
			}
			target.Min = &minV
		}
		if maxR, has := m[FileImageConstraintRatioMax]; has && maxR != nil {
			maxV, ok := maxR.(float32)
			if !ok {
				return RatioSpec{}, errors.New("invalid max ratio")
			}
			target.Max = &maxV
		}
	}
	return target, nil
}
