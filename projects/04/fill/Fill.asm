// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.
//
//  for {
//      if (keypress) {
//        color black
//        goto for
//      }
//      color white
//  }
//
//

// Put your code here.
@8192
D = A

@R0
M = D

(LOOP)
    @R0
    D = M
    @R1
    M = D

    @KBD
    D = M

    @WHITE
    D;JEQ

    (BLACK)
         @R1
         D = M
         @SCREEN
         A = D + A
         M = -1
         @R1
         D = M
         D = D - 1
         M = D
         @BLACK
         D;JGE

    @LOOP
    0;JMP

    (WHITE)
         @R1
         D = M
         @SCREEN
         A = D + A
         M = 0
         @R1
         D = M
         D = D - 1
         M = D
         @WHITE
         D;JGE

    @LOOP
    0;JMP
