package goval_test

import (
	"fmt"

	"github.com/XHao/goval/pkg/goval"
)

// ExampleEvaluate demonstrates a rule-engine scenario: apply a discount rule
// to an order based on user level and order amount.
func ExampleEvaluate() {
	// Rule: VIP users with orders over 100 get 20% off.
	rule := `
var Order = (amount, userId) -> {
    amount: amount,
    userId: userId,
    discounted: (rate) -> amount * rate
}
var User = (id, level) -> { id: id, level: level }

var users = [
    User("u1", "gold"),
    User("u2", "vip")
]
var userMap = reduce(users, {}, (acc, u) -> put(acc, u.id, u))

var order = Order(500, "u2")
var user = userMap[order.userId]

user.level == "vip" && order.amount > 100
    ? order.discounted(0.8)
    : order.amount
`

	result, err := goval.Evaluate(rule, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("final price: %v\n", result)

	// Output:
	// final price: 400
}

// ExampleEvaluate_context demonstrates passing Go values as rule context.
func ExampleEvaluate_context() {
	// A simple threshold rule driven by external context.
	rule := "amount >= threshold ? amount * discount : amount"

	ctx := map[string]interface{}{
		"amount":   int64(300),
		"threshold": int64(100),
		"discount":  float64(0.9),
	}

	result, err := goval.Evaluate(rule, ctx)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("result: %v\n", result)

	// Output:
	// result: 270
}

// ExampleEvaluate_listOps demonstrates immutable list operations.
func ExampleEvaluate_listOps() {
	rule := `
var nums = [1, 2, 3, 4, 5, 6]
var doubled = map(nums, (x) -> x * 2)
var evens = filter(doubled, (x) -> x % 2 == 0)
var sum = reduce(evens, 0, (acc, x) -> acc + x)
sum
`

	result, err := goval.Evaluate(rule, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("sum of even doubled: %v\n", result)

	// Output:
	// sum of even doubled: 42
}

// ExampleEvaluate_closure demonstrates lambda closures.
func ExampleEvaluate_closure() {
	rule := `
var makeMultiplier = (factor) -> (x) -> x * factor
var triple = makeMultiplier(3)
triple(14)
`

	result, err := goval.Evaluate(rule, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("triple(14) = %v\n", result)

	// Output:
	// triple(14) = 42
}
