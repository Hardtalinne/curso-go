package junina

import (
	"encoding/json"
	"errors"

	"github.com/hardtalinne/curso-go/domain"
)

//Junina retorno da função calcula
type Junina struct {
	TotalPessoas          int `json:"total-pessoas"`
	TotalParesAdultos     int `json:"total-pares-adultos"`
	ParesHomemMulher      int `json:"pares-homem-mulher"`
	ParesAdultosMesmoSexo int `json:"pares-adultos-mesmo-sexo"`
	ParesCriancas         int `json:"pares-criancas"`
	TotalAcompanhamentos  int `json:"total-acompanhamentos"`
	QuentoesNaoAlcoolicos int `json:"quentoes-nao-alcoolicos"`
	QuentoesAlcoolicos    int `json:"quentoes-alcoolicos"`
}

type calculoPares struct {
	totalPares       int
	paresHomemMulher int
	paresDoMesmoSexo int
}

func (c Junina) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

//Service define um serviço de festa junina
type Service struct{}

//NewJunina cria uma nova festa junina
func NewJunina() Service {
	return Service{}
}

//Calcula faz o cálculo da festa junina
func (s Service) Calcula(p domain.Parametros) (domain.Resultado, error) {
	junina := Junina{}

	if p.Homens < 5 || p.Mulheres < 5 {
		return junina, errors.New("Homens e mulheres devem ser igual ou maior que cinco")
	}

	junina.TotalPessoas = p.Mulheres + p.Homens + p.Criancas

	pares := calculaParesAdultos(p.Mulheres, p.Homens)

	junina.ParesHomemMulher = pares.paresHomemMulher
	junina.ParesAdultosMesmoSexo = pares.paresDoMesmoSexo
	junina.TotalParesAdultos = pares.totalPares

	junina.ParesCriancas = p.Criancas / 2

	if impar(p.Criancas) {
		junina.ParesCriancas = (p.Criancas - 1) / 2
	}

	junina.TotalAcompanhamentos = junina.TotalPessoas * 300
	junina.QuentoesNaoAlcoolicos = 400 * junina.TotalPessoas
	junina.QuentoesAlcoolicos = 500 * (p.Mulheres + p.Homens)

	return junina, nil
}

func impar(numero int) bool {
	return numero%2 != 0
}

func calculaParesAdultos(mulheres int, homens int) calculoPares {
	paresDoMesmoSexo := 0
	paresHomemMulher := mulheres
	restantes := homens - mulheres

	if mulheres > homens {
		paresHomemMulher = homens
		restantes = mulheres - homens
	}

	if restantes != 1 {
		paresDoMesmoSexo = restantes / 2

		if impar(restantes) {
			paresDoMesmoSexo = (restantes - 1) / 2
		}
	}

	totalPares := paresHomemMulher + paresDoMesmoSexo

	return calculoPares{totalPares, paresHomemMulher, paresDoMesmoSexo}
}
