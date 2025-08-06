package interp

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetEasingFunction(easingFn Easing) EasingFunction {
	switch easingFn {
	default:
		return Linear
	case EASING_IN_QUAD:
		return InQuad
	case EASING_OUT_QUAD:
		return OutQuad
	case EASING_IN_OUT_QUAD:
		return InOutQuad
	case EASING_IN_CUBIC:
		return InCubic
	case EASING_OUT_CUBIC:
		return OutCubic
	case EASING_IN_OUT_CUBIC:
		return InOutCubic
	case EASING_IN_QUART:
		return InQuart
	case EASING_OUT_QUART:
		return OutQuart
	case EASING_IN_OUT_QUART:
		return InOutQuart
	case EASING_IN_QUINT:
		return InQuint
	case EASING_OUT_QUINT:
		return OutQuint
	case EASING_IN_OUT_QUINT:
		return InOutQuint
	case EASING_IN_SINE:
		return InSine
	case EASING_OUT_SINE:
		return OutSine
	case EASING_IN_OUT_SINE:
		return InOutSine
	case EASING_IN_EXPO:
		return InExpo
	case EASING_OUT_EXPO:
		return OutExpo
	case EASING_IN_OUT_EXPO:
		return InOutExpo
	case EASING_IN_CIRC:
		return InCirc
	case EASING_OUT_CIRC:
		return OutCirc
	case EASING_IN_OUT_CIRC:
		return InOutCirc
	case EASING_IN_ELASTIC:
		return InElastic
	case EASING_OUT_ELASTIC:
		return OutElastic
	case EASING_IN_OUT_ELASTIC:
		return InOutElastic
	case EASING_IN_BACK:
		return InBack
	case EASING_OUT_BACK:
		return OutBack
	case EASING_IN_OUT_BACK:
		return InOutBack
	case EASING_IN_BOUNCE:
		return InBounce
	case EASING_OUT_BOUNCE:
		return OutBounce
	case EASING_IN_OUT_BOUNCE:
		return InOutBounce
	case EASING_IN_SQUARE:
		return InSquare
	case EASING_OUT_SQUARE:
		return OutSquare
	case EASING_IN_OUT_SQUARE:
		return InOutSquare
	}
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type Value[T Number] struct {
	start      T
	end        T
	startTime  float64
	speed      float64
	easingFunc EasingFunction
}

func New[T Number](inittal_value T) *Value[T] {
	return &Value[T]{
		start:      inittal_value,
		end:        inittal_value,
		easingFunc: Linear,
	}
}

func (v *Value[T]) getElapsedTime() float64 {
	return rl.GetTime() - v.startTime
}

func (v *Value[T]) SetTransition(easingFunc EasingFunction) *Value[T] {
	v.easingFunc = easingFunc
	return v
}

func (v *Value[T]) SetDuration(duration float64) *Value[T] {
	v.speed = 1.0 / duration
	return v
}

func (v *Value[T]) SetValue(new_value T) *Value[T] {
	v.start = v.GetValue()
	v.end = new_value
	v.startTime = rl.GetTime()
	return v
}

func (v *Value[T]) GetValue() T {
	elapsed := v.getElapsedTime()
	t := elapsed * v.speed
	if t >= 1.0 {
		return v.end
	}
	delta := v.end - v.start
	// start + delta * t
	return T(float64(v.start) + float64(delta)*v.easingFunc(t))
}
