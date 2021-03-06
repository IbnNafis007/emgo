// +build l476xx

package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type ADC_Periph struct {
	ISR     RISR
	IER     RIER
	CR      RCR
	CFGR    RCFGR
	CFGR2   RCFGR2
	SMPR1   RSMPR1
	SMPR2   RSMPR2
	_       uint32
	TR1     RTR1
	TR2     RTR2
	TR3     RTR3
	_       uint32
	SQR1    RSQR1
	SQR2    RSQR2
	SQR3    RSQR3
	SQR4    RSQR4
	DR      RDR
	_       [2]uint32
	JSQR    RJSQR
	_       [4]uint32
	OFR     [4]ROFR
	_       [4]uint32
	JDR     [4]RJDR
	_       [4]uint32
	AWD2CR  RAWD2CR
	AWD3CR  RAWD3CR
	_       [2]uint32
	DIFSEL  RDIFSEL
	CALFACT RCALFACT
}

func (p *ADC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var ADC1 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC1_BASE)))

//emgo:const
var ADC2 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC2_BASE)))

//emgo:const
var ADC3 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC3_BASE)))

type ISR uint32

func (b ISR) Field(mask ISR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR) J(v int) ISR {
	return ISR(bits.Make32(v, uint32(mask)))
}

type RISR struct{ mmio.U32 }

func (r *RISR) Bits(mask ISR) ISR     { return ISR(r.U32.Bits(uint32(mask))) }
func (r *RISR) StoreBits(mask, b ISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISR) SetBits(mask ISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RISR) ClearBits(mask ISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RISR) Load() ISR             { return ISR(r.U32.Load()) }
func (r *RISR) Store(b ISR)           { r.U32.Store(uint32(b)) }

func (r *RISR) AtomicStoreBits(mask, b ISR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RISR) AtomicSetBits(mask ISR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RISR) AtomicClearBits(mask ISR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMISR struct{ mmio.UM32 }

func (rm RMISR) Load() ISR   { return ISR(rm.UM32.Load()) }
func (rm RMISR) Store(b ISR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) ADRDY() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ADRDY)}}
}

func (p *ADC_Periph) EOSMP() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(EOSMP)}}
}

func (p *ADC_Periph) EOC() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(EOC)}}
}

func (p *ADC_Periph) EOS() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(EOS)}}
}

func (p *ADC_Periph) OVR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(OVR)}}
}

func (p *ADC_Periph) JEOC() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(JEOC)}}
}

func (p *ADC_Periph) JEOS() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(JEOS)}}
}

func (p *ADC_Periph) AWD1() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(AWD1)}}
}

func (p *ADC_Periph) AWD2() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(AWD2)}}
}

func (p *ADC_Periph) AWD3() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(AWD3)}}
}

func (p *ADC_Periph) JQOVF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(JQOVF)}}
}

type IER uint32

func (b IER) Field(mask IER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IER) J(v int) IER {
	return IER(bits.Make32(v, uint32(mask)))
}

type RIER struct{ mmio.U32 }

func (r *RIER) Bits(mask IER) IER     { return IER(r.U32.Bits(uint32(mask))) }
func (r *RIER) StoreBits(mask, b IER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIER) SetBits(mask IER)      { r.U32.SetBits(uint32(mask)) }
func (r *RIER) ClearBits(mask IER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIER) Load() IER             { return IER(r.U32.Load()) }
func (r *RIER) Store(b IER)           { r.U32.Store(uint32(b)) }

func (r *RIER) AtomicStoreBits(mask, b IER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIER) AtomicSetBits(mask IER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIER) AtomicClearBits(mask IER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIER struct{ mmio.UM32 }

func (rm RMIER) Load() IER   { return IER(rm.UM32.Load()) }
func (rm RMIER) Store(b IER) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) ADRDYIEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(ADRDYIEIE)}}
}

func (p *ADC_Periph) EOSMPIEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(EOSMPIEIE)}}
}

func (p *ADC_Periph) EOCIEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(EOCIEIE)}}
}

func (p *ADC_Periph) EOSIEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(EOSIEIE)}}
}

func (p *ADC_Periph) OVRIEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(OVRIEIE)}}
}

func (p *ADC_Periph) JEOCIEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(JEOCIEIE)}}
}

