package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: rename-files <extensión_actual> <nueva_extensión>")
		os.Exit(1)
	}

	currentExtension := os.Args[1]
	newExtension := os.Args[2]
	allFiles, err := filepath.Glob("*." + currentExtension)
	if err != nil {
		fmt.Println("Error al obtener la lista de archivos:", err)
		os.Exit(1)
	}

	for _, fileItem := range allFiles {
		fileWithOutExt := strings.TrimSuffix(fileItem, "."+currentExtension)
		newFileNameWithExt := fileWithOutExt + "." + newExtension
		err := os.Rename(fileItem, newFileNameWithExt)
		if err != nil {
			fmt.Printf("Error al cambiar la extensión de %s a %s: %s\n", fileItem, newFileNameWithExt, err)
		} else {
			fmt.Printf("Se ha cambiado la extensión de %s a %s\n", fileItem, newFileNameWithExt)
		}
	}
}
