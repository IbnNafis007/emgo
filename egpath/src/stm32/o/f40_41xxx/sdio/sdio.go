// Peripheral: SDIO_Periph  SD host Interface.
// Instances:
//  SDIO  mmap.SDIO_BASE
// Registers:
//  0x00 32  POWER  Power control register.
//  0x04 32  CLKCR  SDI clock control register.
//  0x08 32  ARG    Argument register.
//  0x0C 32  CMD    Command register.
//  0x10 32  DTIMER Data timer register.
//  0x14 32  DLEN   Data length register.
//  0x18 32  DCTRL  Data control register.
//  0x1C 32  ICR    Interrupt clear register.
//  0x20 32  MASK   Mask register.
//  0x60 32  FIFO   Data FIFO register.
// Import:
//  stm32/o/f40_41xxx/mmap
package sdio

const (
	PWRCTRL   POWER_Bits = 0x03 << 0 //+ PWRCTRL[1:0] bits (Power supply control bits).
	PWRCTRL_0 POWER_Bits = 0x01 << 0 //  Bit 0.
	PWRCTRL_1 POWER_Bits = 0x02 << 0 //  Bit 1.
)

const (
	CLKDIV   CLKCR_Bits = 0xFF << 0  //+ Clock divide factor.
	CLKEN    CLKCR_Bits = 0x01 << 8  //+ Clock enable bit.
	PWRSAV   CLKCR_Bits = 0x01 << 9  //+ Power saving configuration bit.
	BYPASS   CLKCR_Bits = 0x01 << 10 //+ Clock divider bypass enable bit.
	WIDBUS   CLKCR_Bits = 0x03 << 11 //+ WIDBUS[1:0] bits (Wide bus mode enable bit).
	WIDBUS_0 CLKCR_Bits = 0x01 << 11 //  Bit 0.
	WIDBUS_1 CLKCR_Bits = 0x02 << 11 //  Bit 1.
	NEGEDGE  CLKCR_Bits = 0x01 << 13 //+ SDIO_CK dephasing selection bit.
	HWFC_EN  CLKCR_Bits = 0x01 << 14 //+ HW Flow Control enable.
)

const (
	CMDARG ARG_Bits = 0xFFFFFFFF << 0 //+ Command argument.
)

const (
	CMDINDEX    CMD_Bits = 0x3F << 0  //+ Command Index.
	WAITRESP    CMD_Bits = 0x03 << 6  //+ WAITRESP[1:0] bits (Wait for response bits).
	WAITRESP_0  CMD_Bits = 0x01 << 6  //  Bit 0.
	WAITRESP_1  CMD_Bits = 0x02 << 6  //  Bit 1.
	WAITINT     CMD_Bits = 0x01 << 8  //+ CPSM Waits for Interrupt Request.
	WAITPEND    CMD_Bits = 0x01 << 9  //+ CPSM Waits for ends of data transfer (CmdPend internal signal).
	CPSMEN      CMD_Bits = 0x01 << 10 //+ Command path state machine (CPSM) Enable bit.
	SDIOSUSPEND CMD_Bits = 0x01 << 11 //+ SD I/O suspend command.
	ENCMDCOMPL  CMD_Bits = 0x01 << 12 //+ Enable CMD completion.
	NIEN        CMD_Bits = 0x01 << 13 //+ Not Interrupt Enable.
	CEATACMD    CMD_Bits = 0x01 << 14 //+ CE-ATA command.
)

const (
	DATATIME DTIMER_Bits = 0xFFFFFFFF << 0 //+ Data timeout period..
)

const (
	DATALENGTH DLEN_Bits = 0x1FFFFFF << 0 //+ Data length value.
)

const (
	DTEN         DCTRL_Bits = 0x01 << 0  //+ Data transfer enabled bit.
	DTDIR        DCTRL_Bits = 0x01 << 1  //+ Data transfer direction selection.
	DTMODE       DCTRL_Bits = 0x01 << 2  //+ Data transfer mode selection.
	DMAEN        DCTRL_Bits = 0x01 << 3  //+ DMA enabled bit.
	DBLOCKSIZE   DCTRL_Bits = 0x0F << 4  //+ DBLOCKSIZE[3:0] bits (Data block size).
	DBLOCKSIZE_0 DCTRL_Bits = 0x01 << 4  //  Bit 0.
	DBLOCKSIZE_1 DCTRL_Bits = 0x02 << 4  //  Bit 1.
	DBLOCKSIZE_2 DCTRL_Bits = 0x04 << 4  //  Bit 2.
	DBLOCKSIZE_3 DCTRL_Bits = 0x08 << 4  //  Bit 3.
	RWSTART      DCTRL_Bits = 0x01 << 8  //+ Read wait start.
	RWSTOP       DCTRL_Bits = 0x01 << 9  //+ Read wait stop.
	RWMOD        DCTRL_Bits = 0x01 << 10 //+ Read wait mode.
	SDIOEN       DCTRL_Bits = 0x01 << 11 //+ SD I/O enable functions.
)