func (p *ADC_Periph) JEOSIEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(JEOSIEIE)}}
}

func (p *ADC_Periph) AWD1IEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(AWD1IEIE)}}
}

func (p *ADC_Periph) AWD2IEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(AWD2IEIE)}}
}

func (p *ADC_Periph) AWD3IEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(AWD3IEIE)}}
}

func (p *ADC_Periph) JQOVFIEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(JQOVFIEIE)}}
}

type CR uint32

func (b CR) Field(mask CR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR) J(v int) CR {
	return CR(bits.Make32(v, uint32(mask)))
}

type RCR struct{ mmio.U32 }

func (r *RCR) Bits(mask CR) CR      { return CR(r.U32.Bits(uint32(mask))) }
func (r *RCR) StoreBits(mask, b CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR) SetBits(mask CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR) ClearBits(mask CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR) Load() CR             { return CR(r.U32.Load()) }
func (r *RCR) Store(b CR)           { r.U32.Store(uint32(b)) }

func (r *RCR) AtomicStoreBits(mask, b CR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR) AtomicSetBits(mask CR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR) AtomicClearBits(mask CR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR struct{ mmio.UM32 }

func (rm RMCR) Load() CR   { return CR(rm.UM32.Load()) }
func (rm RMCR) Store(b CR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) ADEN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADEN)}}
}

func (p *ADC_Periph) ADDIS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADDIS)}}
}

func (p *ADC_Periph) ADSTART() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADSTART)}}
}

func (p *ADC_Periph) JADSTART() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(JADSTART)}}
}

func (p *ADC_Periph) ADSTP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADSTP)}}
}

func (p *ADC_Periph) JADSTP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(JADSTP)}}
}

func (p *ADC_Periph) ADVREGEN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADVREGEN)}}
}

func (p *ADC_Periph) DEEPPWD() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DEEPPWD)}}
}

func (p *ADC_Periph) ADCALDIF() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADCALDIF)}}
}

func (p *ADC_Periph) ADCAL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADCAL)}}
}

type CFGR uint32

func (b CFGR) Field(mask CFGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR) J(v int) CFGR {
	return CFGR(bits.Make32(v, uint32(mask)))
}

type RCFGR struct{ mmio.U32 }

func (r *RCFGR) Bits(mask CFGR) CFGR    { return CFGR(r.U32.Bits(uint32(mask))) }
func (r *RCFGR) StoreBits(mask, b CFGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) SetBits(mask CFGR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR) ClearBits(mask CFGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR) Load() CFGR             { return CFGR(r.U32.Load()) }
func (r *RCFGR) Store(b CFGR)           { r.U32.Store(uint32(b)) }

func (r *RCFGR) AtomicStoreBits(mask, b CFGR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) AtomicSetBits(mask CFGR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR) AtomicClearBits(mask CFGR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR struct{ mmio.UM32 }

func (rm RMCFGR) Load() CFGR   { return CFGR(rm.UM32.Load()) }
func (rm RMCFGR) Store(b CFGR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) DMAEN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(DMAEN)}}
}

func (p *ADC_Periph) DMACFG() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(DMACFG)}}
}

func (p *ADC_Periph) RES() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(RES)}}
}

func (p *ADC_Periph) ALIGN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(ALIGN)}}
}

func (p *ADC_Periph) EXTSEL() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(EXTSEL)}}
}

func (p *ADC_Periph) EXTEN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(EXTEN)}}
}

func (p *ADC_Periph) OVRMOD() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(OVRMOD)}}
}

func (p *ADC_Periph) CONT() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(CONT)}}
}

func (p *ADC_Periph) AUTDLY() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(AUTDLY)}}
}

func (p *ADC_Periph) DISCEN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(DISCEN)}}
}

func (p *ADC_Periph) DISCNUM() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(DISCNUM)}}
}

func (p *ADC_Periph) JDISCEN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(JDISCEN)}}
}

func (p *ADC_Periph) JQM() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(JQM)}}
}

func (p *ADC_Periph) AWD1SGL() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(AWD1SGL)}}
}

func (p *ADC_Periph) AWD1EN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(AWD1EN)}}
}

func (p *ADC_Periph) JAWD1EN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(JAWD1EN)}}
}

func (p *ADC_Periph) JAUTO() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(JAUTO)}}
}

