# test-gits

Golang
How to run
Test 1 : go run test_1.go string queries
Example : go run test_1.go abbccc 1,4,9
Test 2 : go run test_2.go number max_number_change
Example : go run test_2.go 3943 1
Test 3 : go run test_3.go "string"
Example : go run test_3.go "{()}"


Test 3 Explanation:
for detect the open and close bracket, iam using map and array
using key and value of the map
open bracket as key of map and close bracket as value
and it func getBracketType will check it is a open or closed bracket
and using array for queueing the unpair bracket
and will slice of the array queue if close bracket found the pair

if queue not empty it will return NO
and return YES if queue is empty
