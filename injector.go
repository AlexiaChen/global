package global

import (
	"fmt"

	"github.com/samber/do"
)

var injector *do.Injector = nil

func InitInjector() {
	if injector == nil {
		injector = do.New()
	}
}

func ProvideService[T any](service do.Provider[T]) error {
	if injector != nil {
		do.Provide(injector, service)
		return nil
	}
	return fmt.Errorf("do not forget invoke InitInjector before doing this")
}

func ProvideWithNamedService(name string, service do.Provider[any]) error {
	if injector != nil {
		do.ProvideNamed(injector, name, service)
		return nil
	}
	return fmt.Errorf("do not forget invoke InitInjector before doing this")
}

func GetService[T any]() (T, error) {
	return do.Invoke[T](injector)
}

func GetNamedService[T any](name string) (T, error) {
	return do.InvokeNamed[T](injector, name)
}

func ShutdownOnSIGTERM() {
	injector.ShutdownOnSIGTERM() // will block until receiving sigterm signal
}

func Shutdown() {
	injector.Shutdown()
}