func (p *ADC_Periph) AWD1CH() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(AWD1CH)}}
}

func (p *ADC_Periph) JQDIS() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(JQDIS)}}
}

type CFGR2 uint32

func (b CFGR2) Field(mask CFGR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR2) J(v int) CFGR2 {
	return CFGR2(bits.Make32(v, uint32(mask)))
}

type RCFGR2 struct{ mmio.U32 }

func (r *RCFGR2) Bits(mask CFGR2) CFGR2   { return CFGR2(r.U32.Bits(uint32(mask))) }
func (r *RCFGR2) StoreBits(mask, b CFGR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR2) SetBits(mask CFGR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR2) ClearBits(mask CFGR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR2) Load() CFGR2             { return CFGR2(r.U32.Load()) }
func (r *RCFGR2) Store(b CFGR2)           { r.U32.Store(uint32(b)) }

func (r *RCFGR2) AtomicStoreBits(mask, b CFGR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR2) AtomicSetBits(mask CFGR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR2) AtomicClearBits(mask CFGR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR2 struct{ mmio.UM32 }

func (rm RMCFGR2) Load() CFGR2   { return CFGR2(rm.UM32.Load()) }
func (rm RMCFGR2) Store(b CFGR2) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) ROVSE() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(ROVSE)}}
}

func (p *ADC_Periph) JOVSE() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(JOVSE)}}
}

func (p *ADC_Periph) OVSR() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(OVSR)}}
}

func (p *ADC_Periph) OVSS() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(OVSS)}}
}

func (p *ADC_Periph) TROVS() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(TROVS)}}
}

func (p *ADC_Periph) ROVSM() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(ROVSM)}}
}

type SMPR1 uint32

func (b SMPR1) Field(mask SMPR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR1) J(v int) SMPR1 {
	return SMPR1(bits.Make32(v, uint32(mask)))
}

type RSMPR1 struct{ mmio.U32 }

func (r *RSMPR1) Bits(mask SMPR1) SMPR1   { return SMPR1(r.U32.Bits(uint32(mask))) }
func (r *RSMPR1) StoreBits(mask, b SMPR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSMPR1) SetBits(mask SMPR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSMPR1) ClearBits(mask SMPR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSMPR1) Load() SMPR1             { return SMPR1(r.U32.Load()) }
func (r *RSMPR1) Store(b SMPR1)           { r.U32.Store(uint32(b)) }

func (r *RSMPR1) AtomicStoreBits(mask, b SMPR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSMPR1) AtomicSetBits(mask SMPR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSMPR1) AtomicClearBits(mask SMPR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSMPR1 struct{ mmio.UM32 }

func (rm RMSMPR1) Load() SMPR1   { return SMPR1(rm.UM32.Load()) }
func (rm RMSMPR1) Store(b SMPR1) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP0() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP0)}}
}

func (p *ADC_Periph) SMP1() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP1)}}
}

func (p *ADC_Periph) SMP2() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP2)}}
}

func (p *ADC_Periph) SMP3() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP3)}}
}

func (p *ADC_Periph) SMP4() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP4)}}
}

func (p *ADC_Periph) SMP5() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP5)}}
}

func (p *ADC_Periph) SMP6() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP6)}}
}

func (p *ADC_Periph) SMP7() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP7)}}
}

func (p *ADC_Periph) SMP8() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP8)}}
}

func (p *ADC_Periph) SMP9() RMSMPR1 {
	return RMSMPR1{mmio.UM32{&p.SMPR1.U32, uint32(SMP9)}}
}

type SMPR2 uint32

func (b SMPR2) Field(mask SMPR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR2) J(v int) SMPR2 {
	return SMPR2(bits.Make32(v, uint32(mask)))
}

type RSMPR2 struct{ mmio.U32 }

func (r *RSMPR2) Bits(mask SMPR2) SMPR2   { return SMPR2(r.U32.Bits(uint32(mask))) }
func (r *RSMPR2) StoreBits(mask, b SMPR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSMPR2) SetBits(mask SMPR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSMPR2) ClearBits(mask SMPR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSMPR2) Load() SMPR2             { return SMPR2(r.U32.Load()) }
func (r *RSMPR2) Store(b SMPR2)           { r.U32.Store(uint32(b)) }

