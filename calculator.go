package main

import "fmt"
import "os"
func inputRaw(){
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() { line = scanner.Text()
            }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
func readNumber(line,index):
		number = 0 
		while index < len(line) and unicode.Isdigit(line[index])
			number = number * 10  + int(line[index])
			index += 1
		if index < len(line) and line[index] == '.' {
			keta = 0.1
			index += 1
		token := map[string]string{ "type":"NUMBER" , "number":number}
		return token, index}
	
func readPlus(line,index)
		token = {'type' : 'PLUS'}
		return token , index + 1

func readMinus(line,index)
		token = {'type' : 'MINUS'}
		return token , index + 1
		
func readTimes(line,index)
		token = {'type' : 'TIMES'}
		return token , index +1
		
func readDivides(line,index)
		token = {'type' : 'DIVIDES'}
		return token , index +1
		
		
func tokenize(line)
		tokens = []
		index = 0
		while index < len (line)
			if line[index].Isdigit	()
				(token, index) = readNumber(line,index)
			else if line[index] == '+'
				(token, index) =readPlus(line,index)
			else if line[index] == '-'
				(token, index) = readMinus(line,index)
			else
				fmt.Printf 'Invalid character found:'+ line[index]
            	exit(1)
        	tokens.append(token)
   		 return tokens
func evaluate(tokens)
		answer = 0 
		tokens.insert(0,{'type': 'PLUS'}) //Insert a dummy '+' token
		index = 1
    while index < len(tokens):
        if tokens[index]['type'] == 'NUMBER':
            if tokens[index - 1]['type'] == 'PLUS':
                answer += tokens[index]['number']
      else if tokens[index - 1]['type'] == 'MINUS':
                answer -= tokens[index]['number']
      else if tokens[index - 1]['type'] == 'TIMES'
      			answer = answer * tokens[index]['number']
      else if tokens[index - 1]['type'] == 'DIVIDES'
      			answer = answer / tokens[index]['number']
            else:
                fmt.Printf 'Invalid syntax'
        index += 1
    return answer
    
	func main(){
fmt.Printf(">")
line = inputRaw
tokens = tokenize(line)
answer = evaluate(tokens)
fmt.Pritnf(answer)}