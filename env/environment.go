package env

// import (
// 	"fmt"
// 	"runtime"

// 	"github.com/PatrickAmbrosso/servd/database"
// 	"github.com/PatrickAmbrosso/servd/env/unix"
// 	"github.com/PatrickAmbrosso/servd/env/windows"
// )

// type PlatformHandlers struct {
// 	// Checkers
// 	IsAdmin func() bool
// 	// Individual Service Actions
// 	InstallService func(service database.Service) error
// 	ModifyService  func(service database.Service) error
// 	StartService   func(serviceName string) (bool, error)
// 	RestartService func(serviceName string) (bool, error)
// 	PauseService   func(serviceName string) (bool, error)
// 	StopService    func(serviceName string) (bool, error)
// 	DeleteService  func(serviceName string) (bool, error)
// 	// Multiple Service Actions
// 	MonitorServices func(watch bool) error
// 	ManageServices  func(action string, path string) error
// 	PollLogs        func(tail bool) error
// }

// var handlers = map[string]PlatformHandlers{
// 	"windows": {
// 		IsAdmin: windows.IsAdmin,
// 	},
// 	"linux": {
// 		IsAdmin: unix.IsAdmin,
// 	},
// 	"darwin": {
// 		IsAdmin: unix.IsAdmin,
// 	},
// }

// const currOS = runtime.GOOS

// func GetPlatformHandlers() (*PlatformHandlers, error) {
// 	platform, exists := handlers[currOS]
// 	if !exists {
// 		return nil, fmt.Errorf("unsupported platform: %s", currOS)
// 	}

// 	return &platform, nil
// }
