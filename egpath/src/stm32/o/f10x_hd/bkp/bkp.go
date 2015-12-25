// Peripheral: BKP_Periph  Backup Registers.
// Instances:
//  BKP  mmap.BKP_BASE
// Registers:
//  0x04 16  DR1
//  0x08 16  DR2
//  0x0C 16  DR3
//  0x10 16  DR4
//  0x14 16  DR5
//  0x18 16  DR6
//  0x1C 16  DR7
//  0x20 16  DR8
//  0x24 16  DR9
//  0x28 16  DR10
//  0x2C 16  RTCCR
//  0x30 16  CR
//  0x34 16  CSR
//  0x40 16  DR11
//  0x44 16  DR12
//  0x48 16  DR13
//  0x4C 16  DR14
//  0x50 16  DR15
//  0x54 16  DR16
//  0x58 16  DR17
//  0x5C 16  DR18
//  0x60 16  DR19
//  0x64 16  DR20
//  0x68 16  DR21
//  0x6C 16  DR22
//  0x70 16  DR23
//  0x74 16  DR24
//  0x78 16  DR25
//  0x7C 16  DR26
//  0x80 16  DR27
//  0x84 16  DR28
//  0x88 16  DR29
//  0x8C 16  DR30
//  0x90 16  DR31
//  0x94 16  DR32
//  0x98 16  DR33
//  0x9C 16  DR34
//  0xA0 16  DR35
//  0xA4 16  DR36
//  0xA8 16  DR37
//  0xAC 16  DR38
//  0xB0 16  DR39
//  0xB4 16  DR40
//  0xB8 16  DR41
//  0xBC 16  DR42
// Import:
//  stm32/o/f10x_hd/mmap
package bkp

const (
	D DR1_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR2_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR3_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR4_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR5_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR6_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR7_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR8_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR9_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR10_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	CAL  RTCCR_Bits = 0x7F << 0 //+ Calibration value.
	CCO  RTCCR_Bits = 0x01 << 7 //+ Calibration Clock Output.
	ASOE RTCCR_Bits = 0x01 << 8 //+ Alarm or Second Output Enable.
	ASOS RTCCR_Bits = 0x01 << 9 //+ Alarm or Second Output Selection.
)

const (
	TPE  CR_Bits = 0x01 << 0 //+ TAMPER pin enable.
	TPAL CR_Bits = 0x01 << 1 //+ TAMPER pin active level.
)

const (
	CTE  CSR_Bits = 0x01 << 0 //+ Clear Tamper event.
	CTI  CSR_Bits = 0x01 << 1 //+ Clear Tamper Interrupt.
	TPIE CSR_Bits = 0x01 << 2 //+ TAMPER Pin interrupt enable.
	TEF  CSR_Bits = 0x01 << 8 //+ Tamper Event Flag.
	TIF  CSR_Bits = 0x01 << 9 //+ Tamper Interrupt Flag.
)

const (
	D DR11_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR12_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR13_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR14_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR15_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR16_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR17_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR18_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR19_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR20_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR21_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR22_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR23_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR24_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR25_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR26_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR27_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR28_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR29_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR30_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR31_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR32_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR33_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR34_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR35_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR36_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR37_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR38_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR39_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR40_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR41_Bits = 0xFFFF << 0 //+ Backup data.
)

const (
	D DR42_Bits = 0xFFFF << 0 //+ Backup data.
)