func (r *RSMPR2) AtomicStoreBits(mask, b SMPR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSMPR2) AtomicSetBits(mask SMPR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSMPR2) AtomicClearBits(mask SMPR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSMPR2 struct{ mmio.UM32 }

func (rm RMSMPR2) Load() SMPR2   { return SMPR2(rm.UM32.Load()) }
func (rm RMSMPR2) Store(b SMPR2) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP10() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP10)}}
}

func (p *ADC_Periph) SMP11() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP11)}}
}

func (p *ADC_Periph) SMP12() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP12)}}
}

func (p *ADC_Periph) SMP13() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP13)}}
}

func (p *ADC_Periph) SMP14() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP14)}}
}

func (p *ADC_Periph) SMP15() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP15)}}
}

func (p *ADC_Periph) SMP16() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP16)}}
}

func (p *ADC_Periph) SMP17() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP17)}}
}

func (p *ADC_Periph) SMP18() RMSMPR2 {
	return RMSMPR2{mmio.UM32{&p.SMPR2.U32, uint32(SMP18)}}
}

type TR1 uint32

func (b TR1) Field(mask TR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR1) J(v int) TR1 {
	return TR1(bits.Make32(v, uint32(mask)))
}

type RTR1 struct{ mmio.U32 }

func (r *RTR1) Bits(mask TR1) TR1     { return TR1(r.U32.Bits(uint32(mask))) }
func (r *RTR1) StoreBits(mask, b TR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTR1) SetBits(mask TR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RTR1) ClearBits(mask TR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTR1) Load() TR1             { return TR1(r.U32.Load()) }
func (r *RTR1) Store(b TR1)           { r.U32.Store(uint32(b)) }

func (r *RTR1) AtomicStoreBits(mask, b TR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTR1) AtomicSetBits(mask TR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTR1) AtomicClearBits(mask TR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTR1 struct{ mmio.UM32 }

func (rm RMTR1) Load() TR1   { return TR1(rm.UM32.Load()) }
func (rm RMTR1) Store(b TR1) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT1() RMTR1 {
	return RMTR1{mmio.UM32{&p.TR1.U32, uint32(LT1)}}
}

func (p *ADC_Periph) HT1() RMTR1 {
	return RMTR1{mmio.UM32{&p.TR1.U32, uint32(HT1)}}
}

type TR2 uint32

func (b TR2) Field(mask TR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR2) J(v int) TR2 {
	return TR2(bits.Make32(v, uint32(mask)))
}

type RTR2 struct{ mmio.U32 }

func (r *RTR2) Bits(mask TR2) TR2     { return TR2(r.U32.Bits(uint32(mask))) }
func (r *RTR2) StoreBits(mask, b TR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTR2) SetBits(mask TR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RTR2) ClearBits(mask TR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTR2) Load() TR2             { return TR2(r.U32.Load()) }
func (r *RTR2) Store(b TR2)           { r.U32.Store(uint32(b)) }

func (r *RTR2) AtomicStoreBits(mask, b TR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTR2) AtomicSetBits(mask TR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTR2) AtomicClearBits(mask TR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTR2 struct{ mmio.UM32 }

func (rm RMTR2) Load() TR2   { return TR2(rm.UM32.Load()) }
func (rm RMTR2) Store(b TR2) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT2() RMTR2 {
	return RMTR2{mmio.UM32{&p.TR2.U32, uint32(LT2)}}
}

func (p *ADC_Periph) HT2() RMTR2 {
	return RMTR2{mmio.UM32{&p.TR2.U32, uint32(HT2)}}
}

type TR3 uint32

func (b TR3) Field(mask TR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR3) J(v int) TR3 {
	return TR3(bits.Make32(v, uint32(mask)))
}

type RTR3 struct{ mmio.U32 }

func (r *RTR3) Bits(mask TR3) TR3     { return TR3(r.U32.Bits(uint32(mask))) }
func (r *RTR3) StoreBits(mask, b TR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTR3) SetBits(mask TR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RTR3) ClearBits(mask TR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTR3) Load() TR3             { return TR3(r.U32.Load()) }
func (r *RTR3) Store(b TR3)           { r.U32.Store(uint32(b)) }

func (r *RTR3) AtomicStoreBits(mask, b TR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTR3) AtomicSetBits(mask TR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTR3) AtomicClearBits(mask TR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTR3 struct{ mmio.UM32 }

func (rm RMTR3) Load() TR3   { return TR3(rm.UM32.Load()) }
func (rm RMTR3) Store(b TR3) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT3() RMTR3 {
	return RMTR3{mmio.UM32{&p.TR3.U32, uint32(LT3)}}
}

func (p *ADC_Periph) HT3() RMTR3 {
	return RMTR3{mmio.UM32{&p.TR3.U32, uint32(HT3)}}
}

type SQR1 uint32

func (b SQR1) Field(mask SQR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR1) J(v int) SQR1 {
	return SQR1(bits.Make32(v, uint32(mask)))
}

type RSQR1 struct{ mmio.U32 }

func (r *RSQR1) Bits(mask SQR1) SQR1    { return SQR1(r.U32.Bits(uint32(mask))) }
func (r *RSQR1) StoreBits(mask, b SQR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSQR1) SetBits(mask SQR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSQR1) ClearBits(mask SQR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSQR1) Load() SQR1             { return SQR1(r.U32.Load()) }
func (r *RSQR1) Store(b SQR1)           { r.U32.Store(uint32(b)) }

func (r *RSQR1) AtomicStoreBits(mask, b SQR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSQR1) AtomicSetBits(mask SQR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSQR1) AtomicClearBits(mask SQR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSQR1 struct{ mmio.UM32 }

func (rm RMSQR1) Load() SQR1   { return SQR1(rm.UM32.Load()) }
func (rm RMSQR1) Store(b SQR1) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) L() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(L)}}
}

