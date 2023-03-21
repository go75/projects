package config

func init() {
	err := inject(map[string]any {
		"database": DataBase,
		"redis": Redis,
		"jwt": Jwt,
		"log": Log,
		"dispatcher": Dispatcher,
	})

	if err !=nil {
		panic(err)
	}
}