
Remember the Quiz exercise, we had the questions and answers stored in map.
This time will read the questions and answers from CSV file using "encoding/csv" package

csv package has NewReader function which creates csv.Reader using given io.Reader interface 
So it accepts the types which implements io.Reader interface (ex File pointer)


create a game package to read csv file and stores Questions and Answers in []Quiz 
1. create Quiz struct with Question and Answer Fields

2. create func NewQuiz to return []Quiz after reading from csv file

	1. open file using os.Open
	2. create csv reader using file pointer
	3. use Read function
	4. Load the Questions and Answers in Quiz struct and append it to slice

main function call game.NewQuiz function to get the Question and answers
use the questions in []Quiz and compare the given answer with user input

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.
```
What is 5+5?
10
What is 1+1?
2
What is 8+3?
13


You got 10/12 
```
