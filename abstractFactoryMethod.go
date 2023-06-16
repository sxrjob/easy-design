package main

import (
	"fmt"
)

/*
练习：
设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，显卡具有显示功能（display，功能实现只要打印出意义即可），内存具有存储功能（storage），cpu具有计算功能（calculate）。
现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
要求组装两台电脑，
1台（Intel的CPU，Intel的显卡，Intel的内存）
1台（Intel的CPU， nvidia的显卡，Kingston的内存）
用抽象工厂模式实现。
*/

func main() {

	computer1 := new(Computer)
	intelFactory := new(IntelFactory)

	computer1.cpu = intelFactory.CreateCPU()
	computer1.arm = intelFactory.CreateARM()
	computer1.gpu = intelFactory.CreateGPU()
	computer1.show()

	computer2 := new(Computer)
	kingstonFactory := new(KingstonFactory)
	nvidiaFactory := new(NvidiaFactory)

	computer2.cpu = intelFactory.CreateCPU()
	computer2.arm = kingstonFactory.CreateARM()
	computer2.gpu = nvidiaFactory.CreateGPU()

	computer2.show()
}

// 抽象产品-GPU ARM CPU

type AbstractGPU interface {
	show()
}
type AbstractARM interface {
	show()
}
type AbstractCPU interface {
	show()
}

// 抽象工厂

type AbstractComputerFactory interface {
	CreateGPU() AbstractGPU
	CreateARM() AbstractARM
	CreateCPU() AbstractCPU
}

// 抽象的产品-电脑

type Computer struct {
	gpu AbstractGPU
	arm AbstractARM
	cpu AbstractCPU
}

func (c *Computer) show() {
	c.arm.show()
	c.cpu.show()
	c.gpu.show()
}

// 具体的产品
// intel厂商

type IntelCPU struct {
	AbstractCPU
}

func (intelCPU *IntelCPU) show() {
	fmt.Printf("i am intelCPU \n")
}

type IntelGPU struct {
	AbstractGPU
}

func (intelGPU *IntelGPU) show() {
	fmt.Printf("i am intelGPU\n")
}

type IntelARM struct {
	AbstractARM
}

func (intelARM *IntelARM) show() {
	fmt.Printf("i am intelARM\n")
}

// Kingston厂商

type KingstonCPU struct {
	AbstractCPU
}

func (kingstonGPU *KingstonCPU) show() {
	fmt.Printf("i am KingstonCPU\n")
}

type KingstonGPU struct {
	AbstractGPU
}

func (kingstonGPU *KingstonGPU) show() {
	fmt.Printf("i am KingstonGPU\n")
}

type KingstonARM struct {
	AbstractARM
}

func (kingstonARM *KingstonARM) show() {
	fmt.Printf("i am KingstonARM\n")
}

// Nvidia厂商

type NvidiaCPU struct {
	AbstractCPU
}

func (nvidiaGPU *NvidiaCPU) show() {
	fmt.Printf("i am NvidiaCPU\n")
}

type NvidiaGPU struct {
	AbstractGPU
}

func (nvidiaGPU *NvidiaGPU) show() {
	fmt.Printf("i am NvidiaGPU\n")
}

type NvidiaARM struct {
	AbstractARM
}

func (nvidiaARM *NvidiaARM) show() {
	fmt.Printf("i am NvidiaARM\n")
}

// 实体 intel工厂

type IntelFactory struct {
	AbstractComputerFactory
}

func (intelFactory *IntelFactory) CreateGPU() AbstractGPU {
	gpu := new(IntelGPU)
	return gpu
}
func (intelFactory *IntelFactory) CreateCPU() AbstractCPU {
	cpu := new(IntelCPU)
	return cpu
}
func (intelFactory *IntelFactory) CreateARM() AbstractARM {
	arm := new(IntelARM)
	return arm
}

type KingstonFactory struct {
	AbstractComputerFactory
}

func (factory *KingstonFactory) CreateGPU() AbstractGPU {
	gpu := new(KingstonGPU)
	return gpu
}
func (factory *KingstonFactory) CreateCPU() AbstractCPU {
	cpu := new(KingstonCPU)
	return cpu
}
func (factory *KingstonFactory) CreateARM() AbstractARM {
	arm := new(KingstonARM)
	return arm
}

type NvidiaFactory struct {
	AbstractComputerFactory
}

func (factory *NvidiaFactory) CreateGPU() AbstractGPU {
	gpu := new(NvidiaGPU)
	return gpu
}
func (factory *NvidiaFactory) CreateCPU() AbstractCPU {
	cpu := new(NvidiaCPU)
	return cpu
}
func (factory *NvidiaFactory) CreateARM() AbstractARM {
	arm := new(NvidiaARM)
	return arm
}
