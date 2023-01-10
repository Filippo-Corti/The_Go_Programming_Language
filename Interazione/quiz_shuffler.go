package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
)

type Question struct {
	question string
	answers []string //First one is always the correct one
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) < 2 {
		fmt.Println("Passare da riga di comando il nome del file contenente le domande")
		os.Exit(1)
	}
	fileName := os.Args[1]
	questions := getQuestions(fileName)
	fmt.Println("Rispondere alle domande con A, B, C. Se si inserisce EXIT il test finisce e ti dà i risultati, altrimenti finisce da solo quando finiscono le domande")
	startQuiz(questions)
}

func getQuestions(fileName string) (questions []Question) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File non trovato")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //Question
		q := Question{}
		q.question = scanner.Text()
		for i := 0; i < 3; i++ {
			scanner.Scan()
			q.answers = append(q.answers, scanner.Text()[3:]) //Answers
		}
		scanner.Scan() //ANSWER: A
		scanner.Scan() //Empty
		questions = append(questions, q)
	}
	return
}

func startQuiz(questions []Question) {
	var rights, wrongs int
	var indexesOfWrongs []int
	for i, question := range questions {
		fmt.Println("---------- Domanda", i + 1, "----------")
		correct, end := askQuestion(question)
		if end {
			break
		}
		if correct {
			rights++
		} else {
			wrongs++
			indexesOfWrongs = append(indexesOfWrongs, i)
		}
	}
	fmt.Println("FINE TEST!")
	fmt.Printf("Corrette: %d/%d\n", rights, rights + wrongs)
	stampaErrate(indexesOfWrongs, questions)
}

func stampaErrate(indexes []int, questions []Question) {
	fmt.Println("Domande sbagliate: (Fai CTRL+F del testo della domanda nel file delle domande originali per capire quale fosse la risposta giusta)")
	for _, ind := range indexes {
		fmt.Println("-", questions[ind].question)
	}
}

//First bool is if the answer is correct or not, second one is if the user asked to stop the test
func askQuestion(q Question) (bool, bool) {
	var ans string
	var ansInt int
	var valid bool
	fmt.Println(q.question)
	correctAnswer, order := getRandomOrder()
	for i, answerIndex := range order {
		fmt.Printf("%c) %s\n", rune(int('A') + i), q.answers[answerIndex])
	}
	for {
		fmt.Scan(&ans)
		if ans == "EXIT" {
			return true, true
		}
		ansInt, valid = convertAnswer(ans)
		if valid {
			break
		} else {
			fmt.Println("Risposta non valida, riprova con A, B o C")
		}
	}
	return correctAnswer == ansInt, false
}

func getRandomOrder() (zero int, order []int) {
	order = []int{0, 1, 2}
	rand.Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i]})
	for i, v := range order {
		if v == 0 {
			zero = i
		}
	}
	return
}

func convertAnswer(answer string) (int, bool) {
	switch answer {
	case "A":
		return 0, true
	case "B":
		return 1, true
	case "C":
		return 2, true
	default:
		return -1, false
	}
}