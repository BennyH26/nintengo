package nes

import (
	"github.com/nwidger/rp2ago3"
)

type NROM struct {
	*ROMFile
}

func NewNROM(romf *ROMFile) *NROM {
	return &NROM{ROMFile: romf}
}

func (nrom *NROM) Mappings(which rp2ago3.Mapping) (fetch, store []uint16) {
	fetch = []uint16{}
	store = []uint16{}

	switch which {
	case rp2ago3.PPU:
		// CHR bank 1
		for i := uint32(0x0000); i <= 0x0fff; i++ {
			fetch = append(fetch, uint16(i))
		}

		// CHR bank 2
		for i := uint32(0x1000); i <= 0x1fff; i++ {
			fetch = append(fetch, uint16(i))
		}
	case rp2ago3.CPU:
		// PRG bank 1
		for i := uint32(0x8000); i <= 0xbfff; i++ {
			fetch = append(fetch, uint16(i))
		}

		// PRG bank 2
		for i := uint32(0xc000); i <= 0xffff; i++ {
			fetch = append(fetch, uint16(i))
		}
	}

	return
}

func (nrom *NROM) Reset() {

}

func (nrom *NROM) Fetch(address uint16) (value uint8) {
	switch {
	// PPU only
	case address >= 0x0000 && address <= 0x1fff:
		index := address & 0x0fff

		switch {
		// CHR bank 1
		case address >= 0x0000 && address <= 0x0fff:
			if nrom.ROMFile.chrBanks > 0 {
				value = nrom.ROMFile.vromBanks[0][index]
			}
		// CHR bank 2
		case address >= 0x1000 && address <= 0x1fff:
			if nrom.ROMFile.chrBanks > 0 {
				value = nrom.ROMFile.vromBanks[nrom.ROMFile.chrBanks-1][index]
			}
		}
	// CPU only
	case address >= 0x8000 && address <= 0xffff:
		index := address & 0x3fff

		switch {
		// PRG bank 1
		case address >= 0x8000 && address <= 0xbfff:
			if nrom.ROMFile.prgBanks > 0 {
				value = nrom.ROMFile.romBanks[0][index]
			}
		// PRG bank 2
		case address >= 0xc000 && address <= 0xffff:
			if nrom.ROMFile.prgBanks > 0 {
				value = nrom.ROMFile.romBanks[nrom.ROMFile.prgBanks-1][index]
			}
		}
	}

	return
}

func (nrom *NROM) Store(address uint16, value uint8) (oldValue uint8) {
	index := address & 0x0fff

	// PPU only
	switch {
	// CHR bank 1
	case address >= 0x0000 && address <= 0x0fff:
		if nrom.ROMFile.chrBanks > 0 {
			oldValue = nrom.ROMFile.vromBanks[0][index]
			nrom.ROMFile.vromBanks[0][index] = value
		}
	// CHR bank 2
	case address >= 0x1000 && address <= 0x1fff:
		if nrom.ROMFile.chrBanks > 0 {
			oldValue = nrom.ROMFile.vromBanks[nrom.ROMFile.chrBanks-1][index]
			nrom.ROMFile.vromBanks[nrom.ROMFile.chrBanks-1][index] = value
		}
	}

	return
}