func (p *ADC_Periph) SQ1() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(SQ1)}}
}

func (p *ADC_Periph) SQ2() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(SQ2)}}
}

func (p *ADC_Periph) SQ3() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(SQ3)}}
}

func (p *ADC_Periph) SQ4() RMSQR1 {
	return RMSQR1{mmio.UM32{&p.SQR1.U32, uint32(SQ4)}}
}

type SQR2 uint32

func (b SQR2) Field(mask SQR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR2) J(v int) SQR2 {
	return SQR2(bits.Make32(v, uint32(mask)))
}

type RSQR2 struct{ mmio.U32 }

func (r *RSQR2) Bits(mask SQR2) SQR2    { return SQR2(r.U32.Bits(uint32(mask))) }
func (r *RSQR2) StoreBits(mask, b SQR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSQR2) SetBits(mask SQR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSQR2) ClearBits(mask SQR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSQR2) Load() SQR2             { return SQR2(r.U32.Load()) }
func (r *RSQR2) Store(b SQR2)           { r.U32.Store(uint32(b)) }

func (r *RSQR2) AtomicStoreBits(mask, b SQR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSQR2) AtomicSetBits(mask SQR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSQR2) AtomicClearBits(mask SQR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSQR2 struct{ mmio.UM32 }

func (rm RMSQR2) Load() SQR2   { return SQR2(rm.UM32.Load()) }
func (rm RMSQR2) Store(b SQR2) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ5() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ5)}}
}

func (p *ADC_Periph) SQ6() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ6)}}
}

func (p *ADC_Periph) SQ7() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ7)}}
}

func (p *ADC_Periph) SQ8() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ8)}}
}

func (p *ADC_Periph) SQ9() RMSQR2 {
	return RMSQR2{mmio.UM32{&p.SQR2.U32, uint32(SQ9)}}
}

type SQR3 uint32

func (b SQR3) Field(mask SQR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR3) J(v int) SQR3 {
	return SQR3(bits.Make32(v, uint32(mask)))
}

type RSQR3 struct{ mmio.U32 }

func (r *RSQR3) Bits(mask SQR3) SQR3    { return SQR3(r.U32.Bits(uint32(mask))) }
func (r *RSQR3) StoreBits(mask, b SQR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSQR3) SetBits(mask SQR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RSQR3) ClearBits(mask SQR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSQR3) Load() SQR3             { return SQR3(r.U32.Load()) }
func (r *RSQR3) Store(b SQR3)           { r.U32.Store(uint32(b)) }

func (r *RSQR3) AtomicStoreBits(mask, b SQR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSQR3) AtomicSetBits(mask SQR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSQR3) AtomicClearBits(mask SQR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSQR3 struct{ mmio.UM32 }

func (rm RMSQR3) Load() SQR3   { return SQR3(rm.UM32.Load()) }
func (rm RMSQR3) Store(b SQR3) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ10() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ10)}}
}

