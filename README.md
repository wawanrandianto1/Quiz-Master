# Quiz Master

### Quiz master is simple game using bash script and golang. 
<br>
<br>
 
## Installation Instructions
1. go to main user directory in linux
 
       cd ~
 
2. clone the repository
  
       git clone https://github.com/wawanrandianto1/Quiz-Master.git ~/quiz_master
  
   this will clone repository to "Quiz-Master" folder
 
3. create bin folder, and go inside bin folder
 
       mkdir bin
       cd bin/
 
4. copy 2 script executable file in bash_script folder to bin folder and add permission to executable file
 
       sudo chmod 755 setup quiz_master
 
5. run command, to compile the code

       bin/setup

6. last step, run command and enjoy the quiz

       bin/quiz_master


## How to play
- you can input question and answer the question, example:

> $ create_question 1 “How many letters are there in the English alphabet?” 26

> $ create_question 2 “How many vowels are there in the English alphabet?” 5

> $ answer_question 2 5

> Correct!

> $ exit

- see "help" command for better experience