package main

import (
	"fmt"
  //"time"
)

type DataNasc struct{
  ano int 
  mes string 
  dia int
}

type Notas struct{
  nota1 float32
  nota2 float32
}

type Aluno struct{
  nome string 
  dtNascimento DataNasc
  notas Notas
}

func (aluno Aluno) calcIdade(){
  //var data := time.Now(), 
  //idade := data.Year()-aluno.dtNascimento.ano
  var anoAtual int 
  fmt.Print("Insira o ano atual: ")
  fmt.Scan(&anoAtual)
  idade := anoAtual-aluno.dtNascimento.ano
  fmt.Print("Idade: ", idade) 
}

func (aluno Aluno) calcMedia(){
  fmt.Println("\nA média das notas deste aluno é: ", (aluno.notas.nota1+aluno.notas.nota2)/2)
}


func main() {

  var nome, mes string 
  var nt1, nt2 float32
  var ano, dia int 

  fmt.Print("Nome do aluno: ")
  fmt.Scan(&nome)
  fmt.Print("Sobre a data de nascimento, insira\n: dia, mês (por extenso) e ano: ")
  fmt.Scan(&dia)
  fmt.Scan(&mes)
  fmt.Scan(&ano)
  fmt.Print("Sobre as notas, insira\na primeira e a segunda nota deste aluno: ")
  fmt.Scan(&nt1)
  fmt.Scan(&nt2)

  notasAluno := Notas{
    nota1: nt1,
    nota2: nt2,
  }

  dt_nascAluno := DataNasc{
    dia: dia, 
    mes: mes, 
    ano: ano, 
  }
  
  aluno := Aluno{
    nome: nome,
    notas: notasAluno, 
    dtNascimento: dt_nascAluno,
  }

  aluno.calcIdade()
  aluno.calcMedia()

}
