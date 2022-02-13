@17
D = A
@SP
AM = M + 1
A = A - 1
M = D
@17
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE2
D; JEQ
@SP
A = M - 1
M = 0
@END2
D; JMP
(TRUE2)
@SP
A = M - 1
M = -1
(END2)
@17
D = A
@SP
AM = M + 1
A = A - 1
M = D
@16
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE5
D; JEQ
@SP
A = M - 1
M = 0
@END5
D; JMP
(TRUE5)
@SP
A = M - 1
M = -1
(END5)
@16
D = A
@SP
AM = M + 1
A = A - 1
M = D
@17
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE8
D; JEQ
@SP
A = M - 1
M = 0
@END8
D; JMP
(TRUE8)
@SP
A = M - 1
M = -1
(END8)
@892
D = A
@SP
AM = M + 1
A = A - 1
M = D
@891
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE11
D; JLT
@SP
A = M - 1
M = 0
@END11
D; JMP
(TRUE11)
@SP
A = M - 1
M = -1
(END11)
@891
D = A
@SP
AM = M + 1
A = A - 1
M = D
@892
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE14
D; JLT
@SP
A = M - 1
M = 0
@END14
D; JMP
(TRUE14)
@SP
A = M - 1
M = -1
(END14)
@891
D = A
@SP
AM = M + 1
A = A - 1
M = D
@891
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE17
D; JLT
@SP
A = M - 1
M = 0
@END17
D; JMP
(TRUE17)
@SP
A = M - 1
M = -1
(END17)
@32767
D = A
@SP
AM = M + 1
A = A - 1
M = D
@32766
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE20
D; JGT
@SP
A = M - 1
M = 0
@END20
D; JMP
(TRUE20)
@SP
A = M - 1
M = -1
(END20)
@32766
D = A
@SP
AM = M + 1
A = A - 1
M = D
@32767
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE23
D; JGT
@SP
A = M - 1
M = 0
@END23
D; JMP
(TRUE23)
@SP
A = M - 1
M = -1
(END23)
@32766
D = A
@SP
AM = M + 1
A = A - 1
M = D
@32766
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
D = M - D
@TRUE26
D; JGT
@SP
A = M - 1
M = 0
@END26
D; JMP
(TRUE26)
@SP
A = M - 1
M = -1
(END26)
@57
D = A
@SP
AM = M + 1
A = A - 1
M = D
@31
D = A
@SP
AM = M + 1
A = A - 1
M = D
@53
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
M = D + M
@112
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
M = M - D
@SP
A = M - 1
M = -M
@SP
AM = M - 1
D = M
@SP
A = M - 1
M = D & M
@82
D = A
@SP
AM = M + 1
A = A - 1
M = D
@SP
AM = M - 1
D = M
@SP
A = M - 1
M = D | M
@SP
A = M - 1
M = !M
