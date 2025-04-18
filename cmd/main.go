package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type Response struct {
	Data  any   
	Error error
}

func main() {
	cep := "01153000"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resCh := make(chan Response, 2)

	go buscarEnderecoViaCEP(ctx, cep, resCh)
	go buscarEnderecoBrasilAPI(ctx, cep, resCh)

	// Aguarda as respostas
	select {
	case res := <-resCh:
		if res.Error != nil {
			slog.Error("request failed", "error", res.Error)
			return
		}
		showResponse(res)
	case <-ctx.Done():
		slog.Error("request timed out", "error", ctx.Err())
	}

}

// buscarUsandoBrasilAPI busca informações de endereço usando a API ViaCEP
func buscarEnderecoViaCEP(ctx context.Context, cep string, resCh chan Response) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		resCh <- Response{Error: err}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		resCh <- Response{Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resCh <- Response{Error: fmt.Errorf("viaCEP api returned status code %d", resp.StatusCode)}
		return
	}

	var viaCEPResponse ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResponse); err != nil {
		resCh <- Response{Error: err}
		return
	}

	resCh <- Response{Data: viaCEPResponse}

}

// buscarUsandoBrasilAPI busca informações de endereço usando a API BrasilAPI
func buscarEnderecoBrasilAPI(ctx context.Context, cep string, resCh chan Response) {

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		resCh <- Response{Error: err}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		resCh <- Response{Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resCh <- Response{Error: fmt.Errorf("brasilAPI returned status code %d", resp.StatusCode)}
		return
	}

	var brasilAPIResponse BrasilAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&brasilAPIResponse); err != nil {
		resCh <- Response{Error: err}
		return
	}

	resCh <- Response{Data: brasilAPIResponse}

}

// showResponse exibe a resposta da API
// de acordo com o tipo de dado recebido
func showResponse(res Response) {
	switch data := res.Data.(type) {
	case ViaCEPResponse:
		slog.Info("ViaCEP response", "data", data)
	case BrasilAPIResponse:
		slog.Info("BrasilAPI response", "data", data)
	}
}