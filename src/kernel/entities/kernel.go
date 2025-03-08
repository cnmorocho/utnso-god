package entities

import (
	"bufio"
	"cnmorocho/utnso-god/src/utils"
	"fmt"
	"os"
	"strings"
)

type Kernel struct {
	Scheduler *Scheduler
}

func GetKernelInstance(codePath string) *Kernel {
	return &Kernel{
		Scheduler: &Scheduler{},
	}
}

func DecodeProgram(codePath string) (utils.Queue[Instruction], error) {
	file, err := os.Open(codePath)
	if err != nil {
		return utils.Queue[Instruction]{}, fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer file.Close()

	code := bufio.NewScanner(file)
	instructions := utils.Queue[Instruction]{}

	for code.Scan() {
		parts := strings.Fields(code.Text())
		if len(parts) == 0 {
			continue
		}

		operator := parts[0]
		arguments := []string{}

		if len(parts) > 1 {
			arguments = strings.Split(strings.Join(parts[1:], ""), ",")
			for i := range arguments {
				arguments[i] = strings.TrimSpace(arguments[i])
			}
		}

		instructions.Enqueue(Instruction{Operator: operator, Arguments: arguments})
	}

	if err := code.Err(); err != nil {
		return utils.Queue[Instruction]{}, fmt.Errorf("error al leer el archivo: %w", err)
	}

	return instructions, nil
}
