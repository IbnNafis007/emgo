// Peripheral: LTDC_Periph  LCD-TFT Display Controller.
// Instances:
//  LTDC  mmap.LTDC_BASE
// Registers:
//  0x08 32  SSCR  Synchronization Size Configuration Register.
//  0x0C 32  BPCR  Back Porch Configuration Register.
//  0x10 32  AWCR  Active Width Configuration Register.
//  0x14 32  TWCR  Total Width Configuration Register.
//  0x18 32  GCR   Global Control Register.
//  0x24 32  SRCR  Shadow Reload Configuration Register.
//  0x2C 32  BCCR  Background Color Configuration Register.
//  0x34 32  IER   Interrupt Enable Register.
//  0x38 32  ISR   Interrupt Status Register.
//  0x3C 32  ICR   Interrupt Clear Register.
//  0x40 32  LIPCR Line Interrupt Position Configuration Register.
//  0x44 32  CPSR  Current Position Status Register.
//  0x48 32  CDSR  Current Display Status Register.
// Import:
//  stm32/o/f40_41xxx/mmap
package ltdc

const (
	VSH SSCR_Bits = 0x7FF << 0  //+ Vertical Synchronization Height.
	HSW SSCR_Bits = 0xFFF << 16 //+ Horizontal Synchronization Width.
)

const (
	AVBP BPCR_Bits = 0x7FF << 0  //+ Accumulated Vertical Back Porch.
	AHBP BPCR_Bits = 0xFFF << 16 //+ Accumulated Horizontal Back Porch.
)

const (
	AAH AWCR_Bits = 0x7FF << 0  //+ Accumulated Active heigh.
	AAW AWCR_Bits = 0xFFF << 16 //+ Accumulated Active Width.
)

const (
	TOTALH TWCR_Bits = 0x7FF << 0  //+ Total Heigh.
	TOTALW TWCR_Bits = 0xFFF << 16 //+ Total Width.
)

const (
	LTDCEN GCR_Bits = 0x01 << 0  //+ LCD-TFT controller enable bit.
	DBW    GCR_Bits = 0x07 << 4  //+ Dither Blue Width.
	DGW    GCR_Bits = 0x07 << 8  //+ Dither Green Width.
	DRW    GCR_Bits = 0x07 << 12 //+ Dither Red Width.
	DTEN   GCR_Bits = 0x01 << 16 //+ Dither Enable.
	PCPOL  GCR_Bits = 0x01 << 28 //+ Pixel Clock Polarity.
	DEPOL  GCR_Bits = 0x01 << 29 //+ Data Enable Polarity.
	VSPOL  GCR_Bits = 0x01 << 30 //+ Vertical Synchronization Polarity.
	HSPOL  GCR_Bits = 0x01 << 31 //+ Horizontal Synchronization Polarity.
)

const (
	IMR SRCR_Bits = 0x01 << 0 //+ Immediate Reload.
	VBR SRCR_Bits = 0x01 << 1 //+ Vertical Blanking Reload.
)

const (
	BCBLUE  BCCR_Bits = 0xFF << 0  //+ Background Blue value.
	BCGREEN BCCR_Bits = 0xFF << 8  //+ Background Green value.
	BCRED   BCCR_Bits = 0xFF << 16 //+ Background Red value.
)

const (
	LIE    IER_Bits = 0x01 << 0 //+ Line Interrupt Enable.
	FUIE   IER_Bits = 0x01 << 1 //+ FIFO Underrun Interrupt Enable.
	TERRIE IER_Bits = 0x01 << 2 //+ Transfer Error Interrupt Enable.
	RRIE   IER_Bits = 0x01 << 3 //+ Register Reload interrupt enable.
)

const (
	LIF    ISR_Bits = 0x01 << 0 //+ Line Interrupt Flag.
	FUIF   ISR_Bits = 0x01 << 1 //+ FIFO Underrun Interrupt Flag.
	TERRIF ISR_Bits = 0x01 << 2 //+ Transfer Error Interrupt Flag.
	RRIF   ISR_Bits = 0x01 << 3 //+ Register Reload interrupt Flag.
)

const (
	CLIF    ICR_Bits = 0x01 << 0 //+ Clears the Line Interrupt Flag.
	CFUIF   ICR_Bits = 0x01 << 1 //+ Clears the FIFO Underrun Interrupt Flag.
	CTERRIF ICR_Bits = 0x01 << 2 //+ Clears the Transfer Error Interrupt Flag.
	CRRIF   ICR_Bits = 0x01 << 3 //+ Clears Register Reload interrupt Flag.
)

const (
	LIPOS LIPCR_Bits = 0x7FF << 0 //+ Line Interrupt Position.
)

const (
	CYPOS CPSR_Bits = 0xFFFF << 0  //+ Current Y Position.
	CXPOS CPSR_Bits = 0xFFFF << 16 //+ Current X Position.
)

const (
	VDES   CDSR_Bits = 0x01 << 0 //+ Vertical Data Enable Status.
	HDES   CDSR_Bits = 0x01 << 1 //+ Horizontal Data Enable Status.
	VSYNCS CDSR_Bits = 0x01 << 2 //+ Vertical Synchronization Status.
	HSYNCS CDSR_Bits = 0x01 << 3 //+ Horizontal Synchronization Status.
)
