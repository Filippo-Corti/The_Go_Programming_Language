/**
 * "libreria" di test per gli esami, non modificare questo file!
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func LanciaGenerica(
	t *testing.T,
	nomeProg string,
	inputString string,
	expectedOutString string,
	args ...string) {

	fmt.Println()
	fmt.Println("[ Questo test confronta l'output con l'output atteso ]")
	fmt.Println("[ Presuppone che l'eseguibile da testare (", nomeProg, ") sia già stato compilato! ]")
	fmt.Println()

	subproc := exec.Command(nomeProg, args...)
	subproc.Stdin = strings.NewReader(inputString)
	stdout, err := subproc.CombinedOutput()

	if err != nil {
		// verificare se uscito per esplicito os.Exit()
		//t.Errorf("Fallito: %s\n", stderr)

		//fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se l'esercizio richiedeva uscita esplicita con exit code)\n", err)
		fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se si termina il programma con un esplicito os.Exit\n", err)
	}

	fmt.Println()
	fmt.Printf("/// Args:\n%s\n", args)
	fmt.Printf("\n/// Input:\n%s\n", inputString)
	fmt.Printf("\n/// Output:\n%s\n", string(stdout))
	fmt.Printf("\n/// Output atteso:\n%s\n", expectedOutString)
	if string(stdout) != expectedOutString {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}

	subproc.Process.Kill()
	fmt.Println()
}

func LanciaGenericaConFileOutAtteso(
	t *testing.T,
	nomeProg string,
	inputString string,
	expectedOutFilename string,
	args ...string) {

	content, err := ioutil.ReadFile(expectedOutFilename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	//fmt.Println(text)

	LanciaGenerica(t, nomeProg, inputString, text, args...)
}

func LanciaGenericaConFileInOutAtteso(
	t *testing.T,
	nomeProg string,
	inputFilename string,
	expectedOutFilename string,
	args ...string) {

	input, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	in := string(input)

	exout, err := ioutil.ReadFile(expectedOutFilename)
	if err != nil {
		log.Fatal(err)
	}
	out := string(exout)
	//fmt.Println(text)

	LanciaGenerica(t, nomeProg, in, out, args...)
}
