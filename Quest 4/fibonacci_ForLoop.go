package piscine

func Fibonacci1(index int) int {
    // Handle edge cases
    if index < 0 {
        return -1
    }
    if index == 0 {
        return 0
    }
    if index == 1 {
        return 1
    }

    // Initialize the first two numbers of the sequence
    a := 0 // represents Fibonacci(0)
    b := 1 // represents Fibonacci(1)

    // Start looping from 2 up to the target index
    for i := 2; i <= index; i++ {
        next := a + b // Calculate the next Fibonacci number
        a = b         // Move 'b' into the 'a' position for the next round
        b = next      // Move 'next' into the 'b' position
    }

    return b
}