const (
	CCRCFAILC ICR_Bits = 0x01 << 0  //+ CCRCFAIL flag clear bit.
	DCRCFAILC ICR_Bits = 0x01 << 1  //+ DCRCFAIL flag clear bit.
	CTIMEOUTC ICR_Bits = 0x01 << 2  //+ CTIMEOUT flag clear bit.
	DTIMEOUTC ICR_Bits = 0x01 << 3  //+ DTIMEOUT flag clear bit.
	TXUNDERRC ICR_Bits = 0x01 << 4  //+ TXUNDERR flag clear bit.
	RXOVERRC  ICR_Bits = 0x01 << 5  //+ RXOVERR flag clear bit.
	CMDRENDC  ICR_Bits = 0x01 << 6  //+ CMDREND flag clear bit.
	CMDSENTC  ICR_Bits = 0x01 << 7  //+ CMDSENT flag clear bit.
	DATAENDC  ICR_Bits = 0x01 << 8  //+ DATAEND flag clear bit.
	STBITERRC ICR_Bits = 0x01 << 9  //+ STBITERR flag clear bit.
	DBCKENDC  ICR_Bits = 0x01 << 10 //+ DBCKEND flag clear bit.
	SDIOITC   ICR_Bits = 0x01 << 22 //+ SDIOIT flag clear bit.
	CEATAENDC ICR_Bits = 0x01 << 23 //+ CEATAEND flag clear bit.
)

const (
	CCRCFAILIE MASK_Bits = 0x01 << 0  //+ Command CRC Fail Interrupt Enable.
	DCRCFAILIE MASK_Bits = 0x01 << 1  //+ Data CRC Fail Interrupt Enable.
	CTIMEOUTIE MASK_Bits = 0x01 << 2  //+ Command TimeOut Interrupt Enable.
	DTIMEOUTIE MASK_Bits = 0x01 << 3  //+ Data TimeOut Interrupt Enable.
	TXUNDERRIE MASK_Bits = 0x01 << 4  //+ Tx FIFO UnderRun Error Interrupt Enable.
	RXOVERRIE  MASK_Bits = 0x01 << 5  //+ Rx FIFO OverRun Error Interrupt Enable.
	CMDRENDIE  MASK_Bits = 0x01 << 6  //+ Command Response Received Interrupt Enable.
	CMDSENTIE  MASK_Bits = 0x01 << 7  //+ Command Sent Interrupt Enable.
	DATAENDIE  MASK_Bits = 0x01 << 8  //+ Data End Interrupt Enable.
	STBITERRIE MASK_Bits = 0x01 << 9  //+ Start Bit Error Interrupt Enable.
	DBCKENDIE  MASK_Bits = 0x01 << 10 //+ Data Block End Interrupt Enable.
	CMDACTIE   MASK_Bits = 0x01 << 11 //+ CCommand Acting Interrupt Enable.
	TXACTIE    MASK_Bits = 0x01 << 12 //+ Data Transmit Acting Interrupt Enable.
	RXACTIE    MASK_Bits = 0x01 << 13 //+ Data receive acting interrupt enabled.
	TXFIFOHEIE MASK_Bits = 0x01 << 14 //+ Tx FIFO Half Empty interrupt Enable.
	RXFIFOHFIE MASK_Bits = 0x01 << 15 //+ Rx FIFO Half Full interrupt Enable.
	TXFIFOFIE  MASK_Bits = 0x01 << 16 //+ Tx FIFO Full interrupt Enable.
	RXFIFOFIE  MASK_Bits = 0x01 << 17 //+ Rx FIFO Full interrupt Enable.
	TXFIFOEIE  MASK_Bits = 0x01 << 18 //+ Tx FIFO Empty interrupt Enable.
	RXFIFOEIE  MASK_Bits = 0x01 << 19 //+ Rx FIFO Empty interrupt Enable.
	TXDAVLIE   MASK_Bits = 0x01 << 20 //+ Data available in Tx FIFO interrupt Enable.
	RXDAVLIE   MASK_Bits = 0x01 << 21 //+ Data available in Rx FIFO interrupt Enable.
	SDIOITIE   MASK_Bits = 0x01 << 22 //+ SDIO Mode Interrupt Received interrupt Enable.
	CEATAENDIE MASK_Bits = 0x01 << 23 //+ CE-ATA command completion signal received Interrupt Enable.
)

const (
	FIFODATA FIFO_Bits = 0xFFFFFFFF << 0 //+ Receive and transmit FIFO data.
)
