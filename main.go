package main

import (
	advancedtasks "tasks/advancedtasks/develop/dev_01"
	"tasks/basetasks"
	"tasks/pattern"
)

func main() {
	// --- basetasks functions call --- you should uncomment one of the function
	basetasks.TaskOne()
	//basetasks.TaskTwo([]int{1, 2, 3, 4})
	//basetasks.TaskThree()
	//basetasks.TaskFive()
	//basetasks.TaskSeventeen([]int{1, 3, 5, 8, 15}, 8)
	//basetasks.TaskSixteen([]int{83, 32, 4, 7, 1, 3, 5, 8, 15})
	//basetasks.TaskNineteen("главрыба")
	//basetasks.TaskTwenty("snow dog sun")
	//basetasks.TaskTwentyThee([]int{1, 2, 3, 4, 5}, 2)
	//basetasks.TaskTwentySix("THeMessage")
	//basetasks.TaskTwentyFive()
	//basetasks.TaskTwentyFour()
	//basetasks.TaskTwentyTwo()
	//basetasks.TaskTwentyOne()
	//basetasks.TaskEighteen()
	//basetasks.TaskFifteen()
	//basetasks.TaskFourteen()
	//basetasks.TaskThirteen()
	//basetasks.TaskFour()
	//basetasks.TaskSeven()
	//basetasks.TaskEight()
	//basetasks.TaskNine()
	//basetasks.TaskTen()
	//basetasks.TaskEleven()
	//basetasks.TaskSix()
	//basetasks.TaskTwelve()

	// --- advancedtasks functions call --- you should uncomment one of the function
	advancedtasks.TaskOne()
	//advancedtasks2.TaskTwo()
	//advancedtasks3.TaskThree()
	//advancedtasks4.TaskFour()
	//advancedtasks5.TaskFive()
	//advancedtasks6.TaskSix()
	//advancedtasks7.TaskSeven()
	//advancedtasks8.TaskEight()
	//advancedtasks9.TaskNine()
	//advancedtasks10.TaskTen()

	// --- pattern call --- you should uncomment one of the pattern
	//01 - Facade
	facade := pattern.NewStart()
	facade.Start()
	//02 - Builder
	//builder := pattern.NewCarBuilder()
	//car := builder.Engine(623).Body("BMW").Wheels("Yokohama").Build()
	//fmt.Println(car)
	//03 - Visitor
	//pattern.VisitorPatternStart()
	//04 - Command
	//pattern.CommandPatternStart()
	//05 - Chain of responsibility
	//pattern.ChainPatternStart()
	//06 - Factory method
	//pattern.FactoryMethodPatternStart()
	//07 - Strategy
	//pattern.StrategyPatternStart()
	//08 - State
	//pattern.StatePatternStart()
}
