/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-04
 * @fileoverview How long it will take the user to pay off post-secondary education.
 */

// declare variables
let years: number = 0;

// input
// starting amount in bank account
const startingAmountString: string = prompt("Enter the starting bank account amount: ") || ("\n");
let startingAmountNumber: number = parseInt(startingAmountString);

// interest rate
const rateString: string = prompt("Enter the yearly interest rate (as a percentage): ") || ("\n");
let rateNumber: number = parseInt(rateString);

// amount needed for post-secondary
const requiredAmountString: string = prompt("Enter the amount needed for post-secondary education: ") || ("\n");
const requiredAmountNumber: number = parseInt(requiredAmountString);

// processing
// turn rate into decimal
rateNumber = rateNumber / 100;

// loop starts
while (startingAmountNumber < requiredAmountNumber) {
  startingAmountNumber = startingAmountNumber + (startingAmountNumber * rateNumber);
  years++;
	};

// output
console.log(`It will take ${years} for the starting bank account to reach the required amount for post-secondary education. `);

console.log("\nDone.");