func (p *ADC_Periph) SQ11() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ11)}}
}

func (p *ADC_Periph) SQ12() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ12)}}
}

func (p *ADC_Periph) SQ13() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ13)}}
}

func (p *ADC_Periph) SQ14() RMSQR3 {
	return RMSQR3{mmio.UM32{&p.SQR3.U32, uint32(SQ14)}}
}

type SQR4 uint32

func (b SQR4) Field(mask SQR4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR4) J(v int) SQR4 {
	return SQR4(bits.Make32(v, uint32(mask)))
}

type RSQR4 struct{ mmio.U32 }

func (r *RSQR4) Bits(mask SQR4) SQR4    { return SQR4(r.U32.Bits(uint32(mask))) }
func (r *RSQR4) StoreBits(mask, b SQR4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSQR4) SetBits(mask SQR4)      { r.U32.SetBits(uint32(mask)) }
func (r *RSQR4) ClearBits(mask SQR4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSQR4) Load() SQR4             { return SQR4(r.U32.Load()) }
func (r *RSQR4) Store(b SQR4)           { r.U32.Store(uint32(b)) }

func (r *RSQR4) AtomicStoreBits(mask, b SQR4) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSQR4) AtomicSetBits(mask SQR4)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSQR4) AtomicClearBits(mask SQR4)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSQR4 struct{ mmio.UM32 }

func (rm RMSQR4) Load() SQR4   { return SQR4(rm.UM32.Load()) }
func (rm RMSQR4) Store(b SQR4) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ15() RMSQR4 {
	return RMSQR4{mmio.UM32{&p.SQR4.U32, uint32(SQ15)}}
}

func (p *ADC_Periph) SQ16() RMSQR4 {
	return RMSQR4{mmio.UM32{&p.SQR4.U32, uint32(SQ16)}}
}

type DR uint32

func (b DR) Field(mask DR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR) J(v int) DR {
	return DR(bits.Make32(v, uint32(mask)))
}

type RDR struct{ mmio.U32 }

func (r *RDR) Bits(mask DR) DR      { return DR(r.U32.Bits(uint32(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDR) SetBits(mask DR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDR) Load() DR             { return DR(r.U32.Load()) }
func (r *RDR) Store(b DR)           { r.U32.Store(uint32(b)) }

func (r *RDR) AtomicStoreBits(mask, b DR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDR) AtomicSetBits(mask DR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDR) AtomicClearBits(mask DR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDR struct{ mmio.UM32 }

func (rm RMDR) Load() DR   { return DR(rm.UM32.Load()) }
func (rm RMDR) Store(b DR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) RDATA() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(RDATA)}}
}

type JSQR uint32

func (b JSQR) Field(mask JSQR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JSQR) J(v int) JSQR {
	return JSQR(bits.Make32(v, uint32(mask)))
}

type RJSQR struct{ mmio.U32 }

func (r *RJSQR) Bits(mask JSQR) JSQR    { return JSQR(r.U32.Bits(uint32(mask))) }
func (r *RJSQR) StoreBits(mask, b JSQR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RJSQR) SetBits(mask JSQR)      { r.U32.SetBits(uint32(mask)) }
func (r *RJSQR) ClearBits(mask JSQR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RJSQR) Load() JSQR             { return JSQR(r.U32.Load()) }
func (r *RJSQR) Store(b JSQR)           { r.U32.Store(uint32(b)) }

func (r *RJSQR) AtomicStoreBits(mask, b JSQR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RJSQR) AtomicSetBits(mask JSQR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RJSQR) AtomicClearBits(mask JSQR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMJSQR struct{ mmio.UM32 }

func (rm RMJSQR) Load() JSQR   { return JSQR(rm.UM32.Load()) }
func (rm RMJSQR) Store(b JSQR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JL() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JL)}}
}

func (p *ADC_Periph) JEXTSEL() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JEXTSEL)}}
}

func (p *ADC_Periph) JEXTEN() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JEXTEN)}}
}

func (p *ADC_Periph) JSQ1() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JSQ1)}}
}

func (p *ADC_Periph) JSQ2() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JSQ2)}}
}

