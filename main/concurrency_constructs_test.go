package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoRoutineDefinitions(t *testing.T) {
	goRoutineDefinitions()
}

func TestWaitGroupSimpleExample(t *testing.T) {
	waitGroupSimpleExample()
}

func TestWaitGroupCanModifyVariableOutsideItsScope(t *testing.T) {
	waitGroupCanModifyVariableOutsideItsScope()
}

func TestWaitGroupLoopWithWronglyUsedVariable(t *testing.T) {
	waitGroupLoopWithWronglyUsedVariable()
}

func TestWaitGroupLoopWithCorrectlyUserVariable(t *testing.T) {
	waitGroupLoopWithCorrectlyUserVariable()
}

func TestMutex(t *testing.T) {
	mutex()
}

func TestRwMutex(t *testing.T) {
	rwmutex()
}

func TestWaitingForGoRoutineToFinish(t *testing.T) {
	waitingForGoRoutineToFinish()
}

func TestCondSimpleCase(t *testing.T) {
	conditionWithSignal()
}

func TestCondWithMultipleGoRoutines(t *testing.T) {
	conditionWithBroadcast()
}

func TestOnce(t *testing.T) {
	once()
}

func TestPool(t *testing.T) {
	pool()
}

func TestUnbufferedChannelsTwoWay(t *testing.T) {
	unbufferedChannelsTwoWay()
}

func TestUnbufferedChannelsOneWay(t *testing.T) {
	unbufferedChannelsOneWay()
}

func TestChannelCloseExampleToGiveSignalToBlockedGoroutines(t *testing.T) {
	channelCloseExampleToGiveSignalToBlockedGoroutines()
}

func TestBufferedChannels(t *testing.T) {
	bufferedChannels()
}

func TestRangeOverChannel1(t *testing.T) {
	rangeOverChannel1()
}

func TestRangeOverChannel2(t *testing.T) {
	rangeOverChannel2()
}

func TestSelectStatement(t *testing.T) {
	selectStatement()
}

func TestSetRuntimeGOMAXPROCS(t *testing.T) {
	setRuntimeGOMAXPROCS()
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
