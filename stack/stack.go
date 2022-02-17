// Package "stack" is an implementation of the stack data structure
package stack

import (
    "sync"
    "errors"
)

// Sructure describing the stack
type Stack struct {
	lock sync.Mutex
	buff []interface{}
}

// Stack creates a new stack instance
// Return *Stack
func NewStack() *Stack {
	return &Stack { sync.Mutex{}, make([]interface{}, 0)}	
}

// Push adds a new element to the top of the stack
func (stk *Stack) Push(val interface{}) {

	stk.lock.Lock()
	defer stk.lock.Unlock()

	stk.buff = append(stk.buff, val)
}

// Pop removes an element at the top of the stack
// Top returns the element at the top of the stack and the value of error
func (stk *Stack) Pop() (interface{}, error) {

	stk.lock.Lock()
	defer stk.lock.Unlock()

	length := len(stk.buff)

	if length == 0 { 
	    return 0, errors.New("Pop function: Stack is empty")
	}

	lastValue := stk.buff[length - 1]
	stk.buff   = stk.buff[:length - 1]

	return lastValue, nil
}

// Top returns the element at the top of the stack and the value of error
func (stk *Stack) Top() (interface{}, error) {

	stk.lock.Lock()
	defer stk.lock.Unlock()

	length := len(stk.buff)

	if length == 0 { 
	    return 0, errors.New("Top function: Stack is empty")
	}

	return stk.buff[length - 1], nil

}

// Len returns the current length of stack
func (stk *Stack) Len() int {

	stk.lock.Lock()
	defer stk.lock.Unlock()

	return len(stk.buff)
}

// Empty checks if the stack is empty
// Returns the bool value. The True - is empty, The False - is not empty
func (stk *Stack) Empty() bool {
 
	stk.lock.Lock()
	defer stk.lock.Unlock()

	return (!len(stk.buff))
}

// Clear clears the stack. All values are deleted
func (stk *Stack) Clear() {
  
  stk.lock.Lock()
  defer stk.lock.Unlock()
  
  stk.buff = nil
}