func (p *ADC_Periph) JSQ3() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JSQ3)}}
}

func (p *ADC_Periph) JSQ4() RMJSQR {
	return RMJSQR{mmio.UM32{&p.JSQR.U32, uint32(JSQ4)}}
}

type OFR uint32

func (b OFR) Field(mask OFR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OFR) J(v int) OFR {
	return OFR(bits.Make32(v, uint32(mask)))
}

type ROFR struct{ mmio.U32 }

func (r *ROFR) Bits(mask OFR) OFR     { return OFR(r.U32.Bits(uint32(mask))) }
func (r *ROFR) StoreBits(mask, b OFR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROFR) SetBits(mask OFR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROFR) ClearBits(mask OFR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROFR) Load() OFR             { return OFR(r.U32.Load()) }
func (r *ROFR) Store(b OFR)           { r.U32.Store(uint32(b)) }

func (r *ROFR) AtomicStoreBits(mask, b OFR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROFR) AtomicSetBits(mask OFR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROFR) AtomicClearBits(mask OFR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOFR struct{ mmio.UM32 }

func (rm RMOFR) Load() OFR   { return OFR(rm.UM32.Load()) }
func (rm RMOFR) Store(b OFR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) OFFSET1(n int) RMOFR {
	return RMOFR{mmio.UM32{&p.OFR[n].U32, uint32(OFFSET1)}}
}

func (p *ADC_Periph) OFFSET1_CH(n int) RMOFR {
	return RMOFR{mmio.UM32{&p.OFR[n].U32, uint32(OFFSET1_CH)}}
}

func (p *ADC_Periph) OFFSET1_EN(n int) RMOFR {
	return RMOFR{mmio.UM32{&p.OFR[n].U32, uint32(OFFSET1_EN)}}
}

type JDR uint32

func (b JDR) Field(mask JDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JDR) J(v int) JDR {
	return JDR(bits.Make32(v, uint32(mask)))
}

type RJDR struct{ mmio.U32 }

func (r *RJDR) Bits(mask JDR) JDR     { return JDR(r.U32.Bits(uint32(mask))) }
func (r *RJDR) StoreBits(mask, b JDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RJDR) SetBits(mask JDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RJDR) ClearBits(mask JDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RJDR) Load() JDR             { return JDR(r.U32.Load()) }
func (r *RJDR) Store(b JDR)           { r.U32.Store(uint32(b)) }

func (r *RJDR) AtomicStoreBits(mask, b JDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RJDR) AtomicSetBits(mask JDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RJDR) AtomicClearBits(mask JDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMJDR struct{ mmio.UM32 }

func (rm RMJDR) Load() JDR   { return JDR(rm.UM32.Load()) }
func (rm RMJDR) Store(b JDR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JDATA(n int) RMJDR {
	return RMJDR{mmio.UM32{&p.JDR[n].U32, uint32(JDATA)}}
}

type AWD2CR uint32

func (b AWD2CR) Field(mask AWD2CR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AWD2CR) J(v int) AWD2CR {
	return AWD2CR(bits.Make32(v, uint32(mask)))
}

type RAWD2CR struct{ mmio.U32 }

func (r *RAWD2CR) Bits(mask AWD2CR) AWD2CR  { return AWD2CR(r.U32.Bits(uint32(mask))) }
func (r *RAWD2CR) StoreBits(mask, b AWD2CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAWD2CR) SetBits(mask AWD2CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAWD2CR) ClearBits(mask AWD2CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAWD2CR) Load() AWD2CR             { return AWD2CR(r.U32.Load()) }
func (r *RAWD2CR) Store(b AWD2CR)           { r.U32.Store(uint32(b)) }

func (r *RAWD2CR) AtomicStoreBits(mask, b AWD2CR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAWD2CR) AtomicSetBits(mask AWD2CR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAWD2CR) AtomicClearBits(mask AWD2CR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAWD2CR struct{ mmio.UM32 }

func (rm RMAWD2CR) Load() AWD2CR   { return AWD2CR(rm.UM32.Load()) }
func (rm RMAWD2CR) Store(b AWD2CR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) AWD2CH() RMAWD2CR {
	return RMAWD2CR{mmio.UM32{&p.AWD2CR.U32, uint32(AWD2CH)}}
}

type AWD3CR uint32

func (b AWD3CR) Field(mask AWD3CR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AWD3CR) J(v int) AWD3CR {
	return AWD3CR(bits.Make32(v, uint32(mask)))
}

type RAWD3CR struct{ mmio.U32 }

func (r *RAWD3CR) Bits(mask AWD3CR) AWD3CR  { return AWD3CR(r.U32.Bits(uint32(mask))) }
func (r *RAWD3CR) StoreBits(mask, b AWD3CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAWD3CR) SetBits(mask AWD3CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAWD3CR) ClearBits(mask AWD3CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAWD3CR) Load() AWD3CR             { return AWD3CR(r.U32.Load()) }
func (r *RAWD3CR) Store(b AWD3CR)           { r.U32.Store(uint32(b)) }

func (r *RAWD3CR) AtomicStoreBits(mask, b AWD3CR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAWD3CR) AtomicSetBits(mask AWD3CR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAWD3CR) AtomicClearBits(mask AWD3CR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAWD3CR struct{ mmio.UM32 }

func (rm RMAWD3CR) Load() AWD3CR   { return AWD3CR(rm.UM32.Load()) }
func (rm RMAWD3CR) Store(b AWD3CR) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) AWD3CH() RMAWD3CR {
	return RMAWD3CR{mmio.UM32{&p.AWD3CR.U32, uint32(AWD3CH)}}
}

type DIFSEL uint32

func (b DIFSEL) Field(mask DIFSEL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIFSEL) J(v int) DIFSEL {
	return DIFSEL(bits.Make32(v, uint32(mask)))
}

type RDIFSEL struct{ mmio.U32 }

func (r *RDIFSEL) Bits(mask DIFSEL) DIFSEL  { return DIFSEL(r.U32.Bits(uint32(mask))) }
func (r *RDIFSEL) StoreBits(mask, b DIFSEL) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIFSEL) SetBits(mask DIFSEL)      { r.U32.SetBits(uint32(mask)) }
func (r *RDIFSEL) ClearBits(mask DIFSEL)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDIFSEL) Load() DIFSEL             { return DIFSEL(r.U32.Load()) }
func (r *RDIFSEL) Store(b DIFSEL)           { r.U32.Store(uint32(b)) }

func (r *RDIFSEL) AtomicStoreBits(mask, b DIFSEL) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDIFSEL) AtomicSetBits(mask DIFSEL)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDIFSEL) AtomicClearBits(mask DIFSEL)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDIFSEL struct{ mmio.UM32 }

