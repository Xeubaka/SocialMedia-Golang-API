package models

// Password presents the response format of password update
type Password struct {
	New     string `json: new`
	Current string `json: current`
}
