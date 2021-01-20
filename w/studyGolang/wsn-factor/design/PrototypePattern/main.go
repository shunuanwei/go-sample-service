package main

func main() {
	var workerCreate WorkerCreator

	workerCreate = new(ProgrammerCreator)
	p := workerCreate.Create()
	taskProject := "Project"
	p.Work(&taskProject)

	workerCreate = new(FarmerCreator)
	f := workerCreate.Create()
	taskFarmland := "farmland"
	f.Work(&taskFarmland)
}
