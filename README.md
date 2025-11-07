# superman-chicken-rescue

# Instructions
## Commands
- make run
    - run the prorgam
- make test 
    - run the test

# Solution Approach
- As the first solution which comes to my mind is to use iteration approach, that solution uses brute force with nested iteration, resulting in the time complexity of O(n^2). I try to figure how to reduce the time complexity to just O(n). The solution is to use recursion. This reduces the time comlexity from O(n^2) to O(n). However, the trade-off of the recursive approach is that it uses more resources such as overhead.

# Solution Explanation
## First Block (line 6 - 11)
- The first condition is for a base case. If the length of the position list is equal to one, stop the recursion and return maxCoveredChicken.
- Inside the codition, there is another condition checking if it is the first stack call and the length of an input is one. Then stop recursion and return the maxCoveredChicken which is one.

## Second Block (line 13 - 27)
- The first condition in this section is to check whether the next position in position list is beyond the coverRange or not.
    - If it is beyond the coverRange (no chickens within the coverRange), slice the current position list, including only positions starting from the the next position (the first index of the position list) onwards.
    - If it is not, execute the second of this section (else block) which is a loop.

- The loop will execute for n times of roofSize because the maximum number of chickens that could be covered is within the roofSize.
- The first condition in the loop is a logic checking whether the current index if the iteration will not exceed the length of the current sliced position. This prevents 'out of range' error when the provided roofSize is less than the length of current sliced position.
    - If yes, stop the loop.
- The second condition determining how many chickens are covered.
    - If the current position of the chicken is less than the coverRange (the chicken in the current position is covered), add 1 to chickenCount.
    
- At the end of this section, prepare the next sliced position list by slicing the current position list with positions starting from the position index 1 onwards.

## Last Block (line 29 - 31)
- The condition checks if the current chickenCount is greater than the current maxCoveredChicken.
    - If yes, assign a new value of chickenCount to maxCoveredChicken.
## Finally, return the function by calling itself (with the new sliced position list).

# Time & Space Explanation
## Time Complexity
- O(1)
    - The time complexity can be O(1) for the best case scenario in which an input's length is one and the roof the Superman can carries has at least 1 in length. When the recusion starts, the program stops as there is no more data to process.

    - Therefore, for the best case, this program can achieve O(1).

- O(n)
    - For the average and worst case scenarios, this program can achieve O(n) as the time processing depends upon the size of the input.

    - To check if the current position of the Superman who is carrying the roof can cover how many chickens requires some steps of logic for each call stack.

    - In addition, there is a logic checking if current position of the chicken where the Superman can stand and carry the roof is less than the next position, the loop will continue to the next position because there is no chicken within the range. This logic can reduce the time of processing but the time complexity is still O(n).

    - Therefore, O(n) is the time complexity of this program for average and worst cases.

## Space Complexity
- O(n)
    - The space complexity of this solution is O(n). In this recursive function, a slice is used to store chickens' postions. One of the advantages of such data type is that it is a refernce type. When passing a slice to the next call stack, what is passed is a pointer to that data. So, no data is copied, just a few resouces are used here. 

    - Even though the input is long, the space used here in this recursion does not increase as the the next stack is called.

    - Therefore, the space complexity of this solution is O(n).