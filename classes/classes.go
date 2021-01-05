package main

import (
	"fmt"
	"os"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutBR  = "02/01/2006"
)

func gerarDatas(inicio time.Time, fim time.Time) []time.Time {
	var datas []time.Time

	for d := inicio; d.Before(fim); d = d.AddDate(0, 0, 7) {
		datas = append(datas, d)
		datas = append(datas, d.AddDate(0, 0, 1))
	}

	return datas
}

func main() {
	argumentos := os.Args[1:]
	if argumentos[0] == "datas" {
		inicio, err := time.Parse(layoutBR, os.Args[2])
		if err != nil {
			fmt.Println("Problema no formato da data.")
			os.Exit(1)
		}

		fim, err := time.Parse(layoutBR, os.Args[3])
		if err != nil {
			fmt.Println("Problema no formato da data.")
		}

		for _, d := range gerarDatas(inicio, fim) {
			fmt.Println(d.Format(layoutBR))
		}
	} else if argumentos[0] == "pagina" {
		fmt.Println("Ainda não implementado.")
	}
	/*
		https://gowebexamples.com/templates/
		https://github.com/360EntSecGroup-Skylar/excelize
		data := PaginaDisciplina{
			Nome: "Programação de Scripts 2020.2",
			Links: []Link{
				{Url: "https://join.slack.com/t/scripts20202/shared_invite/zt-jcjzxokl-7ZemQZm6DpytEmIG~3Cdfg", Descricao: "Grupo do Slack"},
			},
			Aulas: []Aula{
				{
					Data1:           "23/11/2020",
					Assunto:         "Introdução",
					Comentario1:     "O que é Shell Scripting, história, diferentes tipos de Shells, etc.",
					Video:           "http://teste.com.br",
					Atividade:       "http://teste.com.br",
					NumeroAtividade: "Atividade 01",
					Data2:           "24/11/2020",
					Comentario2:     "Atividade de como usar o SSH/SCP",
				},
			},
		}
	*/
	return
}
