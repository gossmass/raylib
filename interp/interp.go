package interp

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
