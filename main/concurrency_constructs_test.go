package main

import (
	"fmt"
	"runtime"
	"testing"
)

/**********************************************************************
                           go routine
**********************************************************************/

func TestGoRoutineDefinitions(t *testing.T) {
	goRoutineDefinitions()
}

/**********************************************************************
                           wait group
**********************************************************************/

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

/**********************************************************************
                                mutex
**********************************************************************/

func TestMutex(t *testing.T) {
	mutex()
}

/**********************************************************************
                               rwmutex
**********************************************************************/

func TestRwMutex(t *testing.T) {
	rwmutex()
}

/**********************************************************************
                                cond
**********************************************************************/

func TestWaitingForGoRoutineToFinish(t *testing.T) {
	waitingForGoRoutineToFinish()
}

func TestCondSimpleCase(t *testing.T) {
	conditionWithSignal()
}

func TestCondWithMultipleGoRoutines(t *testing.T) {
	conditionWithBroadcast()
}

/**********************************************************************
                                once
**********************************************************************/

func TestOnce(t *testing.T) {
	once()
}

/**********************************************************************
                                pool
**********************************************************************/

func TestPool(t *testing.T) {
	pool()
}

/**********************************************************************
                              channels
**********************************************************************/

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

/**********************************************************************
                              GOMAXPROCS
**********************************************************************/

func TestSetRuntimeGOMAXPROCS(t *testing.T) {
	setRuntimeGOMAXPROCS()
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
