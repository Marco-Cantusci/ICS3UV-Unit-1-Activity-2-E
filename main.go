/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-09-25
 * @fileoverview Math problem
 */

package main

import "fmt"


func main() {

	// print question
	fmt.Println("The dogs weight is 34.5 Kilograms, what is its weight in pounds?");
	
	// conversion of 1kg to pounds
	fmt.Println("1kg = 2.2 pounds");

	fmt.Println("To do this, we will need to multiply the dogs weight by 2.2.");

	// The math
	fmt.Println("34.5 * 2.2 = " + fmt.Sprint(34.5 * 2.2));

	fmt.Println("\nDone.");
}
