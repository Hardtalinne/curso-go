package junina

import (
	"encoding/json"
	"testing"

	"github.com/hardtalinne/curso-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculaJuninaSemQuadrilha(t *testing.T) {
	s := NewJunina()
	_, err := s.Calcula(domain.Parametros{
		Homens:          2,
		Mulheres:        1,
		Criancas:        2,
		Acompanhamentos: true,
	})
	assert.Equal(t, err.Error(), "Homens e mulheres devem ser igual ou maior que cinco")
}

func TestCalculaJuninaBombando(t *testing.T) {
	s := NewJunina()
	ch, err := s.Calcula(domain.Parametros{
		Homens:          5,
		Mulheres:        8,
		Criancas:        15,
		Acompanhamentos: true,
	})
	assert.Nil(t, err)

	esperado := Junina{
		TotalPessoas:          28,
		TotalParesAdultos:     6,
		ParesHomemMulher:      5,
		ParesAdultosMesmoSexo: 1,
		ParesCriancas:         7,
		TotalAcompanhamentos:  8400,
		QuentoesNaoAlcoolicos: 11200,
		QuentoesAlcoolicos:    6500,
	}

	j, err := ch.ToJSON()
	jEsperado, err := json.Marshal(esperado)

	assert.Equal(t, ch, esperado)
	assert.Equal(t, j, jEsperado)
}
