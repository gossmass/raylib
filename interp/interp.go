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

func (i *Value[T]) getElapsedTime() float64 {
	return rl.GetTime() - i.startTime
}

func (i *Value[T]) SetTransition(easingFunc EasingFunction) *Value[T] {
	i.easingFunc = easingFunc
	return i
}

func (i *Value[T]) SetDuration(duration float64) *Value[T] {
	i.speed = 1.0 / duration
	return i
}

func (i *Value[T]) SetValue(new_value T) *Value[T] {
	i.start = i.GetValue()
	i.end = new_value
	i.startTime = rl.GetTime()
	return i
}

func (i *Value[T]) GetValue() T {
	elapsed := i.getElapsedTime()
	t := elapsed * i.speed
	if t >= 1.0 {
		return i.end
	}
	delta := i.end - i.start
	// start + delta * t
	return T(float64(i.start) + float64(delta)*i.easingFunc(t))
}
