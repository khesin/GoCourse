package carsstruct

//Я не стал создавать несколько структур, т.к. это бессмыслено для машины. Легче указать ее категорию. Простите, ленив:(
type Carssrtuct struct {
	Category    string
	Model       string
	StartDate   string
	TrunkValue  int
	EngineOn    bool
	WindowsOpen bool
	Color       string
}
