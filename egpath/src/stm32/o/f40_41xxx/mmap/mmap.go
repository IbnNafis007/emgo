// Package mmap provides base memory adresses for all peripherals.
package mmap

const (
	FLASH_BASE         uintptr = 0x08000000 // FLASH(up to 1 MB) base address in the alias region
	CCMDATARAM_BASE    uintptr = 0x10000000 // CCM(core coupled memory) data RAM(64 KB) base address in the alias region
	SRAM1_BASE         uintptr = 0x20000000 // SRAM1(112 KB) base address in the alias region
	SRAM2_BASE         uintptr = 0x2001C000 // SRAM2(16 KB) base address in the alias region
	SRAM3_BASE         uintptr = 0x20020000 // SRAM3(64 KB) base address in the alias region
	PERIPH_BASE        uintptr = 0x40000000 // Peripheral base address in the alias region
	BKPSRAM_BASE       uintptr = 0x40024000 // Backup SRAM(4 KB) base address in the alias region
	FSMC_R_BASE        uintptr = 0xA0000000 // FSMC registers base address
	CCMDATARAM_BB_BASE uintptr = 0x12000000 // CCM(core coupled memory) data RAM(64 KB) base address in the bit-band region
	SRAM1_BB_BASE      uintptr = 0x22000000 // SRAM1(112 KB) base address in the bit-band region
	SRAM2_BB_BASE      uintptr = 0x2201C000 // SRAM2(16 KB) base address in the bit-band region
	SRAM3_BB_BASE      uintptr = 0x22400000 // SRAM3(64 KB) base address in the bit-band region
	PERIPH_BB_BASE     uintptr = 0x42000000 // Peripheral base address in the bit-band region
	BKPSRAM_BB_BASE    uintptr = 0x42024000 // Backup SRAM(4 KB) base address in the bit-band region
	SRAM_BASE          uintptr = SRAM1_BASE
	SRAM_BB_BASE       uintptr = SRAM1_BB_BASE
)

// Peripheral memory map
const (
	APB1PERIPH_BASE uintptr = PERIPH_BASE
	APB2PERIPH_BASE uintptr = PERIPH_BASE + 0x00010000
	AHB1PERIPH_BASE uintptr = PERIPH_BASE + 0x00020000
	AHB2PERIPH_BASE uintptr = PERIPH_BASE + 0x10000000
)

// APB1 peripherals
const (
	TIM2_BASE    uintptr = APB1PERIPH_BASE + 0x0000
	TIM3_BASE    uintptr = APB1PERIPH_BASE + 0x0400
	TIM4_BASE    uintptr = APB1PERIPH_BASE + 0x0800
	TIM5_BASE    uintptr = APB1PERIPH_BASE + 0x0C00
	TIM6_BASE    uintptr = APB1PERIPH_BASE + 0x1000
	TIM7_BASE    uintptr = APB1PERIPH_BASE + 0x1400
	TIM12_BASE   uintptr = APB1PERIPH_BASE + 0x1800
	TIM13_BASE   uintptr = APB1PERIPH_BASE + 0x1C00
	TIM14_BASE   uintptr = APB1PERIPH_BASE + 0x2000
	RTC_BASE     uintptr = APB1PERIPH_BASE + 0x2800
	WWDG_BASE    uintptr = APB1PERIPH_BASE + 0x2C00
	IWDG_BASE    uintptr = APB1PERIPH_BASE + 0x3000
	I2S2ext_BASE uintptr = APB1PERIPH_BASE + 0x3400
	SPI2_BASE    uintptr = APB1PERIPH_BASE + 0x3800
	SPI3_BASE    uintptr = APB1PERIPH_BASE + 0x3C00
	I2S3ext_BASE uintptr = APB1PERIPH_BASE + 0x4000
	USART2_BASE  uintptr = APB1PERIPH_BASE + 0x4400
	USART3_BASE  uintptr = APB1PERIPH_BASE + 0x4800
	UART4_BASE   uintptr = APB1PERIPH_BASE + 0x4C00
	UART5_BASE   uintptr = APB1PERIPH_BASE + 0x5000
	I2C1_BASE    uintptr = APB1PERIPH_BASE + 0x5400
	I2C2_BASE    uintptr = APB1PERIPH_BASE + 0x5800
	I2C3_BASE    uintptr = APB1PERIPH_BASE + 0x5C00
	CAN1_BASE    uintptr = APB1PERIPH_BASE + 0x6400
	CAN2_BASE    uintptr = APB1PERIPH_BASE + 0x6800
	PWR_BASE     uintptr = APB1PERIPH_BASE + 0x7000
	DAC_BASE     uintptr = APB1PERIPH_BASE + 0x7400
	UART7_BASE   uintptr = APB1PERIPH_BASE + 0x7800
	UART8_BASE   uintptr = APB1PERIPH_BASE + 0x7C00
)

