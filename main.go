package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

// Estruturas e constantes
type Moeda struct {
	Codigo string
	Nome   string
	Prompt string
}

var moedas = []Moeda{
	{"USD", "Dólar Americano", "$"},
	{"EUR", "Euro", "€"},
	{"GBP", "Libra Esterlina", "£"},
	{"JPY", "Iene Japonês", "¥"},
	{"BRL", "Real Brasileiro", "R$"},
}

const (
	apiPrimaria   = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies"
	apiSecundaria = "https://latest.currency-api.pages.dev/v1/currencies"
)

// Função principal
func main() {
	for {
		moedaOrigem := selecionarMoeda("Selecione a moeda de origem")
		moedaDestino := selecionarMoeda("Selecione a moeda de destino")
		valor := obterValor(getCurrencyPrompt(moedaOrigem))

		realizarConversao(moedaOrigem, moedaDestino, valor)

		if !desejaConverterNovamente() {
			break
		}
	}
}

// Funções auxiliares
func selecionarMoeda(titulo string) string {
	var moedaSelecionada string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(titulo).
				Options(moedasParaOpcoes()...).
				Value(&moedaSelecionada),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	return moedaSelecionada
}

func obterValor(prompt string) float64 {
	var valorString string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Digite o valor a ser convertido").
				Placeholder("Ex: 100").
				Prompt(prompt).
				Validate(validarValor).
				Value(&valorString),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	valor, _ := strconv.ParseFloat(valorString, 64)
	return valor
}

func realizarConversao(origem, destino string, valor float64) {
	taxa, err := obterTaxaDeCambio(origem, destino)
	if err != nil {
		log.Fatal(err)
	}

	resultado := valor * taxa

	fmt.Printf("%.2f %s = %.2f %s\n", valor, origem, resultado, destino)
	fmt.Printf("Taxa de câmbio: %.4f\n", taxa)
}

func desejaConverterNovamente() bool {
	var continuar string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Deseja converter outro valor?").
				Options(huh.NewOption("Sim", "s"), huh.NewOption("Não", "n")).
				Value(&continuar),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	return continuar == "s"
}

func getCurrencyPrompt(moeda string) string {
	for _, m := range moedas {
		if m.Codigo == moeda {
			return m.Prompt
		}
	}
	return ""
}

func moedasParaOpcoes() []huh.Option[string] {
	opcoes := make([]huh.Option[string], len(moedas))
	for i, moeda := range moedas {
		opcoes[i] = huh.NewOption(moeda.Nome, moeda.Codigo)
	}
	return opcoes
}

func obterTaxaDeCambio(origem, destino string) (float64, error) {
	taxa, err := fetchExchangeRate(origem, destino, apiPrimaria)
	if err != nil {
		taxa, err = fetchExchangeRate(origem, destino, apiSecundaria)
		if err != nil {
			return 0, err
		}
	}
	return taxa, nil
}

func fetchExchangeRate(origem, destino, baseURL string) (float64, error) {
	url := fmt.Sprintf("%s/%s.json", baseURL, strings.ToLower(origem))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", "CurrencyConverter/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("erro na requisição: status %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	taxaInterface, ok := data[strings.ToLower(origem)].(map[string]interface{})[strings.ToLower(destino)]
	if !ok {
		return 0, fmt.Errorf("não foi possível encontrar a taxa de câmbio para %s", destino)
	}

	taxa, ok := taxaInterface.(float64)
	if !ok {
		return 0, fmt.Errorf("valor da taxa não é um número válido")
	}

	return taxa, nil
}

func validarValor(s string) error {
	if s == "" {
		return fmt.Errorf("o valor não pode estar vazio")
	}
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("valor inválido: %v", err)
	}
	return nil
}
