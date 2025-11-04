# superman-chicken-rescue

# instructions
- commands
    - make run
        - run the prorgam
    - make test 
        - run the test

# solution approach
- As the first solution which comes to my mind is to use iteration approach, that solution uses brute force with nested iteration, resulting in the time complexity of O(n^2). I try to figure how to reduce the time complexity to just O(n). The solution is to use recursion. This reduces the time comlexity from O(n^2) to O(n). However, the trade-off of the recursive approach is that it uses more resources.

# solution explanation
- first block (line 6 - 11)
    - the first condition is for a base case which will stop the recursion.
    - inside the codition, there is another condition checking if it is the first stack call and the length of an input is one. Then stop recursion and return the maxCoveredChicken which is one.

- second block (line 13 - 27)
    - the first condition in this section is to check whether the next position is beyond the coverRange or not.
    - if it is beyond the coverRange, 


    - the 'if' checks if next position is beyond the coverRange.
    - if yes (no chickens within the coverRange), slice the positions by starting from the next position onwards.
    - if no (execute the loop), check how many chikens are in the coverRange.
        - inside the loop, there is another 'if' checking if the chikens is within the coverRange so the loop will iterate only n times (roofSize)
        - the logic above prevents 'out of range' error when the provided roofSize is less than the length of current sliced position.
            - if yes, stop the loop.
        - the next 'if' does actual logic determining how many chickens are covered.
            - if the position is less than the coverRange, it is covered
            - add 1 to chickenCount
    - at the end of this block, prepare the next sliced position by slicing the positions starting from the next position onwards.
- last block (line 29 - 31)
    - check if current chickenCount is greater than the current maxCover
    - if yes, assign a new value of chickenCount to maxCover
- call recursive function

# time & space explanation
1. time complexity
    - O(1)
        - The time complexity can be O(1) for the best case scenario in which an input's length is one and the roof the Superman can carries has at least 1 in length. When the recusion starts, the program stops as there is no more data to process.

        Therefore, for the best case, this program can achieve O(1)

    - O(n)
        - For the average and worst case scenarios, this program can achieve O(n) as the time processing depends upon the size of the input

        To check if the current position of the Superman who is carrying the roof can cover how many chickens requires some steps of logic for each call stack.

        In addition, there is a logic checking if current position of the chicken where the Superman can stand and carry the roof is less than the next position, the loop will continue to the next position because there is no chicken within the range. This logic can reduce the time of processing but the time complexity is still O(n).

        Therefore, O(n) is the time complexity of this program for average and worst cases

2. space complexity
    - O(n)
        - The space complexity of this solution is O(n). In this recursive function, a slice is used to store chickens' postions. One of the advantages of such data type is that it is a refernce type. When passing a slice to the next call stack, what is passed is a pointer to that data. So, no data is copied, just a few resouces are used here. 

        Even though the input is long, the space used here in this recursion does not increase as the the next stack is called.

        Therefore, the space complexity of this solution is O(n)

# input constraints