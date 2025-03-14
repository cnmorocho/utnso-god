package entities

import (
	"cnmorocho/utnso-god/src/utils"
	"os"
	"reflect"
	"testing"
)

func createTempFile(t *testing.T, content string) *os.File {
	tmpFile, err := os.CreateTemp("", "mockInstructions")
	if err != nil {
		t.Fatalf("error creando archivo temporal: %v", err)
	}

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("error escribiendo en el archivo temporal: %v", err)
	}

	return tmpFile
}

func TestDecodeProgram(t *testing.T) {
	t.Run("Dado una ruta de archivo, cuando el archivo existe, entonces devuelve Queue[Instruction] con las instrucciones", func(t *testing.T) {
		mockFileContent := "mov bx, ax\nadd ax, cx"

		tmpFile := createTempFile(t, mockFileContent)
		defer os.Remove(tmpFile.Name())

		codePath := tmpFile.Name()
		got, err := DecodeProgram(codePath)
		firstInstruction := Instruction{Operator: "mov", Arguments: []string{"bx", "ax"}}
		secondInstruction := Instruction{Operator: "add", Arguments: []string{"ax", "cx"}}
		expectedQueue := utils.NewQueue[Instruction]()
		expectedQueue.Enqueue(firstInstruction)
		expectedQueue.Enqueue(secondInstruction)

		if err != nil || !reflect.DeepEqual(got, expectedQueue) {
			t.Errorf("got %v, want %v, err %v", got, expectedQueue, err)
		}
	})

	t.Run("Dado una ruta de archivo, cuando el archivo no existe, entonces devuelve Queue[Instruction] vacio", func(t *testing.T) {
		codePath := "noexiste.txt"
		got, err := DecodeProgram(codePath)
		firstInstruction := Instruction{Operator: "mov", Arguments: []string{"bx", "ax"}}
		expectedQueue := utils.NewQueue[Instruction]()
		expectedQueue.Enqueue(firstInstruction)

		if err == nil {
			t.Errorf("got %v, want %v, err %v", got, expectedQueue, err)
		}
	})
}