func (rm RMDIFSEL) Load() DIFSEL   { return DIFSEL(rm.UM32.Load()) }
func (rm RMDIFSEL) Store(b DIFSEL) { rm.UM32.Store(uint32(b)) }

type CALFACT uint32

func (b CALFACT) Field(mask CALFACT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CALFACT) J(v int) CALFACT {
	return CALFACT(bits.Make32(v, uint32(mask)))
}

type RCALFACT struct{ mmio.U32 }

func (r *RCALFACT) Bits(mask CALFACT) CALFACT { return CALFACT(r.U32.Bits(uint32(mask))) }
func (r *RCALFACT) StoreBits(mask, b CALFACT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCALFACT) SetBits(mask CALFACT)      { r.U32.SetBits(uint32(mask)) }
func (r *RCALFACT) ClearBits(mask CALFACT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCALFACT) Load() CALFACT             { return CALFACT(r.U32.Load()) }
func (r *RCALFACT) Store(b CALFACT)           { r.U32.Store(uint32(b)) }

func (r *RCALFACT) AtomicStoreBits(mask, b CALFACT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCALFACT) AtomicSetBits(mask CALFACT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCALFACT) AtomicClearBits(mask CALFACT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCALFACT struct{ mmio.UM32 }

func (rm RMCALFACT) Load() CALFACT   { return CALFACT(rm.UM32.Load()) }
func (rm RMCALFACT) Store(b CALFACT) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) CALFACT_S() RMCALFACT {
	return RMCALFACT{mmio.UM32{&p.CALFACT.U32, uint32(CALFACT_S)}}
}

func (p *ADC_Periph) CALFACT_D() RMCALFACT {
	return RMCALFACT{mmio.UM32{&p.CALFACT.U32, uint32(CALFACT_D)}}
}