// APB2 peripherals
const (
	TIM1_BASE         uintptr = APB2PERIPH_BASE + 0x0000
	TIM8_BASE         uintptr = APB2PERIPH_BASE + 0x0400
	USART1_BASE       uintptr = APB2PERIPH_BASE + 0x1000
	USART6_BASE       uintptr = APB2PERIPH_BASE + 0x1400
	ADC1_BASE         uintptr = APB2PERIPH_BASE + 0x2000
	ADC2_BASE         uintptr = APB2PERIPH_BASE + 0x2100
	ADC3_BASE         uintptr = APB2PERIPH_BASE + 0x2200
	ADC_BASE          uintptr = APB2PERIPH_BASE + 0x2300
	SDIO_BASE         uintptr = APB2PERIPH_BASE + 0x2C00
	SPI1_BASE         uintptr = APB2PERIPH_BASE + 0x3000
	SPI4_BASE         uintptr = APB2PERIPH_BASE + 0x3400
	SYSCFG_BASE       uintptr = APB2PERIPH_BASE + 0x3800
	EXTI_BASE         uintptr = APB2PERIPH_BASE + 0x3C00
	TIM9_BASE         uintptr = APB2PERIPH_BASE + 0x4000
	TIM10_BASE        uintptr = APB2PERIPH_BASE + 0x4400
	TIM11_BASE        uintptr = APB2PERIPH_BASE + 0x4800
	SPI5_BASE         uintptr = APB2PERIPH_BASE + 0x5000
	SPI6_BASE         uintptr = APB2PERIPH_BASE + 0x5400
	SAI1_BASE         uintptr = APB2PERIPH_BASE + 0x5800
	SAI1_Block_A_BASE uintptr = SAI1_BASE + 0x004
	SAI1_Block_B_BASE uintptr = SAI1_BASE + 0x024
	LTDC_BASE         uintptr = APB2PERIPH_BASE + 0x6800
	LTDC_Layer1_BASE  uintptr = LTDC_BASE + 0x84
	LTDC_Layer2_BASE  uintptr = LTDC_BASE + 0x104
)

// AHB1 peripherals
const (
	GPIOA_BASE        uintptr = AHB1PERIPH_BASE + 0x0000
	GPIOB_BASE        uintptr = AHB1PERIPH_BASE + 0x0400
	GPIOC_BASE        uintptr = AHB1PERIPH_BASE + 0x0800
	GPIOD_BASE        uintptr = AHB1PERIPH_BASE + 0x0C00
	GPIOE_BASE        uintptr = AHB1PERIPH_BASE + 0x1000
	GPIOF_BASE        uintptr = AHB1PERIPH_BASE + 0x1400
	GPIOG_BASE        uintptr = AHB1PERIPH_BASE + 0x1800
	GPIOH_BASE        uintptr = AHB1PERIPH_BASE + 0x1C00
	GPIOI_BASE        uintptr = AHB1PERIPH_BASE + 0x2000
	GPIOJ_BASE        uintptr = AHB1PERIPH_BASE + 0x2400
	GPIOK_BASE        uintptr = AHB1PERIPH_BASE + 0x2800
	CRC_BASE          uintptr = AHB1PERIPH_BASE + 0x3000
	RCC_BASE          uintptr = AHB1PERIPH_BASE + 0x3800
	FLASH_R_BASE      uintptr = AHB1PERIPH_BASE + 0x3C00
	DMA1_BASE         uintptr = AHB1PERIPH_BASE + 0x6000
	DMA1_Stream0_BASE uintptr = DMA1_BASE + 0x010
	DMA1_Stream1_BASE uintptr = DMA1_BASE + 0x028
	DMA1_Stream2_BASE uintptr = DMA1_BASE + 0x040
	DMA1_Stream3_BASE uintptr = DMA1_BASE + 0x058
	DMA1_Stream4_BASE uintptr = DMA1_BASE + 0x070
	DMA1_Stream5_BASE uintptr = DMA1_BASE + 0x088
	DMA1_Stream6_BASE uintptr = DMA1_BASE + 0x0A0
	DMA1_Stream7_BASE uintptr = DMA1_BASE + 0x0B8
	DMA2_BASE         uintptr = AHB1PERIPH_BASE + 0x6400
	DMA2_Stream0_BASE uintptr = DMA2_BASE + 0x010
	DMA2_Stream1_BASE uintptr = DMA2_BASE + 0x028
	DMA2_Stream2_BASE uintptr = DMA2_BASE + 0x040
	DMA2_Stream3_BASE uintptr = DMA2_BASE + 0x058
	DMA2_Stream4_BASE uintptr = DMA2_BASE + 0x070
	DMA2_Stream5_BASE uintptr = DMA2_BASE + 0x088
	DMA2_Stream6_BASE uintptr = DMA2_BASE + 0x0A0
	DMA2_Stream7_BASE uintptr = DMA2_BASE + 0x0B8
	ETH_BASE          uintptr = AHB1PERIPH_BASE + 0x8000
	ETH_MAC_BASE      uintptr = ETH_BASE
	ETH_MMC_BASE      uintptr = ETH_BASE + 0x0100
	ETH_PTP_BASE      uintptr = ETH_BASE + 0x0700
	ETH_DMA_BASE      uintptr = ETH_BASE + 0x1000
	DMA2D_BASE        uintptr = AHB1PERIPH_BASE + 0xB000
)

// AHB2 peripherals
const (
	DCMI_BASE        uintptr = AHB2PERIPH_BASE + 0x50000
	CRYP_BASE        uintptr = AHB2PERIPH_BASE + 0x60000
	HASH_BASE        uintptr = AHB2PERIPH_BASE + 0x60400
	HASH_DIGEST_BASE uintptr = AHB2PERIPH_BASE + 0x60710
	RNG_BASE         uintptr = AHB2PERIPH_BASE + 0x60800
)
