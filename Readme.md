# Emulator 6502

![Test Workflow](https://github.com/kabi175/6502/actions/workflows/test.yml/badge.svg)

This software component emulates the behaviour of mos6502 processor.

### Introduction

The mos6502 is a processor that operates on 8-bit data at a time and is capable of 16-bit addressing. It is based on Little Endian architecture.

The mos6502 has

-   3 8-bit general purpose Registers
-   a 8-bit Stack Pointer Register
-   a 8-bit Flag Register
-   a 16-bit Program Counter Register

### Registers

#### Accumulator Register

-   8-bit general purpose register
-   Only register that performs math
-   can copy the content of accumulator into memory location.
-   also can copy the content from memory locations.

#### X-Index Register

-   8-bit general purpose register
-   can copy the content of accumulator into memory location.
-   also can copy the content from memory locations.

#### Y-Index Register

-   8-bit general purpose register
-   can copy the content of accumulator into memory location.
-   also can copy the content from memory locations.

#### Status Registers

-   8-bit register, each bit corresponds to a flag.
-   bits are altered depending on the arithimetic and logical operations performed by processor.

| Bit | Flag                  | Usage                                                                                                                                              |
| --- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| 0   | Carry Flag (C)        | This flag contains the carry of MBS in any logical operations.                                                                                     |
| 1   | Zero Flag (Z)         | Zero flag is set if the operation produces 0 as result and not set for non-zero results.                                                           |
| 2   | Intrrupt Flag (I)     | This is an interrupt enable/disable flag. If it is set, interrupts are disabled. If it is cleared, interrupts are enabled.                         |
| 3   | Decimal Mode Flag (D) | When set, and an Add with Carry or Subtract with Carry instruction is executed, the source values are treated as valid BCD (Binary Coded Decimal). |
| 4   | Break Flag (B)        | Set flag when software intrupt is executed.                                                                                                        |
| 5   |                       | Not used. Always remains set.                                                                                                                      |
| 6   | Overflow Flag (V)     | when an arithmetic operation produces a result too large to be represented in a byte, V is set.                                                    |
| 7   | Sign Flag (S)         | Sign Flag is set if the result of an operation is negative, cleared if positive.                                                                   |

#### Stack Pointer

Stack Pointer is a 8-bit register. This register contains the location of the first empty place on the stack. The stack is used for temporary storage by machine language programs, and by the computer.

#### Program Counter

Program Counter is a 16-bit register, that contains the location of the instruction that will be executed next.
